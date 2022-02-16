FROM golang:latest as server-building-stage

WORKDIR /app

COPY ../server/go.mod .
COPY ../server/go.sum .

RUN go mod download

COPY ../server .

RUN go build -o ./binary/server ./cmd/temi_rpc

EXPOSE 8000