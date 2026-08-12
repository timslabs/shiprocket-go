## Orders

### Create adhoc order

```go
data := map[string]interface{}{
  "order_id":              "ORD-1001",
  "order_date":            "2026-08-13",
  "pickup_location":       "Primary",
  "billing_customer_name": "Jane",
  "billing_last_name":     "Doe",
  "billing_address":       "221B Baker Street",
  "billing_city":          "Mumbai",
  "billing_pincode":       "400001",
  "billing_state":         "Maharashtra",
  "billing_country":       "India",
  "billing_email":         "jane@example.com",
  "billing_phone":         "9999999999",
  "shipping_is_billing":   true,
  "payment_method":        "Prepaid",
  "sub_total":             499,
  "length":                10,
  "breadth":               10,
  "height":                10,
  "weight":                0.5,
  "order_items": []map[string]interface{}{
    {
      "name":          "Widget",
      "sku":           "W-1",
      "units":         1,
      "selling_price": 499,
    },
  },
}
body, err := client.Order.CreateAdhoc(data, nil)
```

**Parameters:**

| Name                    | Type    | Description                              |
|-------------------------|---------|------------------------------------------|
| order_id\*              | string  | Your channel order ID                    |
| order_date\*            | string  | Order date                               |
| pickup_location\*       | string  | Registered pickup location name          |
| billing_customer_name\* | string  | Billing first name                       |
| billing_address\*       | string  | Billing address                          |
| billing_city\*          | string  | Billing city                             |
| billing_pincode\*       | string  | Billing pincode                          |
| billing_state\*         | string  | Billing state                            |
| billing_country\*       | string  | Billing country                          |
| billing_email\*         | string  | Billing email                            |
| billing_phone\*         | string  | Billing phone                            |
| payment_method\*        | string  | `Prepaid` or `COD`                       |
| sub_total\*             | number  | Order subtotal                           |
| length\* / breadth\* / height\* | number | Dimensions (cm)                   |
| weight\*                | number  | Weight (kg)                              |
| order_items\*           | array   | Line items                               |
| shipping_is_billing     | boolean | Use billing address for shipping         |

**Response:**

```json
{
  "order_id": 1001,
  "channel_order_id": "ORD-1001",
  "shipment_id": 2001,
  "status": "NEW",
  "status_code": 1
}
```

-------------------------------------------------------------------------------------------------------

### Create order

```go
body, err := client.Order.Create(data, nil)
```

**Parameters:** same family of fields as Create adhoc (channel order payload).

**Response:** same shape as Create adhoc.

-------------------------------------------------------------------------------------------------------

### Create return order

```go
body, err := client.Order.CreateReturn(data, nil)
```

**Parameters:** return-order payload as per Shiprocket docs.

**Response:**

```json
{
  "order_id": 1001,
  "shipment_id": 2001,
  "status": "NEW"
}
```

-------------------------------------------------------------------------------------------------------

### Update adhoc order

```go
data := map[string]interface{}{
  "order_id": "ORD-1001",
}
body, err := client.Order.UpdateAdhoc(data, nil)
```

**Parameters:**

| Name       | Type   | Description                 |
|------------|--------|-----------------------------|
| order_id\* | string | Order / channel order ID    |

Plus fields to update.

**Response:** updated order object.

-------------------------------------------------------------------------------------------------------

### Cancel orders

```go
data := map[string]interface{}{
  "ids": []interface{}{1001},
}
body, err := client.Order.Cancel(data, nil)
```

**Parameters:**

| Name  | Type  | Description              |
|-------|-------|--------------------------|
| ids\* | array | Shiprocket order IDs     |

**Response:**

```json
{
  "message": "Orders cancelled successfully"
}
```

-------------------------------------------------------------------------------------------------------

### Fetch all orders

```go
query := map[string]interface{}{
  "page": 1,
}
body, err := client.Order.All(query, nil)
```

**Parameters:**

| Name | Type    | Description        |
|------|---------|--------------------|
| page | integer | Page number        |

Other list filters supported by Shiprocket may be passed in `query`.

**Response:**

```json
{
  "data": [
    {
      "id": 1001,
      "channel_order_id": "ORD-1001",
      "status": "NEW"
    }
  ],
  "meta": {
    "pagination": {
      "total": 2,
      "count": 2,
      "per_page": 15,
      "current_page": 1,
      "total_pages": 1
    }
  }
}
```

-------------------------------------------------------------------------------------------------------

### Fetch particular order

```go
body, err := client.Order.Fetch("1001", nil, nil)
```

**Parameters:**

| Name       | Type   | Description              |
|------------|--------|--------------------------|
| orderID\*  | string | Shiprocket order ID      |

**Response:**

```json
{
  "order_id": 1001,
  "channel_order_id": "ORD-1001",
  "shipment_id": 2001,
  "status": "NEW",
  "status_code": 1
}
```

-------------------------------------------------------------------------------------------------------

### Processing returns

```go
body, err := client.Order.ProcessingReturns(nil, nil)
```

**Parameters:** optional list query params.

**Response:** collection of return orders in processing.

-------------------------------------------------------------------------------------------------------

### Update address

```go
body, err := client.Order.UpdateAddress(data, nil)
```

**Parameters:** address update payload (`order_id`, shipping/billing fields).

**Response:**

```json
{
  "message": "Address updated successfully"
}
```

-------------------------------------------------------------------------------------------------------

### Update pickup address

```go
data := map[string]interface{}{
  "order_id":        []interface{}{1001},
  "pickup_location": "Primary",
}
body, err := client.Order.UpdatePickupAddress(data, nil)
```

**Parameters:**

| Name              | Type  | Description                 |
|-------------------|-------|-----------------------------|
| order_id\*        | array | Order IDs                   |
| pickup_location\* | string| Pickup location name        |

**Response:**

```json
{
  "message": "Pickup address updated successfully"
}
```

-------------------------------------------------------------------------------------------------------

### Fulfill orders

```go
body, err := client.Order.Fulfill(data, nil)
```

**Parameters:** fulfill payload as per Shiprocket docs.

**Response:**

```json
{
  "message": "Orders fulfilled successfully"
}
```

-------------------------------------------------------------------------------------------------------

### Map orders

```go
body, err := client.Order.Mapping(data, nil)
```

**Parameters:** mapping payload as per Shiprocket docs.

**Response:**

```json
{
  "message": "Orders mapped successfully"
}
```

-------------------------------------------------------------------------------------------------------

### Export orders

```go
body, err := client.Order.Export(nil, nil)
```

**Parameters:** optional export filters in body.

**Response:**

```json
{
  "import_id": 77,
  "status": "completed"
}
```

-------------------------------------------------------------------------------------------------------

### Import orders (CSV)

```go
f, err := os.Open("orders.csv")
if err != nil {
  log.Fatal(err)
}
defer f.Close()

body, err := client.Order.Import(requests.FileUploadParams{
  File: f,
  Fields: map[string]string{
    "type": "orders",
  },
}, nil)
```

**Parameters:**

| Name   | Type   | Description                    |
|--------|--------|--------------------------------|
| File\* | \*os.File | CSV/Excel file handle        |
| Fields | map    | Extra multipart form fields    |

**Response:**

```json
{
  "import_id": 77,
  "status": "completed",
  "errors": []
}
```

-------------------------------------------------------------------------------------------------------

### Print invoice

```go
data := map[string]interface{}{
  "ids": []interface{}{1001},
}
body, err := client.Order.PrintInvoice(data, nil)
```

**Parameters:**

| Name  | Type  | Description           |
|-------|-------|-----------------------|
| ids\* | array | Order IDs             |

**Response:**

```json
{
  "invoice_url": "https://example.com/invoice.pdf"
}
```

-------------------------------------------------------------------------------------------------------

**PN: \* indicates mandatory fields**

**For reference click [here](https://apidocs.shiprocket.in/)**
