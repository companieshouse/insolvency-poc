package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/companieshouse/insolvency-poc/models"
	"net/http"
)

func HandlePostLiquidator(w http.ResponseWriter, req *http.Request) {

	appointLiquidatorResponse := models.AppointLiquidatorResponse{
		Id: "123456789",
		Links:  models.AppointLiquidatorResponseLinks{Self: "/transactions/987654321/insolvency/liquidator/123456789"},
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(appointLiquidatorResponse)
	if err != nil {
		fmt.Println(fmt.Errorf("error encoding insolvency practitioner to response: [%v]", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func HandleDeleteLiquidator(w http.ResponseWriter, req *http.Request) {

	// Delete liquidator

	w.WriteHeader(http.StatusOK)
}

func HandlePostStatementAffairs(w http.ResponseWriter, req *http.Request) {

	appointStatementAffairsResponse := models.NoticeStatementAffairsResponse{
		Id:    "3fa85f64-5717-4562-b3fc-2c963f66afa6",
		Links: models.NoticeStatementAffairsResponseLinks{Self: "/transactions/987654321/insolvency/statement-of-affairs/3fa85f64-5717-4562-b3fc-2c963f66afa6"},
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(appointStatementAffairsResponse)
	if err != nil {
		fmt.Println(fmt.Errorf("error encoding appointment of statement of affairs to response: [%v]", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func HandleDeleteStatementAffairs(w http.ResponseWriter, req *http.Request) {

	// Delete statement of affairs

	w.WriteHeader(http.StatusOK)
}
