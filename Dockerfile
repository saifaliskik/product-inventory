FROM ubuntu:latest AS builder

RUN apt-get update && apt-get install -y golang

WORKDIR /app

COPY . /app

RUN go build -o myapp cmd/main.go

EXPOSE 8080

CMD ["./myapp"]