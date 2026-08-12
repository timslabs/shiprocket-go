package resources_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/utils"
)

func TestAccountWalletBalance(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ACCOUNT_WALLET_BALANCE_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_wallet_balance")
	defer teardown()

	body, err := utils.Client.Account.WalletBalance(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestAccountStatement(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ACCOUNT_STATEMENT_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_statement")
	defer teardown()

	body, err := utils.Client.Account.Statement(map[string]interface{}{
		"from": "2026-01-01",
		"to":   "2026-01-31",
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}

func TestAccountDiscrepancy(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.BILLING_DISCREPANCY_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_discrepancy")
	defer teardown()

	body, err := utils.Client.Account.Discrepancy(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(body)
	utils.TestResponse(b, []byte(fixture), t)
}
