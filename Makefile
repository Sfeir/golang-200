# Makefile for Todolist : Go-200

.DEFAULT_GOAL := all

# -----------------------------------------------------------------
#
#        ENV VARIABLE
#
# -----------------------------------------------------------------

# check os for path params
ifeq ($(OS),Windows_NT)
	PATHSEP=;
	FOLDERSEP=\\
	EXTENSION=.exe
else
	PATHSEP=:
	FOLDERSEP=/
	EXTENSION=""
endif

# go env vars
GO=$(firstword $(subst $(PATHSEP), ,$(GOPATH)))
# list of pkgs for the project without vendor
PKGS=$(shell go list ./... | grep -v /vendor/)
DOCKER_IP=$(shell if [ -z "$(DOCKER_MACHINE_NAME)" ]; then echo 'localhost'; else docker-machine ip $(DOCKER_MACHINE_NAME); fi)
export GO15VENDOREXPERIMENT=1

# -----------------------------------------------------------------
#        Version
# -----------------------------------------------------------------

# version
VERSION=0.0.2
BUILDDATE=$(shell date -u '+%s')
BUILDHASH=$(shell git rev-parse --short HEAD)
VERSION_FLAG=-ldflags "-X main.Version=$(VERSION) -X main.GitHash=$(BUILDHASH) -X main.BuildStmp=$(BUILDDATE)"

# -----------------------------------------------------------------
#        Main targets
# -----------------------------------------------------------------

all: clean build ## Clean and build the project

clean: ## Clean the project
	@go clean
	@rm -Rf .tmp .DS_Store *.log *.out *.mem *.test build/

build: format ## Build all libraries and binaries
	@go build -v $(VERSION_FLAG) -o "$(GO)$(FOLDERSEP)bin$(FOLDERSEP)todolist$(EXTENSION)" todolist.go

format: ## Format all packages
	@go fmt $(PKGS)

teardownTest: ## Tear down databases for integration tests
	@docker-compose -f docker/databases.yml down

setupTest: teardownTest ## Start databases for integration tests
	@docker-compose -f docker/databases.yml up -d

test: setupTest ## Start tests with a databases docker image
	@export DB_HOST=$(DOCKER_IP); sleep 2; go test -v $(PKGS); make teardownTest

bench: setupTest ## Start benchmark
	@go test -v -run nothing -bench=. -memprofile=prof.mem github.com/Sfeir/golang-200/web ; make teardownTest

benchTool: bench ## Start benchmark tool
	@echo "### TIP : type 'top 5' and 'list path_of_the_first_item'"
	@go tool pprof --alloc_space web.test prof.mem

lint: ## Lint all packages
	@golint dao/...
	@golint model/...
	@golint web/...
	@golint utils/...
	@golint ./.
	@go vet $(PKGS)

# -----------------------------------------------------------------
#        Docker targets
# -----------------------------------------------------------------

dockerBuild: ## Build a docker image of the program
	docker build -t -f docker/DockerFile sfeir/todolist:latest .

dockerBuildMulti: ## Build a docker multistep image of the program
	docker build -f docker/Dockerfile.multi -t sfeir/todolist:latest .

dockerClean: ## Remove the docker image of the program
	docker rmi -f sfeir/todolist:latest

dockerUp: ## Start the program instances with their databases
	docker-compose -f docker/docker-compose.yml up -d

dockerDown: ## Stop the program instances, their databases and remove the containers
	docker-compose -f docker/docker-compose.yml down

dockerBuildUp: dockerDown dockerBuild dockerUp ## Stop, build and launch the docker images of the program

dockerBuildUpMulti: dockerDown dockerBuildMulti dockerUp ## Stop, build multi step and launch the docker images of the program

dockerWatch: ## Watch the status of the docker container
	@watch -n1 'docker ps | grep todolist'

dockerLogs: ## Print the logs of the container
	docker-compose logs -f

help: ## Print this message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: all test clean teardownTest setupTest
