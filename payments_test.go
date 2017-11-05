package mercadopago_test

import (
	"fmt"
	"testing"

	"github.com/gpascual2/mp-sdk-go"
)

// TestCreatePreference - A checkout preference should be obtained from MercadoPago API
func TestCreatePreference(t *testing.T) {
	fmt.Println("mp_test : CreatePreference")

	pref := &mercadopago.Preference{
		ExternalReference: "ExRef",
	}
	pref.Payer.Name = "Jon"
	pref.Payer.Surname = "Snow"
	pref.Payer.Email = "jonsnow@winterfell.north"
	pref.Items = append(pref.Items, mercadopago.Item{
		ID:         "Item1_ID",
		Title:      "Item1_title",
		Quantity:   1,
		CurrencyID: "ARS",
		UnitPrice:  10.2,
	})

	prefRes, err := mp.CreatePreference(pref)
	if err != nil {
		t.Fatalf("Error creating a checkout preference: %v", err)
	}
	if prefRes.InitPoint == "" {
		t.Errorf("Expected InitPoint to contain a value and is empty")
	}
	if prefRes.Items[0] != pref.Items[0] {
		t.Errorf("Expected Preference Item to equal the requested one. Sent: %v / Got: %v", pref.Items[0], prefRes.Items[0])
	}
}
