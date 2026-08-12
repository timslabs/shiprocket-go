package resources_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/utils"
)

func TestWarehouseSrfServiceability(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.WAREHOUSE_PREFIX, constants.WAREHOUSE_SRF_SERVICEABILITY)
	teardown, fixture := utils.StartMockServer(url, "fake_warehouse_srf")
	defer teardown()

	body, err := utils.Client.Warehouse.SrfServiceability(map[string]interface{}{
		"postcode": "110030",
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}
