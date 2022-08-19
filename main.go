package main

import (
	"time"

	"github.com/MangKong-coder/daily-quote-scraper/email"
	"github.com/MangKong-coder/daily-quote-scraper/scraper"
)

func main() {
	for {
		currentTime := time.Now().Local().Format("15:04:05")
		// sends email at 8AM in the morning
		if currentTime == "12:00:00" {
			images := scraper.ScrapeQuoteImages()
			var htmlString string
			for _, j := range images {
				htmlString += "<img src=\"" + j + "\" > \n"
			}
			email.EmailQuotes(htmlString)
		}
	}
}
