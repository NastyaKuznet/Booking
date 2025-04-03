FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/authorizationservice ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/authorizationservice /app/
RUN chmod +x /app/authorizationservice
EXPOSE 50053
CMD ["/app/authorizationservice"]