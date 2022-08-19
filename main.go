package main

import (
	"log"
	"time"

	"github.com/MangKong-coder/daily-quote-scraper/email"
	"github.com/MangKong-coder/daily-quote-scraper/scraper"
)

func main() {
	for {
		// print time
		loc, err := time.LoadLocation("Asia/Manila")
		if err != nil {
			log.Fatal(err)
		}
		currentTime := time.Now().In(loc).Format("15:04:05")
		// sends email at 8:00:00AM in the morning
		if currentTime == "8:00:00" {
			images := scraper.ScrapeQuoteImages()
			var htmlString string
			for _, j := range images {
				htmlString += "<img src=\"" + j + "\" > \n"
			}
			email.EmailQuotes(htmlString)
		}
	}
}
