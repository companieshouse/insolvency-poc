package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Register(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/healthcheck", healthCheck).Methods(http.MethodGet).Name("healthcheck")

	// Create a router for attachments
	attachmentRouter := mainRouter.PathPrefix("/transactions/{transaction_id}/insolvency/attachment").Subrouter()

	attachmentRouter.HandleFunc("", HandlePostAttachment).Methods(http.MethodPost).Name("add-attachment")
	attachmentRouter.HandleFunc("/{attachment_id}", HandleGetAttachment).Methods(http.MethodGet).Name("get-attachment")
	attachmentRouter.HandleFunc("/{attachment_id}", HandleDeleteAttachment).Methods(http.MethodDelete).Name("delete-attachment")
	attachmentRouter.HandleFunc("/{attachment_id}/download", HandleDownloadAttachment).Methods(http.MethodGet).Name("download_attachment")

	// Create a router for appointments of liquidation
	liquidationRouter := mainRouter.PathPrefix("/transactions/{transaction_id}/insolvency/liquidation").Subrouter()

	liquidationRouter.HandleFunc("/appointment", HandlePostLiquidator).Methods(http.MethodPost).Name("add-insolvency-practitioner")
	liquidationRouter.HandleFunc("/appointment/{liquidator_id}", HandleDeleteLiquidator).Methods(http.MethodDelete).Name("delete-insolvency-practitioner")
	liquidationRouter.HandleFunc("/statement-of-affairs", HandlePostStatementAffairs).Methods(http.MethodPost).Name("add-statement-of-affairs")
	liquidationRouter.HandleFunc("/statement-of-affairs/{statement_id}", HandleDeleteStatementAffairs).Methods(http.MethodDelete).Name("delete-statement-of-affairs")
}

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
