# shiprocket-go

Go SDK for the [Shiprocket External API](https://apidocs.shiprocket.in/) (v1).

## Requirements

- Go **1.21+**

```bash
go get github.com/timslabs/shiprocket-go
```

## Authentication

Create an API user in the Shiprocket panel: **Settings → API → Configure → Create an API User** (not your panel login).

Authenticate with email and password to receive a JWT (valid for **10 days** / 240 hours). The SDK attaches it as `Authorization: Bearer {token}` on subsequent requests.

```go
import (
	shiprocket "github.com/timslabs/shiprocket-go"
)

client := shiprocket.NewClient(
	os.Getenv("SHIPROCKET_EMAIL"),
	os.Getenv("SHIPROCKET_PASSWORD"),
)

// JWT is fetched automatically on the first authenticated call.
// Or authenticate explicitly:
if err := client.Authenticate(); err != nil {
	log.Fatal(err)
}
```

Or manage the token yourself:

```go
client := shiprocket.NewClientWithToken(os.Getenv("SHIPROCKET_TOKEN"))
```

## Quick start

```go
package main

import (
	"fmt"
	"log"
	"os"

	shiprocket "github.com/timslabs/shiprocket-go"
)

func main() {
	client := shiprocket.NewClient(
		os.Getenv("SHIPROCKET_EMAIL"),
		os.Getenv("SHIPROCKET_PASSWORD"),
	)

	rates, err := client.Courier.Serviceability(map[string]interface{}{
		"pickup_postcode":   "110030",
		"delivery_postcode": "122001",
		"weight":            0.5,
		"cod":               0,
	}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", rates)

	order, err := client.Order.CreateAdhoc(map[string]interface{}{
		"order_id":               "ORD-1001",
		"order_date":             "2026-08-13",
		"pickup_location":        "Primary",
		"billing_customer_name":  "Jane",
		"billing_last_name":      "Doe",
		"billing_address":        "221B Baker Street",
		"billing_city":           "Mumbai",
		"billing_pincode":        "400001",
		"billing_state":          "Maharashtra",
		"billing_country":        "India",
		"billing_email":          "jane@example.com",
		"billing_phone":          "9999999999",
		"shipping_is_billing":    true,
		"payment_method":         "Prepaid",
		"sub_total":              499,
		"length":                 10,
		"breadth":                10,
		"height":                 10,
		"weight":                 0.5,
		"order_items": []map[string]interface{}{
			{
				"name":          "Widget",
				"sku":           "W-1",
				"units":         1,
				"selling_price": 499,
			},
		},
	}, nil)
	if err != nil {
		log.Fatal(err)
	}

	awb, err := client.Courier.AssignAwb(map[string]interface{}{
		"shipment_id": order["shipment_id"],
	}, nil)
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Courier.GeneratePickup(map[string]interface{}{
		"shipment_id": []interface{}{order["shipment_id"]},
	}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("AWB response: %v\n", awb)
}
```

Successful responses are `map[string]interface{}` (decoded JSON). Failed HTTP responses return typed errors from the `errors` package (`BadRequestError`, `ServerError`, `AuthError`).

All resource methods accept an optional `extraHeaders map[string]string` as the last argument.

## Configuration

```go
client := shiprocket.NewClient(email, password)
client.SetTimeout(60)                 // seconds
client.SetUserAgent("MyApp/1.0")
client.AddHeaders(map[string]string{
	"X-Custom": "value",
})
client.SetBaseURL("https://apiv2.shiprocket.in") // default
```

Base URL: `https://apiv2.shiprocket.in`  
External API prefix: `/v1/external`

## API groups

| Resource | Coverage |
|----------|----------|
| [`Auth`](documents/auth.md) | Login, logout |
| [`Order`](documents/order.md) | Create / update / cancel / address / fulfill / mapping / import / export / invoice / returns |
| [`Courier`](documents/courier.md) | Serviceability, courier list, AWB, pickup, label, manifest, track (AWB / shipment / order / bulk) |
| [`Shipment`](documents/shipment.md) | List / show / create forward & return |
| [`Pickup`](documents/pickup.md) | List / add pickup locations |
| [`Product`](documents/product.md) | Catalogue + bulk import / sample |
| [`Inventory`](documents/inventory.md) | List / update |
| [`Listing`](documents/listing.md) | Channel catalog mappings (list / link / import / export) |
| [`Channel`](documents/channel.md) | Channels, countries, zones, postcode details |
| [`Account`](documents/account.md) | Wallet balance, statement, billing discrepancy |
| [`Ndr`](documents/ndr.md) | NDR list / details / action |
| [`Import`](documents/import.md) | Bulk import error / status check |
| [`International`](documents/international.md) | International orders, couriers, manifest, KYC, bank details |
| [`Warehouse`](documents/warehouse.md) | Warehouse SRF serviceability (`/v1/warehouse`, not under `/v1/external`) |

## Package layout

```text
shiprocket-go/
  client.go          # NewClient, resource wiring
  version.go
  constants/         # URLs, headers, status codes
  errors/            # BadRequestError, ServerError, AuthError
  requests/          # HTTP Get/Post/Put/Patch/Delete/File
  resources/         # API resource methods
  documents/         # Per-resource usage examples
  testdata/          # JSON fixtures for tests
  utils/             # Test helpers
```

## Tests

```bash
go test ./...
go test ./... -v
go test ./resources -run TestAccountWalletBalance -v
```

## Related

- [Shiprocket API docs](https://apidocs.shiprocket.in/)
- PHP SDK: [`tims/shiprocket-php-sdk`](https://github.com/timslabs/shiprocket-php-sdk)
- Laravel package: [`tims/laravel-shiprocket`](https://github.com/timslabs/laravel-shiprocket)

## License

MIT
