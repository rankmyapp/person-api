FROM golang:1.21.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/bin/app ./cmd/api/main.go

FROM alpine:latest

ENV GIN_MODE=release

WORKDIR /root/

COPY --from=builder /app/bin/app .

EXPOSE 8080

CMD ["./app"]