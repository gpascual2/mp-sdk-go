package mercadopago

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

// CreatePayment Creates a payment
//	@param preference
//	@return json
func (mp *MP) CreatePayment(payment *Payment) (*Payment, error) {
	res := &Payment{}
	uri := fmt.Sprintf("/v1/payments")
	// Call POST method
	r, err := mp.post(uri, payment, 2)
	if err != nil {
		return nil, err
	}
	// Check response status
	if r.StatusCode != 200 && r.StatusCode != 201 {
		return nil, fmt.Errorf("Bad status received in HTTP response: %v", r.Status)
	}
	// Read response Body and unmarshall content
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetPayment Get a payment by ID
//	@param id
//	@return json
func (mp *MP) GetPayment(id string) (*Payment, error) {
	res := &Payment{}
	uri := fmt.Sprintf("/v1/payments/%v", id)
	// Call GET method
	r, err := mp.jget(uri, nil, 2)
	if err != nil {
		return nil, err
	}
	// Check response status
	if r.StatusCode != 200 {
		return nil, fmt.Errorf("Bad status received in HTTP response: %v", r.Status)
	}
	// Read response Body and unmarshall content
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetPaymentsByRef Get a payment by External Reference
//	@param id
//	@return json
func (mp *MP) GetPaymentsByRef(externalReference string) (*PaymentSearch, error) {
	res := &PaymentSearch{}
	uri := fmt.Sprintf("/v1/payments/search")
	data := &url.Values{}
	data.Add("external_reference", externalReference)
	// Call GET method
	r, err := mp.get(uri, data, 2)
	if err != nil {
		return nil, err
	}
	// Check response status
	if r.StatusCode != 200 {
		return nil, fmt.Errorf("Bad status received in HTTP response: %v", r.Status)
	}
	// Read response Body and unmarshall content
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	return res, nil
}

// PaymentsSearch Search for payments using a filter set
//	@param filters in url.Values object
//	@return json
func (mp *MP) PaymentsSearch(filters *url.Values) (*PaymentSearch, error) {
	res := &PaymentSearch{}
	uri := fmt.Sprintf("/v1/payments/search")
	// Call GET method
	r, err := mp.get(uri, filters, 2)
	if err != nil {
		return nil, err
	}
	// Check response status
	if r.StatusCode != 200 {
		return nil, fmt.Errorf("Bad status received in HTTP response: %v", r.Status)
	}
	// Read response Body and unmarshall content
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	return res, nil
}
