# Makefile

PROTOC_GEN_GO := $(shell go env GOPATH)/bin/protoc-gen-go
PROTOC_GEN_GO_GRPC := $(shell go env GOPATH)/bin/protoc-gen-go-grpc
PROTO_DIR := internal/proto
OUT_DIR := internal/proto

.PHONY: all proto

all: proto

# Install the protobuf compiler and Go plugins
install:
	@echo "Installing protobuf compiler and Go plugins..."
	GO111MODULE=on go get google.golang.org/protobuf/cmd/protoc-gen-go
	GO111MODULE=on go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

# Compile the protobuf files
proto:
	@echo "Compiling protobuf files..."
	@mkdir -p $(OUT_DIR)
	protoc -I=$(PROTO_DIR) \
		--go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/*.proto

# Clean generated files
clean:
	@echo "Cleaning generated files..."
	rm -f $(OUT_DIR)/*.pb.go
