package session

import (
	"log"

	"mahasanbkk-webscraper/pkg/config"

	"github.com/bwmarrin/discordgo"
)

var DiscordSession *discordgo.Session

func InitSession() (err error) {
	// Create new Discord Session
	DiscordSession, err = discordgo.New("Bot " + config.ConfigData.DiscordToken)
	if err != nil {
		log.Fatalln("cannot create discord session. err ", err)
	}

	return
}
