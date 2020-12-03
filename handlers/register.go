package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Register(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/healthcheck", healthCheck).Methods(http.MethodGet).Name("healthcheck")
	mainRouter.HandleFunc("/insolvency-practitioners/{insolvency-practitioner-id}", HandleGetInsolvencyPractitioner).Methods(http.MethodGet).Name("get-insolvency-practitioner")
}

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
