package resources

import (
	"fmt"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/requests"
)

type Auth struct {
	Request *requests.Request
}

func (a *Auth) Login(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.AUTH_LOGIN_URL)
	return a.Request.PostUnauthenticated(url, data, extraHeaders)
}

func (a *Auth) Logout(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.AUTH_LOGOUT_URL)
	return a.Request.Post(url, data, extraHeaders)
}
