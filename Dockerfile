FROM golang:1.23-alpine

WORKDIR /app

# Copiar apenas o go.mod e go.sum primeiro
COPY go.mod go.sum ./
RUN go mod download

# Copiar o resto do código
COPY . .

# Compilar a aplicação
RUN go build -o main ./cmd/ordersystem

# Copiar o .env para o diretório raiz
COPY cmd/ordersystem/.env .

EXPOSE 8000 8080 50051

CMD ["./main"]
