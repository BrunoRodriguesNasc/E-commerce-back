# Estágio de build
FROM golang:1.23-alpine AS builder

# Adiciona git e dependências necessárias
RUN apk add --no-cache git

# Define o diretório de trabalho
WORKDIR /app

# Copia apenas o go.mod inicialmente
COPY go.mod ./

# Baixa as dependências (se houver)
RUN go mod download

# Copia o resto do código fonte
COPY . .

# Compila o aplicativo
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Estágio final
FROM alpine:latest

# Adiciona certificados SSL/TLS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copia o binário compilado do estágio anterior
COPY --from=builder /app/main .

# Expõe a porta que sua aplicação usa
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./main"]
