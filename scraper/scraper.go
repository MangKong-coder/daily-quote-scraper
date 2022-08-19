package scraper

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeQuoteImages() []string {
	var images []string
	res, err := http.Get("https://www.insightoftheday.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Find the quote images
	doc.Find(".img-responsive").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band, _ := s.Attr("src")
		images = append(images, band)
	})
	return images
}
