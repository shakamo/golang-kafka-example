FROM golang:latest

ENV GO111MODULE=on
ENV BROKER_List=kafka:9092

ARG repository

WORKDIR /go/src/$repository
ADD . /go

CMD ["make", "producer"]
