NONACOIN = cmd/nonacoin/main.go
SRC_DIR = src
BIN_DIR = bin
BIN_NAME = nonacoin
IMAGE_NAME = 12152004/nonacoin:latest

docker-build:
	sudo docker build . -t $(IMAGE_NAME)

docker-push:
	sudo docker push $(IMAGE_NAME)

build:
	go build -o $(BIN_DIR)/$(BIN_NAME) $(SRC_DIR)/$(NONACOIN)

run:
	go run $(SRC_DIR)/$(NONACOIN)