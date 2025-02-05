.PHONY: all test build clean tag-version release

# Version information
CURRENT_VERSION := $(shell git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
MAJOR := $(word 1,$(subst ., ,$(patsubst v%,%,$(CURRENT_VERSION))))
MINOR := $(word 2,$(subst ., ,$(patsubst v%,%,$(CURRENT_VERSION))))
PATCH := $(word 3,$(subst ., ,$(patsubst v%,%,$(CURRENT_VERSION))))

# Default target
all: test

# Test the package
test:
	@echo "Running tests..."
	@go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf dist/
	@go clean

# Build the package
build: clean test
	@echo "Building..."
	@mkdir -p dist
	@go build -o dist/

# Tag a new version
tag-version:
	@echo "Current version: $(CURRENT_VERSION)"
	@echo "Select version type to bump:"
	@echo "1) Major ($(MAJOR) -> $$(($(MAJOR)+1)).0.0)"
	@echo "2) Minor ($(MINOR) -> $(MAJOR).$$(($(MINOR)+1)).0)"
	@echo "3) Patch ($(PATCH) -> $(MAJOR).$(MINOR).$$(($(PATCH)+1)))"
	@read -p "Enter choice (1-3): " choice; \
	case $$choice in \
		1) new_version="v$$(($(MAJOR)+1)).0.0" ;; \
		2) new_version="v$(MAJOR).$$(($(MINOR)+1)).0" ;; \
		3) new_version="v$(MAJOR).$(MINOR).$$(($(PATCH)+1))" ;; \
		*) echo "Invalid choice" && exit 1 ;; \
	esac; \
	echo "Tagging new version: $$new_version"; \
	git tag -a $$new_version -m "Release $$new_version"

# Release the package
release: test
	@if [ -z "$(VERSION)" ]; then \
		echo "Error: VERSION is required. Use 'make release VERSION=vX.Y.Z'"; \
		exit 1; \
	fi
	@echo "Preparing release $(VERSION)..."
	@if ! git rev-parse $(VERSION) >/dev/null 2>&1; then \
		echo "Creating tag $(VERSION)..."; \
		git tag -a $(VERSION) -m "Release $(VERSION)"; \
	else \
		echo "Tag $(VERSION) already exists."; \
		exit 1; \
	fi
	@echo "Pushing tag to remote..."
	@git push origin $(VERSION)
	@echo "Release $(VERSION) completed!"

# Show help
help:
	@echo "Available targets:"
	@echo "  all          : Run tests (default)"
	@echo "  test         : Run all tests"
	@echo "  clean        : Clean build artifacts"
	@echo "  build        : Build the package"
	@echo "  tag-version  : Interactive version bumping"
	@echo "  release      : Create and push a new release (use VERSION=vX.Y.Z)"
	@echo "  help         : Show this help message"
	@echo ""
	@echo "Current version: $(CURRENT_VERSION)"

# Version information display
version:
	@echo $(CURRENT_VERSION)
