package mercadopago

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
)

// General API configuration values
const (
	MPVersion   string = "0.1"
	MPUserAgent string = "MercadoPago Go SDK v" + MPVersion
	APIBaseURL  string = "https://api.mercadopago.com"
	MIMEJSON    string = "application/json"
	MIMEForm    string = "application/x-www-form-urlencoded"
)

// MP is the implementation to consume Mercado Pago API services
type MP struct {
	CustomAccessToken string
	BasicAccessToken  string
	ClientID          string
	clientSecret      string
	Sandbox           bool
	Debug             bool
}

// TokenResponse is the structure of data obtained from the MP Auth Token service
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	LiveMode     bool   `json:"live_mode"`
	UserID       int32  `json:"user_id"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int32  `json:"expires_in"`
	Scope        string `json:"scope"`
}

// NewMP returns a new instance of the MP service library
func NewMP(clientID string, clientSecret string, customAccessToken string, sandbox bool, debug bool) MP {
	mp := MP{}
	mp.BasicAccessToken = ""
	mp.CustomAccessToken = customAccessToken
	mp.ClientID = clientID
	mp.clientSecret = clientSecret
	mp.Sandbox = sandbox
	mp.Debug = debug
	return mp
}

// GetAccessToken returns an Access Token obtained from MP API
func (mp *MP) GetAccessToken() (string, error) {
	if mp.BasicAccessToken == "" {
		err := mp.obtainAccessToken()
		if err != nil {
			return "", err
		}
	}
	return mp.BasicAccessToken, nil
}

func (mp *MP) obtainAccessToken() error {
	data := &url.Values{}
	data.Set("client_id", mp.ClientID)
	data.Add("client_secret", mp.clientSecret)
	data.Add("grant_type", "client_credentials")
	r, err := mp.restFormCall("POST", "/oauth/token", data, 0)
	if err != nil {
		return err
	}
	// Check response status
	if r.StatusCode != 200 && r.StatusCode != 201 {
		return fmt.Errorf("Bad status received in HTTP response: %v", r.Status)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	var token TokenResponse
	if err = json.Unmarshal(body, &token); err != nil {
		return err
	}
	mp.BasicAccessToken = token.AccessToken
	return nil
}

// GET HTTP method wrapper for authentication (Form)
func (mp *MP) get(resource string, values *url.Values, auth int) (*http.Response, error) {
	return mp.restFormCall("GET", resource, values, auth)
}

// JGET HTTP method wrapper for authentication (JSON)
func (mp *MP) jget(resource string, data interface{}, auth int) (*http.Response, error) {
	dataBuffer := new(bytes.Buffer)
	if data != nil {
		json.NewEncoder(dataBuffer).Encode(data)
	}
	return mp.restJSONCall("GET", resource, dataBuffer, auth)
}

// POST HTTP method wrapper for authentication (JSON)
func (mp *MP) post(resource string, data interface{}, auth int) (*http.Response, error) {
	dataBuffer := new(bytes.Buffer)
	if data != nil {
		json.NewEncoder(dataBuffer).Encode(data)
	}
	return mp.restJSONCall("POST", resource, dataBuffer, auth)
}

// PUT HTTP method wrapper for authentication (JSON)
func (mp *MP) put(resource string, data interface{}, auth int) (*http.Response, error) {
	dataBuffer := new(bytes.Buffer)
	if data != nil {
		json.NewEncoder(dataBuffer).Encode(data)
	}
	return mp.restJSONCall("PUT", resource, dataBuffer, auth)
}

// generic API REST call with Mercado Pago preferences
func (mp *MP) restFormCall(method string, resource string, values *url.Values, auth int) (*http.Response, error) {
	// Build resource URL
	u, err := url.ParseRequestURI(APIBaseURL)
	if err != nil {
		return nil, err
	}
	u.Path = resource
	urlStr := fmt.Sprintf("%v", u)
	// If no values passed, then initialize an empty object
	if values == nil {
		values = &url.Values{}
	}
	// If authed method, then add a form entry for the MP Access Token (Basic Workflow)
	if auth == 1 {
		if mp.BasicAccessToken == "" {
			err := mp.obtainAccessToken()
			if err != nil {
				return nil, err
			}
		}
		values.Add("access_token", mp.BasicAccessToken)
	}
	// If authed method, then add a form entry for the MP Access Token (Custom Workflow)
	if auth == 2 {
		values.Add("access_token", mp.CustomAccessToken)
	}
	// Create HTTP Request
	r, _ := http.NewRequest(method, urlStr, bytes.NewBufferString(values.Encode()))
	if err == nil {
		r.Header.Add("Content-Length", strconv.Itoa(len(values.Encode())))
		r.Header.Set("User-Agent", MPUserAgent)
		r.Header.Add("accept", MIMEJSON)
		r.Header.Add("content-type", MIMEForm)
		// DEBUG only - Print full request
		if mp.Debug {
			debug(httputil.DumpRequestOut(r, true))
		}
	}
	client := &http.Client{}
	resp, respErr := client.Do(r)
	if respErr == nil {
		// DEBUG only - Print full response
		if mp.Debug {
			debug(httputil.DumpResponse(resp, true))
		}
	}
	return resp, respErr
}

// generic API REST call with Mercado Pago preferences
func (mp *MP) restJSONCall(method string, resource string, data *bytes.Buffer, auth int) (*http.Response, error) {
	// Build resource URL
	u, err := url.ParseRequestURI(APIBaseURL)
	if err != nil {
		return nil, err
	}
	u.Path = resource
	urlStr := fmt.Sprintf("%v", u)
	// If authed method, then add a form entry for the MP Access Token (Basic Workflow)
	if auth == 1 {
		if mp.BasicAccessToken == "" {
			err := mp.obtainAccessToken()
			if err != nil {
				return nil, err
			}
		}
		if strings.Contains(urlStr, "?") {
			urlStr += "&access_token=" + mp.BasicAccessToken
		} else {
			urlStr += "?access_token=" + mp.BasicAccessToken
		}
	}
	// If authed method, then add a form entry for the MP Access Token (Custom Workflow)
	if auth == 2 {
		if strings.Contains(urlStr, "?") {
			urlStr += "&access_token=" + mp.CustomAccessToken
		} else {
			urlStr += "?access_token=" + mp.CustomAccessToken
		}
	}

	// Create HTTP Request
	r, err := http.NewRequest(method, urlStr, data)
	if err == nil {
		r.Header.Add("Content-Length", strconv.Itoa(data.Len()))
		r.Header.Set("User-Agent", MPUserAgent)
		r.Header.Set("Content-Type", MIMEJSON)
		r.Header.Add("Accept", MIMEJSON)
		// DEBUG only - Print full request
		if mp.Debug {
			debug(httputil.DumpRequestOut(r, true))
		}
	}

	client := &http.Client{}
	resp, respErr := client.Do(r)
	if respErr == nil {
		// DEBUG only - Print full response
		if mp.Debug {
			debug(httputil.DumpResponse(resp, true))
		}
	}
	return resp, respErr
}

func debug(data []byte, err error) {
	if err == nil {
		fmt.Printf("%s\n\n", data)
	} else {
		log.Fatalf("%s\n\n", err)
	}
}
