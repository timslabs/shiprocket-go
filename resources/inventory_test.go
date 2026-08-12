package resources_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/utils"
)

func TestInventoryAll(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.INVENTORY_URL)
	teardown, fixture := utils.StartMockServer(url, "inventory_collection")
	defer teardown()

	body, err := utils.Client.Inventory.All(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestInventoryUpdate(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s/update", constants.EXTERNAL_PREFIX, constants.INVENTORY_URL, "55")
	teardown, fixture := utils.StartMockServer(url, "fake_inventory_update")
	defer teardown()

	body, err := utils.Client.Inventory.Update("55", map[string]interface{}{"quantity": 10}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}
