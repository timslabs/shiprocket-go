package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"

	shiprocket "github.com/timslabs/shiprocket-go"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	Client *shiprocket.Client
)

const TestToken = "fake_token"

func testSetup() func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	Client = shiprocket.NewClientWithToken(TestToken)
	Client.Request.BaseURL = server.URL
	Client.Request.HTTPClient = server.Client()
	return func() {
		server.Close()
	}
}

func testdataDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), "..", "testdata")
}

func getFixture(name string) string {
	b, err := os.ReadFile(filepath.Join(testdataDir(), name+".json"))
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}

func jsonCompare(b1, b2 []byte) (bool, error) {
	var o1, o2 interface{}
	if err := json.Unmarshal(b1, &o1); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b2, &o2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(o1, o2), nil
}

// TestResponse asserts the API response JSON matches the fixture JSON.
func TestResponse(jsonBody []byte, fixture []byte, t *testing.T) {
	t.Helper()
	ok, err := jsonCompare(jsonBody, fixture)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatalf("response mismatch\ngot:  %s\nwant: %s", string(jsonBody), string(fixture))
	}
}

// StartMockServer serves a testdata JSON fixture at url and returns teardown + fixture body.
func StartMockServer(url string, fixtureName string) (func(), string) {
	teardown := testSetup()
	fixture := getFixture(fixtureName)
	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture)
	})
	return teardown, fixture
}
