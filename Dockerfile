FROM golang:1.16-alpine

WORKDIR /src

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./bin/nonacoin src/cmd/nonacoin/main.go

CMD ["./bin/nonacoin"]