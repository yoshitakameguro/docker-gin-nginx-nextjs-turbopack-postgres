FROM golang:1.20.0-alpine3.17
WORKDIR /app
COPY ./app/. .
ENV GO111MODULE=on
RUN apk update && apk add git
RUN go install github.com/cosmtrek/air@v1.41.0
CMD air -c .air.toml
