FROM golang:1.15-alpine

LABEL maintainer="rytsh@devusage.com"

RUN apk add --no-cache \
    make zip git bash

RUN go get golang.org/x/tools/cmd/cover && \
    go get github.com/mattn/goveralls

WORKDIR /app
