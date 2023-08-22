package handler

import (
	"net/http"
	"fmt"

	"mahasanbkk-webscraper/pkg/config"
	"mahasanbkk-webscraper/pkg/session"

	"github.com/gorilla/mux"
)

func ServiceStatusHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
	service := vars["service"]
	if service == "discord" {
		w.WriteHeader(http.StatusOK)
    	fmt.Fprintf(w, "Discord webhook status: " + config.ConfigData.DiscordStatus)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Service "+service+ " not found")
	}
}

func ServiceActionHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
	service := vars["service"]
	action := vars["action"]
	if service == "discord" {
		w.WriteHeader(http.StatusOK)
		if action == "start" {
			config.ConfigData.DiscordStatus = "on"
			fmt.Fprintf(w, "Discord webhook " + action + "ed\n")
		} else if action == "stop" {
			config.ConfigData.DiscordStatus = "off"
			fmt.Fprintf(w, "Discord webhook " + action + "ped\n")
		}
		session.ResetPreviousTableIDs()
    	fmt.Fprintf(w, "Current discord service status: " + config.ConfigData.DiscordStatus)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Service "+service+ " not found")
	}
}