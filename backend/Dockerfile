# Build stage
FROM golang:1.24 as builder

WORKDIR /app
COPY . .
COPY ./migrations ./migrations

RUN go mod tidy
RUN CGO_ENABLED=0 go build -o app .

# Final stage
FROM debian:bullseye-slim

WORKDIR /root/
COPY --from=builder /app/app .
COPY --from=builder /app/migrations ./migrations

EXPOSE 8080

CMD ["./app"]
