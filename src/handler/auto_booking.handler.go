package handler

import (
	"net/http"
	"fmt"

	"mahasanbkk-webscraper/pkg/config"
	"mahasanbkk-webscraper/src/webscraper"

	"github.com/gorilla/mux"
)

func AutoBookingHandler(w http.ResponseWriter, r *http.Request) {
	if config.ConfigData.ServerEnv == "local" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	} else {
		w.Header().Set("Access-Control-Allow-Origin", config.ConfigData.ServerHost)
	}
    vars := mux.Vars(r)
	tableId := vars["id"]
	resp := webscraper.AutoBooking(tableId)
	if resp.StatusCode == http.StatusOK {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Successfully Auto-Booking Table ID: %v\n", tableId)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Sorry, Auto-Booking Table ID: %v is failed.\nTry again in a minute.", tableId)
	}
}

