package shiprocket_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	shiprocket "github.com/timslabs/shiprocket-go"
	"github.com/timslabs/shiprocket-go/errors"
	"github.com/timslabs/shiprocket-go/requests"
)

func TestNewClient(t *testing.T) {
	client := shiprocket.NewClient("api@example.com", "secret")

	if client == nil {
		t.Fatal("client is nil")
	}
	if client.Request.Auth.Email != "api@example.com" {
		t.Fatalf("email = %q", client.Request.Auth.Email)
	}
	if client.Request.Auth.Password != "secret" {
		t.Fatalf("password = %q", client.Request.Auth.Password)
	}
	if client.Request.Auth.Token != "" {
		t.Fatalf("token should be empty before auth, got %q", client.Request.Auth.Token)
	}
}

func TestNewClientWithToken(t *testing.T) {
	client := shiprocket.NewClientWithToken("jwt-abc")
	if client.Request.Auth.Token != "jwt-abc" {
		t.Fatalf("token = %q", client.Request.Auth.Token)
	}
}

func TestNewClient_ResourcesInitialized(t *testing.T) {
	client := shiprocket.NewClient("a@b.c", "x")

	checks := []struct {
		name string
		v    any
	}{
		{"Auth", client.Auth},
		{"Order", client.Order},
		{"Courier", client.Courier},
		{"Shipment", client.Shipment},
		{"Pickup", client.Pickup},
		{"Product", client.Product},
		{"Inventory", client.Inventory},
		{"Listing", client.Listing},
		{"Channel", client.Channel},
		{"Account", client.Account},
		{"Ndr", client.Ndr},
		{"Import", client.Import},
		{"International", client.International},
		{"Warehouse", client.Warehouse},
	}
	for _, c := range checks {
		if c.v == nil {
			t.Fatalf("%s resource is nil", c.name)
		}
	}
}

func TestAddHeadersAndTimeoutAndUserAgent(t *testing.T) {
	client := shiprocket.NewClient("a@b.c", "x")
	client.AddHeaders(map[string]string{"X-Custom": "1"})
	client.AddHeaders(map[string]string{"X-Another": "2"})
	client.SetTimeout(30)
	client.SetUserAgent("MyApp/1.0")

	if client.Request.Headers["X-Custom"] != "1" {
		t.Fatalf("X-Custom = %q", client.Request.Headers["X-Custom"])
	}
	if client.Request.Headers["X-Another"] != "2" {
		t.Fatalf("X-Another = %q", client.Request.Headers["X-Another"])
	}
	if client.Request.HTTPClient.Timeout != 30*time.Second {
		t.Fatalf("timeout = %v", client.Request.HTTPClient.Timeout)
	}
	if client.Request.GetUserAgent() != "MyApp/1.0" {
		t.Fatalf("userAgent = %q", client.Request.GetUserAgent())
	}
}

func TestAuthenticateAndServiceability(t *testing.T) {
	mux := http.NewServeMux()
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)

	mux.HandleFunc("/v1/external/auth/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("login method = %s", r.Method)
		}
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"token": "test-jwt"})
	})

	mux.HandleFunc("/v1/external/courier/serviceability/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("serviceability method = %s", r.Method)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer test-jwt" {
			t.Fatalf("Authorization = %q", got)
		}
		q := r.URL.Query()
		if q.Get("pickup_postcode") != "110030" || q.Get("delivery_postcode") != "122001" {
			t.Fatalf("unexpected query: %s", r.URL.RawQuery)
		}
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"status": 200,
			"data": map[string]interface{}{
				"available_courier_companies": []interface{}{
					map[string]interface{}{"courier_name": "Demo", "rate": 45.5},
				},
			},
		})
	})

	client := shiprocket.NewClient("api@example.com", "secret")
	client.SetBaseURL(srv.URL)
	client.HTTPClient = srv.Client()

	out, err := client.Courier.Serviceability(map[string]interface{}{
		"pickup_postcode":   "110030",
		"delivery_postcode": "122001",
		"weight":            0.5,
		"cod":               0,
	}, nil)
	if err != nil {
		t.Fatalf("Serviceability: %v", err)
	}
	if out["status"] != float64(200) {
		t.Fatalf("status = %#v", out["status"])
	}
	if client.Token() != "test-jwt" {
		t.Fatalf("token after lazy auth = %q", client.Token())
	}
}

func TestAPIErrorOn4xx(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"message":"Unauthenticated"}`))
	}))
	t.Cleanup(srv.Close)

	client := shiprocket.NewClientWithToken("bad")
	client.SetBaseURL(srv.URL)
	client.HTTPClient = srv.Client()

	_, err := client.Courier.Serviceability(map[string]interface{}{
		"pickup_postcode":   "110030",
		"delivery_postcode": "122001",
		"weight":            0.5,
	}, nil)
	if err == nil {
		t.Fatal("expected error")
	}
	apiErr, ok := err.(*errors.BadRequestError)
	if !ok {
		t.Fatalf("err type = %T", err)
	}
	if apiErr.StatusCode != http.StatusUnauthorized {
		t.Fatalf("StatusCode = %d", apiErr.StatusCode)
	}
	if !strings.Contains(apiErr.Body, "Unauthenticated") {
		t.Fatalf("Body = %q", apiErr.Body)
	}
}

func TestAPIErrorOn5xx(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
		_, _ = w.Write([]byte(`{"message":"bad gateway"}`))
	}))
	t.Cleanup(srv.Close)

	client := shiprocket.NewClientWithToken("tok")
	client.SetBaseURL(srv.URL)
	client.HTTPClient = srv.Client()

	_, err := client.Account.WalletBalance(nil, nil)
	if err == nil {
		t.Fatal("expected error")
	}
	if _, ok := err.(*errors.ServerError); !ok {
		t.Fatalf("err type = %T", err)
	}
}

type recordedCall struct {
	Method string
	Path   string
}

type recorder struct {
	mu    sync.Mutex
	calls []recordedCall
}

func (r *recorder) add(method, path string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.calls = append(r.calls, recordedCall{Method: method, Path: path})
}

func (r *recorder) snapshot() []recordedCall {
	r.mu.Lock()
	defer r.mu.Unlock()
	out := make([]recordedCall, len(r.calls))
	copy(out, r.calls)
	return out
}

func newRecordingClient(t *testing.T) (*shiprocket.Client, *recorder) {
	t.Helper()
	rec := &recorder{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec.add(r.Method, r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	t.Cleanup(srv.Close)

	client := shiprocket.NewClientWithToken("token")
	client.SetBaseURL(srv.URL)
	client.HTTPClient = srv.Client()
	return client, rec
}

func mustCall(t *testing.T, fn func() (map[string]interface{}, error)) {
	t.Helper()
	out, err := fn()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out["ok"] != true {
		t.Fatalf("response = %#v", out)
	}
}

func assertCalls(t *testing.T, rec *recorder, want []recordedCall) {
	t.Helper()
	got := rec.snapshot()
	if len(got) != len(want) {
		t.Fatalf("call count = %d, want %d\ngot:  %#v\nwant: %#v", len(got), len(want), got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("call[%d] = %#v, want %#v", i, got[i], want[i])
		}
	}
}

func TestAccountEndpoints(t *testing.T) {
	client, rec := newRecordingClient(t)

	mustCall(t, func() (map[string]interface{}, error) { return client.Account.WalletBalance(nil, nil) })
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Account.Statement(map[string]interface{}{"from": "2026-01-01"}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) { return client.Account.Discrepancy(nil, nil) })

	assertCalls(t, rec, []recordedCall{
		{"GET", "/v1/external/account/details/wallet-balance"},
		{"GET", "/v1/external/account/details/statement"},
		{"GET", "/v1/external/billing/discrepancy"},
	})
}

func TestResourcePathCoverage(t *testing.T) {
	client, rec := newRecordingClient(t)

	mustCall(t, func() (map[string]interface{}, error) { return client.Auth.Logout(nil, nil) })

	mustCall(t, func() (map[string]interface{}, error) {
		return client.Order.CreateAdhoc(map[string]interface{}{"order_id": "1"}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Order.Create(map[string]interface{}{"order_id": "1"}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Order.CreateReturn(map[string]interface{}{"order_id": "1"}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Order.UpdateAdhoc(map[string]interface{}{"order_id": "1"}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Order.Cancel(map[string]interface{}{"ids": []interface{}{1}}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Order.All(map[string]interface{}{"page": 1}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) { return client.Order.Fetch("123", nil, nil) })
	mustCall(t, func() (map[string]interface{}, error) { return client.Order.ProcessingReturns(nil, nil) })
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Order.UpdateAddress(map[string]interface{}{}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Order.UpdatePickupAddress(map[string]interface{}{}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Order.Fulfill(map[string]interface{}{}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Order.Mapping(map[string]interface{}{}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) { return client.Order.Export(nil, nil) })
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Order.PrintInvoice(map[string]interface{}{}, nil)
	})

	mustCall(t, func() (map[string]interface{}, error) { return client.Courier.Serviceability(nil, nil) })
	mustCall(t, func() (map[string]interface{}, error) { return client.Courier.ListWithCounts(nil, nil) })
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Courier.AssignAwb(map[string]interface{}{}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Courier.GeneratePickup(map[string]interface{}{}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Courier.GenerateLabel(map[string]interface{}{}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Courier.GenerateManifest(map[string]interface{}{}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Courier.PrintManifest(map[string]interface{}{}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Courier.CancelShipment(map[string]interface{}{}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) { return client.Courier.TrackByAwb("AWB1", nil, nil) })
	mustCall(t, func() (map[string]interface{}, error) { return client.Courier.TrackByShipment("99", nil, nil) })
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Courier.TrackByOrder(map[string]interface{}{"order_id": 123}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Courier.TrackByAwbs(map[string]interface{}{"awbs": []interface{}{"A1"}}, nil)
	})

	mustCall(t, func() (map[string]interface{}, error) { return client.Shipment.All(nil, nil) })
	mustCall(t, func() (map[string]interface{}, error) { return client.Shipment.Fetch("99", nil, nil) })
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Shipment.CreateForward(map[string]interface{}{}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Shipment.CreateReturn(map[string]interface{}{}, nil)
	})

	mustCall(t, func() (map[string]interface{}, error) { return client.Pickup.All(nil, nil) })
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Pickup.Add(map[string]interface{}{}, nil)
	})

	mustCall(t, func() (map[string]interface{}, error) { return client.Product.All(nil, nil) })
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Product.Create(map[string]interface{}{}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) { return client.Product.Fetch("55", nil, nil) })
	mustCall(t, func() (map[string]interface{}, error) { return client.Product.Sample(nil, nil) })

	mustCall(t, func() (map[string]interface{}, error) { return client.Inventory.All(nil, nil) })
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Inventory.Update("55", map[string]interface{}{"quantity": 10}, nil)
	})

	mustCall(t, func() (map[string]interface{}, error) { return client.Listing.All(nil, nil) })
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Listing.Link(map[string]interface{}{}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) { return client.Listing.ExportMapped(nil, nil) })
	mustCall(t, func() (map[string]interface{}, error) { return client.Listing.ExportUnmapped(nil, nil) })
	mustCall(t, func() (map[string]interface{}, error) { return client.Listing.Sample(nil, nil) })

	mustCall(t, func() (map[string]interface{}, error) { return client.Channel.All(nil, nil) })
	mustCall(t, func() (map[string]interface{}, error) { return client.Channel.Countries(nil, nil) })
	mustCall(t, func() (map[string]interface{}, error) { return client.Channel.CountryZones("1", nil, nil) })
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Channel.PostcodeDetails(map[string]interface{}{"postcode": "110001"}, nil)
	})

	mustCall(t, func() (map[string]interface{}, error) { return client.Account.WalletBalance(nil, nil) })
	mustCall(t, func() (map[string]interface{}, error) { return client.Account.Statement(nil, nil) })
	mustCall(t, func() (map[string]interface{}, error) { return client.Account.Discrepancy(nil, nil) })

	mustCall(t, func() (map[string]interface{}, error) { return client.Ndr.All(nil, nil) })
	mustCall(t, func() (map[string]interface{}, error) { return client.Ndr.Fetch("AWB1", nil, nil) })
	mustCall(t, func() (map[string]interface{}, error) {
		return client.Ndr.Action("AWB1", map[string]interface{}{"action": "re-attempt"}, nil)
	})

	mustCall(t, func() (map[string]interface{}, error) { return client.Import.Check("77", nil, nil) })

	mustCall(t, func() (map[string]interface{}, error) {
		return client.International.CreateAdhoc(map[string]interface{}{}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) {
		return client.International.UpdateAdhoc(map[string]interface{}{}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) { return client.International.Track(nil, nil) })
	mustCall(t, func() (map[string]interface{}, error) { return client.International.Serviceability(nil, nil) })
	mustCall(t, func() (map[string]interface{}, error) {
		return client.International.AssignAwb(map[string]interface{}{}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) {
		return client.International.GenerateManifest(map[string]interface{}{}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) {
		return client.International.CreateForwardShipment(map[string]interface{}{}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) {
		return client.International.AddBankDetails(map[string]interface{}{}, nil)
	})
	mustCall(t, func() (map[string]interface{}, error) {
		return client.International.Kyc(map[string]interface{}{}, nil)
	})

	mustCall(t, func() (map[string]interface{}, error) {
		return client.Warehouse.SrfServiceability(map[string]interface{}{"postcode": "110030"}, nil)
	})

	assertCalls(t, rec, []recordedCall{
		{"POST", "/v1/external/auth/logout"},
		{"POST", "/v1/external/orders/create/adhoc"},
		{"POST", "/v1/external/orders/create"},
		{"POST", "/v1/external/orders/create/return"},
		{"POST", "/v1/external/orders/update/adhoc"},
		{"POST", "/v1/external/orders/cancel"},
		{"GET", "/v1/external/orders"},
		{"GET", "/v1/external/orders/show/123"},
		{"GET", "/v1/external/orders/processing/return"},
		{"POST", "/v1/external/orders/address/update"},
		{"PATCH", "/v1/external/orders/address/pickup"},
		{"PATCH", "/v1/external/orders/fulfill"},
		{"PATCH", "/v1/external/orders/mapping"},
		{"POST", "/v1/external/orders/export"},
		{"POST", "/v1/external/orders/print/invoice"},
		{"GET", "/v1/external/courier/serviceability/"},
		{"GET", "/v1/external/courier/courierListWithCounts"},
		{"POST", "/v1/external/courier/assign/awb"},
		{"POST", "/v1/external/courier/generate/pickup"},
		{"POST", "/v1/external/courier/generate/label"},
		{"POST", "/v1/external/manifests/generate"},
		{"POST", "/v1/external/manifests/print"},
		{"POST", "/v1/external/orders/cancel/shipment/awbs"},
		{"GET", "/v1/external/courier/track/awb/AWB1"},
		{"GET", "/v1/external/courier/track/shipment/99"},
		{"GET", "/v1/external/courier/track"},
		{"POST", "/v1/external/courier/track/awbs"},
		{"GET", "/v1/external/shipments"},
		{"GET", "/v1/external/shipments/99"},
		{"POST", "/v1/external/shipments/create/forward-shipment"},
		{"POST", "/v1/external/shipments/create/return-shipment"},
		{"GET", "/v1/external/settings/company/pickup"},
		{"POST", "/v1/external/settings/company/addpickup"},
		{"GET", "/v1/external/products"},
		{"POST", "/v1/external/products"},
		{"GET", "/v1/external/products/show/55"},
		{"GET", "/v1/external/products/sample"},
		{"GET", "/v1/external/inventory"},
		{"PUT", "/v1/external/inventory/55/update"},
		{"GET", "/v1/external/listings"},
		{"POST", "/v1/external/listings/link"},
		{"GET", "/v1/external/listings/export/mapped"},
		{"GET", "/v1/external/listings/export/unmapped"},
		{"GET", "/v1/external/listings/sample"},
		{"GET", "/v1/external/channels"},
		{"GET", "/v1/external/countries"},
		{"GET", "/v1/external/countries/show/1"},
		{"GET", "/v1/external/open/postcode/details"},
		{"GET", "/v1/external/account/details/wallet-balance"},
		{"GET", "/v1/external/account/details/statement"},
		{"GET", "/v1/external/billing/discrepancy"},
		{"GET", "/v1/external/ndr/all"},
		{"GET", "/v1/external/ndr/AWB1"},
		{"POST", "/v1/external/ndr/AWB1/action"},
		{"GET", "/v1/external/errors/77/check"},
		{"POST", "/v1/external/international/orders/create/adhoc"},
		{"POST", "/v1/external/international/orders/update/adhoc"},
		{"GET", "/v1/external/international/orders/track"},
		{"GET", "/v1/external/international/courier/serviceability"},
		{"POST", "/v1/external/international/courier/assign/awb"},
		{"POST", "/v1/external/international/manifests/generate"},
		{"POST", "/v1/external/international/shipments/create/forward-shipment"},
		{"POST", "/v1/external/international/settings/add-bank-details"},
		{"POST", "/v1/external/international/settings/international_kyc"},
		{"GET", "/v1/warehouse/srf-serviceability"},
	})
}

func TestOrderImportMultipart(t *testing.T) {
	var gotCT string
	var gotAuth string
	var body []byte

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/external/orders/import" {
			t.Fatalf("path = %s", r.URL.Path)
		}
		gotCT = r.Header.Get("Content-Type")
		gotAuth = r.Header.Get("Authorization")
		var err error
		body, err = io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read body: %v", err)
		}
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"import_id": 1})
	}))
	t.Cleanup(srv.Close)

	tmp, err := os.CreateTemp("", "orders-*.csv")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Remove(tmp.Name()) })
	if _, err := tmp.WriteString("order_id,sku\n1,ABC\n"); err != nil {
		t.Fatal(err)
	}
	if _, err := tmp.Seek(0, 0); err != nil {
		t.Fatal(err)
	}

	client := shiprocket.NewClientWithToken("token")
	client.SetBaseURL(srv.URL)
	client.HTTPClient = srv.Client()

	out, err := client.Order.Import(requests.FileUploadParams{
		File:   tmp,
		Fields: map[string]string{"type": "orders"},
	}, nil)
	if err != nil {
		t.Fatalf("Import: %v", err)
	}
	if out["import_id"] != float64(1) {
		t.Fatalf("import_id = %#v", out["import_id"])
	}
	if gotAuth != "Bearer token" {
		t.Fatalf("Authorization = %q", gotAuth)
	}
	if !strings.HasPrefix(gotCT, "multipart/form-data") {
		t.Fatalf("Content-Type = %q", gotCT)
	}
	if !strings.Contains(string(body), "order_id") {
		t.Fatalf("multipart body missing file content: %q", string(body))
	}
}

func TestAuthLoginUnauthenticated(t *testing.T) {
	var gotAuth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"token": "new-jwt"})
	}))
	t.Cleanup(srv.Close)

	client := shiprocket.NewClient("a@b.c", "x")
	client.SetBaseURL(srv.URL)
	client.HTTPClient = srv.Client()

	out, err := client.Auth.Login(map[string]interface{}{
		"email":    "a@b.c",
		"password": "x",
	}, nil)
	if err != nil {
		t.Fatalf("Login: %v", err)
	}
	if out["token"] != "new-jwt" {
		t.Fatalf("token = %#v", out["token"])
	}
	if gotAuth != "" {
		t.Fatalf("login should not send Authorization, got %q", gotAuth)
	}
}
