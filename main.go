package main

import (
	"daily-quote-scraper/scraper"
	"fmt"
)

func main() {
	images := scraper.ScrapeQuoteImages()
	for i, j := range images {
		images[i] = "<img src=\"" + j + "\" >"
	}
	fmt.Println(images)
}
