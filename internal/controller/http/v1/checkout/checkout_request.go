package checkout

import "errors"

type CheckoutRequest struct {
	Items []string `json:"items"`
}

func (data CheckoutRequest) validationCheckoutRequest() (err error) {
	if len(data.Items) < 1 {
		return errors.New("order item can't empty")
	}

	return nil
}