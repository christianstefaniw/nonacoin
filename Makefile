NONACOIN = cmd/nonacoin/main.go
SRC_DIR = src
BIN_DIR = bin
NONA_BIN_NAME = nonacoin
CLIENT_BIN_NAME = client
IMAGE_NAME = 12152004/nonacoin:latest
PROTO_PATH = /home/christianstefaniw/Desktop/code/src/github.com/christianstefaniw/nonacoin/nonacoin-protobufs
COVERAGE_DIR = coverage
TEST_CLIENT = cmd/temp_client/main.go

.PHONY: blockchainpb test test-blockchain cov

cov:
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

docker-run:
	sudo docker run -p 8080:8080 $(IMAGE_NAME)

nonacoin-build:
	go build -o $(BIN_DIR)/$(NONA_BIN_NAME) $(SRC_DIR)/$(NONACOIN)

nonacoin-run:
	go run $(SRC_DIR)/$(NONACOIN)

client-run:
	go run $(SRC_DIR)/$(TEST_CLIENT)

client-build:
	go build -o $(BIN_DIR)/$(NONA_BIN_NAME) $(SRC_DIR)/$(TEST_CLIENT)

peer2peerpb:
	protoc --go_out=plugins=grpc:src/peer2peer --proto_path=$(PROTO_PATH) peer2peer.proto

bootstrappb:
	protoc --go_out=plugins=grpc:src/peer2peer --proto_path=$(PROTO_PATH) boot_node.proto