# Dockerfile

FROM golang:latest

RUN apt-get update && apt-get install -y gcc

WORKDIR /app

COPY . .

RUN go build -o out

CMD ["./out"]
