## Inventory

### Fetch all inventory

```go
body, err := client.Inventory.All(nil, nil)
```

**Parameters:**

Optional list query params.

**Response:**

```json
{
  "data": [
    {
      "product_id": 55,
      "sku": "W-1",
      "quantity": 10
    }
  ]
}
```

-------------------------------------------------------------------------------------------------------

### Update inventory

```go
data := map[string]interface{}{
  "quantity": 10,
}
body, err := client.Inventory.Update("55", data, nil)
```

**Parameters:**

| Name         | Type    | Description              |
|--------------|---------|--------------------------|
| productID\*  | string  | Product ID               |
| quantity\*   | integer | Updated available qty    |

**Response:**

```json
{
  "message": "Inventory updated successfully"
}
```

-------------------------------------------------------------------------------------------------------

**PN: \* indicates mandatory fields**

**For reference click [here](https://apidocs.shiprocket.in/)**
