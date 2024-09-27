# syntax=docker/dockerfile:1
FROM golang:1.23 AS build-updater
COPY ./src/ /go/app/
WORKDIR /go/app
RUN go install github.com/cosmtrek/air@latest

RUN go mod download
WORKDIR /go/app/key_management

CMD [ "air", "-c", ".air.toml"]
