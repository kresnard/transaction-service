package checkout

import (
	"encoding/json"
	"net/http"
	http_response "transaction-service/internal/controller/response"
)

func (c CheckoutRoutes) Checkout(w http.ResponseWriter, r *http.Request) {
	checkout := CheckoutRequest{}

	if err := json.NewDecoder(r.Body).Decode(&checkout); err != nil {
		http_response.HttpErrorResponse(w, false, http.StatusBadRequest, "400", err.Error())
		return
	}

	res, err := c.cu.Checkout(r.Context(), checkout.Items)
	if err != nil {
		http_response.HttpErrorResponse(w, false, http.StatusBadRequest, "500", err.Error())
		return
	}

	http_response.HttpSuccessResponse(w, true, http.StatusCreated, "201", "success create checkout", res)

}