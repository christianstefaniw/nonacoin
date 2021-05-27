NONACOIN = cmd/nonacoin/main.go
SRC_DIR = src
BIN_DIR = bin
BIN_NAME = nonacoin
IMAGE_NAME = 12152004/nonacoin:latest
PROTO_PATH = /home/christianstefaniw/Desktop/code/src/github.com/christianstefaniw/nonacoin/nonacoin-protobufs
COVERAGE_DIR = coverage

.PHONY: blockchain test test-blockchain test-cov

test-cov:
	go test ./... -coverprofile $(COVERAGE_DIR)/coverage.txt
	go tool cover -func $(COVERAGE_DIR)/coverage.txt

test:
	go test ./...

test-blockchain:
	go test ./$(SRC_DIR)/apps/blockchain

docker-build:
	sudo docker build . -t $(IMAGE_NAME)

docker-push:
	sudo docker push $(IMAGE_NAME)

build:
	go build -o $(BIN_DIR)/$(BIN_NAME) $(SRC_DIR)/$(NONACOIN)

run:
	go run $(SRC_DIR)/$(NONACOIN)

blockchain-pb:
	protoc --go_out=plugins=grpc:src/apps/blockchain --proto_path=$(PROTO_PATH) blockchain.proto