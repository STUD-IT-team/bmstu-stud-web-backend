GOROOT ?= /usr/local/go
GOPATH := $(shell go env GOPATH)

BINARIES_DIR := cmd
SERVICES_LIST := $(shell find $(BINARIES_DIR) -maxdepth 1 \( ! -iname "$(BINARIES_DIR)" \) -type d -exec basename {} \;)
SERVICES_RUN_TARGETS_LIST := $(addprefix run-, $(SERVICES_LIST))

ENV := $(if $(ENV),$(ENV),local)

_gen:
	go generate ./...
.PHONY: _gen

gen: ## run code generation on host machine
	@echo "+ $@"
	@$(MAKE) _gen
.PHONY: local-gen

gen_proto:
	protoc -I internal/ports/grpc internal/ports/grpc/guard.proto --go_out=./internal/ports/grpc --go_opt=paths=source_relative --go-grpc_out=./internal/ports/grpc --go-grpc_opt=paths=source_relative

$(SERVICES_RUN_TARGETS_LIST): run-%: ## run service from $(BINARIES_DIR)
	go run ./cmd/$* --config-path=./cmd/$*/infra/$(ENV)/application.conf
.PHONY: $(SERVICES_RUN_TARGETS_LIST)

# need v1.54.2 of golangci-lint
lint:
	golangci-lint run -v -c golangci.yml ./...

migration-create-sql:
	goose -dir=./migrations create $(NAME) sql
