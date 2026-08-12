# Auth

```go
client := shiprocket.NewClient(email, password)

// Prefer lazy auth on first API call, or:
_ = client.Authenticate()

data, err := client.Auth.Login(map[string]interface{}{
	"email":    email,
	"password": password,
}, nil)

_, err = client.Auth.Logout(nil, nil)
```
