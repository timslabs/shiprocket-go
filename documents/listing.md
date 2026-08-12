## Listings

### Fetch all listings

```go
body, err := client.Listing.All(nil, nil)
```

**Parameters:**

Optional list query params.

**Response:**

```json
{
  "data": [
    {
      "channel_sku": "CH-1",
      "sku": "W-1",
      "status": "mapped"
    }
  ]
}
```

-------------------------------------------------------------------------------------------------------

### Link listing

```go
data := map[string]interface{}{
  "channel_sku": "CH-1",
  "sku":         "W-1",
}
body, err := client.Listing.Link(data, nil)
```

**Parameters:**

| Name          | Type   | Description           |
|---------------|--------|-----------------------|
| channel_sku\* | string | Channel SKU           |
| sku\*         | string | Catalogue SKU to link |

**Response:**

```json
{
  "success": true,
  "message": "Listing linked successfully"
}
```

-------------------------------------------------------------------------------------------------------

### Import listings (CSV)

```go
f, err := os.Open("listings.csv")
if err != nil {
  log.Fatal(err)
}
defer f.Close()

body, err := client.Listing.Import(requests.FileUploadParams{
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

### Export mapped listings

```go
body, err := client.Listing.ExportMapped(nil, nil)
```

**Parameters:**

Optional export query params.

**Response:** export payload / file metadata.

-------------------------------------------------------------------------------------------------------

### Export unmapped listings

```go
body, err := client.Listing.ExportUnmapped(nil, nil)
```

**Parameters:**

Optional export query params.

**Response:** export payload / file metadata.

-------------------------------------------------------------------------------------------------------

### Sample listings file

```go
body, err := client.Listing.Sample(nil, nil)
```

**Parameters:**

None / optional query params.

**Response:** sample file metadata or download payload.

-------------------------------------------------------------------------------------------------------

**PN: \* indicates mandatory fields**

**For reference click [here](https://apidocs.shiprocket.in/)**
