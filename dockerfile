# Etapa de build com Go 1.23 (base Debian)
FROM golang:1.23 AS builder

WORKDIR /app

# Copia os arquivos de dependência
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código
COPY . .

# Compila o binário
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Etapa final com imagem leve
FROM alpine:latest

WORKDIR /app

# Instala certificados SSL
RUN apk --no-cache add ca-certificates

# Copia o binário da etapa anterior
COPY --from=builder /app/app .

# Cria usuário não-root por segurança
RUN adduser -D -g '' appuser
USER appuser

# Expõe a porta da aplicação
EXPOSE 9091

# Comando para iniciar o app
CMD ["./app"]
