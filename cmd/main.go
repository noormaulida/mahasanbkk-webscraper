package main

import (
	"mahasanbkk-webscraper/pkg/config"
	"mahasanbkk-webscraper/src/webscraper"
)

func main() {
	config.Load(".")
	webscraper.DoMagic()
}
