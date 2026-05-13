FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN go build -o api-geteway ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/a .

EXPOSE 50055

CMD ["./ai-service"]