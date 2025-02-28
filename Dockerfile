# Dockerfile
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod  ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main ./main

EXPOSE 8080

CMD ["./main"]