package routes

import (
	"atm-teste/controller"
	"encoding/json"
	"log"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("saque")
		response := controller.Saque(param)
		responseJson, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJson)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
