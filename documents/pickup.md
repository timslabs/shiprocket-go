# Pickup

```go
client := shiprocket.NewClient(email, password)

data, err := client.Pickup.All(nil, nil)
data, err = client.Pickup.Add(map[string]interface{}{
	"pickup_location": "Primary",
	"name":            "Jane",
	"email":           "jane@example.com",
	"phone":           "9999999999",
	"address":         "221B Baker Street",
	"city":            "Mumbai",
	"state":           "Maharashtra",
	"country":         "India",
	"pin_code":        "400001",
}, nil)
```
