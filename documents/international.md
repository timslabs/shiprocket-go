# International

```go
client := shiprocket.NewClient(email, password)

data, err := client.International.CreateAdhoc(map[string]interface{}{ /* ... */ }, nil)
data, err = client.International.UpdateAdhoc(map[string]interface{}{ /* ... */ }, nil)
data, err = client.International.Track(nil, nil)
data, err = client.International.Serviceability(nil, nil)
data, err = client.International.AssignAwb(map[string]interface{}{ /* ... */ }, nil)
data, err = client.International.GenerateManifest(map[string]interface{}{ /* ... */ }, nil)
data, err = client.International.CreateForwardShipment(map[string]interface{}{ /* ... */ }, nil)
data, err = client.International.AddBankDetails(map[string]interface{}{ /* ... */ }, nil)
data, err = client.International.Kyc(map[string]interface{}{ /* ... */ }, nil)
```
