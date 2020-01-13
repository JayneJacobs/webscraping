package jjscraper

import (
	"fmt"
	"github.com/gocolly/colly"
)

// JayneScraper collects info from web
func JayneScraper() {
	//Instantiate default collector
	fmt.Println("Starting Scrape")
	c := colly.NewCollector(
		colly.AllowedDomains("jaynejacobs.com", "udemy.com"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		e.Request.Visit(link)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("http://jaynejacobs.com")
	c.Visit("https://www.udemy.com/")
}
