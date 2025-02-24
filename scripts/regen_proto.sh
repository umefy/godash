#!/bin/bash

BASE_PATH="$(cd "$(dirname "$0")/.." && pwd)"

# Define output directories for different languages
GO_OUT_DIR=$BASE_PATH/internal/testdata

rm -rf $GO_OUT_DIR

# Create output directories if they don't exist
mkdir -p $GO_OUT_DIR

# Find all .proto files
PROTO_FILES=$(find $BASE_PATH/internal/proto -name "*.proto")

# Generate Go files
protoc -I $BASE_PATH/internal/proto \
--go_out=$GO_OUT_DIR --go_opt=paths=source_relative \
--go-grpc_out=$GO_OUT_DIR --go-grpc_opt=paths=source_relative \
$PROTO_FILES

