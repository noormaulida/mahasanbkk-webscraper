package config

import (
    "log"
    "github.com/bwmarrin/discordgo"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var (
	DiscordSession *discordgo.Session
	LineSession *linebot.Client
	err            error
)


func InitSession() (err error) {
	// Create new Discord Session
	DiscordSession, err = discordgo.New("Bot " + ConfigData.DiscordToken)
	if err != nil {
		log.Fatalln("cannot create discord session. err ", err)
	}

	// Create new Linebot Session
	LineSession, err = linebot.New(
		ConfigData.LineSecret,
		ConfigData.LineAccessToken,
	)
    if err != nil {
		log.Fatalln("cannot create line session. err ", err)
	}

	return
}
