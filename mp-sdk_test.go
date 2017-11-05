package mercadopago_test

import (
	"fmt"
	"testing"

	"github.com/gpascual2/mp-sdk-go"
)

var mp mercadopago.MP

// Test - An instance of MP should be created
func TestMPInstance(t *testing.T) {
	fmt.Println("mp_test : MPInstance")
	mp = mercadopago.NewMP(TestClientID, TestClientSecret, true)
	if mp.ClientID != TestClientID {
		t.Fatalf("Error creating MP instance. Expected ClientID id to be %s and got %v", TestClientID, mp.ClientID)
	}
	if mp.Sandbox != true {
		t.Fatalf("Error creating MP instance. Expected Sandbox mode to be true")
	}
	if mp.AccessToken != "" {
		t.Errorf("Expected AccessToken to be empty and got %s", mp.AccessToken)
	}
}

// TestAccessToken - An Access Token should be obtained from MercadoPago API
func TestAccessToken(t *testing.T) {
	fmt.Println("mp_test : GetAccessToken")
	at, err := mp.GetAccessToken()
	if err != nil {
		t.Fatalf("Error requesting an Access Token: %v", err)
	}
	if at == "" {
		t.Errorf("Expected AccessToken to contain a value and is empty")
	}
}
