#!/bin/bash

BASE_PATH="$(cd "$(dirname "$0")/.." && pwd)"
GO_PROJECT="." # replace with go project

SEARCH_PROTO_FOLDERS=("jsonkit")

for SEARCH_PROTO_FOLDER in "${SEARCH_PROTO_FOLDERS[@]}"; do
    # Define output directories for different languages
    GO_OUT_DIR=$BASE_PATH/$GO_PROJECT/$SEARCH_PROTO_FOLDER/testdata

    rm -rf $GO_OUT_DIR/*.pb.go

    # Create output directories if they don't exist
    mkdir -p $GO_OUT_DIR

    # Find all .proto files
    PROTO_FILES=$(find $BASE_PATH/$SEARCH_PROTO_FOLDER/testdata/proto -name "*.proto")

    # Generate Go files
    protoc -I $BASE_PATH/$SEARCH_PROTO_FOLDER/testdata/proto \
    --go_out=$GO_OUT_DIR --go_opt=paths=source_relative \
    --go-grpc_out=$GO_OUT_DIR --go-grpc_opt=paths=source_relative \
    $PROTO_FILES
done
