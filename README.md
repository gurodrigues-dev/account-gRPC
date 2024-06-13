<h1 align="center"> account-gRPC </h1>

## ⚙️ API gRPC

Para realizar as requisições é necessário importar o protobuffer `account/account.proto`

### RPC CreateBill

Cria uma nova mensalidade ou despesa

**Parâmetros**

| Nome     | Local | Tipo   | Descrição            |
|----------|-------|--------|----------------------|
| `name`   | body  | string | Nome e titulo da despesa.      |
| `description`   | body  | string | Descrição da despesa.      |
| `amount`  | body  | string | Valor das parcelas.    |
| `installment` | body | int | Quantidade de parcelas.    | 

**Resposta**

```json
{
  "bill": {
    "name": "Conta de energia",
    "description": "Descrição simples",
    "amount": "64",
    "installment": 12,
  }
}
```

---

### RPC DeleteBill

Deleta uma conta ou despesa mensal,

**Parâmetros**

| Nome     | Local | Tipo   | Descrição            |
|----------|-------|--------|----------------------|
| `id`   | body  | string | Identificador do conta.      |

**Resposta**

```json
{
    "message": "deleted w successfully"
}
```

