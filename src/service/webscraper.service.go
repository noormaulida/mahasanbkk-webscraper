package service

import (
	"fmt"
	"time"

	"mahasanbkk-webscraper/src/webscraper"
)

func WebScraper() {
	fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " WebScraper is running.\n")
	webscraper.DoMagic(false)
}
