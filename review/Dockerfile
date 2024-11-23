FROM golang:1.22.7 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY auth_service/ ./auth_service
WORKDIR /app/auth_service

RUN go build cmd/main.go
EXPOSE 8081

CMD [ "./main" ]
