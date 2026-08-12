package resources_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/utils"
)

const TestProductID = "55"

func TestProductAll(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.PRODUCTS_URL)
	teardown, fixture := utils.StartMockServer(url, "product_collection")
	defer teardown()

	body, err := utils.Client.Product.All(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestProductCreate(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.PRODUCTS_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_product")
	defer teardown()

	body, err := utils.Client.Product.Create(map[string]interface{}{"name": "Widget", "sku": "W-1"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestProductFetch(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s", constants.EXTERNAL_PREFIX, constants.PRODUCTS_SHOW_URL, TestProductID)
	teardown, fixture := utils.StartMockServer(url, "fake_product")
	defer teardown()

	body, err := utils.Client.Product.Fetch(TestProductID, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestProductSample(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.PRODUCTS_SAMPLE_URL)
	teardown, fixture := utils.StartMockServer(url, "product_collection")
	defer teardown()

	body, err := utils.Client.Product.Sample(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}
