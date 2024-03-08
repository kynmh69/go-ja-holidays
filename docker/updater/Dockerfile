# syntax=docker/dockerfile:1
FROM golang:1.22 AS build-updater
COPY src/ /go/app/
WORKDIR /go/app
RUN go mod download
WORKDIR /go/app/updater
RUN CGO_ENABLED=0 go build -o bin/main

FROM alpine:latest
COPY --from=build-updater /go/app/updater/bin/main /usr/local/bin/

CMD [ "/usr/local/bin/main" ]
