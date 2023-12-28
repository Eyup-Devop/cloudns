package cloudns

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

//
// Public constants
//

const (
	// APIVersion is the currently supported API version
	APIVersion string = apiVersion

	// APIURL is the URL of the API service backend.
	APIURL string = "https://api.cloudns.net"
)

//
// Public variables
//

var (
	// AuthId Api user ID
	AuthId *string = nil

	// SubAuthId Api sub user ID
	SubAuthId *int = nil

	// SubAuthUser Api sub user name
	SubAuthUser *string = nil

	// AuthPassword is the password for API user ID or for API sub user ID
	AuthPassword *string = nil

	// DefaultMaxNetworkRetries is the default maximum number of retries made
	// by a Cloudns client.
	DefaultMaxNetworkRetries int64 = 2
)

var (
	appInfo  *AppInfo
	backends Backends

	encodedUserAgent string

	httpClient = &http.Client{
		Timeout: defaultHTTPTimeout,
		//	Transport: &http.Transport{
		//		TLSNextProto: make(map[string]func(string, *tls.Conn) http.RoundTripper),
		//	},
	}

	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

const (
	clientversion      = "00.00.10"
	defaultHTTPTimeout = 80 * time.Second
)

type AppInfo struct {
	Name      string `json:"name"`
	PartnerID string `json:"partner_id"`
	URL       string `json:"url"`
	Version   string `json:"version"`
}

func (a *AppInfo) formatUserAgent() string {
	str := a.Name
	if a.Version != "" {
		str += "/" + a.Version
	}
	if a.URL != "" {
		str += " (" + a.URL + ")"
	}
	return str
}

func SetAppInfo(info *AppInfo) {
	if info != nil && info.Name == "" {
		panic(fmt.Errorf("app info name cannot be empty"))
	}
	appInfo = info

	initUserAgent()
}

func initUserAgent() {
	encodedUserAgent = "Cloudns/v1 GoBindings/" + clientversion
	if appInfo != nil {
		encodedUserAgent += " " + appInfo.formatUserAgent()
	}
}

func init() {
	ValidatorInit()

	initUserAgent()
}

type Backend interface {
	Call(method, path string, params interface{}, v interface{}, isPlain bool) error
}

type BackendConfig struct {
	HTTPClient *http.Client

	MaxNetworkRetries *int64

	URL *string
}

func extractParams(params interface{}) ([]byte, error) {
	val := reflect.ValueOf(params)

	authIdField := val.Elem().FieldByName("AuthId")
	subAuthId := val.Elem().FieldByName("SubAuthId")
	subAuthUser := val.Elem().FieldByName("SubAuthUser")
	authPassword := val.Elem().FieldByName("AuthPassword")

	if authIdField.CanSet() && AuthId != nil {
		authIdField.Set(reflect.ValueOf(AuthId))
	}
	if subAuthId.CanSet() && SubAuthId != nil {
		subAuthId.Set(reflect.ValueOf(SubAuthId))
	}
	if subAuthUser.CanSet() && SubAuthUser != nil {
		subAuthUser.Set(reflect.ValueOf(SubAuthUser))
	}
	if authPassword.CanSet() && AuthPassword != nil {
		authPassword.Set(reflect.ValueOf(AuthPassword))
	}

	return json.Marshal(val.Interface())
}

func String(v string) *string {
	return &v
}

func Int(v int) *int {
	return &v
}

func Int64(v int64) *int64 {
	return &v
}

type BackendImplementation struct {
	URL                 string
	HTTPClient          *http.Client
	MaxNetworkRetries   int64
	networkRetriesSleep bool
}

func (s *BackendImplementation) Call(method, path string, params interface{}, v interface{}, isPlain bool) error {
	body, err := extractParams(params)
	if err != nil {
		return err
	}
	return s.CallRaw(method, path, body, v, isPlain)
}

func (s *BackendImplementation) CallRaw(method, path string, params []byte, v interface{}, isPlain bool) error {
	bodyBuffer := bytes.NewBuffer(params)

	req, err := s.NewRequest(method, path, "application/json", bodyBuffer)
	if err != nil {
		return err
	}

	if err := s.Do(req, bodyBuffer, v, isPlain); err != nil {
		return err
	}

	return nil
}

func (s *BackendImplementation) NewRequest(method, path, contentType string, params *bytes.Buffer) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = s.URL + path

	req, err := http.NewRequest(method, path, params)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	req.Header.Add("User-Agent", encodedUserAgent)

	return req, nil
}

func (s *BackendImplementation) Do(req *http.Request, body *bytes.Buffer, v interface{}, isPlain bool) error {
	handleResponse := func(res *http.Response, err error) (interface{}, error) {
		var resBody []byte
		if err == nil {
			resBody, err = io.ReadAll(res.Body)
			res.Body.Close()
		}

		return resBody, err
	}

	res, result, _, err := s.requestWithRetries(req, body, handleResponse)
	if err != nil {
		return err
	}
	resBody := result.([]byte)

	if isPlain {
		err = s.ReadPlainResponse(resBody, v)
	} else {
		err = s.UnmarshalJSONVerbose(res.StatusCode, resBody, v)
	}
	//	v.SetLastResponse(newAPIResponse(res, resBody, requestDuration))
	return err
}

func (s *BackendImplementation) ReadPlainResponse(resBody []byte, v interface{}) error {
	t := reflect.TypeOf(v)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	switch t.Kind() {
	case reflect.Int:
		intValue, err := strconv.Atoi(string(resBody))
		if err != nil {
			return errors.New(string(resBody))
		}
		reflect.ValueOf(v).Elem().SetInt(int64(intValue))
	case reflect.String:
		reflect.ValueOf(v).Elem().SetString(string(resBody))
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(string(resBody))
		if err != nil {
			boolValue = false
		}
		reflect.ValueOf(v).Elem().SetBool(boolValue)
	}
	return nil
}

func (s *BackendImplementation) requestWithRetries(req *http.Request, body *bytes.Buffer,
	handleResponse func(*http.Response, error) (interface{}, error),
) (*http.Response, interface{}, *time.Duration, error) {
	var resp *http.Response
	var err error
	var requestDuration time.Duration
	var result interface{}
	for retry := 0; ; {
		start := time.Now()
		resetBodyReader(body, req)

		resp, err = s.HTTPClient.Do(req)

		requestDuration = time.Since(start)

		result, err = handleResponse(resp, err)

		//		shouldRetry, noRetryReason := s.shouldRetry(err, req, resp, retry)
		shouldRetry, _ := s.shouldRetry(err, req, resp, retry)

		if !shouldRetry {
			//	fmt.Printf("Not retrying request: %v", noRetryReason)
			break
		}

		sleepDuration := s.sleepTime(retry)
		retry++

		time.Sleep(sleepDuration)
	}

	if err != nil {
		return nil, nil, nil, err
	}

	return resp, result, &requestDuration, nil
}

func (s *BackendImplementation) shouldRetry(err error, req *http.Request, resp *http.Response, numRetries int) (bool, string) {
	if numRetries >= int(s.MaxNetworkRetries) {
		return false, "max retries exceeded"
	}

	if resp.StatusCode >= http.StatusInternalServerError && req.Method == http.MethodPost {
		return true, ""
	}

	if resp.StatusCode == http.StatusServiceUnavailable {
		return true, ""
	}

	return false, "response not known to be safe for retry"
}

const (
	maxNetworkRetriesDelay = 5000 * time.Millisecond
	minNetworkRetriesDelay = 500 * time.Millisecond
)

func (s *BackendImplementation) sleepTime(numRetries int) time.Duration {
	if !s.networkRetriesSleep {
		return 0 * time.Second
	}

	delay := minNetworkRetriesDelay + minNetworkRetriesDelay*time.Duration(numRetries*numRetries)

	if delay > maxNetworkRetriesDelay {
		delay = maxNetworkRetriesDelay
	}

	jitter := rand.Int63n(int64(delay / 4))
	delay -= time.Duration(jitter)

	if delay < minNetworkRetriesDelay {
		delay = minNetworkRetriesDelay
	}

	return delay
}

func (s *BackendImplementation) UnmarshalJSONVerbose(statusCode int, body []byte, v interface{}) error {
	err := json.Unmarshal(body, v)
	if err != nil {

		bodySample := string(body)
		if len(bodySample) > 500 {
			bodySample = bodySample[0:500] + " ..."
		}

		bodySample = strings.Replace(bodySample, "\n", "\\n", -1)

		newErr := fmt.Errorf("couldn't deserialize JSON (response status: %v, body sample: '%s'): %v",
			statusCode, bodySample, err)
		return newErr
	}

	return nil
}

func resetBodyReader(body *bytes.Buffer, req *http.Request) {
	if body != nil {
		reader := bytes.NewReader(body.Bytes())

		req.Body = nopReadCloser{reader}

		req.GetBody = func() (io.ReadCloser, error) {
			reader := bytes.NewReader(body.Bytes())
			return nopReadCloser{reader}, nil
		}
	}
}

type nopReadCloser struct {
	io.Reader
}

func (nopReadCloser) Close() error { return nil }

type Backends struct {
	API Backend
	mu  sync.RWMutex
}

func GetBackend() Backend {
	var backend Backend

	backends.mu.RLock()
	backend = backends.API
	backends.mu.RUnlock()
	if backend != nil {
		return backend
	}

	backend = GetBackendWithConfig(
		&BackendConfig{
			HTTPClient:        httpClient,
			MaxNetworkRetries: nil, // Set by GetBackendWithConfiguation when nil
			URL:               nil, // Set by GetBackendWithConfiguation when nil
		},
	)

	SetBackend(backend)

	return backend
}

func GetBackendWithConfig(config *BackendConfig) Backend {
	if config.HTTPClient == nil {
		config.HTTPClient = httpClient
	}

	if config.MaxNetworkRetries == nil {
		config.MaxNetworkRetries = Int64(DefaultMaxNetworkRetries)
	}

	if config.URL == nil {
		config.URL = String(APIURL)
	}

	config.URL = String(normalizeURL(*config.URL))

	return newBackendImplementation(config)
}

func normalizeURL(url string) string {
	url = strings.TrimSuffix(url, "/")
	return url
}

func newBackendImplementation(config *BackendConfig) Backend {
	return &BackendImplementation{
		HTTPClient:          config.HTTPClient,
		MaxNetworkRetries:   *config.MaxNetworkRetries,
		URL:                 *config.URL,
		networkRetriesSleep: true,
	}
}

func SetBackend(b Backend) {
	backends.mu.Lock()
	defer backends.mu.Unlock()
	backends.API = b
}
