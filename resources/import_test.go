package resources_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/utils"
)

func TestImportCheck(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s/check", constants.EXTERNAL_PREFIX, constants.ERRORS_CHECK_URL, "77")
	teardown, fixture := utils.StartMockServer(url, "fake_import_check")
	defer teardown()

	body, err := utils.Client.Import.Check("77", nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}
