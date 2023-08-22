package service

import (
	"fmt"
    "time"

    "mahasanbkk-webscraper/src/discord"
	"mahasanbkk-webscraper/pkg/config"
)

func DiscordWebhook() {
	if config.ConfigData.DiscordStatus == "on" {
		fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " Discord Webhook is running.\n")
		discord.Webhook()
	}
}
