FROM golang:1.17.8-alpine

ENV ROOT=/go/src/app
ENV CGO_ENABLED=0

RUN apk update && \
    apk add --no-cache git build-base

WORKDIR ${ROOT}
COPY . ${ROOT}

EXPOSE 8080
