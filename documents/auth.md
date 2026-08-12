## Auth

### Login

```go
data := map[string]interface{}{
  "email":    "api@example.com",
  "password": "your-api-password",
}
body, err := client.Auth.Login(data, nil)
```

Prefer `shiprocket.NewClient(email, password)` (lazy login) or `client.Authenticate()` for normal use.

**Parameters:**

| Name       | Type   | Description                                      |
|------------|--------|--------------------------------------------------|
| email\*    | string | Shiprocket API user email                        |
| password\* | string | Shiprocket API user password                     |

**Response:**

```json
{
  "token": "fake_jwt_token",
  "email": "api@example.com",
  "company_id": 12345,
  "first_name": "API",
  "last_name": "User"
}
```

-------------------------------------------------------------------------------------------------------

### Logout

```go
body, err := client.Auth.Logout(nil, nil)
```

**Parameters:**

None (uses current bearer token).

**Response:**

```json
{
  "message": "Logged out successfully"
}
```

-------------------------------------------------------------------------------------------------------

**PN: \* indicates mandatory fields**

**For reference click [here](https://apidocs.shiprocket.in/)**
