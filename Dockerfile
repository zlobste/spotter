FROM golang:1.15

WORKDIR /go/src/github.com/zlobste/spotter
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/spotter github.com/zlobste/spotter

ENTRYPOINT ["spotter"]
