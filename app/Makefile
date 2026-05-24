.PHONY: build install clean test run

APP_NAME=toris

build:
	@echo "Building $(APP_NAME)..."
	go build -o $(APP_NAME) .

install:
	@echo "Installing $(APP_NAME) globally..."
	go build -o $$(go env GOPATH)/bin/$(APP_NAME) .

clean:
	@echo "Cleaning up..."
	go clean
	rm -f $(APP_NAME)

test:
	@echo "Running tests..."
	go test -v ./...
