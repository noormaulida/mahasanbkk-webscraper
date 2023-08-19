package main

import (
	"mahasanbkk-webscraper/pkg/config"
	"mahasanbkk-webscraper/pkg/session"
	"mahasanbkk-webscraper/src/discord"
	"mahasanbkk-webscraper/src/webscraper"

	"fmt"
	// "log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	cron "github.com/robfig/cron/v3"
)

func main() {
	config.Load(".")
	session.InitSession()

	ApplyScheduler()
	DiscordWebhook()
	ApplyRouter()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}

func ApplyRouter() {
	r := mux.NewRouter()
    r.HandleFunc("/mahasan-bot/status", GlobalStatusHandler)
    r.HandleFunc("/mahasan-bot/status/{type}", ServiceStatusHandler)
    r.HandleFunc("/mahasan-bot/auto-booking/{id}", AutoBookingHandler)
	fmt.Println("🚀 Server is up at port 3000 🚀")
	http.ListenAndServe(":"+config.ConfigData.ServerPort, r)
}

func GlobalStatusHandler(w http.ResponseWriter, _ *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Hello, your mahasan bot dev is up 🚀")
}

func ServiceStatusHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
	service := vars["type"]
	if service == "discord" {
		w.WriteHeader(http.StatusOK)
    	fmt.Fprintf(w, "Discord webhook status: " + config.ConfigData.DiscordStatus)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Service "+service+ " not found")
	}
}

func AutoBookingHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
	tableId := vars["id"]
	resp := webscraper.AutoBooking(tableId)
	w.WriteHeader(http.StatusOK)
	if resp.StatusCode == http.StatusOK {
		fmt.Fprintf(w, "Successfully Auto-Booking Table ID: %v\n", tableId)
	} else {
		fmt.Fprintf(w, "Sorry, Auto-Booking Table ID: %v is Failed\nTry again in a minute.", tableId)
	}
}

func WebScraper() {
	fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " WebScraper is running.\n")
	webscraper.DoMagic(false)
}

func ApplyScheduler() {
	scheduler := cron.New()
	defer scheduler.Stop()

	scheduler.AddFunc("*/1 * * * *", WebScraper)
	go scheduler.Start()
}

func DiscordWebhook() {
	if config.ConfigData.DiscordStatus == "on" {
		fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " Discord Webhook is running.\n")
		discord.Webhook()
	}
}
