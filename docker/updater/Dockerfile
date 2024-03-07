FROM golang:1.22 as build-stage

COPY ./src/updater /go/app/updater
WORKDIR /go/app

RUN go mod download

RUN go build updater/main.go -o updater/bin/main

FROM alpine

COPY --from=build-stage /go/app/updater/bin/main /usr/local/bin/main

CMD [ "main" ]
