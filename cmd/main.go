package main

import (
	"mahasanbkk-webscraper/pkg/config"
	"mahasanbkk-webscraper/src/discord"
	"mahasanbkk-webscraper/src/line"
	"mahasanbkk-webscraper/src/webscraper"

	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"net/http"
	cron "github.com/robfig/cron/v3"
)

func main() {
	config.Load(".")
	config.InitSession()
	
	
	wg := &sync.WaitGroup{}
	wg.Add(1)

	AddScheduler()
	DiscordWebhook()
	ApplyRouter()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	wg.Wait()
}

func ApplyRouter() {
	http.HandleFunc("/mahasan-bot-status", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "Hello, your mahasan bot dev is up 🚀")
	})
	http.HandleFunc("/line-webhook", func(writer http.ResponseWriter, req *http.Request) {
		LineWebhook(writer, req)
	})

	fmt.Println("🚀 Server is up at port 3000 🚀")
	log.Fatal(http.ListenAndServe(":"+ config.ConfigData.ServerPort, nil))
}

func AddScheduler() {
	scheduler := cron.New()
	defer scheduler.Stop()

	if config.ConfigData.ServerEnv == "local" {
		scheduler.AddFunc("*/1 * * * *", WebScraper)
	} else {
		scheduler.AddFunc("*/2 * * * *", WebScraper)
	}

	go scheduler.Start()
}

func WebScraper() {
	fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " WebScraper is running.\n")
	webscraper.DoMagic(false)
}

func DiscordWebhook() {
	if config.ConfigData.DiscordStatus == "on" {	
		fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " Discord webhook is running.\n")
		discord.AvailableCommand()
	}
}

func LineWebhook(writer http.ResponseWriter, req *http.Request) {
	if config.ConfigData.LineStatus == "on" {	
		fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " Line webhook is running.\n")
		line.Webhook(writer, req)
	}
}
