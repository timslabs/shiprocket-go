# Inventory

```go
client := shiprocket.NewClient(email, password)

data, err := client.Inventory.All(nil, nil)
data, err = client.Inventory.Update("123", map[string]interface{}{
	"quantity": 10,
}, nil)
```
