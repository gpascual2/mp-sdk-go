package mercadopago

// Payment is the data struct for payment MP API
type Payment struct {
	ID               int    `json:"id,omitempty"`
	DateCreated      string `json:"date_created,omitempty"`
	DateApproved     string `json:"date_approved,omitempty"`
	DateLastUpdated  string `json:"date_last_updated,omitempty"`
	MoneyReleaseDate string `json:"money_release_date,omitempty"`
	CollectorID      int    `json:"collector_id,omitempty"`
	OperationType    string `json:"operation_type,omitempty"`
	Payer            struct {
		EntityType string `json:"entity_type,omitempty"`
		Type       string `json:"type,omitempty"`
		// ID             string `json:"id,omitempty"`   // Issue on MP API - sometimes this is string, other is int :(
		Email          string `json:"email,omitempty"`
		Identification struct {
			Type   string `json:"type,omitempty"`
			Number string `json:"number,omitempty"`
		} `json:"identification,omitempty"`
		Phone struct {
			AreaCode  string `json:"area_code,omitempty"`
			Number    string `json:"number,omitempty"`
			Extension string `json:"extension,omitempty"`
		} `json:"phone,omitempty"`
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
	} `json:"payer,omitempty"`
	BinaryMode bool `json:"binary_mode,omitempty"`
	LiveMode   bool `json:"live_mode,omitempty"`
	Order      struct {
		Type string `json:"type,omitempty"`
		// ID   string `json:"id,omitempty"`        // Issue on MP API - sometimes this is string, other is int :(
	} `json:"order,omitempty"`
	ExternalReference         string  `json:"external_reference,omitempty"`
	Description               string  `json:"description,omitempty"`
	CurrencyID                string  `json:"currency_id,omitempty"`
	TransactionAmount         float32 `json:"transaction_amount,omitempty"`
	TransactionAmountRefunded float32 `json:"transaction_amount_refunded,omitempty"`
	CouponAmount              float32 `json:"coupon_amount,omitempty"`
	CampaignID                int     `json:"campaign_id,omitempty"`
	CouponCode                string  `json:"coupon_code,omitempty"`
	TransactionDetails        struct {
		FinancialInstitution   string  `json:"financial_institution,omitempty"`
		NetReceivedAmount      float32 `json:"net_received_amount,omitempty"`
		TotalPaidAmount        float32 `json:"total_paid_amount,omitempty"`
		InstallmentAmount      float32 `json:"installment_amount,omitempty"`
		OverpaidAmount         float32 `json:"overpaid_amount,omitempty"`
		PaymentMethodReference string  `json:"payment_method_reference,omitempty"`
	} `json:"transaction_details,omitempty"`
	FeeDetails []struct {
		Type     string  `json:"type,omitempty"`
		FeePayer string  `json:"fee_payer,omitempty"`
		Amount   float32 `json:"amount,omitempty"`
	} `json:"fee_details,omitempty"`
	DifferentialPricingID int     `json:"differential_pricing_id,omitempty"`
	ApplicationFee        float32 `json:"application_fee,omitempty"`
	Status                string  `json:"status,omitempty"`
	StatusDetail          string  `json:"status_detail,omitempty"`
	Capture               bool    `json:"capture,omitempty"`
	Captured              bool    `json:"captured,omitempty"`
	CallForAuthorizeID    string  `json:"call_for_authorize_id,omitempty"`
	PaymentMethodID       string  `json:"payment_method_id,omitempty"`
	// IssuerID              string     `json:"issuer_id,omitempty"`        // Issue on MP API - sometimes this is string, other is int :(
	PaymentTypeID string `json:"payment_type_id,omitempty"`
	Token         string `json:"token,omitempty"`
	Card          struct {
		ID              int    `json:"id,omitempty"`
		LastFourDigits  string `json:"last_four_digits,omitempty"`
		FirstSixDigits  string `json:"first_six_digits,omitempty"`
		ExpirationYear  int    `json:"expiration_year,omitempty"`
		ExpirationMonth int    `json:"expiration_month,omitempty"`
		DateCreated     string `json:"date_created,omitempty"`
		DateLastUpdated string `json:"date_last_updated,omitempty"`
		Cardholder      struct {
			Name           string `json:"name,omitempty"`
			Identification struct {
				Type   string `json:"type,omitempty"`
				Number string `json:"number,omitempty"`
			} `json:"identification,omitempty"`
		} `json:"cardholder,omitempty"`
	} `json:"card,omitempty"`
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
	Installments        int    `json:"installments,omitempty"`
	NotificationURL     string `json:"notification_url,omitempty"`
	CallbackURL         string `json:"callback_url,omitempty"`
	Refunds             []struct {
		ID        int     `json:"id,omitempty"`
		PaymentID int     `json:"payment_id,omitempty"`
		Amount    float32 `json:"amount,omitempty"`
		Source    struct {
			ID   string `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"source,omitempty"`
	} `json:"refunds,omitempty"`
	AdditionalInfo struct {
		IPAddress string `json:"ip_address,omitempty"`
		Items     []Item `json:"items,omitempty"`
		Payer     struct {
			FirstName string `json:"first_name,omitempty"`
			LastName  string `json:"last_name,omitempty"`
			Phone     struct {
				AreaCode string `json:"area_code,omitempty"`
				Number   string `json:"number,omitempty"`
			} `json:"phone,omitempty"`
			Address struct {
				ZipCode      string `json:"zip_code,omitempty"`
				StreetName   string `json:"street_name,omitempty"`
				StreetNumber int    `json:"street_number,omitempty"`
			} `json:"address,omitempty"`
			RegistrationDate string `json:"registration_date,omitempty"`
		} `json:"payer,omitempty"`
		Shipments struct {
			ReceiverAddress struct {
				ZipCode      string `json:"zip_code,omitempty"`
				StreetName   string `json:"street_name,omitempty"`
				StreetNumber int    `json:"street_number,omitempty"`
				Floor        string `json:"floor,omitempty"`
				Apartment    string `json:"apartment,omitempty"`
			} `json:"receiver_address,omitempty"`
		} `json:"shipments,omitempty"`
		Barcode struct {
			Type    string `json:"type,omitempty"`
			Content string `json:"content,omitempty"`
			Width   int    `json:"width,omitempty"`
			Height  int    `json:"height,omitempty"`
		} `json:"barcode,omitempty"`
	} `json:"additional_info,omitempty"`
}

// PaymentSearch is the data struct for payment MP API
type PaymentSearch struct {
	Paging struct {
		Total  int `json:"total,omitempty"`
		Limit  int `json:"limit,omitempty"`
		Offset int `json:"offset,omitempty"`
	} `json:"paging,omitempty"`
	Results []Payment `json:"results,omitempty"`
}
