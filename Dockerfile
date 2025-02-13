FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o consumption-tracker ./cmd/app/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/consumption-tracker .

RUN apk add --no-cache ca-certificates
EXPOSE 8080

CMD ["./consumption-tracker"]