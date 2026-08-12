## Account

### Wallet balance

```go
body, err := client.Account.WalletBalance(nil, nil)
```

**Parameters:**

None.

**Response:**

```json
{
  "currency_code": "INR",
  "currency_symbol": "₹",
  "balance_amount": 1500.5,
  "available_balance": 1500.5
}
```

-------------------------------------------------------------------------------------------------------

### Account statement

```go
query := map[string]interface{}{
  "from": "2026-01-01",
  "to":   "2026-01-31",
}
body, err := client.Account.Statement(query, nil)
```

**Parameters:**

| Name | Type   | Description                          |
|------|--------|--------------------------------------|
| from | string | Start date (`YYYY-MM-DD`)            |
| to   | string | End date (`YYYY-MM-DD`)              |

**Response:**

```json
{
  "status": 200,
  "data": [
    {
      "type": "debit",
      "amount": 45.5,
      "created_at": "2026-01-15 10:00:00"
    }
  ]
}
```

-------------------------------------------------------------------------------------------------------

### Billing discrepancy

```go
body, err := client.Account.Discrepancy(nil, nil)
```

**Parameters:**

Optional query filters as supported by Shiprocket.

**Response:**

```json
{
  "status": 200,
  "data": []
}
```

-------------------------------------------------------------------------------------------------------

**PN: \* indicates mandatory fields**

**For reference click [here](https://apidocs.shiprocket.in/)**
