package v1

import (
	"transaction-service/config"
	checkoutHttp "transaction-service/internal/controller/http/v1/checkout"
	checkoutUsecase "transaction-service/internal/usecase/checkout"
	"transaction-service/pkg/logger"

	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router, l *logger.Logger, cfg *config.Config, cu checkoutUsecase.ICheckoutUsecase) {
	{
		checkoutHttp.NewCheckoutRoutes(r,l, cfg, cu)
	}
}