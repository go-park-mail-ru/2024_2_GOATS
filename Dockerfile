FROM golang:1.22.2
WORKDIR /
COPY . ./app
WORKDIR /app
RUN go mod download
RUN go mod tidy
RUN go build cmd/main.go

EXPOSE 8080

CMD [ "./main" ]
