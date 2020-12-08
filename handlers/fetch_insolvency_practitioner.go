package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/companieshouse/insolvency-poc/models"
	"net/http"
)

func HandleGetInsolvencyPractitioner(w http.ResponseWriter, req *http.Request) {
	//vars := mux.Vars(req)
	//id := vars["insolvency_practitioner_id"]
	insolvencyPractitioner := models.InsolvencyPractitioner{
		Forename: "Joe",
		Surname:  "Bloggs",
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(insolvencyPractitioner)
	if err != nil {
		fmt.Println(fmt.Errorf("error encoding insolvency practitioner to response: [%v]", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}