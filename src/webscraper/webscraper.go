package webscraper

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"mahasanbkk-webscraper/pkg/config"
	"mahasanbkk-webscraper/pkg/session"
	"mahasanbkk-webscraper/src/entities"

	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly"
)

func DoMagic(forcely bool) (string, []*discordgo.MessageEmbed) {
	url := config.ConfigData.MahasanUrl + config.ConfigData.MahasanSubUrl
	mahasanChannelID := config.ConfigData.MahasanChannelID
	availableScheds := []entities.Schedule{}
	var messageEmbeds []*discordgo.MessageEmbed
	var message string

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
		// Set prefix of notif
		var prefix string
		if config.ConfigData.ServerEnv == "production" {
			prefix = "@everyone"
		}

		txt := h.ChildText("h4.available-table-sorry")
		if txt != "Sorry, there is no table available at the moment." {
			h.ForEach("a", func(_ int, el *colly.HTMLElement) {
				hrefLink := el.Attr("href")
				link := config.ConfigData.MahasanUrl + hrefLink[1:]
				guest := el.ChildText("div.available-table-content > div:nth-child(1)")
				time := el.ChildText("div.available-table-content > div:nth-child(2)")

				guestString := strings.Replace(guest, " GUESTS", "", -1)
				schedule := entities.Schedule{}
				schedule.Guest, _ = strconv.Atoi(guestString)
				schedule.DateTime = time
				schedule.Link = strings.ReplaceAll(link, "#navbar", "")
				schedule.Notes = guest + " : " + time + " " + link
				availableScheds = append(availableScheds, schedule)
			})
			message = prefix + " Available Schedule: \n"
			for _, sched := range availableScheds {
				message += (sched.Notes + "\n")
				autoBookLink := config.ConfigData.WebAutoBookURL +"?id="+ GetTableID(sched.Link)
				link := "🌹 Manual Booking: "+sched.Link+"\n"+"🌷 Auto Booking: "+autoBookLink
				embed := discordgo.MessageEmbed{
					Type:        discordgo.EmbedTypeRich,
					Title:       sched.DateTime + " - " + strconv.Itoa(sched.Guest) + " Guests",
					Description: link,
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:   "Date",
							Value:  sched.DateTime,
							Inline: true,
						},
						{
							Name:   "Guest",
							Value:  strconv.Itoa(sched.Guest),
							Inline: true,
						},
					},
				}
				messageEmbeds = append(messageEmbeds, &embed)
			}
			SendDiscordMessage(forcely, mahasanChannelID, message, messageEmbeds)
			fmt.Println(message)
		} else {
			message = "No available schedule \n"
			SendDiscordMessage(forcely, mahasanChannelID, message, nil)
			fmt.Println(message)
		}
	})

	c.Visit(url)

	fmt.Println("--- Scrapping Ended ---")

	return message, messageEmbeds
}

func GetTableID(url string) (tableID string) {
	data := strings.SplitN(url, "booking-form.php?id=", 2)
	if (len(data) == 2) {
		tableID = data[1]
	}
	return
}

func SendDiscordMessage(forcely bool, channelID string, _ string, messageEmbeds []*discordgo.MessageEmbed) {
	if config.ConfigData.DiscordStatus == "on" && !forcely {
		if messageEmbeds != nil {
			var prefix string
			if config.ConfigData.ServerEnv == "production" {
				prefix = "@everyone "
			}
			session.DiscordSession.ChannelMessageSend(channelID, prefix+"Available Schedule:")
			session.DiscordSession.ChannelMessageSendEmbeds(channelID, messageEmbeds)
		}
	}
}
