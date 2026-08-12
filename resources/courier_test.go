package resources_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/utils"
)

const TestAwb = "AWB123456"
const TestShipmentTrackID = "99"

func TestCourierServiceability(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.COURIER_SERVICEABILITY_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_serviceability")
	defer teardown()

	body, err := utils.Client.Courier.Serviceability(map[string]interface{}{
		"pickup_postcode":   "110030",
		"delivery_postcode": "122001",
		"weight":            0.5,
		"cod":               0,
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestCourierListWithCounts(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.COURIER_LIST_COUNTS_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_serviceability")
	defer teardown()

	body, err := utils.Client.Courier.ListWithCounts(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestCourierAssignAwb(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.COURIER_ASSIGN_AWB_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_assign_awb")
	defer teardown()

	body, err := utils.Client.Courier.AssignAwb(map[string]interface{}{
		"shipment_id": 2001,
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestCourierGeneratePickup(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.COURIER_GENERATE_PICKUP_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_pickup_add")
	defer teardown()

	body, err := utils.Client.Courier.GeneratePickup(map[string]interface{}{
		"shipment_id": []interface{}{2001},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestCourierGenerateLabel(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.COURIER_GENERATE_LABEL_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_pickup_add")
	defer teardown()

	body, err := utils.Client.Courier.GenerateLabel(map[string]interface{}{
		"shipment_id": []interface{}{2001},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestCourierGenerateManifest(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.MANIFESTS_GENERATE_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_pickup_add")
	defer teardown()

	body, err := utils.Client.Courier.GenerateManifest(map[string]interface{}{
		"shipment_id": []interface{}{2001},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestCourierPrintManifest(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.MANIFESTS_PRINT_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_pickup_add")
	defer teardown()

	body, err := utils.Client.Courier.PrintManifest(map[string]interface{}{
		"shipment_id": []interface{}{2001},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestCourierCancelShipment(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.CANCEL_SHIPMENT_AWBS_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_pickup_add")
	defer teardown()

	body, err := utils.Client.Courier.CancelShipment(map[string]interface{}{
		"awbs": []interface{}{TestAwb},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestCourierTrackByAwb(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s", constants.EXTERNAL_PREFIX, constants.COURIER_TRACK_AWB_URL, TestAwb)
	teardown, fixture := utils.StartMockServer(url, "fake_tracking")
	defer teardown()

	body, err := utils.Client.Courier.TrackByAwb(TestAwb, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestCourierTrackByShipment(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s", constants.EXTERNAL_PREFIX, constants.COURIER_TRACK_SHIPMENT_URL, TestShipmentTrackID)
	teardown, fixture := utils.StartMockServer(url, "fake_tracking")
	defer teardown()

	body, err := utils.Client.Courier.TrackByShipment(TestShipmentTrackID, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestCourierTrackByOrder(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.COURIER_TRACK_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_tracking")
	defer teardown()

	body, err := utils.Client.Courier.TrackByOrder(map[string]interface{}{
		"order_id": 123,
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestCourierTrackByAwbs(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.COURIER_TRACK_AWBS_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_tracking")
	defer teardown()

	body, err := utils.Client.Courier.TrackByAwbs(map[string]interface{}{
		"awbs": []interface{}{TestAwb},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}
