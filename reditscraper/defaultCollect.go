package reditscraper

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type redititem struct {
	StoryURL    string
	Source      string
	comments    string
	CrawledTime time.Time
	Comments    string
	Title       string
}

// UserCrawl uses an argument passed in during the call
func UserCrawl() {
	stories := []redititem{}
	// Start default Collector
	c := colly.NewCollector(
		//Visit old domains
		colly.AllowedDomains("old.Redit.com"),
		colly.Async(true),
	)
	c.OnHTML(".top-matter", func(e *colly.HTMLElement) {
		temp := redititem{}
		temp.StoryURL = e.ChildAttr("a[data-event-action=title]", "href")
		temp.Source = "https://old.reddit.com/r/programming/"
		temp.Title = e.ChildText("a[data-event-action=comments]")
		temp.CrawledTime = time.Now()
		stories = append(stories, temp)
	})

	//On span tag with class next-button

	c.OnHTML("span.next.button", func(h *colly.HTMLElement) {
		t := h.ChildAttr("a", "href")
		c.Visit(t)
	})

	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	// Crawl redit the user passes in
	reddits := os.Args[1:]
	for _, reddit := range reddits {
		c.Visit(reddit)
	}

	c.Wait()
	fmt.Println(stories)

}

func ReditWords() {
	stories := []redititem{}

	//Start Default Collector
	c := colly.NewCollector(
		//visit old.reddit.com
		colly.AllowedDomains("old.reddit.com"),
		colly.Async(true),
	)
	// for every element with .top-matter attribute call call back.
	c.OnHTML(".top-matter", func(e *colly.HTMLElement) {
		temp := redititem{}
		temp.StoryURL = e.ChildAttr("a[data-event-action=title]", "hrefß")
		temp.Source = "https://old.reddit.com/r/programming/"
		temp.Title = e.ChildText("a[data-event-action=title]")
		temp.Comments = e.ChildAttr("a[data-event-action=comments]", "htref")
		temp.CrawledTime = time.Now()
		stories = append(stories, temp)

	})

	c.OnHTML("span.next-bbutton", func(h *colly.HTMLElement) {
		t := h.ChildAttr("a", "href")
		c.Visit(t)
	})

	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	// Print Visiting before makin request
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Crawl redits from user
	reddits := os.Args[1:]
	for _, reddit := range reddits {
		c.Visit(reddit)
	}

	c.Wait()

	m := map[string]int{}

	for _, story := range stories {
		words := strings.Split(strings.ToLower(story.Title), " ")
		for _, word := range words {
			m[word]++
		}
	}

	type wordCount struct {
		word  string
		count int
	}

	xwc := []wordCount{}

	for w, c := range m {
		xwc = append(xwc, wordCount{
			word:  w,
			count: c,
		})
	}

	sort.Slice(xwc, func(i, j int) bool {
		return xwc[i].count > xwc[j].count
	})

	for _, v := range xwc {
		fmt.Println(v.word, v.count)
	}
}
