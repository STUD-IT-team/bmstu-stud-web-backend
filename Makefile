GO := go
SWAGGER := swag

APP := ./cmd/app
SWAGMAINFILE := cmd/app/main.go
DOCS := cmd/app/docs

DOCKER:=docker
COMPOSE:=deployments/docker-compose.yaml

.PHONY: run

run:
	$(GO) run $(APP)

.PHONY: swag-generate

swag-generate:
	$(SWAGGER) init --parseDependency --parseInternal -o $(DOCS) -g $(SWAGMAINFILE)

.PHONY: lint

lint:
	golangci-lint run

.PHONY: up

up:
	$(DOCKER) compose -f $(COMPOSE) up  -d

.PHONY: upd

upd:
	$(DOCKER) compose -f $(COMPOSE) up -d --build

upda: 
	$(DOCKER) compose -f $(COMPOSE) up  --build

.PHONY: down

down:
	$(DOCKER) compose -f $(COMPOSE) down