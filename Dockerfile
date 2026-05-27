FROM golang:1.25-alpine AS builder


WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o vibecheck ./cmd/api/

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/vibecheck .
# COPY --from=builder /app/db/migrations ./db/migrations


EXPOSE 8080
CMD ["./vibecheck"]
