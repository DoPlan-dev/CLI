.PHONY: build build-all clean test install help version

# Version (can be overridden)
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")

# Build directory
BUILD_DIR = dist
BINARY_NAME = doplan

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

version: ## Show version information
	@echo "Version: $(VERSION)"

build: ## Build for current platform
	@echo "Building DoPlan CLI v$(VERSION) for $(shell go env GOOS)/$(shell go env GOARCH)..."
	@go build -ldflags "-X github.com/doplan/cli/internal/version.Version=$(VERSION)" -o $(BINARY_NAME) ./cmd/doplan
	@echo "✓ Built $(BINARY_NAME)"

build-all: ## Build for all platforms
	@echo "Building DoPlan CLI v$(VERSION) for all platforms..."
	@bash scripts/build.sh

clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@rm -f $(BINARY_NAME)
	@rm -f $(BINARY_NAME).exe
	@echo "✓ Cleaned"

test: ## Run tests
	@echo "Running tests..."
	@go test ./... -v

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html
	@echo "✓ Coverage report generated: coverage.html"

install: build ## Install to GOPATH/bin
	@echo "Installing DoPlan CLI..."
	@go install -ldflags "-X github.com/doplan/cli/internal/version.Version=$(VERSION)" ./cmd/doplan
	@echo "✓ Installed to $(GOPATH)/bin/$(BINARY_NAME)"

lint: ## Run linter
	@echo "Running linter..."
	@golangci-lint run || echo "golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"

fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...
	@echo "✓ Formatted"

vet: ## Run go vet
	@echo "Running go vet..."
	@go vet ./...
	@echo "✓ No issues found"

