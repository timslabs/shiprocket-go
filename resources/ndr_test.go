package resources_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/utils"
)

func TestNdrAll(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.NDR_ALL_URL)
	teardown, fixture := utils.StartMockServer(url, "ndr_collection")
	defer teardown()

	body, err := utils.Client.Ndr.All(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestNdrFetch(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s", constants.EXTERNAL_PREFIX, constants.NDR_URL, "AWB1")
	teardown, fixture := utils.StartMockServer(url, "fake_ndr")
	defer teardown()

	body, err := utils.Client.Ndr.Fetch("AWB1", nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestNdrAction(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s/action", constants.EXTERNAL_PREFIX, constants.NDR_URL, "AWB1")
	teardown, fixture := utils.StartMockServer(url, "fake_ndr_action")
	defer teardown()

	body, err := utils.Client.Ndr.Action("AWB1", map[string]interface{}{"action": "re-attempt"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}
