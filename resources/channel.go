package resources

import (
	"fmt"
	"net/url"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/requests"
)

type Channel struct {
	Request *requests.Request
}

func (c *Channel) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.CHANNELS_URL)
	return c.Request.Get(path, queryParams, extraHeaders)
}

func (c *Channel) Countries(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.COUNTRIES_URL)
	return c.Request.Get(path, queryParams, extraHeaders)
}

func (c *Channel) CountryZones(countryID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s/%s", constants.EXTERNAL_PREFIX, constants.COUNTRIES_SHOW_URL, url.PathEscape(countryID))
	return c.Request.Get(path, queryParams, extraHeaders)
}

func (c *Channel) PostcodeDetails(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.POSTCODE_DETAILS_URL)
	return c.Request.Get(path, queryParams, extraHeaders)
}
