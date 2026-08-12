# Shipment

```go
client := shiprocket.NewClient(email, password)

data, err := client.Shipment.All(map[string]interface{}{"page": 1}, nil)
data, err = client.Shipment.Fetch("123", nil, nil)
data, err = client.Shipment.CreateForward(map[string]interface{}{ /* ... */ }, nil)
data, err = client.Shipment.CreateReturn(map[string]interface{}{ /* ... */ }, nil)
```
