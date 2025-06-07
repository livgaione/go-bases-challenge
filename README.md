# go-bases-challenge

Desafio final do módulo de bases do Go.

## :rocket: Como rodar

```
go run ./cmd
```

A API estará disponível em:

```
http://localhost:8080
```

## :package: Endpoints

### :mag:  Buscar por país
Retorna todos os tickets com destino ao país especificado.



Request: 
```
GET /tickets/{Brazil}
```

Response:
```json

{
  "country": "Brazil",
  "count": 1,
  "data": [
        {
            "ID": "5",
            "Name": "Ana Doe",
            "Email": "adoe@google.com.br",
            "Country": "Brazil",
            "Hour": "0:31",
            "Price": 1398
        },  ]
}
```

### :mag: Buscar por horário

Retorna todos os tickets comprados para um determinado horário.



Request: 
```
GET /tickets/{Morning}
```

Response:
```json

{
  "count": 1,
  "data": [
           {
            "ID": "5",
            "Name": "Saree Nobes",
            "Email": "snobes4@google.com.au",
            "Country": "Czech Republic",
            "Hour": "0:31",
            "Price": 1398
            },
          ]   
}
```

### :mag: Percentual de viagens para o país

Retorna o percentual de pessoas que viajaram para um determinado país.


Request: 
```
GET /tickets/{Argentina}/percentage
```

Response:

```json
{
  "count": 1,
  "data": [
           {
            "ID": "188",
            "Name": "Sara Lindus",
            "Email": "sali@google.com.ar",
            "Country": "Argentina",
            "Hour": "18:36",
            "Price": 1238
        },
  ]
}
```

## :test_tube: Testes

```
go test ./...

```

## :sparkles: Tecnologias 

- Go

- Chi router

- Testify

