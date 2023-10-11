FROM golang:1.21.1-alpine

WORKDIR /friend-management

COPY go.mod go.mod

COPY go.sum go.sum

RUN apk update

RUN go mod download

CMD [ "go", "version" ]
