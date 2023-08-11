package webscraper

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"mahasanbkk-webscraper/pkg/config"
	"mahasanbkk-webscraper/src/entities"

	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly"
)

func DoMagic() {
	url := config.ConfigData.MahasanUrl

	availableScheds := []entities.Schedule{}

	fmt.Println("--- Scrapping Started ---")
	c := colly.NewCollector(
		colly.AllowURLRevisit(),
	)

	c.WithTransport(&http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   90 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("URL: ", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Status:", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Limit(&colly.LimitRule{
		Delay:       1 * time.Second,
		RandomDelay: 1 * time.Second,
	})

	c.OnHTML("div.available-table-body", func(h *colly.HTMLElement) {
		// Create new Discord Session
		discord, err := discordgo.New("Bot " + config.ConfigData.DiscordToken)
		mahasanChannelId := config.ConfigData.MahasanChannelID

		txt := h.ChildText("h4.available-table-sorry")
		if txt != "Sorry, there is no table available at the moment." {

			if err != nil {
				log.Fatalln("cannot create discord session. err ", err)
			}

			h.ForEach("a", func(_ int, el *colly.HTMLElement) {
				note := el.ChildText("div.available-table-content")
				schedule := entities.Schedule{}
				schedule.Notes = note
				availableScheds = append(availableScheds, schedule)
			})
			words := "@everyone Available Schedule: " + time.Now().Format("2006-01-02 15:04:05") + "\n"
			fmt.Println(words)
			for _, sched := range availableScheds {
				fmt.Println(sched.Notes)
				words += (sched.Notes + "\n")
			}
			words += "Sent from " + config.ConfigData.ServerEnv + " environment\n"
			words += "--------------------------------------"
			discord.ChannelMessageSend(mahasanChannelId, words)
		} else {
			nowords := "No available schedule " + time.Now().Format("2006-01-02 15:04:05") + "\n"
			nowords += "Sent from " + config.ConfigData.ServerEnv + " environment\n"
			nowords += "--------------------------------------"
			discord.ChannelMessageSend(mahasanChannelId, nowords)
			fmt.Println(nowords)
		}
	})

	c.Visit(url)

	fmt.Println("--- Scrapping Ended ---")
}
