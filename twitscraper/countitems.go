package twitscraper

import "fmt"

import "strings"

import "sort"

//MyCounter Counts words ToDo
func MyCounter() {
	fmt.Println("This is my counter")

}

// WordInfo is a type for word count
type WordInfo struct {
	word  string
	count int
}

// AnalysisScrape takes a url for a tweet and counts the words and sorts them in decending order
func AnalysisScrape(url string) {
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

	sortCount := wordCounts(jtweets)
	for _, v := range sortCount {
		fmt.Printf("%s: %d\n", v.word, v.count)
	}

}

func wordCounts(tweets []jtweets) []WordInfo {
	wordMap := map[string]int{}

	for _, t := range tweets {
		words := strings.Split(t.Tweet, " ")
		for _, w := range words {
			wordMap[strings.ToLower(w)]++
		}
	}

	wis := []WordInfo{}
	for w, c := range wordMap {
		wis = append(wis, WordInfo{
			word:  w,
			count: c,
		})
	}

	sort.Slice(wis, func(i, j int) bool {
		return wis[i].count > wis[j].count
	})
	return wis
}

// func makeTweetRequest(next string, twURL string) (*tweetResponse, error) {
// 	urlparam := url.Values{}
// 	urlparam.Set("include_available_features", "1")
// 	urlparam.Set("include_entities", "1")
// 	urlparam.Set("max_position", next)
// 	urlparam.Set("reset_error_state", "false")
// 	twURL = twURL + "?" + urlparam.Encode()
// 	resp, err := http.Get(twURL)

// 	if err != nil {
// 		return nil, fmt.Errorf("Error while gettng conversation data %w", err)

// 	}
// 	defer resp.Body.Close()
// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("Incorrect status code %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))

// 	}

// 	tr := tweetResponse{}
// 	err = json.NewDecoder(resp.Body).Decode(tr)
// 	if err != nil {
// 		return nil, fmt.Errorf("Error decoding response %w", err)

// 	}
// 	return tr, nil
// }
