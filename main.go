package main

import (
	"fmt"

	jjscraper "github.com/JayneJacobs/webscraping/jjscraper"
	reditscraper "github.com/JayneJacobs/webscraping/reditscraper"
	twitterscraper "github.com/JayneJacobs/webscraping/twitscraper"
)

func main() {
	var twURL string

	jjscraper.JayneScraper()
	fmt.Println("Please enter the conversation URL here:   ")
	fmt.Scan(&twURL)
	reditscraper.Myscraper()
	twitterscraper.AnalysisScrape(twURL)
	fmt.Println("Please enter the conversation URL here:   ")
	fmt.Scan(&twURL)
	twitterscraper.MyDirectScraper(twURL)
	reditscraper.UserCrawl()
	reditscraper.ReditWords()

	//"https://old.reddit.com/r/programming/"

}
