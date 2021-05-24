package main

import (
	aggregator "github.com/wwhtrbbtt/PersonalNewsletter/aggregator"
	sender "github.com/wwhtrbbtt/PersonalNewsletter/sender"
)

func main() {
	var feed sender.Feed

	feed.Time = "00:00"
	feed.Email = "test@test.com"
	feed.Feedname = "Test feed"
	feed.Greetingname = "Test"

	// RSS feed
	rss, _ := aggregator.FetchRssFeed("https://github.com/wwhtrbbtt/PersonalNewsletter/commits.atom", 10)
	feed.Modules = append(feed.Modules, rss)
}
