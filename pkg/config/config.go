package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerHost string `mapstructure:"SERVER_HOST"`
	ServerPort string `mapstructure:"SERVER_PORT"`

	MahasanUrl       string `mapstructure:"MAHASAN_URL"`
	MahasanChannelID string `mapstructure:"MAHASAN_CHANNEL_ID"`
	DiscordToken     string `mapstructure:"DISCORD_TOKEN"`
}

var ConfigData *Config

func Load(path string) (err error) {
	var config Config
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	ConfigData = &config
	return
}
