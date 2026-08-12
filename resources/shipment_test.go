package resources_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/utils"
)

const TestShipmentID = "2001"

func TestShipmentAll(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.SHIPMENTS_URL)
	teardown, fixture := utils.StartMockServer(url, "shipment_collection")
	defer teardown()

	body, err := utils.Client.Shipment.All(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestShipmentFetch(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s", constants.EXTERNAL_PREFIX, constants.SHIPMENTS_URL, TestShipmentID)
	teardown, fixture := utils.StartMockServer(url, "fake_shipment")
	defer teardown()

	body, err := utils.Client.Shipment.Fetch(TestShipmentID, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestShipmentCreateForward(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.SHIPMENTS_CREATE_FORWARD_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_shipment")
	defer teardown()

	body, err := utils.Client.Shipment.CreateForward(map[string]interface{}{"order_id": 1}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestShipmentCreateReturn(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.SHIPMENTS_CREATE_RETURN_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_shipment")
	defer teardown()

	body, err := utils.Client.Shipment.CreateReturn(map[string]interface{}{"order_id": 1}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}
