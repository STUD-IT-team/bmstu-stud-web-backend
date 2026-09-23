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

.PHONY: front-back-network

# stud-front-back is declared external in the compose file so this backend can
# share a network with a separately-run frontend stack; docker compose won't
# create it for you, so make sure it exists before bringing the stack up.
front-back-network:
	$(DOCKER) network inspect stud-front-back >/dev/null 2>&1 || $(DOCKER) network create stud-front-back

.PHONY: up

up: front-back-network
	$(DOCKER) compose -f $(COMPOSE) up  -d

.PHONY: upd

upd: front-back-network
	$(DOCKER) compose -f $(COMPOSE) up -d --build

upda: front-back-network
	$(DOCKER) compose -f $(COMPOSE) up  --build

.PHONY: build

build: front-back-network
	$(DOCKER) compose -f $(COMPOSE) build

.PHONY: down

down:
	$(DOCKER) compose -f $(COMPOSE) down