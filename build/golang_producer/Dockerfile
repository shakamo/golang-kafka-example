FROM golang:latest

ENV GO111MODULE=on

ARG repository

WORKDIR /go/src/$repository
ADD . /go

CMD ["make", "producer"]
