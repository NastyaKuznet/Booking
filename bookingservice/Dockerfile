FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bookingservice ./order

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bookingservice /app/
RUN chmod +x /app/bookingservice
EXPOSE 50052
CMD ["/app/bookingservice"]