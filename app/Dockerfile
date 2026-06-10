# Build Stage
FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o observability-app .

# Runtime Stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/observability-app .

EXPOSE 8080

CMD ["./observability-app"]
