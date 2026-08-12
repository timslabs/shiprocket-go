package resources_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/utils"
)

func TestPickupAll(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.PICKUP_LIST_URL)
	teardown, fixture := utils.StartMockServer(url, "pickup_collection")
	defer teardown()

	body, err := utils.Client.Pickup.All(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestPickupAdd(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.PICKUP_ADD_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_pickup_add")
	defer teardown()

	body, err := utils.Client.Pickup.Add(map[string]interface{}{
		"pickup_location": "Primary",
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}
