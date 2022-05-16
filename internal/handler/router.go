package handler

import (
	"github.com/gorilla/mux"
	"github.hdfcbank.com/HDFCBANK/mb-microservices-utils/middlewares"
)

func (e *Controller) Router() *mux.Router {
	r := mux.NewRouter()
	// r.Use(middlewares.RecoverHandler(e.Logger))
	r.Use(middlewares.Loggerhandler(e.Logger))
	r.HandleFunc("/{device_id}/send", e.GetTokenHandler).Methods("POST")

	return r
}
