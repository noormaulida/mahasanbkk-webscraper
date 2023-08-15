package main

import (
	"mahasanbkk-webscraper/pkg/config"
	"mahasanbkk-webscraper/src/discord"
	"mahasanbkk-webscraper/src/webscraper"

	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
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

	wg := &sync.WaitGroup{}
	wg.Add(1)
	scheduler := cron.New()
	defer scheduler.Stop()

	if config.ConfigData.ServerEnv == "local" {
		scheduler.AddFunc("*/1 * * * *", WebScraper)
	} else {
		scheduler.AddFunc("*/2 * * * *", WebScraper)
	}
	go scheduler.Start()
	DiscordCommand()

	app := fiber.New()
	app.Get("/mahasan-bot-status", func(c *fiber.Ctx) error {
		return c.SendString("Hello, your mahasan bot dev is up 🚀")
	})
	log.Fatal(app.Listen(":" + config.ConfigData.ServerPort))

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	wg.Wait()
}

func WebScraper() {
	fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " WebScraper is running.\n")
	webscraper.DoMagic(discordSession, false)
}

func DiscordCommand() {
	fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " Discord Command is running.\n")
	discord.AvailableCommand(discordSession)
}
