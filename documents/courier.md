## Courier

### Check serviceability

```go
query := map[string]interface{}{
  "pickup_postcode":   "110030",
  "delivery_postcode": "122001",
  "weight":            0.5,
  "cod":               0,
}
body, err := client.Courier.Serviceability(query, nil)
```

**Parameters:**

| Name                 | Type    | Description                                      |
|----------------------|---------|--------------------------------------------------|
| pickup_postcode\*    | string  | Pickup pincode                                    |
| delivery_postcode\*  | string  | Delivery pincode                                  |
| weight\*             | float   | Package weight in kg                             |
| cod                  | integer | `0` prepaid, `1` COD                             |
| declared_value       | float   | Declared value of shipment                       |
| mode                 | string  | e.g. `Surface`, `Air`                            |
| is_return            | integer | `1` for return serviceability                    |

**Response:**

```json
{
  "status": 200,
  "data": {
    "available_courier_companies": [
      {
        "courier_company_id": 1,
        "courier_name": "Demo Courier",
        "rate": 45.5,
        "etd": "2 days",
        "cod": 0
      }
    ]
  }
}
```

-------------------------------------------------------------------------------------------------------

### List couriers with counts

```go
query := map[string]interface{}{
  "type": "active",
}
body, err := client.Courier.ListWithCounts(query, nil)
```

**Parameters:**

| Name | Type   | Description                         |
|------|--------|-------------------------------------|
| type | string | Optional filter (e.g. `active`)     |

**Response:**

```json
{
  "status": 200,
  "data": {
    "available_courier_companies": []
  }
}
```

-------------------------------------------------------------------------------------------------------

### Assign AWB

```go
data := map[string]interface{}{
  "shipment_id": 2001,
}
body, err := client.Courier.AssignAwb(data, nil)
```

**Parameters:**

| Name          | Type    | Description                          |
|---------------|---------|--------------------------------------|
| shipment_id\* | integer | Shipment ID to assign AWB            |
| courier_id    | integer | Preferred courier company ID         |

**Response:**

```json
{
  "success": true,
  "message": "AWB assigned successfully",
  "response": {
    "data": {
      "awb_code": "AWB123456",
      "courier_name": "Demo Courier"
    }
  }
}
```

-------------------------------------------------------------------------------------------------------

### Generate pickup

```go
data := map[string]interface{}{
  "shipment_id": []interface{}{2001},
}
body, err := client.Courier.GeneratePickup(data, nil)
```

**Parameters:**

| Name          | Type  | Description                         |
|---------------|-------|-------------------------------------|
| shipment_id\* | array | One or more shipment IDs            |

**Response:**

```json
{
  "success": true,
  "message": "Pickup generated successfully"
}
```

-------------------------------------------------------------------------------------------------------

### Generate label

```go
data := map[string]interface{}{
  "shipment_id": []interface{}{2001},
}
body, err := client.Courier.GenerateLabel(data, nil)
```

**Parameters:**

| Name          | Type  | Description                         |
|---------------|-------|-------------------------------------|
| shipment_id\* | array | One or more shipment IDs            |

**Response:**

```json
{
  "label_created": 1,
  "label_url": "https://example.com/label.pdf"
}
```

-------------------------------------------------------------------------------------------------------

### Generate manifest

```go
data := map[string]interface{}{
  "shipment_id": []interface{}{2001},
}
body, err := client.Courier.GenerateManifest(data, nil)
```

**Parameters:**

| Name          | Type  | Description                         |
|---------------|-------|-------------------------------------|
| shipment_id\* | array | One or more shipment IDs            |

**Response:**

```json
{
  "manifest_url": "https://example.com/manifest.pdf"
}
```

-------------------------------------------------------------------------------------------------------

### Print manifest

```go
data := map[string]interface{}{
  "shipment_id": []interface{}{2001},
}
body, err := client.Courier.PrintManifest(data, nil)
```

**Parameters:**

| Name          | Type  | Description                         |
|---------------|-------|-------------------------------------|
| shipment_id\* | array | One or more shipment IDs            |

**Response:**

```json
{
  "manifest_url": "https://example.com/manifest.pdf"
}
```

-------------------------------------------------------------------------------------------------------

### Cancel shipment by AWB

```go
data := map[string]interface{}{
  "awbs": []interface{}{"AWB123456"},
}
body, err := client.Courier.CancelShipment(data, nil)
```

**Parameters:**

| Name   | Type  | Description                |
|--------|-------|----------------------------|
| awbs\* | array | AWB codes to cancel        |

**Response:**

```json
{
  "message": "Shipment cancelled successfully"
}
```

-------------------------------------------------------------------------------------------------------

### Track by AWB

```go
body, err := client.Courier.TrackByAwb("AWB123456", nil, nil)
```

**Parameters:**

| Name      | Type   | Description     |
|-----------|--------|-----------------|
| awbCode\* | string | AWB code        |

**Response:**

```json
{
  "tracking_data": {
    "track_status": 1,
    "shipment_status": 7,
    "shipment_track": [
      {
        "awb_code": "AWB123456",
        "courier_name": "Demo Courier",
        "current_status": "Delivered"
      }
    ]
  }
}
```

-------------------------------------------------------------------------------------------------------

### Track by shipment

```go
body, err := client.Courier.TrackByShipment("2001", nil, nil)
```

**Parameters:**

| Name          | Type   | Description   |
|---------------|--------|---------------|
| shipmentID\*  | string | Shipment ID   |

**Response:** same shape as Track by AWB.

-------------------------------------------------------------------------------------------------------

### Track by order

```go
query := map[string]interface{}{
  "order_id": "ORD-1001",
}
body, err := client.Courier.TrackByOrder(query, nil)
```

**Parameters:**

| Name       | Type   | Description                          |
|------------|--------|--------------------------------------|
| order_id\* | string | Channel / order reference ID         |

**Response:** same shape as Track by AWB.

-------------------------------------------------------------------------------------------------------

### Track by multiple AWBs

```go
data := map[string]interface{}{
  "awbs": []interface{}{"AWB123456", "AWB654321"},
}
body, err := client.Courier.TrackByAwbs(data, nil)
```

**Parameters:**

| Name   | Type  | Description           |
|--------|-------|-----------------------|
| awbs\* | array | List of AWB codes     |

**Response:**

```json
{
  "tracking_data": []
}
```

-------------------------------------------------------------------------------------------------------

**PN: \* indicates mandatory fields**

**For reference click [here](https://apidocs.shiprocket.in/)**
