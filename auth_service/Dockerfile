FROM golang:1.22.2
WORKDIR /
COPY . ./auth_service
RUN go mod download
RUN go mod tidy
WORKDIR /auth_service
RUN go build cmd/main.go

EXPOSE 8081

CMD [ "./main" ]
