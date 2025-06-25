# Variables
CLI_BINARY_NAME=glint
BUILD_DIR=bin
EXT=$(if $(filter windows,$(GOOS)),.exe,)
VERSION=$(shell git describe --tags --abbrev=0)

# Targets
.PHONY: all build/* test release install clean lint docs

all: build/$(CLI_BINARY_NAME)

build/%:
	@echo "Building $*..."
	@mkdir -p $(BUILD_DIR)
	@go build -ldflags "-X github.com/Phillezi/common/scripts/glint/cmd/$*/cli.version=$(VERSION)" -o $(BUILD_DIR)/$*$(EXT) .
	@echo "Build complete: $(BUILD_DIR)/$*$(EXT)"

test:
	@go test ./...

release/%:
	@echo "Building the application..."
	@mkdir -p $(BUILD_DIR)
	@go build -mod=readonly -ldflags "-w -s -X github.com/Phillezi/scripts/glint/cmd/$*/cli.version=$(VERSION)" -o $(BUILD_DIR)/$*$(EXT) .
	@echo "Build complete."

install: release
	@echo "installing"
	@./scripts/escalate.sh cp ./$(BUILD_DIR)/$(CLI_BINARY_NAME)$(EXT) /usr/local/bin/$(CLI_BINARY_NAME)$(EXT)
	@./scripts/escalate.sh cp ./$(BUILD_DIR)/$(DAEMON_BINARY_NAME)$(EXT) /usr/local/bin/$(DAEMON_BINARY_NAME)$(EXT)

clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete."

lint: run
	@true

docs:
	@go run ./cmd/docs
