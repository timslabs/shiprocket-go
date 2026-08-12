package resources

import (
	"fmt"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/requests"
)

type Account struct {
	Request *requests.Request
}

func (acc *Account) WalletBalance(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ACCOUNT_WALLET_BALANCE_URL)
	return acc.Request.Get(url, queryParams, extraHeaders)
}

func (acc *Account) Statement(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.ACCOUNT_STATEMENT_URL)
	return acc.Request.Get(url, queryParams, extraHeaders)
}

func (acc *Account) Discrepancy(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.BILLING_DISCREPANCY_URL)
	return acc.Request.Get(url, queryParams, extraHeaders)
}
