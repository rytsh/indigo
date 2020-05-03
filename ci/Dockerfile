FROM golang:1.14-alpine

LABEL maintainer="rytsh@devusage.com"

RUN apk add --no-cache \
    make zip git

RUN go get golang.org/x/tools/cmd/cover && \
    go get github.com/mattn/goveralls

WORKDIR /app
