# Project variables
MODULE=github.com/CuriousHet/D-CAS
BINARY_NAME=D-CAS
BUILD_DIR=bin

# Go commands
GO=go
GO_BUILD=$(GO) build
GO_TEST=$(GO) test
GO_CLEAN=$(GO) clean
GO_FMT=$(GO) fmt
GO_VET=$(GO) vet
GO_MOD=$(GO) mod tidy

# Default target
.PHONY: all
all: fmt vet build

# Build the project
.PHONY: build
build:
	@echo "Building the project..."
	@mkdir -p $(BUILD_DIR)
	$(GO_BUILD) -o $(BUILD_DIR)/$(BINARY_NAME) main.go
	@echo "Build complete! Executable is in $(BUILD_DIR)/$(BINARY_NAME)"

# Run a node with a custom port
.PHONY: node
node: build
	@if [ -z "$(port)" ]; then \
		echo "Usage: make node port=<port-number>"; \
	else \
		echo "Starting node on port $(port)..."; \
		$(BUILD_DIR)/$(BINARY_NAME) -port=$(port); \
	fi

# Run CLI mode (Store & Get operations)
.PHONY: cli
cli:
	@if [ -z "$(cmd)" ] || [ -z "$(arg)" ]; then \
		echo "Usage: make cli cmd=<store|get> arg=<file-path|hash> port=<port>"; \
	else \
		echo "Executing CLI command: $(cmd) $(arg) on port $(port)"; \
		$(BUILD_DIR)/$(BINARY_NAME) $(cmd) $(arg) -port=$(port); \
	fi

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	$(GO_TEST) ./... -v
	@echo "Tests completed."

# Format the code
.PHONY: fmt
fmt:
	@echo "Formatting the code..."
	$(GO_FMT) ./...

# Lint and vet the code
.PHONY: vet
vet:
	@echo "Running go vet..."
	$(GO_VET) ./...

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning up..."
	$(GO_CLEAN)
	rm -rf $(BUILD_DIR)
	@echo "Cleanup complete."

# Update dependencies
.PHONY: tidy
tidy:
	@echo "Tidying up Go modules..."
	$(GO_MOD)

# Show available commands
.PHONY: help
help:
	@echo "Makefile Commands:"
	@echo "  make build       - Compile the project"
	@echo "  make run         - Run the project (default port 3000)"
	@echo "  make node port=3001 - Start a node on a custom port"
	@echo "  make cli cmd=store arg=example.txt port=3000 - Store a file"
	@echo "  make cli cmd=get arg=<hash> port=3001 - Retrieve a file"
	@echo "  make test        - Run tests"
	@echo "  make fmt         - Format the code"
	@echo "  make vet         - Run Go vet"
	@echo "  make clean       - Clean build artifacts"
	@echo "  make tidy        - Tidy Go modules"
