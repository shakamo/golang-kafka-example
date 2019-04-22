FROM golang:latest

ENV GO111MODULE=on
ENV brokerList=kafka:9092

ARG repository

WORKDIR /go/src/$repository
ADD . /go

CMD ["make", "consumer"]
