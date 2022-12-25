FROM golang:1.19.1-alpine

RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD ./api /go/src/app

RUN  go mod download

RUN apk update && \
    apk upgrade && \
    apk add git

RUN go install github.com/cosmtrek/air@latest

EXPOSE 8080

CMD air -c .air.toml