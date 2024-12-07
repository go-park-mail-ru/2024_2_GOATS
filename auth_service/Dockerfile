############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY auth_service/ ./auth_service

WORKDIR /app/auth_service

ENV CGO_ENABLED=0
RUN go build -o /app/auth_service/cmd/main cmd/main.go

############################
# STEP 2 build a small image
############################
FROM scratch

COPY --from=builder /app/auth_service/cmd/main /bin/main
COPY --from=builder /app/auth_service/.env /app/auth_service/.env
COPY --from=builder /app/auth_service/internal/config/config.yml /app/auth_service/internal/config/config.yml

WORKDIR /app/auth_service

ENTRYPOINT ["/bin/main"]
