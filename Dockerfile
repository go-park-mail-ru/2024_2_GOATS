FROM golang:1.22.7 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . ./fascade
WORKDIR /app/fascade

RUN go build cmd/main.go
EXPOSE 8080

CMD ["./main"]
