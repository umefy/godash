.PHONY: default check fmt test lint tidy help

TEST_EXCLUDE_PATHS="protogen/pb"
TEST_PATHS=$(shell go list ./... | grep -v -E "$(TEST_EXCLUDE_PATHS)")

default: check

check: tidy fmt lint test

fmt:
	@echo "⏱️ formatting code now..."
	go fmt ./...
	@echo "✅ formatting finish"

test:
	@echo "⏱️ running tests now... "
	go test -race --parallel=4 -timeout 30s -cover $(ARGS) $(TEST_PATHS) 
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