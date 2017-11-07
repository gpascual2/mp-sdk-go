package mercadopago

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// CreatePreference Creates a checkout preference
//	@param preference
//	@return json
func (mp *MP) CreatePreference(preference *Preference) (*Preference, error) {
	res := &Preference{}
	uri := fmt.Sprintf("/checkout/preferences")
	// Call POST method
	r, err := mp.post(uri, preference, 1)
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

// GetPreference Get a checkout preference
//	@param id
//	@return json
func (mp *MP) GetPreference(id string) (*Preference, error) {
	res := &Preference{}
	uri := fmt.Sprintf("/checkout/preferences/%v", id)
	// Call GET method
	r, err := mp.get(uri, nil, 1)
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
