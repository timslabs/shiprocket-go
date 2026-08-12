package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/timslabs/shiprocket-go/constants"
	"github.com/timslabs/shiprocket-go/errors"
)

// TIMEOUT is the default HTTP client timeout in seconds.
const TIMEOUT = 60

// Auth holds Shiprocket API-user credentials and/or a JWT.
type Auth struct {
	Email    string
	Password string
	Token    string
}

// Request encapsulates HTTP access to Shiprocket's APIs.
type Request struct {
	Auth       Auth
	HTTPClient *http.Client
	Headers    map[string]string
	Version    string
	SDKName    string
	BaseURL    string
	userAgent  string

	mu sync.Mutex
}

// AddHeaders appends headers used on subsequent requests.
func (request *Request) AddHeaders(headers map[string]string) {
	for key, value := range headers {
		request.Headers[key] = value
	}
}

// SetTimeout overrides the HTTP client timeout (seconds).
func (request *Request) SetTimeout(timeout int16) {
	request.HTTPClient = &http.Client{Timeout: time.Duration(timeout) * time.Second}
}

// SetUserAgent sets a custom User-Agent prefix.
func (request *Request) SetUserAgent(userAgent string) {
	request.userAgent = userAgent
}

// GetUserAgent returns the custom User-Agent prefix.
func (request *Request) GetUserAgent() string {
	return request.userAgent
}

// SetToken sets the bearer token used for authenticated calls.
func (request *Request) SetToken(token string) {
	request.mu.Lock()
	defer request.mu.Unlock()
	request.Auth.Token = token
}

// Token returns the current bearer token.
func (request *Request) Token() string {
	request.mu.Lock()
	defer request.mu.Unlock()
	return request.Auth.Token
}

// Authenticate logs in with email/password and stores the JWT.
func (request *Request) Authenticate() error {
	request.mu.Lock()
	email := request.Auth.Email
	password := request.Auth.Password
	request.mu.Unlock()

	if email == "" || password == "" {
		return &errors.AuthError{Message: "email and password are required", StatusCode: http.StatusBadRequest}
	}

	data := map[string]interface{}{
		"email":    email,
		"password": password,
	}
	path := fmt.Sprintf("/%s%s", constants.EXTERNAL_PREFIX, constants.AUTH_LOGIN_URL)
	resp, err := request.do(http.MethodPost, path, nil, data, nil, false)
	if err != nil {
		return err
	}
	token, _ := resp["token"].(string)
	if token == "" {
		return &errors.AuthError{Message: "token missing from Shiprocket login response", StatusCode: http.StatusOK}
	}
	request.SetToken(token)
	return nil
}

func (request *Request) ensureAuth() error {
	if request.Token() != "" {
		return nil
	}
	if request.Auth.Email == "" || request.Auth.Password == "" {
		return &errors.AuthError{Message: "access token is not set; call Authenticate() or NewClientWithToken()", StatusCode: http.StatusUnauthorized}
	}
	return request.Authenticate()
}

// Get ...
func (request *Request) Get(path string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	if err := request.ensureAuth(); err != nil {
		return nil, err
	}
	return request.do(http.MethodGet, path, queryParams, nil, extraHeaders, true)
}

// Post ...
func (request *Request) Post(path string, payload map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	if err := request.ensureAuth(); err != nil {
		return nil, err
	}
	return request.do(http.MethodPost, path, nil, payload, extraHeaders, true)
}

// PostUnauthenticated issues a POST without a bearer token (e.g. login).
func (request *Request) PostUnauthenticated(path string, payload map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return request.do(http.MethodPost, path, nil, payload, extraHeaders, false)
}

// Put ...
func (request *Request) Put(path string, payload map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	if err := request.ensureAuth(); err != nil {
		return nil, err
	}
	return request.do(http.MethodPut, path, nil, payload, extraHeaders, true)
}

// Patch ...
func (request *Request) Patch(path string, payload map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	if err := request.ensureAuth(); err != nil {
		return nil, err
	}
	return request.do(http.MethodPatch, path, nil, payload, extraHeaders, true)
}

// Delete ...
func (request *Request) Delete(path string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	if err := request.ensureAuth(); err != nil {
		return nil, err
	}
	return request.do(http.MethodDelete, path, queryParams, nil, extraHeaders, true)
}

// FileUploadParams describes a multipart file upload.
type FileUploadParams struct {
	File   *os.File
	Fields map[string]string
}

// File posts multipart/form-data (CSV/Excel imports).
func (request *Request) File(path string, params FileUploadParams, extraHeaders map[string]string) (map[string]interface{}, error) {
	if err := request.ensureAuth(); err != nil {
		return nil, err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if params.File != nil {
		part, err := writer.CreateFormFile("file", filepath.Base(params.File.Name()))
		if err != nil {
			return nil, err
		}
		if _, err := io.Copy(part, params.File); err != nil {
			return nil, err
		}
	}
	for fieldName, fieldValue := range params.Fields {
		if err := writer.WriteField(fieldName, fieldValue); err != nil {
			return nil, err
		}
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	requestURL := request.BaseURL + path
	req, err := http.NewRequest(http.MethodPost, requestURL, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set(constants.AuthorizationHeader, constants.BearerPrefix+request.Token())
	request.addRequestHeaders(req, extraHeaders, false)
	req.Header.Set(constants.ContentTypeHeader, writer.FormDataContentType())

	return request.doRequestResponse(req)
}

func (request *Request) do(
	method, path string,
	queryParams map[string]interface{},
	payload map[string]interface{},
	extraHeaders map[string]string,
	requiresAuth bool,
) (map[string]interface{}, error) {
	requestURL := request.BaseURL + path
	if len(queryParams) > 0 {
		requestURL = buildURLWithParams(requestURL, queryParams)
	}

	var body io.Reader
	if payload != nil {
		b, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(method, requestURL, body)
	if err != nil {
		return nil, err
	}

	if requiresAuth {
		req.Header.Set(constants.AuthorizationHeader, constants.BearerPrefix+request.Token())
	}
	request.addRequestHeaders(req, extraHeaders, payload != nil)

	return request.doRequestResponse(req)
}

func buildURLWithParams(requestURL string, data map[string]interface{}) string {
	u, err := url.Parse(requestURL)
	if err != nil {
		panic(err)
	}
	parameters := url.Values{}
	for k, v := range data {
		if v == nil {
			continue
		}
		if arr, ok := v.([]string); ok {
			for _, item := range arr {
				parameters.Add(k, item)
			}
			continue
		}
		parameters.Add(k, fmt.Sprintf("%v", v))
	}
	u.RawQuery = parameters.Encode()
	return u.String()
}

func (request *Request) addRequestHeaders(req *http.Request, headers map[string]string, hasJSONBody bool) {
	req.Header.Set(constants.AcceptHeader, "application/json")
	req.Header.Set(constants.UserAgentHeader, request.getUserAgentHeaderValue())
	if hasJSONBody {
		req.Header.Set(constants.ContentTypeHeader, "application/json")
	}
	for key, value := range request.Headers {
		if key == constants.ContentTypeHeader || key == constants.UserAgentHeader {
			continue
		}
		req.Header.Set(key, value)
	}
	for key, value := range headers {
		if key == constants.ContentTypeHeader || key == constants.UserAgentHeader {
			continue
		}
		req.Header.Set(key, value)
	}
}

func (request *Request) getUserAgentHeaderValue() string {
	goSDKVersion := fmt.Sprintf("%s/%s", request.SDKName, request.Version)
	if request.userAgent == "" {
		return goSDKVersion
	}
	return fmt.Sprintf("%s (%s)", request.userAgent, goSDKVersion)
}

func (request *Request) doRequestResponse(req *http.Request) (map[string]interface{}, error) {
	response, err := request.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	raw, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode >= constants.HTTP_STATUS_OK && response.StatusCode < constants.HTTP_STATUS_REDIRECT {
		if len(raw) == 0 {
			return map[string]interface{}{}, nil
		}
		resp := make(map[string]interface{})
		if err := json.Unmarshal(raw, &resp); err != nil {
			return nil, err
		}
		return resp, nil
	}

	msg := strings.TrimSpace(string(raw))
	if msg == "" {
		msg = "API request failed"
	}
	if response.StatusCode >= 500 {
		return nil, &errors.ServerError{Message: msg, StatusCode: response.StatusCode, Body: string(raw)}
	}
	return nil, &errors.BadRequestError{Message: msg, StatusCode: response.StatusCode, Body: string(raw)}
}
