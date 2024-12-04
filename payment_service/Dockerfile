FROM golang:1.22.7 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY payment_service/ ./payment_service
WORKDIR /app/payment_service

RUN go build cmd/main.go
EXPOSE 8084

CMD [ "./main" ]
