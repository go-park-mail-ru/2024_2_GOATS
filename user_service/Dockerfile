FROM golang:1.22.2
WORKDIR /
COPY . ./user_service
RUN go mod download
RUN go mod tidy
WORKDIR /user_service
RUN go build cmd/main.go

EXPOSE 8082

CMD [ "./main" ]
