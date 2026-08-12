## International

### Create international adhoc order

```go
data := map[string]interface{}{
  "order_id": "INT-1",
}
body, err := client.International.CreateAdhoc(data, nil)
```

**Parameters:** international order payload as per Shiprocket docs.

**Response:**

```json
{
  "order_id": "INT-1",
  "shipment_id": 3001,
  "status": "NEW"
}
```

-------------------------------------------------------------------------------------------------------

### Update international adhoc order

```go
data := map[string]interface{}{
  "order_id": "INT-1",
}
body, err := client.International.UpdateAdhoc(data, nil)
```

**Parameters:**

| Name       | Type   | Description |
|------------|--------|-------------|
| order_id\* | string | Order ID    |

Plus fields to update.

**Response:** updated international order object.

-------------------------------------------------------------------------------------------------------

### Track international order

```go
query := map[string]interface{}{
  "order_id": "INT-1",
}
body, err := client.International.Track(query, nil)
```

**Parameters:**

| Name       | Type   | Description |
|------------|--------|-------------|
| order_id\* | string | Order ID    |

**Response:**

```json
{
  "tracking_data": {
    "track_status": 1,
    "shipment_status": 7,
    "shipment_track": []
  }
}
```

-------------------------------------------------------------------------------------------------------

### International serviceability

```go
query := map[string]interface{}{
  "delivery_country": "US",
}
body, err := client.International.Serviceability(query, nil)
```

**Parameters:**

| Name               | Type   | Description        |
|--------------------|--------|--------------------|
| delivery_country\* | string | Destination country|

Plus weight/dimensions fields from Shiprocket docs.

**Response:**

```json
{
  "status": 200,
  "data": {
    "available_courier_companies": [
      {
        "courier_name": "Intl Express",
        "rate": 1200
      }
    ]
  }
}
```

-------------------------------------------------------------------------------------------------------

### Assign international AWB

```go
data := map[string]interface{}{
  "shipment_id": 3001,
}
body, err := client.International.AssignAwb(data, nil)
```

**Parameters:**

| Name          | Type    | Description   |
|---------------|---------|---------------|
| shipment_id\* | integer | Shipment ID   |

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

### Generate international manifest

```go
data := map[string]interface{}{
  "shipment_id": []interface{}{3001},
}
body, err := client.International.GenerateManifest(data, nil)
```

**Parameters:**

| Name          | Type  | Description              |
|---------------|-------|--------------------------|
| shipment_id\* | array | One or more shipment IDs |

**Response:**

```json
{
  "manifest_url": "https://example.com/manifest.pdf"
}
```

-------------------------------------------------------------------------------------------------------

### Create international forward shipment

```go
data := map[string]interface{}{
  "order_id": 1,
}
body, err := client.International.CreateForwardShipment(data, nil)
```

**Parameters:**

| Name       | Type    | Description |
|------------|---------|-------------|
| order_id\* | integer | Order ID    |

**Response:** shipment object.

-------------------------------------------------------------------------------------------------------

### Add bank details

```go
data := map[string]interface{}{
  "account_number": "1",
}
body, err := client.International.AddBankDetails(data, nil)
```

**Parameters:** bank detail fields as per Shiprocket docs.

**Response:**

```json
{
  "success": true,
  "message": "Bank details added successfully"
}
```

-------------------------------------------------------------------------------------------------------

### International KYC

```go
data := map[string]interface{}{
  "document_type": "PASSPORT",
}
body, err := client.International.Kyc(data, nil)
```

**Parameters:**

| Name            | Type   | Description     |
|-----------------|--------|-----------------|
| document_type\* | string | KYC document type|

Plus document fields from Shiprocket docs.

**Response:**

```json
{
  "success": true,
  "message": "KYC submitted successfully"
}
```

-------------------------------------------------------------------------------------------------------

**PN: \* indicates mandatory fields**

**For reference click [here](https://apidocs.shiprocket.in/)**
