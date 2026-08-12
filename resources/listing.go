package resources

import (
	"fmt"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/requests"
)

type Listing struct {
	Request *requests.Request
}

func (l *Listing) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.LISTINGS_URL)
	return l.Request.Get(path, queryParams, extraHeaders)
}

func (l *Listing) Link(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.LISTINGS_LINK_URL)
	return l.Request.Post(path, data, extraHeaders)
}

func (l *Listing) Import(params requests.FileUploadParams, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.LISTINGS_IMPORT_URL)
	return l.Request.File(path, params, extraHeaders)
}

func (l *Listing) ExportMapped(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.LISTINGS_EXPORT_MAPPED_URL)
	return l.Request.Get(path, queryParams, extraHeaders)
}

func (l *Listing) ExportUnmapped(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.LISTINGS_EXPORT_UNMAPPED_URL)
	return l.Request.Get(path, queryParams, extraHeaders)
}

func (l *Listing) Sample(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.LISTINGS_SAMPLE_URL)
	return l.Request.Get(path, queryParams, extraHeaders)
}
