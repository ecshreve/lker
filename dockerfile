# syntax=docker/dockerfile:1

FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o bin/lker github.com/ecshreve/lker/cmd/lker

EXPOSE 8880

CMD [ "./bin/lker" ]
