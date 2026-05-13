FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN go build -o food-service ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/u .

EXPOSE 50053

CMD ["./user-product-service"]