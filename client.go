package shiprocket

import (
	"net/http"
	"time"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/requests"
	"github.com/timslabs/shiprocket-go/resources"
)

// Client provides helper methods to call Shiprocket's External APIs.
type Client struct {
	*requests.Request
	Auth          *resources.Auth
	Order         *resources.Order
	Courier       *resources.Courier
	Shipment      *resources.Shipment
	Pickup        *resources.Pickup
	Product       *resources.Product
	Inventory     *resources.Inventory
	Listing       *resources.Listing
	Channel       *resources.Channel
	Account       *resources.Account
	Ndr           *resources.Ndr
	Import        *resources.Import
	International *resources.International
	Warehouse     *resources.Warehouse
}

func createClientFromRequest(request *requests.Request) *Client {
	auth := resources.Auth{Request: request}
	order := resources.Order{Request: request}
	courier := resources.Courier{Request: request}
	shipment := resources.Shipment{Request: request}
	pickup := resources.Pickup{Request: request}
	product := resources.Product{Request: request}
	inventory := resources.Inventory{Request: request}
	listing := resources.Listing{Request: request}
	channel := resources.Channel{Request: request}
	account := resources.Account{Request: request}
	ndr := resources.Ndr{Request: request}
	imp := resources.Import{Request: request}
	international := resources.International{Request: request}
	warehouse := resources.Warehouse{Request: request}

	return &Client{
		Request:       request,
		Auth:          &auth,
		Order:         &order,
		Courier:       &courier,
		Shipment:      &shipment,
		Pickup:        &pickup,
		Product:       &product,
		Inventory:     &inventory,
		Listing:       &listing,
		Channel:       &channel,
		Account:       &account,
		Ndr:           &ndr,
		Import:        &imp,
		International: &international,
		Warehouse:     &warehouse,
	}
}

// NewClient creates a Shiprocket client using API user email/password.
// The JWT is obtained lazily on the first authenticated request (or via Authenticate).
func NewClient(email, password string) *Client {
	request := &requests.Request{
		Auth: requests.Auth{
			Email:    email,
			Password: password,
		},
		HTTPClient: &http.Client{Timeout: requests.TIMEOUT * time.Second},
		Version:    getVersion(),
		SDKName:    getSDKName(),
		BaseURL:    constants.BASE_URL,
		Headers:    make(map[string]string),
	}
	return createClientFromRequest(request)
}

// NewClientWithToken creates a client that uses an existing JWT bearer token.
func NewClientWithToken(token string) *Client {
	request := &requests.Request{
		Auth:       requests.Auth{Token: token},
		HTTPClient: &http.Client{Timeout: requests.TIMEOUT * time.Second},
		Version:    getVersion(),
		SDKName:    getSDKName(),
		BaseURL:    constants.BASE_URL,
		Headers:    make(map[string]string),
	}
	return createClientFromRequest(request)
}

// Authenticate obtains a JWT using the credentials passed to NewClient.
func (client *Client) Authenticate() error {
	return client.Request.Authenticate()
}

// AddHeaders adds additional headers to all subsequent requests.
func (client *Client) AddHeaders(headers map[string]string) {
	client.Request.AddHeaders(headers)
}

// SetTimeout sets the HTTP timeout in seconds.
func (client *Client) SetTimeout(timeout int16) {
	client.Request.SetTimeout(timeout)
}

// SetUserAgent sets a custom User-Agent prefix.
func (client *Client) SetUserAgent(userAgent string) {
	client.Request.SetUserAgent(userAgent)
}

// SetBaseURL overrides the API host (useful in tests).
func (client *Client) SetBaseURL(baseURL string) {
	client.Request.BaseURL = baseURL
}

func getVersion() string {
	return SDKVersion
}

func getSDKName() string {
	return SDKName
}
