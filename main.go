package main

import (
	"fmt"
	"github.com/companieshouse/chs.go/log"
	"github.com/companieshouse/insolvency-poc/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	namespace := "insolvency-api"
	log.Namespace = namespace

	mainRouter := mux.NewRouter()
	handlers.Register(mainRouter)

	fmt.Println("Starting Insolvency API")

	err := http.ListenAndServe(":5010", mainRouter)
	if err != nil {
		fmt.Println(err)
	}
}