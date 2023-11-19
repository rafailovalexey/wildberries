#!/bin/bash

if [ "$#" -lt 2 ]; then
  echo "Usage: $0 <PROTO_SOURCE_DIRECTORY> <PROTO_OUTPUT_DIRECTORY> <PROTO_FILES>"
  exit 1
fi

PROTO_SOURCE_DIRECTORY="$1"
PROTO_OUTPUT_DIRECTORY="$2"
PROTO_FILES="${*:3}"

for proto_file in $PROTO_FILES; do

  PROTO_FILE_DIRECTORY=$(dirname "$proto_file")

  mkdir -p "$PROTO_OUTPUT_DIRECTORY/$PROTO_FILE_DIRECTORY"

  echo "Generating proto file for $proto_file..."

  protoc \
    --proto_path="$PROTO_SOURCE_DIRECTORY" \
    --go_out="$PROTO_OUTPUT_DIRECTORY" \
    --go_opt=paths=source_relative "$proto_file" \
    --go-grpc_out="$PROTO_OUTPUT_DIRECTORY" \
    --go-grpc_opt=paths=source_relative "$proto_file"

done
