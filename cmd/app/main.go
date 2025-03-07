package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/go-co-op/gocron/v2"
	"github.com/sirupsen/logrus"
	"go.uber.org/automaxprocs/maxprocs"
	"golang.org/x/sync/errgroup"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/cmd/configer/appconfig"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infrastructure/cache"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infrastructure/miniostorage"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infrastructure/postgres"
	internalhttp "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/ports"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/storage"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/handler"
)

//go:generate go run ../configer --apps api --envs local,prod,dev

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// Storage
	appPostgres, err := postgres.NewPostgres(
		os.Getenv("PG_CONNECT"),
	)
	if err != nil {
		logger.WithError(err).Errorf("can`t connect to postgres: %s", os.Getenv("PG_CONNECT"))
	}

	endpoint := os.Getenv("END_POINT")
	user := os.Getenv("MINIO_ROOT_USER")
	password := os.Getenv("MINIO_ROOT_PASSWORD")
	useSSL := false

	minioStorage, err := miniostorage.NewMinioStorage(endpoint, user, password, useSSL)

	if err != nil {
		log.Fatalf("Upload err: " + err.Error())
	}

	sessionCache := cache.NewSessionCache()
	appStorage := storage.NewStorage(*appPostgres, sessionCache, minioStorage)

	// services
	mediaService := app.NewMediaService(appStorage)
	clubService := app.NewClubService(appStorage, mediaService)
	feedService := app.NewFeedService(appStorage)
	eventsService := app.NewEventsService(appStorage)
	membersService := app.NewMembersService(appStorage)
	guardService := app.NewGuardService(appStorage)
	documentsService := app.NewDocumentsService(appStorage)
	documentCategoriesSevice := app.NewDocumentCategoriesService(appStorage)
	apiService := app.NewAPI(logger, feedService, eventsService, membersService, clubService, guardService, documentsService, documentCategoriesSevice)

	var mainGroupHandler *handler.GroupHandler
	// Main API router.
	if cfg.Log.Level == "debug" {
		mainGroupHandler = handler.NewGroupHandler("/",
			internalhttp.NewAPIHandler(jsonRenderer, apiService),
			internalhttp.NewGuardHandler(jsonRenderer, *guardService, logger),
			internalhttp.NewClubsHandler(jsonRenderer, *clubService, logger, guardService),
			internalhttp.NewFeedHandler(jsonRenderer, *feedService, logger, guardService),
			internalhttp.NewEventsHandler(jsonRenderer, *eventsService, logger, guardService),
			internalhttp.NewMembersHandler(jsonRenderer, *membersService, logger, guardService),
			internalhttp.NewMediaHandler(jsonRenderer, mediaService, logger, guardService),
			internalhttp.NewDocumentsHandler(jsonRenderer, *documentsService, *documentCategoriesSevice, logger, guardService),
			internalhttp.NewSwagHandler(jsonRenderer),
		)
	} else {
		mainGroupHandler = handler.NewGroupHandler("/",
			internalhttp.NewAPIHandler(jsonRenderer, apiService),
			internalhttp.NewGuardHandler(jsonRenderer, *guardService, logger),
			internalhttp.NewClubsHandler(jsonRenderer, *clubService, logger, guardService),
			internalhttp.NewFeedHandler(jsonRenderer, *feedService, logger, guardService),
			internalhttp.NewEventsHandler(jsonRenderer, *eventsService, logger, guardService),
			internalhttp.NewMembersHandler(jsonRenderer, *membersService, logger, guardService),
			internalhttp.NewMediaHandler(jsonRenderer, mediaService, logger, guardService),
			internalhttp.NewDocumentsHandler(jsonRenderer, *documentsService, *documentCategoriesSevice, logger, guardService),
			internalhttp.NewSwagHandler(jsonRenderer),
		)
	}

	mainHandler := handler.New(handler.MakePublicRoutes(
		router,
		handler.RoutesCfg{
			BasePath: cfg.Servers.Public.BasePath,
		},
		mainGroupHandler))

	// servers = append(servers, &http.Server{
	// 	Addr:     cfg.Servers.Public.ListenAddr,
	// 	Handler:  mainHandler,
	// 	ErrorLog: log.New(logger.Out, "api", 0),
	// })

	logger.Debugf("Listing actual routes:\n")

	_ = chi.Walk(
		router,
		func(
			method string,
			route string,
			handler http.Handler,
			middlewares ...func(http.Handler) http.Handler,
		) error {
			logger.Debugf("[%s]: /%s%s\n", method, appconfig.APIAppName, route)
			return nil
		})

	server := &http.Server{
		Addr:     cfg.Servers.Public.ListenAddr,
		Handler:  mainHandler,
		ErrorLog: log.New(logger.Out, "api", 0),
	}
	logger.Info(server.Addr)
	go func() {
		logger.Infof("starting server, listening on %s", server.Addr)

		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			logger.WithError(err).Errorf("server can't listen and serve requests")
		}
	}()

	s, err := gocron.NewScheduler()
	if err != nil {
		logger.Fatalf("failed to create gocron scheduler")
	}

	cleanupTime := gocron.NewAtTime(2, 0, 0)
	cleanupJob := gocron.DailyJob(1, gocron.NewAtTimes(cleanupTime))

	_, err = s.NewJob(
		cleanupJob,
		gocron.NewTask(mediaService.ClearUpMedia, context.Background(), logger),
	)
	if err != nil {
		logger.Fatalf("gocron обдристался at ClearMediaStorages...")
	}
	_, err = s.NewJob(
		cleanupJob,
		gocron.NewTask(documentsService.CleanupDocuments, context.Background(), logger),
	)
	if err != nil {
		logger.Fatalf("gocron обдристался at ClearUnknownDocuments...")
	}
	s.Start()
	logger.Infof("Cron started")

	// for i := range servers {
	// 	srv := servers[i]
	// 	go func() {
	// 		logger.Infof("starting server, listening on %s", srv.Addr)

	// 		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
	// 			logger.WithError(err).Errorf("server can't listen and serve requests")
	// 		}
	// 	}()
	// }

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

	if err := server.Shutdown(timeoutCtx); err != nil {
		logger.WithError(err).Fatalf("can't close server listening on '%s'", server.Addr)
	}

	logger.Info("app has been terminated")
}
