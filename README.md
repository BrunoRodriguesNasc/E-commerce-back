# E-commerce API em Go

API RESTful de e-commerce desenvolvida em Go, utilizando PostgreSQL como banco de dados.

## 🚀 Tecnologias

- Go 1.21
- PostgreSQL 15
- Docker & Docker Compose
- Air (Hot Reload)

## 📁 Estrutura do Projeto

```
├──main.go                      # Ponto de entrada da aplicação
├── internal/
│   ├── controllers/            # Lógica de negócios
│   ├── models/                 # Estruturas de dados
│   ├── repositories/           # Camada de acesso ao banco
│   └── routes/                 # Configuração de rotas
├── pkg/
│   └── database/              # Configuração do banco
├── migrations/                # Scripts SQL
├── Dockerfile                # Configuração do container
└── docker-compose.yml        # Orquestração dos serviços
```

## 🛠️ Instalação

### Pré-requisitos

- Docker
- Docker Compose
- Go 1.21+ (para desenvolvimento local)

### Configuração

1. Clone o repositório:
```bash
git clone https://github.com/seu-usuario/e-commerce.git
cd e-commerce
```

2. Inicie os containers:
```bash
docker-compose up -d
```

3. A API estará disponível em:
```
http://localhost:3000
```

## 📡 Endpoints

### Produtos

#### Listar todos os produtos
```http
GET /products
```

#### Criar novo produto
```http
POST /products
Content-Type: application/json

{
    "name": "Product Name",
    "description": "Product Description",
    "price": 29.99,
    "rating": 4.5,
    "size": "M",
    "reviews": ["Great product!", "Love it!"],
    "image": "https://example.com/image.jpg"
}
```

## 🗄️ Banco de Dados

### Estrutura da tabela

```sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    rating DECIMAL(3,2) DEFAULT 0.0,
    size VARCHAR(255) DEFAULT 'M',
    reviews TEXT[] DEFAULT '{}',
    image VARCHAR(255)
);
```

### Comandos Úteis

```bash
# Conectar ao banco
docker-compose exec db psql -U postgres -d ecommerce

# Listar tabelas
\dt

# Sair do psql
\q
```

## 🔧 Desenvolvimento

### Hot Reload

O projeto usa Air para hot reload durante o desenvolvimento:

```bash
# Iniciar com hot reload
docker-compose up
```

### Logs

```bash
# Ver logs em tempo real
docker-compose logs -f

# Ver logs específicos
docker-compose logs -f backend
docker-compose logs -f db
```

## 🧪 Testes

```bash
# Executar testes
go test ./...

# Com cobertura
go test -cover ./...
```

## 📦 Build

```bash
# Build local
go build -o app cmd/main.go

# Build com Docker
docker build -t ecommerce-api .
```

## 🔐 Variáveis de Ambiente

```env
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=ecommerce
```

## 🤝 Contribuindo

1. Fork o projeto
2. Crie sua branch de feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 👥 Autores

* **Seu Nome** - *Trabalho Inicial* - [SeuUsuario](https://github.com/SeuUsuario)

## 🎉 Agradecimentos

* Documentação Go
* Comunidade Go
* Contribuidores