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
```sh
	host     = "localhost"
	port     = 5432
	user     = "nome-user"
	password = "senha-user"
	dbname   = "nome-do-banco"
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
<br>
github.com/lib/pq
<br>
github.com/gorilla/mux
<br>
github.com/google/uuid

<br>
## GET => /superherosapi/id
Cria um superhero no banco de dados utilizando a api superheroapi.com, passando o ID do superhero da API SuperHero.
<br> <strong>Exemplo: /superherosapi/1</strong>
## GET => /superheros
Lista todos os superheros gravados no Banco de Dados.
## GET => /superheros/id
Lista o superhero gravado no banco de dados da aplicação referente ao parametro id.
<br> <strong>Exemplo: /superheros/1</strong>
## GET => /search/query
Busca no Banco de dados o superhero a partir de seu name ou UUID.
<br> <strong>Exemplo: /superheros/spider-man ou /superheros/0bd81b2d-4fa2-40da-a6b9-fa066f4a7b44</strong>
## PUT => /superheros/id
Atualiza o name do superhero referente ao id passado e o name no Body.
<br><strong>Exemplo Body: {"name": "Spider-Man"}</strong>
## DELETE => /superheros/id
Deleta o superhero referente ao ID passado como parametro.
<strong><br>Exemplo: /superheros/1</strong>
