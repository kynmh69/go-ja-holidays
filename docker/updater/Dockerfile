FROM golang:1.22 AS build-stage

COPY ./src /go/app
WORKDIR /go/app

RUN go mod download

WORKDIR /go/app/updater
RUN go build -o ./bin/main

FROM alpine:latest

COPY --from=build-stage /go/app/updater/bin/main /usr/local/bin/main

CMD [ "main" ]
