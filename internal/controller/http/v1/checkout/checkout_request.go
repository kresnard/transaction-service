package checkout

import "errors"

type OrderRequest struct {
	Items []string `json:"items"`
}

func (data OrderRequest) validationOrderRequest() (err error) {
	if len(data.Items) < 1 {
		return errors.New("order item can't empty")
	}

	return nil
}