package resources

import (
	"fmt"
	"net/url"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/requests"
)

type Shipment struct {
	Request *requests.Request
}

func (s *Shipment) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.SHIPMENTS_URL)
	return s.Request.Get(path, queryParams, extraHeaders)
}

func (s *Shipment) Fetch(shipmentID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s/%s", constants.EXTERNAL_PREFIX, constants.SHIPMENTS_URL, url.PathEscape(shipmentID))
	return s.Request.Get(path, queryParams, extraHeaders)
}

func (s *Shipment) CreateForward(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.SHIPMENTS_CREATE_FORWARD_URL)
	return s.Request.Post(path, data, extraHeaders)
}

func (s *Shipment) CreateReturn(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.SHIPMENTS_CREATE_RETURN_URL)
	return s.Request.Post(path, data, extraHeaders)
}
