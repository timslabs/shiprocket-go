package resources

import (
	"fmt"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/requests"
)

type International struct {
	Request *requests.Request
}

func (i *International) CreateAdhoc(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.INTERNATIONAL_ORDERS_CREATE_ADHOC_URL)
	return i.Request.Post(path, data, extraHeaders)
}

func (i *International) UpdateAdhoc(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.INTERNATIONAL_ORDERS_UPDATE_ADHOC_URL)
	return i.Request.Post(path, data, extraHeaders)
}

func (i *International) Track(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.INTERNATIONAL_ORDERS_TRACK_URL)
	return i.Request.Get(path, queryParams, extraHeaders)
}

func (i *International) Serviceability(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.INTERNATIONAL_COURIER_SERVICEABILITY)
	return i.Request.Get(path, queryParams, extraHeaders)
}

func (i *International) AssignAwb(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.INTERNATIONAL_COURIER_ASSIGN_AWB)
	return i.Request.Post(path, data, extraHeaders)
}

func (i *International) GenerateManifest(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.INTERNATIONAL_MANIFESTS_GENERATE)
	return i.Request.Post(path, data, extraHeaders)
}

func (i *International) CreateForwardShipment(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.INTERNATIONAL_SHIPMENTS_FORWARD)
	return i.Request.Post(path, data, extraHeaders)
}

func (i *International) AddBankDetails(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.INTERNATIONAL_ADD_BANK_DETAILS)
	return i.Request.Post(path, data, extraHeaders)
}

func (i *International) Kyc(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.INTERNATIONAL_KYC)
	return i.Request.Post(path, data, extraHeaders)
}
