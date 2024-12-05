############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . ./facade_app

WORKDIR /app/facade_app

ENV CGO_ENABLED=0
RUN go build -o /app/facade_app/cmd/main cmd/main.go

############################
# STEP 2 build a small image
############################
FROM scratch

COPY --from=builder /app/facade_app/cmd/main /bin/main
COPY --from=builder /app/facade_app/.env /app/facade_app/.env
COPY --from=builder /app/facade_app/internal/config/config.yml /app/facade_app/internal/config/config.yml
COPY --from=builder /app/facade_app/internal/db /app/facade_app/internal/db

WORKDIR /app/facade_app

ENTRYPOINT ["/bin/main"]
