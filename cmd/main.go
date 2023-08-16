package main

import (
	"mahasanbkk-webscraper/pkg/config"
	"mahasanbkk-webscraper/src/discord"
	"mahasanbkk-webscraper/src/webscraper"

	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	cron "github.com/robfig/cron/v3"
)

var (
	discordSession *discordgo.Session
	err            error
)

func main() {
	config.Load(".")

	// Create new Discord Session
	discordSession, err = discordgo.New("Bot " + config.ConfigData.DiscordToken)
	if err != nil {
		log.Fatalln("cannot create discord session. err ", err)
	}

	ApplyScheduler()
	DiscordWebhook()
	ApplyRouter()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}

func ApplyRouter() {
	http.HandleFunc("/mahasan-bot-status", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "Hello, your mahasan bot dev is up 🚀")
	})

	fmt.Println("🚀 Server is up at port 3000 🚀")
	log.Fatal(http.ListenAndServe(":"+config.ConfigData.ServerPort, nil))
}

func WebScraper() {
	fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " WebScraper is running.\n")
	webscraper.DoMagic(discordSession, false)
}

func ApplyScheduler() {
	scheduler := cron.New()
	defer scheduler.Stop()

	scheduler.AddFunc("*/1 * * * *", WebScraper)
	go scheduler.Start()
}

func DiscordWebhook() {
	fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " Discord Webhook is running.\n")
	discord.AvailableCommand(discordSession)
}
