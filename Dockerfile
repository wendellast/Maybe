# Dockerfile

# Usar a imagem base do Go
FROM golang:latest

# Instalar GCC e outras dependências
RUN apt-get update && apt-get install -y gcc

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar o código fonte para o contêiner
COPY . .

# Build do aplicativo
RUN go build -o out

# Comando padrão para executar o aplicativo quando o contêiner for iniciado
CMD ["./out"]
