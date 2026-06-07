FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o notifaction-service ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/notifaction-service .

# ✅ EXPOSE gRPC PORT
EXPOSE 50059

CMD ["./notifaction-service"]