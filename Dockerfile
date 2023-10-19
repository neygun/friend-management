FROM golang:1.21.1-alpine

WORKDIR /friend-management

COPY . .

RUN apk update

RUN go mod download

CMD [ "go", "version" ]
