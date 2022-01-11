#pull base image
FROM golang:latest

RUN mdkir app

WORKDIR /app

COPY go.mod .
COPY go.sum .
# ENV GOPATH=/

RUN go mod download
# copy project to current app dir
COPY . .

RUN CGO_ENABLED=0 GOOS = linux  go build -o /app/cmd/main main.go

# CMD ["/app/cmd/main"]
ENTRYPOINT ["/app/cmd/main"]
