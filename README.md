# API Go com Docker e PostgreSQL

Uma API RESTful desenvolvida em Go, utilizando Docker para containerização e PostgreSQL como banco de dados.

## 📁 Estrutura do Projeto

```
├── controllers/ # Controladores da aplicação
├── models/ # Modelos de dados
├── routes/ # Definição das rotas da API
├── database/ # Configurações do banco de dados
├── main.go # Arquivo principal
├── Dockerfile # Configuração do container
└── docker-compose.yml
```

## 🚀 Tecnologias Utilizadas

- [Go](https://golang.org/) (1.22)
- [Docker](https://www.docker.com/)
- [PostgreSQL](https://www.postgresql.org/) (15)
- [Docker Compose](https://docs.docker.com/compose/)

## 🔧 Requisitos

- Docker
- Docker Compose
- Go (para desenvolvimento local)

## ⚡ Início Rápido

1. **Clone o repositório**
```
git clone [URL_DO_SEU_REPOSITORIO]
cd [NOME_DO_PROJETO]
```

2. **Inicie os containers**
```
docker-compose up --build
```

3. **Acesse a API**
```
http://localhost:3000
```

## 🗄️ Configuração do Banco de Dados

### PostgreSQL
- **Host:** localhost
- **Porta:** 5432
- **Usuário:** postgres
- **Senha:** postgres
- **Database:** ecommerce

## 🛠️ Comandos Docker

```
# Iniciar os serviços
docker-compose up -d

# Parar os serviços
docker-compose down

# Visualizar logs
docker-compose logs -f

# Reiniciar um serviço específico
docker-compose restart backend

# Remover containers e volumes
docker-compose down -v
```

## 💻 Desenvolvimento

### Hot Reload
O código fonte está montado como volume no container para permitir hot reload durante o desenvolvimento:
```
volumes:
  - ./:/app
```

### Endpoints da API
- `GET /`: Hello World (teste de conexão)
- [Adicione outros endpoints conforme desenvolver]

## 🧪 Testes

[Adicione instruções para executar testes quando implementados]

## 📦 Build

Para construir a aplicação localmente:
```
go build -o main .
```

Para construir a imagem Docker:
```
docker build -t nome-da-sua-app .
```

## 🔐 Variáveis de Ambiente

[Liste as variáveis de ambiente necessárias quando implementadas]

## 📚 Documentação da API

[Adicione links ou informações sobre a documentação da API quando disponível]

## 🤝 Contribuindo

1. Faça o fork do projeto
2. Crie sua feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📝 Licença

Este projeto está sob a licença [SUA_LICENCA].

## 👨‍💻 Autor

Seu Nome - [@seu-username](https://github.com/seu-username)

## 🙏 Agradecimentos

- Mencione pessoas ou projetos que ajudaram
- Inspirações
- Referências

---
⌨️ com ❤️ por [seu-nome](https://github.com/seu-username)