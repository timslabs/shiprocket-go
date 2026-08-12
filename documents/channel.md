## Channels

### Fetch all channels

```go
body, err := client.Channel.All(nil, nil)
```

**Parameters:**

Optional list query params.

**Response:**

```json
{
  "data": [
    {
      "id": 1,
      "name": "Custom Channel",
      "base_channel_code": "CS"
    }
  ]
}
```

-------------------------------------------------------------------------------------------------------

### Fetch countries

```go
body, err := client.Channel.Countries(nil, nil)
```

**Parameters:**

None.

**Response:**

```json
{
  "data": [
    {
      "id": 1,
      "name": "India"
    }
  ]
}
```

-------------------------------------------------------------------------------------------------------

### Fetch country zones

```go
body, err := client.Channel.CountryZones("1", nil, nil)
```

**Parameters:**

| Name         | Type   | Description |
|--------------|--------|-------------|
| countryID\*  | string | Country ID  |

**Response:**

```json
{
  "data": []
}
```

-------------------------------------------------------------------------------------------------------

### Postcode details

```go
query := map[string]interface{}{
  "postcode": "110030",
}
body, err := client.Channel.PostcodeDetails(query, nil)
```

**Parameters:**

| Name        | Type   | Description |
|-------------|--------|-------------|
| postcode\*  | string | Pincode     |

**Response:**

```json
{
  "postcode": "110030",
  "city": "New Delhi",
  "state": "Delhi",
  "country": "India"
}
```

-------------------------------------------------------------------------------------------------------

**PN: \* indicates mandatory fields**

**For reference click [here](https://apidocs.shiprocket.in/)**
