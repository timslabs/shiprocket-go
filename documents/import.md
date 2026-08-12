## Import

### Check import status / errors

```go
body, err := client.Import.Check("77", nil, nil)
```

**Parameters:**

| Name        | Type   | Description  |
|-------------|--------|--------------|
| importID\*  | string | Import job ID|

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
