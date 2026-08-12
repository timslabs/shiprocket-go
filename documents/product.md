## Products

### Fetch all products

```go
body, err := client.Product.All(nil, nil)
```

**Parameters:**

Optional list query params (`page`, filters).

**Response:**

```json
{
  "products": [
    {
      "id": 55,
      "name": "Widget",
      "sku": "W-1",
      "quantity": 10
    }
  ]
}
```

-------------------------------------------------------------------------------------------------------

### Create product

```go
data := map[string]interface{}{
  "name": "Widget",
  "sku":  "W-1",
}
body, err := client.Product.Create(data, nil)
```

**Parameters:**

| Name   | Type   | Description     |
|--------|--------|-----------------|
| name\* | string | Product name    |
| sku\*  | string | Product SKU     |

Plus other catalogue fields supported by Shiprocket.

**Response:**

```json
{
  "id": 55,
  "name": "Widget",
  "sku": "W-1",
  "quantity": 10
}
```

-------------------------------------------------------------------------------------------------------

### Fetch particular product

```go
body, err := client.Product.Fetch("55", nil, nil)
```

**Parameters:**

| Name         | Type   | Description |
|--------------|--------|-------------|
| productID\*  | string | Product ID  |

**Response:**

```json
{
  "id": 55,
  "name": "Widget",
  "sku": "W-1",
  "quantity": 10
}
```

-------------------------------------------------------------------------------------------------------

### Sample import file

```go
body, err := client.Product.Sample(nil, nil)
```

**Parameters:**

None / optional query params.

**Response:** sample file metadata or download payload.

-------------------------------------------------------------------------------------------------------

### Import products (CSV)

```go
f, err := os.Open("products.csv")
if err != nil {
  log.Fatal(err)
}
defer f.Close()

body, err := client.Product.Import(requests.FileUploadParams{
  File: f,
}, nil)
```

**Parameters:**

| Name   | Type      | Description                 |
|--------|-----------|-----------------------------|
| File\* | \*os.File | CSV/Excel file handle       |
| Fields | map       | Extra multipart form fields |

**Response:**

```json
{
  "import_id": 77,
  "status": "completed",
  "errors": []
}
```

-------------------------------------------------------------------------------------------------------

**PN: \* indicates mandatory fields**

**For reference click [here](https://apidocs.shiprocket.in/)**
