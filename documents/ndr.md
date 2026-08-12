## NDR

### Fetch all NDR

```go
body, err := client.Ndr.All(nil, nil)
```

**Parameters:**

Optional list query params.

**Response:**

```json
{
  "data": [
    {
      "awb_code": "AWB1",
      "status": "NDR",
      "reason": "Customer unavailable"
    }
  ]
}
```

-------------------------------------------------------------------------------------------------------

### Fetch particular NDR

```go
body, err := client.Ndr.Fetch("AWB1", nil, nil)
```

**Parameters:**

| Name  | Type   | Description |
|-------|--------|-------------|
| awb\* | string | AWB code    |

**Response:**

```json
{
  "awb_code": "AWB1",
  "status": "NDR",
  "reason": "Customer unavailable"
}
```

-------------------------------------------------------------------------------------------------------

### NDR action

```go
data := map[string]interface{}{
  "action": "re-attempt",
}
body, err := client.Ndr.Action("AWB1", data, nil)
```

**Parameters:**

| Name     | Type   | Description                                      |
|----------|--------|--------------------------------------------------|
| awb\*    | string | AWB code                                         |
| action\* | string | Action to apply (e.g. `re-attempt`, `return`)    |

Plus action-specific fields from Shiprocket docs.

**Response:**

```json
{
  "message": "Action applied successfully"
}
```

-------------------------------------------------------------------------------------------------------

**PN: \* indicates mandatory fields**

**For reference click [here](https://apidocs.shiprocket.in/)**
