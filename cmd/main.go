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
		scheduler.AddFunc("*/5 * * * *", WebScraper)
	}
	go scheduler.Start()
	DiscordCommand()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, your bot dev is up 🚀")
	})
	app.Get("/term-and-conditions", func(c *fiber.Ctx) error {
		return c.SendString("These Terms of Use constitute a legally binding agreement made between you, whether personally or on behalf of an entity ('you') and Mahasan, doing business as Mahasan Bot ('Mahasan Bot', 'we', 'us', or 'our'), concerning your access to and use of the Mahasan Bot as well as any other media form, website, media channel, mobile website or mobile application related, linked, or otherwise connected thereto (collectively, the “Bot”). You agree that by accessing the Bot, you have read, understood, and agree to be bound by all of these Terms of Use. IF YOU DO NOT AGREE WITH ALL OF THESE TERMS OF USE, THEN YOU ARE EXPRESSLY PROHIBITED FROM USING THE BOT AND YOU MUST DISCONTINUE USE IMMEDIATELY.")
	})
	app.Get("/privacy-policy", func(c *fiber.Ctx) error {
		return c.SendString("Like many other websites, we also use so-called cookies. Cookies are small text files that are stored on your end device (laptop, tablet, smartphone, etc.) when you visit our website. \nThis gives us certain data such as IP address, browser used and operating system. Cookies cannot be used to start programs or transfer viruses to a computer. Based on the information contained in cookies, we can make navigation easier for you and enable our websites to be displayed correctly.")
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
