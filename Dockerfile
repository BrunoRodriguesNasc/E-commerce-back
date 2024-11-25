# Estágio de desenvolvimento
FROM golang:1.21-alpine

WORKDIR /app

# Instalar dependências necessárias
RUN apk add --no-cache git

# Instalar Air
RUN go install github.com/cosmtrek/air@v1.49.0

# Criar diretório tmp
RUN mkdir -p tmp

# Copiar os arquivos do projeto
COPY . .

# Usar Air para hot-reload
CMD ["air"]
