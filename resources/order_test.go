package resources_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/utils"
)

const TestOrderID = "1001"

func TestOrderAll(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_URL)
	teardown, fixture := utils.StartMockServer(url, "order_collection")
	defer teardown()

	body, err := utils.Client.Order.All(map[string]interface{}{"page": 1}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestOrderFetch(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s", constants.EXTERNAL_PREFIX, constants.ORDERS_SHOW_URL, TestOrderID)
	teardown, fixture := utils.StartMockServer(url, "fake_order")
	defer teardown()

	body, err := utils.Client.Order.Fetch(TestOrderID, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestOrderCreateAdhoc(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_CREATE_ADHOC_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_order")
	defer teardown()

	body, err := utils.Client.Order.CreateAdhoc(map[string]interface{}{
		"order_id": "ORD-1001",
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestOrderCreate(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_CREATE_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_order")
	defer teardown()

	body, err := utils.Client.Order.Create(map[string]interface{}{"order_id": "ORD-1001"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestOrderCreateReturn(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_CREATE_RETURN_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_order")
	defer teardown()

	body, err := utils.Client.Order.CreateReturn(map[string]interface{}{"order_id": "ORD-1001"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestOrderUpdateAdhoc(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_UPDATE_ADHOC_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_order")
	defer teardown()

	body, err := utils.Client.Order.UpdateAdhoc(map[string]interface{}{"order_id": "1"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestOrderCancel(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_CANCEL_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_pickup_add")
	defer teardown()

	body, err := utils.Client.Order.Cancel(map[string]interface{}{"ids": []interface{}{1}}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestOrderProcessingReturns(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_PROCESSING_RETURN_URL)
	teardown, fixture := utils.StartMockServer(url, "order_collection")
	defer teardown()

	body, err := utils.Client.Order.ProcessingReturns(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestOrderUpdateAddress(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_ADDRESS_UPDATE_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_pickup_add")
	defer teardown()

	body, err := utils.Client.Order.UpdateAddress(map[string]interface{}{}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestOrderUpdatePickupAddress(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_ADDRESS_PICKUP_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_pickup_add")
	defer teardown()

	body, err := utils.Client.Order.UpdatePickupAddress(map[string]interface{}{}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestOrderFulfill(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_FULFILL_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_pickup_add")
	defer teardown()

	body, err := utils.Client.Order.Fulfill(map[string]interface{}{}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestOrderMapping(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_MAPPING_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_pickup_add")
	defer teardown()

	body, err := utils.Client.Order.Mapping(map[string]interface{}{}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestOrderExport(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_EXPORT_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_import_check")
	defer teardown()

	body, err := utils.Client.Order.Export(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestOrderPrintInvoice(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ORDERS_PRINT_INVOICE_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_pickup_add")
	defer teardown()

	body, err := utils.Client.Order.PrintInvoice(map[string]interface{}{"ids": []interface{}{1}}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}
