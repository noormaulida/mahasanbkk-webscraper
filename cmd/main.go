package main

import (
	"fmt"
	"strconv"

	"mahasanbkk-webscraper/entities"

	"github.com/gocolly/colly"
)

func main() {
	url := "https://www.mahasanbkk.com/available-table.php"

	fmt.Println("--- Scrapping Started ---")
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("URL: ", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Status:", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit(url)

	c.OnHTML("available-table-body", func(h *colly.HTMLElement) {
		txt := h.Text
		availableScheds := []entities.Schedule{}

		if txt != "Sorry, there is no table available at the moment." {
			h.ForEach("tr", func(_ int, el *colly.HTMLElement) {
				fmt.Println(el.ChildText("td:nth-child(1)"))
				availableGuest, _ := strconv.Atoi(el.ChildAttr("a", "href"))

				schedule := entities.Schedule{}
				schedule.Guest = availableGuest
				schedule.Date = el.ChildText("h2") //el.ChildAttr("img", "src")
				schedule.Time = el.ChildText("h2")

				availableScheds = append(availableScheds, schedule)
			})
		} else {
			fmt.Println("No available schedule")
		}

	})

}
