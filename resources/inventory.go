package resources

import (
	"fmt"
	"net/url"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/requests"
)

type Inventory struct {
	Request *requests.Request
}

func (i *Inventory) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.INVENTORY_URL)
	return i.Request.Get(path, queryParams, extraHeaders)
}

func (i *Inventory) Update(productID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s/%s/update", constants.EXTERNAL_PREFIX, constants.INVENTORY_URL, url.PathEscape(productID))
	return i.Request.Put(path, data, extraHeaders)
}
