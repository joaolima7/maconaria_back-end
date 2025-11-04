# Stage 1: Build
FROM golang:1.21-alpine AS builder

# Instala dependências necessárias
RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

# Copia go mod e sum primeiro (cache de dependências)
COPY go.mod go.sum ./
RUN go mod download

# Copia todo o código
COPY . .

# Build da aplicação
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o /app/server \
    cmd/main.go

# Stage 2: Runtime
FROM alpine:latest

# Instala ca-certificates para HTTPS e tzdata para timezone
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copia binário do stage de build
COPY --from=builder /app/server .

# Copia migrations (necessário para auto-migrate)
COPY --from=builder /app/sql/migrations ./sql/migrations

# Cria usuário não-root
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser && \
    chown -R appuser:appuser /app

USER appuser

# Expõe porta
EXPOSE 8080

# Comando de start
CMD ["./server"]