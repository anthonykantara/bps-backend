FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o ghost-hive-c2 ./cmd/simulator/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/ghost-hive-c2 .
EXPOSE 8080 50051
CMD ["./ghost-hive-c2"]
