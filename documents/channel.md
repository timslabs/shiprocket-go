# Channel

```go
client := shiprocket.NewClient(email, password)

data, err := client.Channel.All(nil, nil)
data, err = client.Channel.Countries(nil, nil)
data, err = client.Channel.CountryZones("1", nil, nil)
data, err = client.Channel.PostcodeDetails(map[string]interface{}{
	"postcode": "110030",
}, nil)
```
