#!/bin/bash

set -euo pipefail

echo "⏱️ running linting now..."
golangci-lint run
echo "✅ passing linting..."
echo "⏱️ running tests now..."
go test -timeout 30s -cover ./...
echo "✅ passing testing..."