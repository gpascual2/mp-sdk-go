package mercadopago_test

import (
	"fmt"
	"net/url"
	"testing"
)

// TestGetPayment - A payment should be obtained from MercadoPago API
func TestGetPayment(t *testing.T) {
	fmt.Println("mp_test : GetPayment")

	pmt, err := mp.GetPayment("8262805")
	if err != nil {
		t.Fatalf("Error getting the payment: %v", err)
	}
	fmt.Println("Payment: ", pmt)
}

// TestGetPaymentsByRef - A list of payments matching an External Reference should be obtained from MercadoPago API
func TestGetPaymentsByRef(t *testing.T) {
	fmt.Println("mp_test : GetPaymentsByRef")

	pmtSearch, err := mp.GetPaymentsByRef("QBPL-E9Z9-CTGW-LK2")
	if err != nil {
		t.Fatalf("Error getting the payment: %v", err)
	}
	fmt.Println("Payments: ", pmtSearch)
}

// TestPaymentsSearch - A list of payments matching a Filter criteria should be obtained from MercadoPago API
func TestPaymentsSearch(t *testing.T) {
	fmt.Println("mp_test : PaymentsSearch")

	filter := &url.Values{}
	filter.Add("range", "date_created")
	filter.Add("begin_date", "NOW-1MONTH")
	filter.Add("end_date", "NOW")
	filter.Add("status", "approved")

	pmtSearch, err := mp.PaymentsSearch(filter)
	if err != nil {
		t.Fatalf("Error getting the payment: %v", err)
	}
	fmt.Println("Payments: ", pmtSearch)
}
