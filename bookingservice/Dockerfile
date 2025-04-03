FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bookingservice ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bookingservice /app/
RUN chmod +x /app/bookingservice
EXPOSE 50052
CMD ["/app/bookingservice"]