FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN go build -o user-product-service ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/user-product-service .

EXPOSE 50053

CMD ["./user-product-service"]