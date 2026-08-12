package resources_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/utils"
)

func TestListingAll(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.LISTINGS_URL)
	teardown, fixture := utils.StartMockServer(url, "listing_collection")
	defer teardown()

	body, err := utils.Client.Listing.All(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestListingLink(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.LISTINGS_LINK_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_listing_link")
	defer teardown()

	body, err := utils.Client.Listing.Link(map[string]interface{}{
		"channel_sku": "CH-1",
		"sku":         "W-1",
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestListingExportMapped(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.LISTINGS_EXPORT_MAPPED_URL)
	teardown, fixture := utils.StartMockServer(url, "listing_collection")
	defer teardown()

	body, err := utils.Client.Listing.ExportMapped(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestListingExportUnmapped(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.LISTINGS_EXPORT_UNMAPPED_URL)
	teardown, fixture := utils.StartMockServer(url, "listing_collection")
	defer teardown()

	body, err := utils.Client.Listing.ExportUnmapped(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestListingSample(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.LISTINGS_SAMPLE_URL)
	teardown, fixture := utils.StartMockServer(url, "listing_collection")
	defer teardown()

	body, err := utils.Client.Listing.Sample(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}
