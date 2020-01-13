package twitscraper

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const apisite = "https://twitter.com/i/Todd_McLeod/conversation/1169751640926146560"

type jtweets struct {
	Title string
	User  string
	Tweet string
}

type tweetResponse struct {
	MinP string `json:"min_position"`
	Next bool   `json:"has_more_items"`
	Body string `json:"items_html"`
}

func tweetRequest(next string) (*tweetResponse, error) {
	items := url.Values{}
	items.Set("include_available_features", "1")
	items.Set("include_entities", "1")
	items.Set("max_position", next)
	url := apisite + "?" + items.Encode()
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error while getting tweet data: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Incorrect status %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
	tr := &tweetResponse{}
	err = json.NewDecoder(resp.Body).Decode(tr)
	if err != nil {
		return nil, fmt.Errorf("Error decoding tweet: %w", err)
	}
	return tr, nil
}

func getTweets() ([]string, error) {
	nextValue := ""
	tweets := []string{}

	for i := 0; i < 100; i++ {
		resp, err := tweetRequest(nextValue)
		if err != nil {
			return nil, fmt.Errorf("Cant make tweet request: %w", err)
		}
		tweets = append(tweets, resp.Body)

		if !resp.Next {
			break
		}
		nextValue = resp.MinP
		time.Sleep(time.Second)
	}
	return tweets, nil
}

func parseBody(jtweet string) ([]jtweets, error) {
	twRdr := strings.NewReader(jtweet)
	doc, err := goquery.NewDocumentFromReader(twRdr)
	if err != nil {
		return nil, fmt.Errorf("Cant read body: %w", err)
	}

	tw := []jtweets{}

	doc.Find(".tweet").Each(func(i int, s *goquery.Selection) {
		tw = append(tw, jtweets{
			Title: s.Find(".account-group .fullname").Text(),
			User:  s.Find(".account-group .username").Text(),
			Tweet: s.Find(".tweet-text").Text(),
		})
	})

	return tw, nil
}

//MyScraper gathers responses to a tweet
func MyScraper() {
	resp, err := getTweets()
	if err != nil {
		panic(err)
	}

	jtweets := []jtweets{}
	for _, twt := range resp {
		tw, err := parseBody(twt)
		if err != nil {
			panic(err)
		}
		jtweets = append(jtweets, tw...)
	}
	bs, err := json.MarshalIndent(jtweets, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))
	fmt.Println("Number of tweets:", len(jtweets))
}
