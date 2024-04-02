
# GO Client Server API

A simple Golang project with http requests using context and database

## Regras da aplicação

Os requisitos para cumprir este desafio são:
 
- [x] O server.go deverá consumir a API contendo o câmbio de Dólar e Real no endereço: https://economia.awesomeapi.com.br/json/last/USD-BRL e em seguida deverá retornar no formato JSON o resultado para o cliente.
 
- [x] Usando o package "context", o server.go deverá registrar no banco de dados SQLite cada cotação recebida, sendo que o timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms e o timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms.

- [x] O client.go deverá realizar uma requisição HTTP no server.go solicitando a cotação do dólar.

- [x] O client.go precisará receber do server.go apenas o valor atual do câmbio (campo "bid" do JSON). Utilizando o package "context", o client.go terá um timeout máximo de 300ms para receber o resultado do server.go.
  
- [x] O client.go terá que salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}

- [x] Os 3 contextos deverão retornar erro nos logs caso o tempo de execução seja insuficiente.
 
O endpoint necessário gerado pelo server.go para este desafio será: /cotacao e a porta a ser utilizada pelo servidor HTTP será a 8080.

## 🛠 Construído com
![SQLite](https://img.shields.io/badge/sqlite-%2307405e.svg?style=for-the-badge&logo=sqlite&logoColor=white)

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
## Rodando localmente

No diretório do projeto

Instale as dependências

```bash
  go mod tidy
```

Inicie o server

```bash
  go run server/server.go
```

Realize a chamada através do client

```bash
  go run client/client.go
```

