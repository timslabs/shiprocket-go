package requests

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/timslabs/shiprocket-go/errors"
)

func newTestRequest(t *testing.T, handler http.HandlerFunc) *Request {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	return &Request{
		Auth:       Auth{Token: "test-token"},
		HTTPClient: srv.Client(),
		Headers:    make(map[string]string),
		Version:    "1.0.1",
		SDKName:    "shiprocket-go",
		BaseURL:    srv.URL,
	}
}

func TestBuildURLWithParams(t *testing.T) {
	tests := []struct {
		name       string
		requestURL string
		data       map[string]interface{}
		expected   string
	}{
		{
			name:       "simple params",
			requestURL: "https://apiv2.shiprocket.in/v1/external/orders",
			data: map[string]interface{}{
				"page": 1,
				"per":  10,
			},
			expected: "https://apiv2.shiprocket.in/v1/external/orders?page=1&per=10",
		},
		{
			name:       "special characters",
			requestURL: "https://apiv2.shiprocket.in/v1/external/orders",
			data: map[string]interface{}{
				"q": "ABC & Co.",
			},
			expected: "https://apiv2.shiprocket.in/v1/external/orders?q=ABC+%26+Co.",
		},
		{
			name:       "string array",
			requestURL: "https://apiv2.shiprocket.in/v1/external/courier/track",
			data: map[string]interface{}{
				"awbs": []string{"A1", "A2"},
			},
			expected: "https://apiv2.shiprocket.in/v1/external/courier/track?awbs=A1&awbs=A2",
		},
		{
			name:       "empty map",
			requestURL: "https://apiv2.shiprocket.in/v1/external/orders",
			data:       map[string]interface{}{},
			expected:   "https://apiv2.shiprocket.in/v1/external/orders",
		},
		{
			name:       "nil values skipped",
			requestURL: "https://apiv2.shiprocket.in/v1/external/orders",
			data: map[string]interface{}{
				"page":  1,
				"extra": nil,
			},
			expected: "https://apiv2.shiprocket.in/v1/external/orders?page=1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildURLWithParams(tt.requestURL, tt.data)
			if got != tt.expected {
				t.Fatalf("got %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestGetUserAgentHeaderValue(t *testing.T) {
	tests := []struct {
		name     string
		request  Request
		expected string
	}{
		{
			name: "default",
			request: Request{
				SDKName: "shiprocket-go",
				Version: "1.0.1",
			},
			expected: "shiprocket-go/1.0.1",
		},
		{
			name: "custom",
			request: Request{
				SDKName:   "shiprocket-go",
				Version:   "1.0.1",
				userAgent: "MyApp/2.0",
			},
			expected: "MyApp/2.0 (shiprocket-go/1.0.1)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.request.getUserAgentHeaderValue()
			if got != tt.expected {
				t.Fatalf("got %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestAddHeadersAndOverwrite(t *testing.T) {
	req := &Request{Headers: map[string]string{"X-Existing": "old"}}
	req.AddHeaders(map[string]string{
		"X-Existing": "new",
		"X-Custom":   "1",
	})
	if req.Headers["X-Existing"] != "new" {
		t.Fatalf("X-Existing = %q", req.Headers["X-Existing"])
	}
	if req.Headers["X-Custom"] != "1" {
		t.Fatalf("X-Custom = %q", req.Headers["X-Custom"])
	}
}

func TestSetTimeoutAndUserAgent(t *testing.T) {
	req := &Request{}
	req.SetTimeout(15)
	req.SetUserAgent("App/1")

	if req.HTTPClient == nil || req.HTTPClient.Timeout != 15*time.Second {
		t.Fatalf("timeout = %v", req.HTTPClient)
	}
	if req.GetUserAgent() != "App/1" {
		t.Fatalf("userAgent = %q", req.GetUserAgent())
	}
}

func TestSetTokenAndToken(t *testing.T) {
	req := &Request{}
	req.SetToken("abc")
	if req.Token() != "abc" {
		t.Fatalf("token = %q", req.Token())
	}
}

func TestGetSendsBearerAndQuery(t *testing.T) {
	var gotMethod, gotPath, gotAuth, gotQuery string
	req := newTestRequest(t, func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		gotAuth = r.Header.Get("Authorization")
		gotQuery = r.URL.RawQuery
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"ok": true})
	})

	out, err := req.Get("/v1/external/orders", map[string]interface{}{"page": 2}, map[string]string{"X-Test": "1"})
	if err != nil {
		t.Fatal(err)
	}
	if out["ok"] != true {
		t.Fatalf("out = %#v", out)
	}
	if gotMethod != http.MethodGet || gotPath != "/v1/external/orders" {
		t.Fatalf("%s %s", gotMethod, gotPath)
	}
	if gotAuth != "Bearer test-token" {
		t.Fatalf("auth = %q", gotAuth)
	}
	if gotQuery != "page=2" {
		t.Fatalf("query = %q", gotQuery)
	}
}

func TestPostPutPatchDelete(t *testing.T) {
	var methods []string
	req := newTestRequest(t, func(w http.ResponseWriter, r *http.Request) {
		methods = append(methods, r.Method)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"ok": true})
	})

	if _, err := req.Post("/v1/external/orders", map[string]interface{}{"a": 1}, nil); err != nil {
		t.Fatal(err)
	}
	if _, err := req.Put("/v1/external/inventory/1/update", map[string]interface{}{"q": 1}, nil); err != nil {
		t.Fatal(err)
	}
	if _, err := req.Patch("/v1/external/orders/fulfill", map[string]interface{}{}, nil); err != nil {
		t.Fatal(err)
	}
	if _, err := req.Delete("/v1/external/orders", nil, nil); err != nil {
		t.Fatal(err)
	}

	want := []string{http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete}
	if len(methods) != len(want) {
		t.Fatalf("methods = %#v", methods)
	}
	for i := range want {
		if methods[i] != want[i] {
			t.Fatalf("methods[%d] = %s, want %s", i, methods[i], want[i])
		}
	}
}

func TestPostUnauthenticatedOmitsBearer(t *testing.T) {
	var gotAuth string
	req := newTestRequest(t, func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"token": "jwt"})
	})

	_, err := req.PostUnauthenticated("/v1/external/auth/login", map[string]interface{}{
		"email": "a@b.c", "password": "x",
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if gotAuth != "" {
		t.Fatalf("auth = %q", gotAuth)
	}
}

func TestAuthenticateStoresToken(t *testing.T) {
	req := newTestRequest(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/external/auth/login" {
			t.Fatalf("path = %s", r.URL.Path)
		}
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"token": "jwt-from-login"})
	})
	req.Auth = Auth{Email: "a@b.c", Password: "secret"}
	req.SetToken("")

	if err := req.Authenticate(); err != nil {
		t.Fatal(err)
	}
	if req.Token() != "jwt-from-login" {
		t.Fatalf("token = %q", req.Token())
	}
}

func TestAuthenticateRequiresCredentials(t *testing.T) {
	req := &Request{HTTPClient: http.DefaultClient, Headers: map[string]string{}}
	err := req.Authenticate()
	if _, ok := err.(*errors.AuthError); !ok {
		t.Fatalf("err = %T %v", err, err)
	}
}

func TestAuthenticateMissingToken(t *testing.T) {
	req := newTestRequest(t, func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"ok": true})
	})
	req.Auth = Auth{Email: "a@b.c", Password: "secret"}
	req.SetToken("")

	err := req.Authenticate()
	if _, ok := err.(*errors.AuthError); !ok {
		t.Fatalf("err = %T %v", err, err)
	}
}

func TestLazyAuthenticateOnGet(t *testing.T) {
	var paths []string
	req := newTestRequest(t, func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		if strings.HasSuffix(r.URL.Path, "/auth/login") {
			_ = json.NewEncoder(w).Encode(map[string]interface{}{"token": "lazy-jwt"})
			return
		}
		if r.Header.Get("Authorization") != "Bearer lazy-jwt" {
			t.Fatalf("auth = %q", r.Header.Get("Authorization"))
		}
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"ok": true})
	})
	req.Auth = Auth{Email: "a@b.c", Password: "secret"}
	req.SetToken("")

	if _, err := req.Get("/v1/external/orders", nil, nil); err != nil {
		t.Fatal(err)
	}
	if len(paths) != 2 || paths[0] != "/v1/external/auth/login" || paths[1] != "/v1/external/orders" {
		t.Fatalf("paths = %#v", paths)
	}
}

func TestGetWithoutAuthErrors(t *testing.T) {
	req := &Request{
		HTTPClient: http.DefaultClient,
		Headers:    map[string]string{},
		BaseURL:    "http://example.invalid",
	}
	_, err := req.Get("/v1/external/orders", nil, nil)
	if _, ok := err.(*errors.AuthError); !ok {
		t.Fatalf("err = %T %v", err, err)
	}
}

func TestBadRequestAndServerError(t *testing.T) {
	t.Run("4xx", func(t *testing.T) {
		req := newTestRequest(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"message":"bad"}`))
		})
		_, err := req.Get("/v1/external/orders", nil, nil)
		br, ok := err.(*errors.BadRequestError)
		if !ok {
			t.Fatalf("err = %T", err)
		}
		if br.StatusCode != http.StatusBadRequest {
			t.Fatalf("status = %d", br.StatusCode)
		}
	})

	t.Run("5xx", func(t *testing.T) {
		req := newTestRequest(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`{"message":"boom"}`))
		})
		_, err := req.Get("/v1/external/orders", nil, nil)
		if _, ok := err.(*errors.ServerError); !ok {
			t.Fatalf("err = %T", err)
		}
	})
}

func TestFileUpload(t *testing.T) {
	var gotCT string
	var body string
	req := newTestRequest(t, func(w http.ResponseWriter, r *http.Request) {
		gotCT = r.Header.Get("Content-Type")
		b, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read body: %v", err)
		}
		body = string(b)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"import_id": 9})
	})

	tmp, err := os.CreateTemp("", "import-*.csv")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Remove(tmp.Name()) })
	_, _ = tmp.WriteString("sku,qty\nA,1\n")
	_, _ = tmp.Seek(0, 0)

	out, err := req.File("/v1/external/orders/import", FileUploadParams{
		File:   tmp,
		Fields: map[string]string{"type": "orders"},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if out["import_id"] != float64(9) {
		t.Fatalf("out = %#v", out)
	}
	if !strings.HasPrefix(gotCT, "multipart/form-data") {
		t.Fatalf("ct = %q", gotCT)
	}
	if !strings.Contains(body, "sku,qty") || !strings.Contains(body, "type") {
		t.Fatalf("body = %q", body)
	}
}
