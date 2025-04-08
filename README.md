# desafio-clean-arch

Este projeto implementa uma API usando Clean Architecture e Api Rest, GraphQL e gRPC. Também usa RabbitMQ.

## Serviços e Portas

- **API REST**: `http://localhost:8000`
  - POST /order - Criar pedido
    ```json
    {
      "id": "123",
      "price": 100.00,
      "tax": 10.00
    }
    ```
  - GET /orders - Listar pedidos

- **GraphQL**: `http://localhost:8080`
  - Playground: `http://localhost:8080`
  - Endpoint: `http://localhost:8080/query`
  - Mutations:
    - createOrder
  - Queries:
    - listOrders

- **gRPC**: `localhost:50051`
  - Serviço: OrderService
  - Métodos:
    - CreateOrder
    - ListOrders

- **Banco de Dados**: `localhost:3306`
  - Driver: MySQL
  - Usuário: root
  - Senha: root
  - Database: orders

- **Mensageria**: `localhost:5672`
  - RabbitMQ Management: `http://localhost:15672`
  - Usuário: guest
  - Senha: guest

## Passos para Execução

1. **Pré-requisitos**
   - Docker
   - Docker Compose

2. **Configuração do Ambiente**
   ```bash
   # Clonar o repositório
   git clone [URL_DO_REPOSITORIO]
   cd desafio-clean-arch

   # Iniciar os containers
   docker-compose up -d
   ```

3. **Testar as APIs**
   - REST: Use o Postman ou curl
   - GraphQL: Acesse o playground em `http://localhost:8080`
   - gRPC: Use o Evans CLI (já instalado no container)

4. **Parar os Containers (quando necessário)**
   ```bash
   docker-compose down
   ```