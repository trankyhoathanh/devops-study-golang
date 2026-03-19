.PHONY: help build build-linux build-darwin build-windows test generate clean dev

# Build binary name
BINARY_NAME=devops-study-golang
MAIN_PATH = cmd/main.go
VERSION?=1.0.0
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
GIT_COMMIT=$(shell git rev-parse --short HEAD)

# Build flags
LDFLAGS=-ldflags "-X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME) -X main.gitCommit=$(GIT_COMMIT)"

generate: ## Generate Entgo code
	@echo "🔧 Generating Entgo code..."
#	go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema
	@echo "✅ Code generation completed"

test: generate ## Run tests
	@echo "🧪 Running tests..."
	go test -v ./...

build: generate ## Build for current platform
	@echo "🏗️ Building $(BINARY_NAME) for $(shell uname -s)/$(shell uname -m)..."
	go build $(LDFLAGS) -o $(BINARY_NAME) .
	@echo "✅ Build completed: ./$(BINARY_NAME)"

build-linux: generate ## Build for Linux
	@echo "🐧 Building for Linux..."
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/linux/$(BINARY_NAME) $(MAIN_PATH)

build-darwin: generate ## Build for macOS
	@echo "🍎 Building for macOS..."
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o bin/darwin/$(BINARY_NAME) $(MAIN_PATH)

clean: ## Clean build artifacts
	@echo "🧹 Cleaning build artifacts..."
	rm -rf bin/ $(BINARY_NAME) $(BINARY_NAME).exe
	@echo "✅ Clean completed"

dev: ## Start development server
	@echo "🚀 Starting development server..."
	go run cmd/main.go

# Install dependencies
deps: ## Install dependencies
	@echo "📥 Installing dependencies..."
	go mod tidy
	go mod download
	go mod verify

# Code quality
lint: ## Run linter
	@echo "🔍 Linting code..."
	@which golangci-lint > /dev/null || (echo "❌ golangci-lint not installed" && exit 1)
	golangci-lint run
