.PHONY: build-linter run-example clean setup

LOCAL_BIN := $(shell pwd)/bin
GOLANGCI_LINT := $(LOCAL_BIN)/golangci-lint

setup:
	@echo "Installing golangci-lint locally to ensure Go version match..."
	mkdir -p $(LOCAL_BIN)
	@if [ ! -f $(GOLANGCI_LINT) ]; then \
		GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	else \
		echo "golangci-lint already installed."; \
	fi

build-linter:
	@echo "Building linter plugin..."
	cd linter && go mod tidy
	cd linter && go build -buildmode=plugin -o ../linter.so plugin/main.go

run-example: setup build-linter
	@echo "Running example..."
	cd example && $(GOLANGCI_LINT) run || true

clean:
	rm -f linter.so
	rm -rf $(LOCAL_BIN)
