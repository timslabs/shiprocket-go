package resources

import (
	"fmt"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/requests"
)

type Pickup struct {
	Request *requests.Request
}

func (p *Pickup) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.PICKUP_LIST_URL)
	return p.Request.Get(path, queryParams, extraHeaders)
}

func (p *Pickup) Add(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.PICKUP_ADD_URL)
	return p.Request.Post(path, data, extraHeaders)
}
