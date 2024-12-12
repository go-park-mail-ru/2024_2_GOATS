############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY payment_service/ ./payment_service

WORKDIR /app/payment_service

ENV CGO_ENABLED=0
RUN go build -o /app/payment_service/cmd/main cmd/main.go

############################
# STEP 2 build a small image
############################
FROM scratch

COPY --from=builder /app/payment_service/cmd/main /bin/main
COPY --from=builder /app/payment_service/.env /app/payment_service/.env
COPY --from=builder /app/payment_service/internal/config/config.yml /app/payment_service/internal/config/config.yml
COPY --from=builder /app/payment_service/internal/db /app/payment_service/internal/db

WORKDIR /app/payment_service
HEALTHCHECK --interval=10s --timeout=5s --start-period=10s --retries=5 CMD wget --spider http://localhost:8084/health || exit 1

ENTRYPOINT ["/bin/main"]
