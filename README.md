# API para gerenciamento de estoque

## Explicação da forma de uso

### Rotas: 
 - /products
 - /product
 - /product/{id}
 - /sale/{id}


### Funcionalidade da rota /products:
 - Lista todos os produtos salvos no banco de dados
 - Método aceito: **GET**
 - Exemplo de uso: `curl localhost:3000/products`

### Funcionalidade da rota /product:
 - Cria um novo produto (a rota passada precisa ser /product, caso seja /product/ dará erro)
 - Método aceito: **POST**
 - Modelo de requisição:
   - KEYS: name(string), price(float), amount(int)
   - `{"name": "foo", "price": 10.50, "amount": 100}`
   - **Todos os campos são obrigatórios**
 - Exemplo de uso: `curl -X POST -d '{"name": "foo", "price": 10.50, "amount": 100}' localhost:3000/product`

### Funcionalidades da rota /product/{id}:
 - Deleta um produto, atualiza/edita e busca um determinado produto
 - Métodos aceitos: **DELETE**, **PUT** e **GET**
 - Modelo de requisição (válido apenas pro metódo PUT):
   - KEYS: name(string), price(float), amount(int)
   - `{"name": "foo", "price": 10.50, "amount": 100}`
   - **Todos os campos são obrigatórios**
  - Exemplo de exclusão: `curl -X DELETE localhost:3000/product/{id}`
  - Exemplo de edição: `curl -X PUT -d '{"name": "foo", "price": 10.50, "amount": 100}' localhost:3000/product/{id}`
  - Exemplo de busca: `curl localhost:3000/product/{id}`

### Funcionalidade da rota /sale/{id}:
 - Atualiza a quantidade de produtos vendidos
 - Método aceito: **PUT**
 - Modelo de requisição:
   - KEY: amount(int)
   - `{"amount": 10}`
   - **TODOS OS CAMPOS SÂO OBRIGATÓRIOS**
 - Exemplo de uso: `curl -X PUT -d '{"amount": 10}' localhost:3000/sale/{id}`
