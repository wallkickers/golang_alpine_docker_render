FROM golang:1.12.5-alpine3.9 AS builder

WORKDIR /app

COPY . .

RUN go build -o main app/main.go

CMD ["./main"]
