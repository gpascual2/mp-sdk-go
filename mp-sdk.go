package mercadopago

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
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
	AccessToken  string
	ClientID     string
	clientSecret string
	Sandbox      bool
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
func NewMP(clientID string, clientSecret string, sandbox bool) MP {
	mp := MP{}
	mp.AccessToken = ""
	mp.ClientID = clientID
	mp.clientSecret = clientSecret
	mp.Sandbox = sandbox
	return mp
}

// GetAccessToken returns an Access Token obtained from MP API
func (mp *MP) GetAccessToken() (string, error) {
	if mp.AccessToken == "" {
		err := mp.obtainAccessToken()
		if err != nil {
			return "", err
		}
	}
	return mp.AccessToken, nil
}

func (mp *MP) obtainAccessToken() error {
	data := url.Values{}
	data.Set("client_id", mp.ClientID)
	data.Add("client_secret", mp.clientSecret)
	data.Add("grant_type", "client_credentials")
	r, err := mp.restFormCall("POST", "/oauth/token", data, false)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	var token TokenResponse
	if err = json.Unmarshal(body, &token); err != nil {
		return err
	}
	mp.AccessToken = token.AccessToken
	return nil
}

// GET HTTP method wrapper for authentication (Form)
func (mp *MP) get(resource string, values url.Values) (*http.Response, error) {
	return mp.restFormCall("GET", resource, values, true)
}

// POST HTTP method wrapper for authentication (JSON)
func (mp *MP) post(resource string, data interface{}) (*http.Response, error) {
	dataBuffer := new(bytes.Buffer)
	json.NewEncoder(dataBuffer).Encode(data)
	return mp.restJSONCall("POST", resource, dataBuffer, true)
}

// PUT HTTP method wrapper for authentication (JSON)
func (mp *MP) put(resource string, data interface{}) (*http.Response, error) {
	dataBuffer := new(bytes.Buffer)
	json.NewEncoder(dataBuffer).Encode(data)
	return mp.restJSONCall("PUT", resource, dataBuffer, true)
}

// generic API REST call with Mercado Pago preferences
func (mp *MP) restFormCall(method string, resource string, values url.Values, auth bool) (*http.Response, error) {
	// Build resource URL
	u, err := url.ParseRequestURI(APIBaseURL)
	if err != nil {
		return nil, err
	}
	u.Path = resource
	urlStr := fmt.Sprintf("%v", u)
	// If authed method, then add a form entry for the MP Access Token
	if auth {
		if mp.AccessToken == "" {
			err := mp.obtainAccessToken()
			if err != nil {
				return nil, err
			}
		}
		values.Add("access_token", mp.AccessToken)
	}
	// Create HTTP Request
	r, _ := http.NewRequest(method, urlStr, bytes.NewBufferString(values.Encode()))
	r.Header.Add("Content-Length", strconv.Itoa(len(values.Encode())))
	r.Header.Set("User-Agent", MPUserAgent)
	r.Header.Add("accept", MIMEJSON)
	r.Header.Add("content-type", MIMEForm)
	client := &http.Client{}
	return client.Do(r)
}

// generic API REST call with Mercado Pago preferences
func (mp *MP) restJSONCall(method string, resource string, data *bytes.Buffer, auth bool) (*http.Response, error) {
	// Build resource URL
	u, err := url.ParseRequestURI(APIBaseURL)
	if err != nil {
		return nil, err
	}
	u.Path = resource
	urlStr := fmt.Sprintf("%v", u)
	// If authed method, then append the MP Access Token to the resource URL
	if auth {
		if mp.AccessToken == "" {
			err := mp.obtainAccessToken()
			if err != nil {
				return nil, err
			}
		}
		urlStr += "?access_token=" + mp.AccessToken
	}
	// Create HTTP Request
	r, _ := http.NewRequest(method, urlStr, data)
	r.Header.Add("Content-Length", strconv.Itoa(data.Len()))
	r.Header.Set("User-Agent", MPUserAgent)
	r.Header.Set("Content-Type", MIMEJSON)
	r.Header.Add("Accept", MIMEJSON)
	client := &http.Client{}
	resp, respErr := client.Do(r)
	return resp, respErr
}
