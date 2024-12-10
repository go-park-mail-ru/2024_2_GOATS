############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY user_service/ ./user_service
COPY .env ./user_service

WORKDIR /app/user_service

ENV CGO_ENABLED=0
RUN go build -o /app/user_service/cmd/main cmd/main.go

############################
# STEP 2 build a small image
############################
FROM scratch

COPY --from=builder /app/user_service/cmd/main /bin/main
COPY --from=builder /app/user_service/.env /app/.env
COPY --from=builder /app/user_service/internal/config/config.yml /app/user_service/internal/config/config.yml
COPY --from=builder /app/user_service/internal/db /app/user_service/internal/db

WORKDIR /app/user_service
HEALTHCHECK --interval=10s --timeout=5s --start-period=10s --retries=5 CMD wget --spider http://localhost:8082/health || exit 1

ENTRYPOINT ["/bin/main"]
