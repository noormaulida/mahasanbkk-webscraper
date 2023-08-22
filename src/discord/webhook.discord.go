package discord

import (
	"fmt"
	"mahasanbkk-webscraper/pkg/config"
	"mahasanbkk-webscraper/pkg/session"
	"mahasanbkk-webscraper/src/webscraper"

	"github.com/bwmarrin/discordgo"
)

func Webhook() {
	appID := config.ConfigData.DiscordAppId
	guildID := config.ConfigData.DiscordGuildId
	_, err := session.DiscordSession.ApplicationCommandBulkOverwrite(appID, guildID, []*discordgo.ApplicationCommand{
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
	session.DiscordSession.AddHandler(func(
		_ *discordgo.Session,
		i *discordgo.InteractionCreate,
	) {
		data := i.ApplicationCommandData()
		switch data.Name {
		case "hello-bot":
			err := session.DiscordSession.InteractionRespond(
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
			msg, msgEmbeds := webscraper.DoMagic(true)
			var err error

			if msgEmbeds == nil {
				err = session.DiscordSession.InteractionRespond(
					i.Interaction,
					&discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "Roger that, checking the schedule right away ❤️\n" + msg,
						},
					},
				)
			} else {
				err = session.DiscordSession.InteractionRespond(
					i.Interaction,
					&discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "Roger that, checking the schedule right away ❤️\nAvailable Schedule:",
							Embeds:  msgEmbeds,
						},
					},
				)
			}

			if err != nil {
				fmt.Println("err interaction force-check response ", err)
			}
		}
	})
	err = session.DiscordSession.Open()
	if err != nil {
		fmt.Println("err open webhook ", err)
	}
}
