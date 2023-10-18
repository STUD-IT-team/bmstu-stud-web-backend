package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"go.uber.org/automaxprocs/maxprocs"
	"golang.org/x/sync/errgroup"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/cmd/configer/appconfig"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app"
	internalhttp "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/ports/http"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/handler"
)

//go:generate go run ../configer --apps api --envs local,prod,dev

func main() {
	cfg := appconfig.MustParseAppConfig[appconfig.APIConfig]()

	logger := logrus.New()

	lvl, err := logrus.ParseLevel(cfg.Log.Level)
	if err != nil {
		logger.Fatal(err)
	}

	logger.SetLevel(lvl)

	if _, err := maxprocs.Set(maxprocs.Logger(logger.Infof)); err != nil {
		logger.WithError(err).Errorf("can't set maxprocs")
	}

	eg, ctx := errgroup.WithContext(context.Background())

	// HTTP servers.
	jsonRenderer := handler.NewJSONRenderer()

	servers := make([]*http.Server, 0)
	router := chi.NewRouter()

	// services
	apiService := app.NewAPI(logger)

	// Main API router.
	mainGroupHandler := handler.NewGroupHandler("/",
		internalhttp.NewAPIHandler(jsonRenderer, apiService),
	)

	mainHandler := handler.New(handler.MakePublicRoutes(
		router,
		handler.RoutesCfg{
			BasePath: cfg.Servers.Public.BasePath,
		},
		mainGroupHandler))

	servers = append(servers, &http.Server{
		Addr:     cfg.Servers.Public.ListenAddr,
		Handler:  mainHandler,
		ErrorLog: log.New(logger.Out, "api", 0),
	})

	fmt.Printf("Listing actual routes:\n")
	chi.Walk(router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("[%s]: /bmstu-stud-web/%s%s\n", method, appconfig.APIAppName, route)
		return nil
	})

	for i := range servers {
		srv := servers[i]
		go func() {
			logger.Infof("starting server, listening on %s", srv.Addr)

			if err := srv.ListenAndServe(); err != http.ErrServerClosed {
				logger.WithError(err).Errorf("server can't listen and serve requests")
			}
		}()
	}

	logger.Infof("app started")

	sigQuit := make(chan os.Signal, 1)

	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGTERM)

	eg.Go(func() error {
		select {
		case s := <-sigQuit:
			return fmt.Errorf("captured signal: %v", s)
		case <-ctx.Done():
			return nil
		}
	})

	if err = eg.Wait(); err != nil {
		logger.WithError(err).Infof("gracefully shutting down the server")
	}

	timeoutCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()
	for _, srv := range servers {
		if err := srv.Shutdown(timeoutCtx); err != nil {
			logger.WithError(err).Fatalf("can't close server listening on '%s'", srv.Addr)
		}
	}

	logger.Info("app has been terminated")
}
