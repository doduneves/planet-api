
# Planet API

Gerenciador de planetas do filme Star Wars.

#### Instalação
Aplicação construida em Go Lang usando MongoDB. Portanto, é necessário ter os dois instalados na máquina.
Para instalar os módulos e dependências da aplicação:

```sh
$ go build
```

#### Execução

Para executar o programa basta rodar o comando:

```sh
$ go run main.go
```

### Endpoints
- http://127.0.0.1:8001/planets **[GET]**
Lista todos os planetas já inseridos.
Query params:
1. "nome" vai listar apenas os planetas de nome = [nome]
Ex.: http://127.0.0.1:8001/planets?nome=Terra
2. "page" para informar a página da procura. A listagem é impressa de 10 em 10 itens
Ex.: http://127.0.0.1:8001/planets?page=2


- http://127.0.0.1:8001/planets/*[id]* **[GET]**
Lista o planeta de determinado [id]

- http://127.0.0.1:8001/planets **[POST]**
Insere um novo planeta.
Parametros a serem passados no corpo da requisição (body):
{
  "nome": string,
  "clima": string,
  "terreno": string
}

- http://127.0.0.1:8001/planets/*[id]* **[DELETE]**
Remove um planeta do banco da aplicação pelo determinado [id]