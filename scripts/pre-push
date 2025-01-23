#!/bin/sh

echo "⏱️ running linting now..."
golangci-lint run
echo "✅ passing linting..."
echo "⏱️ running tests now..."
go test -timeout 30s -cover ./...
echo "✅ passing testing..."