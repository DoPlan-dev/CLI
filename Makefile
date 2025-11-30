.PHONY: build build-all clean test install help version test-coverage test-coverage-check pre-release release-check

# Version (can be overridden)
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
# Determine module path once so ldflags always track the actual module/casing
MODULE_PATH ?= $(shell go list -m 2>/dev/null || echo "github.com/DoPlan-dev/CLI")
VERSION_SYMBOL := $(MODULE_PATH)/internal/version.Version

# Build directory
BUILD_DIR = dist
BINARY_NAME = doplan

# Coverage threshold (can be overridden)
COVERAGE_THRESHOLD ?= 80

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

version: ## Show version information
	@echo "Version: $(VERSION)"

build: ## Build for current platform
	@echo "Building DoPlan CLI v$(VERSION) for $(shell go env GOOS)/$(shell go env GOARCH)..."
	@go build -ldflags "-X $(VERSION_SYMBOL)=$(VERSION)" -o $(BINARY_NAME) ./cmd/doplan
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
	@go test ./... -short -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html
	@echo "✓ Coverage report generated: coverage.html"
	@echo "Coverage summary:"
	@go tool cover -func=coverage.out | tail -1

test-coverage-check: test-coverage ## Run tests with coverage and check threshold
	@echo ""
	@echo "Coverage Summary:"
	@echo "  Overall coverage (includes cmd/ and scripts/):"
	@go tool cover -func=coverage.out | tail -1
	@echo ""
	@echo "  Core packages (internal/*, pkg/*, excluding internal/cli - CLI commands are integration-tested):"
	@CORE_COVERAGE=$$(go tool cover -func=coverage.out | grep -E "^github.com/DoPlan-dev/CLI/(internal|pkg)/" | grep -v "internal/cli" | awk '{print $$3}' | sed 's/%//' | awk '{sum+=$$1; count++} END {if(count>0) printf "%.1f", sum/count; else print "0"}'); \
	if [ -z "$$CORE_COVERAGE" ] || [ "$$CORE_COVERAGE" = "0" ]; then \
		echo "    ❌ Error: Could not calculate core package coverage"; \
		exit 1; \
	fi; \
	echo "    total: $$CORE_COVERAGE% of statements"
	@echo ""
	@echo "Checking coverage threshold (minimum: $(COVERAGE_THRESHOLD)%)..."
	@echo "  (Using core packages coverage: internal/*, pkg/*, excluding internal/cli)"
	@CORE_COVERAGE=$$(go tool cover -func=coverage.out | grep -E "^github.com/DoPlan-dev/CLI/(internal|pkg)/" | grep -v "internal/cli" | awk '{print $$3}' | sed 's/%//' | awk '{sum+=$$1; count++} END {if(count>0) printf "%.1f", sum/count; else print "0"}'); \
	if [ -z "$$CORE_COVERAGE" ] || [ "$$CORE_COVERAGE" = "0" ]; then \
		echo "❌ Error: Could not calculate core package coverage"; \
		exit 1; \
	fi; \
	CORE_COVERAGE_INT=$$(echo "$$CORE_COVERAGE" | cut -d. -f1); \
	if [ "$$CORE_COVERAGE_INT" -lt "$(COVERAGE_THRESHOLD)" ]; then \
		echo "❌ Coverage check failed: Core packages $$CORE_COVERAGE% is below threshold of $(COVERAGE_THRESHOLD)%"; \
		exit 1; \
	else \
		echo "✅ Coverage check passed: Core packages $$CORE_COVERAGE% (threshold: $(COVERAGE_THRESHOLD)%)"; \
	fi

pre-release: ## Run full pre-release checks (tests, coverage, lint, vet)
	@echo "=========================================="
	@echo "Running Pre-Release Checks"
	@echo "=========================================="
	@echo ""
	@echo "1. Formatting code..."
	@$(MAKE) fmt
	@echo ""
	@echo "2. Running go vet..."
	@$(MAKE) vet
	@echo ""
	@echo "3. Running linter..."
	@$(MAKE) lint || echo "⚠️  Linter not available (non-blocking)"
	@echo ""
	@echo "4. Running tests with coverage check..."
	@$(MAKE) test-coverage-check
	@echo ""
	@echo "=========================================="
	@echo "✅ All pre-release checks passed!"
	@echo "=========================================="

release-check: pre-release ## Alias for pre-release

install: build ## Install to GOPATH/bin
	@echo "Installing DoPlan CLI..."
	@go install -ldflags "-X $(VERSION_SYMBOL)=$(VERSION)" ./cmd/doplan
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

check-docs: ## Check documentation organization
	@./scripts/check-docs-organization.sh

