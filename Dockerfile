#pull base image
FROM golang:alpine

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o projectx ./cmd/main.go
