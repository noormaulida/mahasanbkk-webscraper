package main

import (
	"mahasanbkk-webscraper/pkg/config"
	"mahasanbkk-webscraper/src/webscraper"

	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	cron "github.com/robfig/cron/v3"
)

func main() {
	config.Load(".")

	wg := &sync.WaitGroup{}
	wg.Add(1)
	scheduler := cron.New()
	defer scheduler.Stop()

	scheduler.AddFunc("*/15 * * * *", WebScraper)
	go scheduler.Start()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, your bot dev is up 🚀")
	})
	log.Fatal(app.Listen(":3000"))
	wg.Wait()
}

func WebScraper() {
	fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " WebScraper is running.\n")
	webscraper.DoMagic()
}
