FROM golang:1.22.7 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY review/ ./review
WORKDIR /app/review

RUN go build cmd/main.go
EXPOSE 8081

CMD [ "./main" ]
