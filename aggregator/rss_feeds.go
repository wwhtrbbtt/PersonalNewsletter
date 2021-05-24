package aggregator

// Get the first n items of a RSS feed

import (
	"fmt"
	"regexp"

	gofeed "github.com/mmcdole/gofeed"
)

// ex. https://google.com/search -> google.com
func _GetSiteName(url string) string {
	re := regexp.MustCompile(`\/\/(.+?\..+?)\/`)
	matches := re.FindStringSubmatch(url + "/") // + "/" to ensure that there are no errors
	if len(matches) < 2 {
		return url
	}
	return matches[1]

}
func FetchRssFeed(url string, count int) (Module, error) {
	var m Module
	siteName := _GetSiteName(url)

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return m, err
	}
	fmt.Println(feed)
	// TODO - parse feed
	m.Caption = "RSS - " + siteName
	m.Text = fmt.Sprintf("The latest %d results for the RSS feed of %s", count, siteName)

	for tmp, entry := range feed.Items {
		if tmp == count {
			break
		}
		m.Links = append(m.Links, Link{Text: entry.Title, URL: entry.Link})
	}
	return m, nil
}
