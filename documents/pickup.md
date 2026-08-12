## Pickup

### Fetch all pickup locations

```go
body, err := client.Pickup.All(nil, nil)
```

**Parameters:**

Optional list query params.

**Response:**

```json
{
  "data": {
    "shipping_address": [
      {
        "id": 1,
        "pickup_location": "Primary",
        "name": "Jane Doe",
        "phone": "9999999999",
        "city": "Mumbai",
        "pin_code": "400001"
      }
    ]
  }
}
```

-------------------------------------------------------------------------------------------------------

### Add pickup location

```go
data := map[string]interface{}{
  "pickup_location": "Primary",
  "name":            "Jane Doe",
  "email":           "jane@example.com",
  "phone":           "9999999999",
  "address":         "221B Baker Street",
  "city":            "Mumbai",
  "state":           "Maharashtra",
  "country":         "India",
  "pin_code":        "400001",
}
body, err := client.Pickup.Add(data, nil)
```

**Parameters:**

| Name              | Type   | Description                    |
|-------------------|--------|--------------------------------|
| pickup_location\* | string | Unique pickup location name    |
| name\*            | string | Contact name                   |
| email\*           | string | Contact email                  |
| phone\*           | string | Contact phone                  |
| address\*         | string | Street address                 |
| city\*            | string | City                           |
| state\*           | string | State                          |
| country\*         | string | Country                        |
| pin_code\*        | string | Pincode                        |

**Response:**

```json
{
  "success": true,
  "message": "Pickup location added successfully"
}
```

-------------------------------------------------------------------------------------------------------

**PN: \* indicates mandatory fields**

**For reference click [here](https://apidocs.shiprocket.in/)**
