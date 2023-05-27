# Go parameters
GO := go
GOFLAGS :=
GOTEST := $(GO) test
BINARY_NAME := MuscleApp
BUILD_DIR := ./build
SRC_FILES := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

# Default target
.PHONY: all
all: build

# Build target
.PHONY: build
build: $(BUILD_DIR)/$(BINARY_NAME)

$(BUILD_DIR)/$(BINARY_NAME): $(SRC_FILES)
	@mkdir -p $(BUILD_DIR)
	$(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)

# Run target
.PHONY: run
run:
	$(GO) run $(GOFLAGS) main.go


# Clean target
.PHONY: clean
clean:
	@rm -rf $(BUILD_DIR)

# Target to generate GoDoc documentation for all packages
.PHONY: docs
docs:
	@rm -rf docs
	@mkdir -p docs
	@go list ./... | xargs -n1 go doc -all -html > docs/index.html

# Launch Docker
.PHONY: docker
docker:
	@docker-compose pull
	@docker-compose up
	
# Help target
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build    - Build the binary"
	@echo "  run      - Run the application"
	@echo "  clean    - Clean build artifacts"
	@echo "  help     - Show this help message"
