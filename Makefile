# VARIABLES
# =======================================================================================
# GO
include .env
GOCMD=go
GOTEST=$(GOCMD) test ./...
# Machine OS
OS := $(shell uname)
WD := $(shell pwd)
OS_LOWERCASE := $(shell uname | tr '[:upper:]' '[:lower:]')

# INSTALL TARGETS
# =======================================================================================
swagger-install: # Install Go-Swagger
ifeq ($(OS), Darwin)
	brew tap go-swagger/go-swagger
	brew install go-swagger
else
	sudo apt-get install jq

	download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url')
	curl -o /usr/local/bin/swagger -L'#' "$download_url"
	chmod +x /usr/local/bin/swagger

endif
	swagger version

precommit-hook:  # Golang Pre-Commit Hook Installation ##https://pre-commit.com/#cli
ifeq ($(OS), Darwin)
	brew install pre-commit
else	
	sudo pip install pre-commit
endif
	pre-commit --version
	pre-commit install

golangci-lint: # Installing Magic golangci-lint 
ifeq ($(OS), Darwin)
	# Run MacOS commands 
	brew install golangci-lint
	brew upgrade golangci-lint
else
	# check for Linux and run other commands
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.45.2
endif
	golangci-lint --version
	golangci-lint linters -E bodyclose
	golangci-lint linters -E gocyclo
	golangci-lint linters -E gocritic
	golangci-lint linters -E goimports
	golangci-lint linters -E goconst
	golangci-lint linters -E sqlclosecheck
	golangci-lint linters -E lll
	golangci-lint linters -E funlen
	golangci-lint linters -E godot
	golangci-lint linters -E exportloopref
	golangci-lint linters -D scopelint

# GIT PRECOMMIT INSTALL
prehook-install: precommit-hook golangci-lint # Install Prehook along with Linters

# GOTOOLS

vet: # Vet examines Go source code and reports suspicious constructs
	${GOCMD} vet

fmt: # Gofmt is a tool that automatically formats Go source code
	gofmt

test: # GO Test
	$(GOTEST)

cover: # Go Test Coverage
	${GOCMD} test -coverprofile=coverage.out ./... && ${GOCMD} tool cover -html=coverage.out

tidy: # Update Modules and Dependency Consistency
	${GOCMD} mod tidy

build: # Builds the project
	${GOCMD} build main.go

run: # Builds the project
	${GOCMD} run main.go

lint: # Lint the files
	golangci-lint run --skip-dirs docs

# SWAGGER

check-swagger:
	which swagger || (go get -u github.com/go-swagger/go-swagger/internal/swagger)

gen-swagger:  # Generate Swagger API Documentation
	swagger generate spec -o $(WD)/docs/swagger.json  --scan-models
	swagger generate spec -o $(WD)/swagger-ui/swagger.json  --scan-models

serve-swagger: check-swagger  # Serve Swagger API Documentation
	swagger validate $(WD)/docs/swagger.json
	swagger serve -F=swagger $(WD)/docs/swagger.json

swagger: gen-swagger serve-swagger	# Generate & Serve Swagger API Documentation

# MIGRATIONS

migrate-up: 
	docker run -v $(PWD)/db/migrations:/migrations --network host --rm migrate/migrate -path=/migrations/ -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable&search_path=$(POSTGRES_SCHEMA)" up
migrate-down: 
	docker run -v $(PWD)/db/migrations:/migrations --network host --rm migrate/migrate -path=/migrations/ -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable&search_path=$(POSTGRES_SCHEMA)" down --all

linter:
	docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.54.2 golangci-lint run -v