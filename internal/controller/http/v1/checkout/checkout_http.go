package checkout

import (
	"encoding/json"
	"net/http"
	"time"
	http_response "transaction-service/internal/controller/response"
	"transaction-service/pkg/logger"
)

var (
	timeNow = time.Since
)


func (c CheckoutRoutes) Order(w http.ResponseWriter, r *http.Request) {
	checkout := OrderRequest{}

	if err := json.NewDecoder(r.Body).Decode(&checkout); err != nil {
		c.l.CreateLog(&logger.Log{
			Event:			"HTTP|CHECKOUT"+"|Order|DECODE",
			StatusCode:		http.StatusBadRequest,
			ResponseTime: 	timeNow(time.Now()),
			Method: 		r.Method,
			Request: 		checkout,
			Message: 		"error decode",
		}, logger.LVL_ERROR)
		http_response.HttpErrorResponse(w, false, http.StatusBadRequest, "400", err.Error())
		return
	}

	err := checkout.validationOrderRequest()
	if err != nil {
		c.l.CreateLog(&logger.Log{
			Event:			"HTTP|CHECKOUT"+"|Order|VALIDATION",
			StatusCode:		http.StatusBadRequest,
			ResponseTime: 	timeNow(time.Now()),
			Method: 		r.Method,
			Request: 		checkout,
			Message: 		"error validation",
		}, logger.LVL_ERROR)
		http_response.HttpErrorResponse(w, false, http.StatusBadRequest, "400", err.Error())
		return
	}

	res, err := c.cu.Order(r.Context(), checkout.Items)
	if err != nil {
		http_response.HttpErrorResponse(w, false, http.StatusBadRequest, "500", err.Error())
		return
	}

	http_response.HttpSuccessResponse(w, true, http.StatusCreated, "201", "success create checkout", res)

}