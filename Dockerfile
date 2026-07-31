FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY .env .env

RUN go install github.com/pressly/goose/v3/cmd/goose@v3.21.1

RUN CGO_ENABLED=0 GOOS=linux go build -o todo-api ./cmd/api/main.go

FROM alpine:3.20

RUN apk --no-cache add \
    tzdata \
    ca-certificates \
    postgresql-client \
    && update-ca-certificates

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder --chown=appuser:appgroup /app/todo-api .
COPY --from=builder --chown=appuser:appgroup /go/bin/goose /usr/local/bin/goose
COPY --from=builder --chown=appuser:appgroup /app/configs ./configs
COPY --from=builder --chown=appuser:appgroup /app/migrations ./migrations

COPY entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

ENTRYPOINT ["/app/entrypoint.sh"]