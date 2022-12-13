FROM golang:1.19.1-alpine
 
RUN apk update && \
    apk upgrade && \
    apk add git
 
RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app