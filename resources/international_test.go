package resources_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/utils"
)

func TestInternationalCreateAdhoc(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.INTERNATIONAL_ORDERS_CREATE_ADHOC_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_international_order")
	defer teardown()

	body, err := utils.Client.International.CreateAdhoc(map[string]interface{}{
		"order_id": "INT-1",
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestInternationalUpdateAdhoc(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.INTERNATIONAL_ORDERS_UPDATE_ADHOC_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_international_order")
	defer teardown()

	body, err := utils.Client.International.UpdateAdhoc(map[string]interface{}{
		"order_id": "INT-1",
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestInternationalTrack(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.INTERNATIONAL_ORDERS_TRACK_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_tracking")
	defer teardown()

	body, err := utils.Client.International.Track(map[string]interface{}{
		"order_id": "INT-1",
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestInternationalServiceability(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.INTERNATIONAL_COURIER_SERVICEABILITY)
	teardown, fixture := utils.StartMockServer(url, "fake_international_serviceability")
	defer teardown()

	body, err := utils.Client.International.Serviceability(map[string]interface{}{
		"delivery_country": "US",
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestInternationalAssignAwb(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.INTERNATIONAL_COURIER_ASSIGN_AWB)
	teardown, fixture := utils.StartMockServer(url, "fake_assign_awb")
	defer teardown()

	body, err := utils.Client.International.AssignAwb(map[string]interface{}{
		"shipment_id": 3001,
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestInternationalGenerateManifest(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.INTERNATIONAL_MANIFESTS_GENERATE)
	teardown, fixture := utils.StartMockServer(url, "fake_assign_awb")
	defer teardown()

	body, err := utils.Client.International.GenerateManifest(map[string]interface{}{
		"shipment_id": []interface{}{3001},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestInternationalCreateForwardShipment(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.INTERNATIONAL_SHIPMENTS_FORWARD)
	teardown, fixture := utils.StartMockServer(url, "fake_shipment")
	defer teardown()

	body, err := utils.Client.International.CreateForwardShipment(map[string]interface{}{
		"order_id": 1,
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestInternationalAddBankDetails(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.INTERNATIONAL_ADD_BANK_DETAILS)
	teardown, fixture := utils.StartMockServer(url, "fake_pickup_add")
	defer teardown()

	body, err := utils.Client.International.AddBankDetails(map[string]interface{}{
		"account_number": "1",
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestInternationalKyc(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.INTERNATIONAL_KYC)
	teardown, fixture := utils.StartMockServer(url, "fake_pickup_add")
	defer teardown()

	body, err := utils.Client.International.Kyc(map[string]interface{}{
		"document_type": "PASSPORT",
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}
