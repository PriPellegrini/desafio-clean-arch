# desafio-clean-arch

Este projeto implementa uma API usando Clean Architecture e Api Rest, GraphQL e gRPC. Também usa RabbitMQ.

## Serviços e Portas

- **API REST**: `http://localhost:8000`
  - POST /order - Criar pedido
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
   - Go 1.16+
   - Protoc (para gRPC)

2. **Configuração do Ambiente**
   ```bash
   # Clonar o repositório
   git clone [URL_DO_REPOSITORIO]
   cd desafio-clean-arch

   # Iniciar os containers
   docker-compose up -d
   ```

3. **Gerar arquivos gRPC**
   ```bash
   protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/order.proto
   ```

4. **Executar a Aplicação**
   ```bash
   cd cmd/ordersystem
   go run main.go wire.go
   ```

5. **Testar as APIs**
   - REST: Use o Postman ou curl
   - GraphQL: Acesse o playground em `http://localhost:8080`
   - gRPC: Use o Evans CLI