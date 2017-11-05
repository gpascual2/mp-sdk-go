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
	MPUserAgent string = "MercadoPago Node.js SDK v" + MPVersion
	APIBaseURL  string = "https://api.mercadopago.com"
	MIMEJSON    string = "application/json"
	MIMEForm    string = "application/x-www-form-urlencoded"
)

// MP is the implementation to consume Mercado Pago API services
type MP struct {
	accessToken  string
	clientID     string
	clientSecret string
	sandbox      bool
}

type tokensResponse struct {
	accessToken  string `json:"access_token"`
	refreshToken string `json:"refresh_token"`
	liveMode     bool   `json:"live_mode"`
}

// NewMP returns a new instance of the MP service library
func NewMP(clientID string, clientSecret string, sandbox bool) *MP {
	mp := &MP{}
	mp.accessToken = ""
	mp.clientID = clientID
	mp.clientSecret = clientSecret
	mp.sandbox = sandbox
	return mp
}

func (mp *MP) getAccessToken() error {
	data := url.Values{}
	data.Set("client_id", mp.clientID)
	data.Add("client_secret", mp.clientSecret)
	data.Add("grant_type", "client_credentials")

	r, err := restCall("POST", "/oauth/token", data)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	var tokens tokensResponse
	if err = json.Unmarshal(body, &tokens); err != nil {
		return err
	}
	mp.accessToken = tokens.accessToken
	return nil
}

// GET HTTP method wrapper for authentication
func (mp *MP) get(resource string, values url.Values) (*http.Response, error) {
	if mp.accessToken == "" {
		err := mp.getAccessToken()
		if err != nil {
			return nil, err
		}
	}
	values.Add("access_token", mp.accessToken)
	return restCall("get", resource, values)
}

// POST HTTP method wrapper for authentication
func (mp *MP) post(resource string, values url.Values) (*http.Response, error) {
	if mp.accessToken == "" {
		err := mp.getAccessToken()
		if err != nil {
			return nil, err
		}
	}
	values.Add("access_token", mp.accessToken)
	return restCall("post", resource, values)
}

// PUT HTTP method wrapper for authentication
func (mp *MP) put(resource string, values url.Values) (*http.Response, error) {
	if mp.accessToken == "" {
		err := mp.getAccessToken()
		if err != nil {
			return nil, err
		}
	}
	values.Add("access_token", mp.accessToken)
	return restCall("put", resource, values)
}

// DELETE HTTP method wrapper for authentication
func (mp *MP) delete(resource string, values url.Values) (*http.Response, error) {
	if mp.accessToken == "" {
		err := mp.getAccessToken()
		if err != nil {
			return nil, err
		}
	}
	values.Add("access_token", mp.accessToken)
	return restCall("delete", resource, values)
}

// generic API REST call with Mercado Pago preferences
func restCall(method string, resource string, values url.Values) (*http.Response, error) {
	client := &http.Client{}
	u, err := url.ParseRequestURI(APIBaseURL)
	if err != nil {
		return nil, err
	}
	u.Path = resource
	urlStr := fmt.Sprintf("%v", u)
	r, _ := http.NewRequest(method, urlStr, bytes.NewBufferString(values.Encode()))
	r.Header.Add("Content-Length", strconv.Itoa(len(values.Encode())))
	r.Header.Add("user-agent:", MPUserAgent)
	r.Header.Add("accept:", MIMEJSON)
	r.Header.Add("content-type:", MIMEForm)
	return client.Do(r)
}
