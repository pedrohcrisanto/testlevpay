# Levpay Test - Implementação
## Configuração

Clone o projeto

```sh
$ git clone https://github.com/pedrohcrisanto/levpaytest.git
```


Crie o banco
```sh
$ createdb "nome-do-banco"
```
Edite no arquivo main.go suas configurações de banco de dados
	host     = "localhost"
	port     = 5432
	user     = "nome-user"
	password = "senha-user"
	dbname   = "nome-do-banco"
```sh

$ createdb "nome-do-banco"
```


Crie a tabela superheros no seu banco de dados
```sh
CREATE TABLE superheros (
  id SERIAL PRIMARY KEY,
  name TEXT,
  fullname TEXT,
  intelligence INT,
  power INT,
  occupation TEXT,
  image TEXT,
  uuid TEXT
);

```

Para rodar a aplicação(http://localhost:8000/)
```sh
$ go run *.go
```
## Documentação
Versão GO: 1.13.8

Dependencias: 
github.com/lib/pq
github.com/gorilla/mux
github.com/google/uuid

Rotas:
GET
/superherosapi/id
## Cria um superhero no banco de dados utilizando a api superheroapi.com, passando o ID do superhero da API SuperHero.
GET
/superheros
## Lista todos os superheros gravados no Banco de Dados
GET
/superheros/id
## Lista o superhero gravado no banco de dados da aplicação passando o parametro id
GET
/search/query
## Busca na aplicação o superhero a partir de seu name ou UUID
PUT
/superheros/id
## Atualiza o superhero referente ao id passado como parametro e atualiza seu name a partir de um JSON passado no Body da requisição.
DELETE
/superheros/id
## Deleta o superhero referente ao ID passado como parametro.
