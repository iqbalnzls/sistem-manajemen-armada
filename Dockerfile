# Builder
FROM golang:1.24-alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o sistem-manajemen-armada ./cmd/main.go

CMD ["/app/sistem-manajemen-armada"]
