package mercadopago

// MPError is the type returned by the lib in case of errors
type MPError struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Stack   string `json:"stack"`
	Status  int    `json:"status"`
}

func newMercadoPagoError(message string, status int) *MPError {
	mpe := &MPError{}
	mpe.Name = "MercadoPagoError"
	if len(message) > 0 {
		mpe.Message = message
	} else {
		mpe.Message = "MercadoPago Unknown error"
	}
	if status > 0 {
		mpe.Status = status
	} else {
		mpe.Status = 500
	}
	return mpe
}
