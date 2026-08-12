# Ndr

```go
client := shiprocket.NewClient(email, password)

data, err := client.Ndr.All(nil, nil)
data, err = client.Ndr.Fetch("AWB", nil, nil)
data, err = client.Ndr.Action("AWB", map[string]interface{}{
	"action": "re-attempt",
}, nil)
```
