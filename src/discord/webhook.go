package discord

import (
	"fmt"
	"mahasanbkk-webscraper/pkg/config"
	"mahasanbkk-webscraper/src/webscraper"

	"github.com/bwmarrin/discordgo"
)

func AvailableCommand() {
	discordSession := config.DiscordSession
	appID := config.ConfigData.DiscordAppId
	guildID := config.ConfigData.DiscordGuildId
	_, err := discordSession.ApplicationCommandBulkOverwrite(appID, guildID, []*discordgo.ApplicationCommand{
		{
			Name:        "hello-bot",
			Description: "Slash command that can say hello to you",
		},
		{
			Name:        "force-check",
			Description: "Slash command to force the server to check schedule right away",
		},
	})
	if err != nil {
		fmt.Println("err bulk overwrite ", err)
	}
	discordSession.AddHandler(func(
		_ *discordgo.Session,
		i *discordgo.InteractionCreate,
	) {
		data := i.ApplicationCommandData()
		switch data.Name {
		case "hello-bot":
			err := discordSession.InteractionRespond(
				i.Interaction,
				&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Hello there! ❤️",
					},
				},
			)
			if err != nil {
				fmt.Println("err interaction hello-bot response ", err)
			}
		case "force-check":
			err := discordSession.InteractionRespond(
				i.Interaction,
				&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Roger that, checking the schedule right away ❤️",
					},
				},
			)
			if err != nil {
				fmt.Println("err interaction force-check response ", err)
			}
			webscraper.DoMagic(true)
		}
	})
	err = discordSession.Open()
	if err != nil {
		fmt.Println("err open webhook ", err)
	}
}
