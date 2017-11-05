package mercadopago

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Preference is the data struct for payment checkouts
type Preference struct {
	Items []Item `json:"items,omitempty"`
	Payer struct {
		Name    string `json:"name,omitempty"`
		Surname string `json:"surname,omitempty"`
		Email   string `json:"email,omitempty"`
		Phone   struct {
			AreaCode string `json:"area_code,omitempty"`
			Number   string `json:"number,omitempty"`
		} `json:"phone,omitempty"`
		Identification struct {
			Type   string `json:"type,omitempty"`
			Number string `json:"number,omitempty"`
		} `json:"identification,omitempty"`
		Address struct {
			ZipCode string `json:"zip_code,omitempty"`
			Street  string `json:"street,omitempty"`
			Number  int    `json:"number,omitempty"`
		} `json:"address,omitempty"`
		DateCreated string `json:"date_created,omitempty"`
	} `json:"payer,omitempty"`
	PaymentMethods struct {
		ExcludedPaymentMethods []struct {
			ID string `json:"id,omitempty"`
		} `json:"excluded_payment_methods,omitempty"`
		ExcludedPaymentTypes []struct {
			ID string `json:"id,omitempty"`
		} `json:"excluded_payment_types,omitempty"`
		DefaultPaymentMethodID string `json:"default_payment_method_id,omitempty"`
		Installments           int    `json:"installments,omitempty"`
		DefaultInstallments    int    `json:"default_installments,omitempty"`
	} `json:"payment_methods,omitempty"`
	BackUrls struct {
		Success string `json:"success,omitempty"`
		Pending string `json:"pending,omitempty"`
		Failure string `json:"failure,omitempty"`
	} `json:"back_urls,omitempty"`
	NotificationURL   string `json:"notification_url,omitempty"`
	ID                string `json:"id,omitempty"`
	InitPoint         string `json:"init_point,omitempty"`
	SandboxInitPoint  string `json:"sandbox_init_point,omitempty"`
	DateCreated       string `json:"date_created,omitempty"`
	OperationType     string `json:"operation_type,omitempty"`
	AdditionalInfo    string `json:"additional_info,omitempty"`
	AutoReturn        string `json:"auto_return,omitempty"`
	ExternalReference string `json:"esternal_reference,omitempty"`
	CollectorID       int    `json:"collector_id,omitempty"`
	ClientID          string `json:"client_id,omitempty"`
}

// Item information
type Item struct {
	ID          string  `json:"id,omitempty"`
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	PictureURL  string  `json:"picture_url,omitempty"`
	CategoryID  string  `json:"category_id,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
	CurrencyID  string  `json:"currency_id,omitempty"`
	UnitPrice   float32 `json:"unit_price,omitempty"`
}

// CreatePreference Creates a checkout preference
func (mp *MP) CreatePreference(preference *Preference) (*Preference, error) {
	res := &Preference{}
	uri := fmt.Sprintf("/checkout/preferences")
	// Call POST method
	r, err := mp.post(uri, preference)
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
