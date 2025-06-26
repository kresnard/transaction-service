package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"transaction-service/config"
	v1 "transaction-service/internal/controller/http/v1"
	repo "transaction-service/internal/repository"
	"transaction-service/internal/usecase/checkout"
	"transaction-service/pkg/httpserver"
	"transaction-service/pkg/logger"
	"transaction-service/pkg/mysql"

	"github.com/gorilla/mux"
)

func Run(cfg *config.Config) {
	fmt.Println("Running service")

	var err error
	l := logger.New(cfg)

	//MYSQL
	db := mysql.New(cfg, l)

	//Repository
	repo := repo.NewRepository(db, l, cfg)

	//Usecase
	checkoutUsecase := checkout.NewUsecase(repo, l, cfg)
	

	//HTTP Server
	handler := mux.NewRouter()
	v1.NewRouter(handler, l,cfg, checkoutUsecase)
	httpServer := httpserver.New(handler, cfg, httpserver.Port(cfg.HTTPServer.Port))

	//Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Println(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	//Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Println(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
