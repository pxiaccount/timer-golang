FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o app

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .

ENTRYPOINT ["./app"]