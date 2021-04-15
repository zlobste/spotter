FROM golang:1.16 AS build

RUN mkdir ./app
COPY . ./app
WORKDIR ./app

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/spotter ./

FROM alpine:3.9

COPY --from=build /go/bin/spotter /usr/local/bin/spotter
