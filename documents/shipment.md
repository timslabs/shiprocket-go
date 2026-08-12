## Shipments

### Fetch all shipments

```go
query := map[string]interface{}{
  "page": 1,
}
body, err := client.Shipment.All(query, nil)
```

**Parameters:**

| Name | Type    | Description |
|------|---------|-------------|
| page | integer | Page number |

**Response:**

```json
{
  "data": [
    {
      "id": 2001,
      "awb": "AWB123456",
      "status": "PICKED UP"
    }
  ]
}
```

-------------------------------------------------------------------------------------------------------

### Fetch particular shipment

```go
body, err := client.Shipment.Fetch("2001", nil, nil)
```

**Parameters:**

| Name          | Type   | Description  |
|---------------|--------|--------------|
| shipmentID\*  | string | Shipment ID  |

**Response:**

```json
{
  "id": 2001,
  "awb": "AWB123456",
  "courier": "Demo Courier",
  "status": "PICKED UP"
}
```

-------------------------------------------------------------------------------------------------------

### Create forward shipment

```go
data := map[string]interface{}{
  "order_id": 1001,
}
body, err := client.Shipment.CreateForward(data, nil)
```

**Parameters:**

| Name       | Type    | Description              |
|------------|---------|--------------------------|
| order_id\* | integer | Shiprocket order ID      |

Plus other forward-shipment fields from Shiprocket docs.

**Response:**

```json
{
  "id": 2001,
  "awb": "AWB123456",
  "status": "PICKED UP"
}
```

-------------------------------------------------------------------------------------------------------

### Create return shipment

```go
data := map[string]interface{}{
  "order_id": 1001,
}
body, err := client.Shipment.CreateReturn(data, nil)
```

**Parameters:**

| Name       | Type    | Description              |
|------------|---------|--------------------------|
| order_id\* | integer | Shiprocket order ID      |

**Response:** shipment object for the return.

-------------------------------------------------------------------------------------------------------

**PN: \* indicates mandatory fields**

**For reference click [here](https://apidocs.shiprocket.in/)**
