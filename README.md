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


<h2>Rotas</h2>
<h3>POST => /superherosapi/id</h3>
```sh
Cria um superhero no banco de dados utilizando a api superheroapi.com, passando o ID do superhero da API SuperHero.com.
<br> <strong>Exemplo: /superherosapi/1</strong>
```
<h3>GET => /superheros</h3>
Lista todos os superheros gravados no Banco de Dados.
<br>
<h3>GET => /superheros/id</h3>
Lista o superhero gravado no banco de dados da aplicação referente ao parametro id.
<br> <strong>Exemplo: /superheros/1</strong>
<br>
<h3>GET => /search/query</h3>
Busca no Banco de dados o superhero a partir de seu name ou UUID.
<br> <strong>Exemplo: /superheros/spider-man ou /superheros/0bd81b2d-4fa2-40da-a6b9-fa066f4a7b44</strong>
<br>
<h3>PUT => /superheros/id</h3>
Atualiza o name do superhero referente ao id passado e o name no Body.
<br><strong>Exemplo Body: {"name": "Spider-Man"}</strong>
<br>
<h3>DELETE => /superheros/id</h3>
Deleta o superhero referente ao ID passado como parametro.
<strong><br>Exemplo: /superheros/1</strong>
