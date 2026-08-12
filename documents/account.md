# Account

Wallet balance, statement, and billing discrepancy.

## WalletBalance

```go
client := shiprocket.NewClient(email, password)

data, err := client.Account.WalletBalance(nil, nil)
```

## Statement

```go
data, err := client.Account.Statement(map[string]interface{}{
	"from": "2026-01-01",
	"to":   "2026-01-31",
}, nil)
```

## Discrepancy

```go
data, err := client.Account.Discrepancy(nil, nil)
```
