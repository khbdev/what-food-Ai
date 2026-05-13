FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN go build -o mail-service ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/mail-service .

EXPOSE 50057

CMD ["./mail-service"]