FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o todo-api ./cmd/api

FROM alpine:latest

RUN apk --no-cache add \
    tzdata \
    ca-certificates \
    && update-ca-certificates

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder --chown=appuser:appgroup /app/todo-api .
COPY --from=builder --chown=appuser:appgroup /app/configs ./configs
COPY --from=builder --chown=appuser:appgroup /app/migrations ./migrations

USER appuser
EXPOSE 8080
CMD ["./todo-api"]