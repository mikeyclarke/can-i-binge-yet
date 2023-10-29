THIS_FILE := $(lastword $(MAKEFILE_LIST))
PARENT_DIR := $(dir $(MKFILE_PATH))

## ---------
##	Docker images
## ---------

build_docker_image:
	docker build -t can-i-binge-yet-python .

## ---------
##	Container management
## ---------

create_containers: ## Create docker containers
	docker compose up -d

destroy_containers: ## Destroy docker containers
	docker compose down

start_containers: ## Start docker containers
	docker compose start

stop_containers: ## Stop docker containers
	docker compose stop

container_shell: ## Open an interactive shell in the main Go container
	docker exec -it ciby-go /bin/sh

## ---------
##	Testing
## ---------

generate_mocks: ## Generate Go mocks with Mockery
	docker run --rm -ti -v .:/var/app -w /var/app vektra/mockery

test: test_go test_ts ## Run all code tests

test_go: gotest vet staticcheck ## Run all go code tests

test_ts: jest ## Run all TypeScript code tests

## ---------
##	Linting
## ---------

lint_go: staticcheck vet ## Run go code quality checks

## ---------
##	Coding standards
## ---------

vet: ## Vet go code
	docker exec -it ciby-go go vet ./src/go/...

## ---------
##	Static analysis
## ---------

staticcheck: ## Check that go code passes checks
	docker exec -it ciby-go staticcheck ./src/go/...

## ---------
##	Unit tests
## ---------

gotest: ## Run go unit tests
	docker exec -it ciby-go go test ./src/go/...

jest: ## Run TypeScript unit tests
	docker exec -it ciby-node yarn run jest --verbose --silent=false $(TEST_REGEXP)

## ---------
##	Dependencies
## ---------

add_yarn_package: ## Add a yarn package
	docker compose run --rm node yarn add $(PACKAGE)

add_yarn_dev_package: ## Add a yarn package
	docker compose run --rm node yarn add --dev $(PACKAGE)

list_outdated_go_packages: ## List outdated golang dependencies
	docker exec -it ciby-go sh -c "go list -u -m -json all | go-mod-outdated -update -direct"

update_go_packages: ## Update all golang dependencies to their latest minor version
	docker exec -it ciby-go go get -u ./...

## ---------
##	Dependency injection
## ---------

generate_go_di_container: ## Have wire generate the go dependency injection container
	docker exec -it ciby-go wire ./src/go

## ---------
##	Asset compilation
## ---------

webpack_watch: ## Have webpack watch source files and re-compile assets on change
	docker compose run --rm node yarn run webpack --watch --config webpack.dev.js

webpack_compile: ## Have webpack execute a one-off asset compilation
	docker compose run --rm node yarn run webpack --config webpack.dev.js

webpack_production_build: ## Have webpack execute a one-off asset compilation
	docker compose run --rm node yarn run webpack --config webpack.production.js

## ---------
##	Deployment
## ---------

deploy:
	fly deploy --ignorefile .dockerignore.production

## ---------
#@	Make setup
## ---------

_PHONY: help

.DEFAULT_GOAL := help

help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-30s\033[0m %s\n", $$1, $$2 } /^##[[:blank:]]{1,2}[a-zA-Z0-9]+/ { str = $$0; sub(/^##[[:blank:]]{1,2}/, "", str); printf "\n\033[34m%s\033[0m\n", str } ' $(MAKEFILE_LIST)
