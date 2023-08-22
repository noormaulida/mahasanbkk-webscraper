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
		{
			Name:        "status",
			Description: "Slash command to check server status",
		},
		{
			Name:        "discord-status",
			Description: "Slash command to check discord webhook notification status",
		},
		{
			Name:        "discord-start",
			Description: "Slash command to start discord webhook notification",
		},
		{
			Name:        "discord-stop",
			Description: "Slash command to temporarily stop discord webhook notification",
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
		case "status":
			err := session.DiscordSession.InteractionRespond(
				i.Interaction,
				&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Server is up 🚀",
					},
				},
			)
			if err != nil {
				fmt.Println("err interaction status response ", err)
			}
		case "discord-status":
			status := config.ConfigData.DiscordStatus
			message := "Discord webhook is "
			if status == "on" {
				message += " up 🚀"
			} else {
				message += " down 😥"
			}
			err := session.DiscordSession.InteractionRespond(
				i.Interaction,
				&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: message,
					},
				},
			)
			if err != nil {
				fmt.Println("err interaction discord-status response ", err)
			}
		case "discord-start":
			config.ConfigData.DiscordStatus = "on"
			session.ResetPreviousTableIDs()
			message := "Discord webhook is starting 🚀\n Sending notification is active."
			err := session.DiscordSession.InteractionRespond(
				i.Interaction,
				&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: message,
					},
				},
			)
			if err != nil {
				fmt.Println("err interaction discord-start response ", err)
			}
		case "discord-stop":
			config.ConfigData.DiscordStatus = "off"
			session.ResetPreviousTableIDs()
			message := "Discord webhook is stopping 😥\n Sending notification is now disabled. Please use /discord-start to start the webhook again."
			err := session.DiscordSession.InteractionRespond(
				i.Interaction,
				&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: message,
					},
				},
			)
			if err != nil {
				fmt.Println("err interaction discord-start response ", err)
			}
		}
	})
	err = session.DiscordSession.Open()
	if err != nil {
		fmt.Println("err open webhook ", err)
	}
}
