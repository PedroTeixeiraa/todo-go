# Etapa 1: Build
FROM golang:1.23.5-alpine AS builder
WORKDIR /app
COPY . ./
RUN go mod download
RUN go build -o main .

# Etapa 2: Runtime
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]