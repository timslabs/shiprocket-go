package resources

import (
	"fmt"
	"net/url"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/requests"
)

type Courier struct {
	Request *requests.Request
}

func (c *Courier) Serviceability(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.COURIER_SERVICEABILITY_URL)
	return c.Request.Get(path, queryParams, extraHeaders)
}

func (c *Courier) ListWithCounts(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.COURIER_LIST_COUNTS_URL)
	return c.Request.Get(path, queryParams, extraHeaders)
}

func (c *Courier) AssignAwb(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.COURIER_ASSIGN_AWB_URL)
	return c.Request.Post(path, data, extraHeaders)
}

func (c *Courier) GeneratePickup(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.COURIER_GENERATE_PICKUP_URL)
	return c.Request.Post(path, data, extraHeaders)
}

func (c *Courier) GenerateLabel(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.COURIER_GENERATE_LABEL_URL)
	return c.Request.Post(path, data, extraHeaders)
}

func (c *Courier) GenerateManifest(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.MANIFESTS_GENERATE_URL)
	return c.Request.Post(path, data, extraHeaders)
}

func (c *Courier) PrintManifest(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.MANIFESTS_PRINT_URL)
	return c.Request.Post(path, data, extraHeaders)
}

func (c *Courier) CancelShipment(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.CANCEL_SHIPMENT_AWBS_URL)
	return c.Request.Post(path, data, extraHeaders)
}

func (c *Courier) TrackByAwb(awbCode string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s/%s", constants.EXTERNAL_PREFIX, constants.COURIER_TRACK_AWB_URL, url.PathEscape(awbCode))
	return c.Request.Get(path, queryParams, extraHeaders)
}

func (c *Courier) TrackByShipment(shipmentID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s/%s", constants.EXTERNAL_PREFIX, constants.COURIER_TRACK_SHIPMENT_URL, url.PathEscape(shipmentID))
	return c.Request.Get(path, queryParams, extraHeaders)
}

func (c *Courier) TrackByOrder(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.COURIER_TRACK_URL)
	return c.Request.Get(path, queryParams, extraHeaders)
}

func (c *Courier) TrackByAwbs(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.COURIER_TRACK_AWBS_URL)
	return c.Request.Post(path, data, extraHeaders)
}
