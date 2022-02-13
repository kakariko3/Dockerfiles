FROM golang:1.17.7-alpine

RUN apk update && \
    apk upgrade && \
    apk add --no-cache git

WORKDIR /go/src/app
COPY ./app /go/src/app
