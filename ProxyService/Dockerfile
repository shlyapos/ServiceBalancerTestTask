FROM golang:latest AS builder
WORKDIR /build
COPY ./ ./
# COPY ./config/config.yml ./config/config.yml
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/main.go

FROM alpine AS server
WORKDIR /Balancer
COPY --from=builder /build/app .
COPY --from=builder /build/config/config.yml ./config/config.yml
CMD ["./app"]