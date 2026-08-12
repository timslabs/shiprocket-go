# Product

```go
client := shiprocket.NewClient(email, password)

data, err := client.Product.All(nil, nil)
data, err = client.Product.Create(map[string]interface{}{ /* ... */ }, nil)
data, err = client.Product.Fetch("123", nil, nil)
data, err = client.Product.Sample(nil, nil)

f, _ := os.Open("products.csv")
defer f.Close()
data, err = client.Product.Import(requests.FileUploadParams{File: f}, nil)
```
