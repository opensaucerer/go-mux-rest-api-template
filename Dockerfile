FROM golang:1.18-alpine

RUN mkdir -p /go/src/github.com/opensaucerer/gotemplate

WORKDIR /go/src/github.com/opensaucerer/gotemplate

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN make build

RUN chmod +x /go/src/github.com/opensaucerer/gotemplate/bin/gotemplate

ENTRYPOINT make start