FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./balance/internal/app

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 5000

CMD ["./main"]
