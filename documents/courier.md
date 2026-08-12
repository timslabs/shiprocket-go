# Courier

```go
client := shiprocket.NewClient(email, password)

data, err := client.Courier.Serviceability(map[string]interface{}{
	"pickup_postcode":   "110030",
	"delivery_postcode": "122001",
	"weight":            0.5,
	"cod":               0,
}, nil)

data, err = client.Courier.ListWithCounts(nil, nil)
data, err = client.Courier.AssignAwb(map[string]interface{}{"shipment_id": 123}, nil)
data, err = client.Courier.GeneratePickup(map[string]interface{}{"shipment_id": []interface{}{123}}, nil)
data, err = client.Courier.GenerateLabel(map[string]interface{}{"shipment_id": []interface{}{123}}, nil)
data, err = client.Courier.GenerateManifest(map[string]interface{}{"shipment_id": []interface{}{123}}, nil)
data, err = client.Courier.PrintManifest(map[string]interface{}{"shipment_id": []interface{}{123}}, nil)
data, err = client.Courier.CancelShipment(map[string]interface{}{"awbs": []interface{}{"AWB"}}, nil)
data, err = client.Courier.TrackByAwb("AWB", nil, nil)
data, err = client.Courier.TrackByShipment("123", nil, nil)
data, err = client.Courier.TrackByOrder(map[string]interface{}{"order_id": "ORD-1"}, nil)
data, err = client.Courier.TrackByAwbs(map[string]interface{}{"awbs": []interface{}{"AWB"}}, nil)
```
