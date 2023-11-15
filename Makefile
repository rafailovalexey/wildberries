# Variables

PROTO_SOURCE_DIRECTORY = api
PROTO_OUTPUT_DIRECTORY = pkg

PROTO_FILES = \
	employees_v1/employees-grpc.proto \
	orders_v1/orders1-grpc.proto

MOCKS_OUTPUT_DIRECTORY = mocks

MOCKS_FILES = \
	internal/repository/repository.go

# GRPC

grpc-generate:
	bin/grpc-generate.sh $(PROTO_SOURCE_DIRECTORY) $(PROTO_OUTPUT_DIRECTORY) $(PROTO_FILES)

# Mocks

mocks-generate:
	bin/mocks-generate.sh $(MOCKS_OUTPUT_DIRECTORY) $(MOCKS_FILES)

# Generate

generate:
	make grpc-generate
	make mocks-generate

# Download

download:
	go mod download

# Build

build:
	go build -o build/main main.go

# Tests

tests:
	go test -v ./...

.PHONY: grpc-generate, mocks-generate, generate, download, build, tests
