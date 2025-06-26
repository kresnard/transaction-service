package checkout

import (
	"net/http"
	"transaction-service/config"
	"transaction-service/internal/usecase/checkout"
	"transaction-service/pkg/logger"

	"github.com/gorilla/mux"
)

type CheckoutRoutes struct {
	l *logger.Logger
	cfg *config.Config
	cu checkout.ICheckoutUsecase
}

func NewCheckoutRoutes(r *mux.Router, l *logger.Logger, cfg *config.Config, cu checkout.ICheckoutUsecase) {
	c := &CheckoutRoutes{l, cfg, cu}

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("OK Service Running..."))
	}).Methods(http.MethodGet)

	group := r.PathPrefix("/v1/checkout").Subrouter()
	group.HandleFunc("/order", c.Checkout).Methods(http.MethodPost)
}