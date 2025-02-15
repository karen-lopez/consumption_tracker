FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY .env .env

RUN go build -o consumption-tracker ./cmd/app/main.go

RUN apk add --no-cache ca-certificates
EXPOSE 8080

CMD ["./consumption-tracker"]