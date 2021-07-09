FROM golang:1.15-alpine3.13

ENV CGO_ENABLED=0

RUN mkdir /app

ADD . /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go build -o main

EXPOSE 9080 

CMD ["/app/main"]
