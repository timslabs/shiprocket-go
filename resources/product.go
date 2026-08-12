package resources

import (
	"fmt"
	"net/url"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/requests"
)

type Product struct {
	Request *requests.Request
}

func (p *Product) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.PRODUCTS_URL)
	return p.Request.Get(path, queryParams, extraHeaders)
}

func (p *Product) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.PRODUCTS_URL)
	return p.Request.Post(path, data, extraHeaders)
}

func (p *Product) Fetch(productID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s/%s", constants.EXTERNAL_PREFIX, constants.PRODUCTS_SHOW_URL, url.PathEscape(productID))
	return p.Request.Get(path, queryParams, extraHeaders)
}

func (p *Product) Sample(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.PRODUCTS_SAMPLE_URL)
	return p.Request.Get(path, queryParams, extraHeaders)
}

func (p *Product) Import(params requests.FileUploadParams, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.PRODUCTS_IMPORT_URL)
	return p.Request.File(path, params, extraHeaders)
}
