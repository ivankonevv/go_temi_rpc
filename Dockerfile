FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./binary/server ./cmd/temi_rpc

EXPOSE 8000