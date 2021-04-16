FROM golang:1.16 AS build

WORKDIR /app

COPY cmd/ cmd/
COPY internal/ internal/
COPY vendor/ vendor/

COPY go.mod go.sum ./

RUN CGO_ENABLED=0 \
    GOOS=linux \
    go build -o /go/bin/spotter \
       /app/cmd/spotter

FROM alpine:3.9

COPY --from=build /go/bin/spotter /usr/local/bin/spotter
