package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerHost string `mapstructure:"SERVER_HOST"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	ServerEnv  string `mapstructure:"SERVER_ENV"`

	MahasanUrl       string `mapstructure:"MAHASAN_URL"`
	MahasanSubUrl    string `mapstructure:"MAHASAN_SUB_URL"`
	MahasanChannelID string `mapstructure:"MAHASAN_CHANNEL_ID"`

	DiscordToken     string `mapstructure:"DISCORD_TOKEN"`
	DiscordAppId     string `mapstructure:"DISCORD_APP_ID"`
	DiscordGuildId   string `mapstructure:"DISCORD_GUILD_ID"`

	DiscordStatus   string `mapstructure:"DISCORD_STATUS"`
	LineStatus   string `mapstructure:"Line_STATUS"`

	LineAccessToken     string `mapstructure:"LINE_ACCESS_TOKEN"`
	LineSecret     string `mapstructure:"LINE_SECRET"`
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
