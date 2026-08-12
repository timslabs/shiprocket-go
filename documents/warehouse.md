# Warehouse

Uses `/v1/warehouse` (not `/v1/external`).

```go
client := shiprocket.NewClient(email, password)

data, err := client.Warehouse.SrfServiceability(map[string]interface{}{
	// query params per Shiprocket docs
}, nil)
```
