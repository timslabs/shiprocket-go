package resources

import (
	"fmt"
	"net/url"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/requests"
)

type Ndr struct {
	Request *requests.Request
}

func (n *Ndr) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.NDR_ALL_URL)
	return n.Request.Get(path, queryParams, extraHeaders)
}

func (n *Ndr) Fetch(awb string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s/%s", constants.EXTERNAL_PREFIX, constants.NDR_URL, url.PathEscape(awb))
	return n.Request.Get(path, queryParams, extraHeaders)
}

func (n *Ndr) Action(awb string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s/%s/action", constants.EXTERNAL_PREFIX, constants.NDR_URL, url.PathEscape(awb))
	return n.Request.Post(path, data, extraHeaders)
}
