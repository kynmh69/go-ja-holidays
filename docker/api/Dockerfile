# syntax=docker/dockerfile:1
FROM golang:1.22 AS build-updater
COPY ./src/ /go/app/
WORKDIR /go/app
RUN go install github.com/cosmtrek/air@latest

RUN go mod download
WORKDIR /go/app/api

CMD [ "air", "-c", "air.toml"]
