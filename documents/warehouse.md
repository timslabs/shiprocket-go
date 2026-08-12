## Warehouse

Warehouse endpoints use `/v1/warehouse` (not `/v1/external`).

### SRF serviceability

```go
query := map[string]interface{}{
  "postcode": "110030",
  "sku":      "Baby-socks",
  "quantity": 1,
}
body, err := client.Warehouse.SrfServiceability(query, nil)
```

**Parameters:**

| Name       | Type    | Description              |
|------------|---------|--------------------------|
| postcode\* | string  | Delivery pincode         |
| sku\*      | string  | SKU                      |
| quantity\* | integer | Units requested          |

**Response:**

```json
{
  "status": 200,
  "data": {
    "serviceable": true
  }
}
```

-------------------------------------------------------------------------------------------------------

**PN: \* indicates mandatory fields**

**For reference click [here](https://apidocs.shiprocket.in/)**
