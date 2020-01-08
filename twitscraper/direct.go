package twitscraper

import "github.com/gocolly/colly"

import "encoding/json"

import "fmt"

const site = "https://twitter.com/WhiteHouse/status/1214946620053184513"

// MyDirectScraper will use the status link and get the conversation
func MyDirectScraper() {
	c := colly.NewCollector()

	tweets := []jtweets{}

	c.OnHTML(".tweet", func(e *colly.HTMLElement) {
		tweets = append(tweets, jtweets{
			Title: e.ChildText(".account-group .fullname"),
			User:  e.ChildText(".account-group .username"),
			Tweet: e.ChildText(".tweet-text"),
		})
	})

	err := c.Visit(site)
	if err != nil {
		panic(err)
	}

	c.Wait()

	tw, err := json.MarshalIndent(tweets, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(tw))
	fmt.Println("Number of tweets", len(tweets))
}
