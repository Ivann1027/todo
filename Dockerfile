FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN CGO_ENABLED=0 GOOS=linux go build -o todo-api ./cmd/api/main.go

FROM alpine:latest

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

RUN echo '#!/bin/sh\n\
set -e\n\
echo "Waiting for database..."\n\
until pg_isready -h db -p 5432 -U ivan -d todo-db; do\n\
  sleep 2\n\
done\n\
echo "Running migrations..."\n\
export PGPASSWORD=1234\n\
goose -dir ./migrations postgres "host=db user=ivan dbname=todo-db sslmode=disable" up\n\
echo "Starting app..."\n\
exec ./todo-api' > /app/entrypoint.sh && \
chmod +x /app/entrypoint.sh && chown appuser:appgroup /app/entrypoint.sh

USER appuser
EXPOSE 8080
ENTRYPOINT ["/app/entrypoint.sh"]