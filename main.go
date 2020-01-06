package main

import (
	jjscraper "github.com/JayneJacobs/webscraping/jjscraper"
	reditscraper "github.com/JayneJacobs/webscraping/reditscraper"
	twitterscraper "github.com/JayneJacobs/webscraping/twitscraper"
)

func main() {
	jjscraper.JayneScraper()
	reditscraper.Myscraper()
	twitterscraper.MyScraper()
}
