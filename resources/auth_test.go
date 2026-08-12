package resources_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/utils"
)

func TestAuthLogout(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.AUTH_LOGOUT_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_logout")
	defer teardown()

	body, err := utils.Client.Auth.Logout(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}
