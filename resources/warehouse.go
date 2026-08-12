package resources

import (
	"fmt"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/requests"
)

type Warehouse struct {
	Request *requests.Request
}

func (w *Warehouse) SrfServiceability(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.WAREHOUSE_PREFIX, constants.WAREHOUSE_SRF_SERVICEABILITY)
	return w.Request.Get(path, queryParams, extraHeaders)
}
