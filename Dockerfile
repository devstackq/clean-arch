#pull base image
FROM golang:alpine

#install git
RUN apk update && apk add --no-cache git

#Where our file will be in the docker container
