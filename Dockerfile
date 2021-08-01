# Build stage
FROM golang:1.16-alpine3.13 AS builder
WORKDIR /build
COPY . .
RUN go mod tidy
RUN go build -o main main.go

# Run stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /build/main .
COPY .env .

EXPOSE 8081
CMD [ "/app/main" ]