package routes

import (
	"atm-teste/controller"
	"atm-teste/erro"
	"encoding/json"
	"log"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		param := r.URL.Query().Get("saque")
		if param == "" {
			warn, _ := json.Marshal(erro.Warn{"Por gentileza, utilize o endereço correto para a requisição"})
			w.Write(warn)
			return
		}
		response := controller.Saque(param)
		responseJson, _ := json.Marshal(response)

		w.Write(responseJson)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
