# Order

```go
client := shiprocket.NewClient(email, password)

data, err := client.Order.CreateAdhoc(map[string]interface{}{ /* ... */ }, nil)
data, err = client.Order.Create(map[string]interface{}{ /* ... */ }, nil)
data, err = client.Order.CreateReturn(map[string]interface{}{ /* ... */ }, nil)
data, err = client.Order.UpdateAdhoc(map[string]interface{}{ /* ... */ }, nil)
data, err = client.Order.Cancel(map[string]interface{}{"ids": []interface{}{1}}, nil)
data, err = client.Order.All(map[string]interface{}{"page": 1}, nil)
data, err = client.Order.Fetch("123", nil, nil)
data, err = client.Order.ProcessingReturns(nil, nil)
data, err = client.Order.UpdateAddress(map[string]interface{}{ /* ... */ }, nil)
data, err = client.Order.UpdatePickupAddress(map[string]interface{}{ /* ... */ }, nil)
data, err = client.Order.Fulfill(map[string]interface{}{ /* ... */ }, nil)
data, err = client.Order.Mapping(map[string]interface{}{ /* ... */ }, nil)
data, err = client.Order.Export(nil, nil)
data, err = client.Order.PrintInvoice(map[string]interface{}{"ids": []interface{}{1}}, nil)

f, _ := os.Open("orders.csv")
defer f.Close()
data, err = client.Order.Import(requests.FileUploadParams{File: f}, nil)
```
