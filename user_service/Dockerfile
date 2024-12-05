############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY user_service/ ./user_service

WORKDIR /app/user_service

ENV CGO_ENABLED=0
RUN go build -o /app/user_service/cmd/main cmd/main.go

############################
# STEP 2 build a small image
############################
FROM scratch

COPY --from=builder /app/user_service/cmd/main /bin/main
COPY --from=builder /app/user_service/.env /app/user_service/.env
COPY --from=builder /app/user_service/internal/config/config.yml /app/user_service/internal/config/config.yml
COPY --from=builder /app/user_service/internal/db /app/user_service/internal/db

WORKDIR /app/user_service

ENTRYPOINT ["/bin/main"]
