.PHONY: default check fmt test lint tidy help

default: check

check: tidy fmt lint test

fmt:
	@echo "⏱️ formatting code now..."
	go fmt ./...
	@echo "✅ formatting finish"

test:
	@echo "⏱️ running tests now... "
	go test -race --parallel=4 -timeout 30s -cover $(ARGS) ./...
	@echo "✅ passing all tests."

lint:
	@echo "⏱️ running linting now..."
	golangci-lint run $(ARGS)
	@echo "✅ passing linting..."

tidy:
	@echo "⏱️ go mod tidy now..."
	go mod tidy
	@echo "✅ finishing tidy..."

help:
	@echo "make - running make check to verify the code quality"
	@echo "make check - formatting, testing and running lint"
	@echo "make test - running go test"
	@echo "make fmt - formatting go code"
	@echo "make lint - running golangci lint"
	@echo "make tidy - install all dependencies"