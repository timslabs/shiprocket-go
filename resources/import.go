package resources

import (
	"fmt"
	"net/url"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/requests"
)

type Import struct {
	Request *requests.Request
}

func (i *Import) Check(importID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s/%s/check", constants.EXTERNAL_PREFIX, constants.ERRORS_CHECK_URL, url.PathEscape(importID))
	return i.Request.Get(path, queryParams, extraHeaders)
}
