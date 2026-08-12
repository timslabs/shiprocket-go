# Listing

```go
client := shiprocket.NewClient(email, password)

data, err := client.Listing.All(nil, nil)
data, err = client.Listing.Link(map[string]interface{}{ /* ... */ }, nil)
data, err = client.Listing.ExportMapped(nil, nil)
data, err = client.Listing.ExportUnmapped(nil, nil)
data, err = client.Listing.Sample(nil, nil)

f, _ := os.Open("listings.csv")
defer f.Close()
data, err = client.Listing.Import(requests.FileUploadParams{File: f}, nil)
```
