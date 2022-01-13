#pull golang base  image
FROM golang:latest
ENV GO111MODULE=on
LABEL name devstack

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
EXPOSE 6969

RUN go build -o main .

CMD [ "./main" ]
# UUID=56af2403-3f74-4d52-8bdd-2962d00a395e  /boot ext2 auto,noatime    1 2
# UUID=d493144f-ee35-4a3c-8291-8468a3bb9f99 /home g2fs rw,noatime,discard 0 2
