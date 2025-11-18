.PHONY: build install test clean release help

# Variables
BINARY_NAME=doplan
VERSION?=dev
COMMIT?=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

# Build flags
LDFLAGS=-ldflags "-X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(BUILD_TIME)"

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the binary
	@echo "Building $(BINARY_NAME)..."
	go build $(LDFLAGS) -o bin/$(BINARY_NAME) ./cmd/doplan

install: build ## Install the binary
	@echo "Installing $(BINARY_NAME)..."
	go install $(LDFLAGS) ./cmd/doplan

test: ## Run tests
	@echo "Running tests..."
	go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -rf bin/
	rm -f coverage.out coverage.html

release: ## Build release binaries
	@echo "Building release binaries..."
	goreleaser build --snapshot --clean

release-full: ## Full release with goreleaser
	@echo "Creating full release..."
	goreleaser release --clean

dev: ## Run in development mode
	@echo "Running in development mode..."
	go run $(LDFLAGS) ./cmd/doplan

fmt: ## Format code
	@echo "Formatting code..."
	go fmt ./...

vet: ## Run go vet
	@echo "Running go vet..."
	go vet ./...

lint: fmt vet ## Run linters

test-scripts: ## Run all test scripts
	@echo "Running CLI tests..."
	@./scripts/test-cli.sh || true
	@echo ""
	@echo "Running installation tests..."
	@./scripts/test-install.sh || true
	@echo ""
	@echo "Running integration tests..."
	@./scripts/test-integration.sh || true

release-prep: ## Prepare for release (tests, lint, build)
	@./scripts/pre-release.sh

check-production: ## Check if ready for production release
	@./scripts/check-production-ready.sh

check-ready: check-production ## Alias for check-production

release-create: ## Create a new release (usage: make release-create VERSION=v1.0.0)
	@if [ -z "$(VERSION)" ]; then \
		echo "Usage: make release-create VERSION=v1.0.0"; \
		exit 1; \
	fi
	@./scripts/release.sh $(VERSION)

release-snapshot: ## Create a snapshot release
	@goreleaser release --snapshot --clean

.DEFAULT_GOAL := help

