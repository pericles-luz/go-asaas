# go-asaas
consome API Asaas para cobrança

Sinta-se a vontade para contribuir com o projeto, implementando as funcionalidades que achar necessárias.

## Instalação

```bash
go get -u github.com/pericles-luz/go-asaas
```

Funcionalidades implementadas:

### Autenticação

Para autenticação e utilização, basta criar um arquivo `json` como o abaixo e passar seu path como parâmetro ao criar o cliente.


```json
{
    "client_id": "4f025a64-1662-46b2-852b-f3b83172539c",
    "client_secret": "83921a65684f44d6c1a03c19012ac788",
    "link": "https://sandbox.api.parcelamostudo.tech"
}
```

### Utilização

Estude os testes disponíveis e faça suas próprias implementações. Pelos testes é possível entender como utilizar cada funcionalidade.