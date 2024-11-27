FROM golang:1.22.7 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY user_service/ ./user_service

WORKDIR /app/user_service

RUN go build cmd/main.go
EXPOSE 8082

CMD ["./main"]
