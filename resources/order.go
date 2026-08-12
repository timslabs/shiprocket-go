package resources

import (
	"fmt"
	"net/url"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/requests"
)

type Order struct {
	Request *requests.Request
}

func (o *Order) CreateAdhoc(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_CREATE_ADHOC_URL)
	return o.Request.Post(path, data, extraHeaders)
}

func (o *Order) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_CREATE_URL)
	return o.Request.Post(path, data, extraHeaders)
}

func (o *Order) CreateReturn(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_CREATE_RETURN_URL)
	return o.Request.Post(path, data, extraHeaders)
}

func (o *Order) UpdateAdhoc(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_UPDATE_ADHOC_URL)
	return o.Request.Post(path, data, extraHeaders)
}

func (o *Order) Cancel(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_CANCEL_URL)
	return o.Request.Post(path, data, extraHeaders)
}

func (o *Order) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_URL)
	return o.Request.Get(path, queryParams, extraHeaders)
}

func (o *Order) Fetch(orderID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s/%s", constants.EXTERNAL_PREFIX, constants.ORDERS_SHOW_URL, url.PathEscape(orderID))
	return o.Request.Get(path, queryParams, extraHeaders)
}

func (o *Order) ProcessingReturns(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_PROCESSING_RETURN_URL)
	return o.Request.Get(path, queryParams, extraHeaders)
}

func (o *Order) UpdateAddress(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_ADDRESS_UPDATE_URL)
	return o.Request.Post(path, data, extraHeaders)
}

func (o *Order) UpdatePickupAddress(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_ADDRESS_PICKUP_URL)
	return o.Request.Patch(path, data, extraHeaders)
}

func (o *Order) Fulfill(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_FULFILL_URL)
	return o.Request.Patch(path, data, extraHeaders)
}

func (o *Order) Mapping(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_MAPPING_URL)
	return o.Request.Patch(path, data, extraHeaders)
}

func (o *Order) Export(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_EXPORT_URL)
	return o.Request.Post(path, data, extraHeaders)
}

func (o *Order) Import(params requests.FileUploadParams, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_IMPORT_URL)
	return o.Request.File(path, params, extraHeaders)
}

func (o *Order) PrintInvoice(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_PRINT_INVOICE_URL)
	return o.Request.Post(path, data, extraHeaders)
}
