package main

import (
	"errors"

	"github.com/wwhtrbbtt/PersonalNewsletter/aggregator"
)

func FetchData(c ModuleConfig) (aggregator.Module, error) {
	switch c.Name {
	case "rss-feed":
		return aggregator.FetchRssFeed(c.Settings[0].Value.(string), int(c.Settings[1].Value.(int64)))
	case "qotd":
		return aggregator.FetchQuoteOfTheDay()
	default:
		return aggregator.Module{}, errors.New("couldn't find module (FetchData)")
	}
}
