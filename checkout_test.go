package mercadopago_test

import (
	"fmt"
	"testing"

	"github.com/gpascual2/mp-sdk-go"
)

var prefBase *mercadopago.Preference
var prefCreated *mercadopago.Preference

func init() {
	prefBase = &mercadopago.Preference{
		ExternalReference: "ExRef",
	}
	prefBase.Payer.Name = "Jon"
	prefBase.Payer.Surname = "Snow"
	prefBase.Payer.Email = "jonsnow@winterfell.north"
	prefBase.Items = append(prefBase.Items, mercadopago.Item{
		ID:         "Item1_ID",
		Title:      "Item1_title",
		Quantity:   1,
		CurrencyID: "ARS",
		UnitPrice:  10.2,
	})
}

// TestCreatePreference - A checkout preference should be obtained from MercadoPago API
func TestCreatePreference(t *testing.T) {
	fmt.Println("mp_test : CreatePreference")

	prefCreated, err := mp.CreatePreference(prefBase)
	if err != nil {
		t.Fatalf("Error creating a checkout preference: %v", err)
	}
	if prefCreated.InitPoint == "" {
		t.Errorf("Expected InitPoint to contain a value and is empty")
	}
	if prefCreated.Items[0] != prefBase.Items[0] {
		t.Errorf("Expected Preference Item to equal the requested one. Sent: %v / Got: %v", prefBase.Items[0], prefCreated.Items[0])
	}
}

// TestGetPreference - A checkout preference should be obtained from MercadoPago API
func TestGetPreference(t *testing.T) {
	fmt.Println("mp_test : GetPreference")

	prefGet, err := mp.GetPreference(prefCreated.ID)
	if err != nil {
		t.Fatalf("Error getting the checkout preference: %v", err)
	}
	if prefGet.Items[0] != prefBase.Items[0] {
		t.Errorf("Expected Preference Item to equal the requested one. Sent: %v / Got: %v", prefBase.Items[0], prefGet.Items[0])
	}
}
