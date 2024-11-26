FROM golang:1.22.7 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY movie_service/ ./movie_service
WORKDIR /app/movie_service

RUN go build cmd/main.go
EXPOSE 8083

CMD [ "./main" ]
