package main

import (
	"mahasanbkk-webscraper/pkg/config"
	"mahasanbkk-webscraper/src/crontab"
	"mahasanbkk-webscraper/src/handler"
	"mahasanbkk-webscraper/src/service"
	"mahasanbkk-webscraper/pkg/session"

	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	
)

func main() {
	config.Load(".")
	session.InitSession()

	crontab.ApplyScheduler()
	service.DiscordWebhook()
	ApplyRouter()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}

func ApplyRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/mahasan-bot/status", handler.GlobalStatusHandler).Methods("GET")
	r.HandleFunc("/mahasan-bot/status/{service}", handler.ServiceStatusHandler).Methods("GET")
	r.HandleFunc("/mahasan-bot/auto-booking/{id}", handler.AutoBookingHandler).Methods("POST")
	r.HandleFunc("/mahasan-bot/{service}/{action}", handler.ServiceActionHandler).Methods("GET")

	fmt.Println("🚀 Server is up at port 3000 🚀")
	http.ListenAndServe(":"+config.ConfigData.ServerPort, r)
}

