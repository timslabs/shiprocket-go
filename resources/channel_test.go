package resources_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/utils"
)

func TestChannelAll(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.CHANNELS_URL)
	teardown, fixture := utils.StartMockServer(url, "channel_collection")
	defer teardown()

	body, err := utils.Client.Channel.All(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestChannelCountries(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.COUNTRIES_URL)
	teardown, fixture := utils.StartMockServer(url, "channel_collection")
	defer teardown()

	body, err := utils.Client.Channel.Countries(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestChannelCountryZones(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s", constants.EXTERNAL_PREFIX, constants.COUNTRIES_SHOW_URL, "1")
	teardown, fixture := utils.StartMockServer(url, "channel_collection")
	defer teardown()

	body, err := utils.Client.Channel.CountryZones("1", nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestChannelPostcodeDetails(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.POSTCODE_DETAILS_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_postcode_details")
	defer teardown()

	body, err := utils.Client.Channel.PostcodeDetails(map[string]interface{}{
		"postcode": "110030",
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}
