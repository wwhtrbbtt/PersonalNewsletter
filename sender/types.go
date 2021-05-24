package sender

import aggregator "github.com/wwhtrbbtt/PersonalNewsletter/aggregator"

type Feed struct {
	Time         string              `json:"time"`         // ex. "10:00"
	Feedname     string              `json:"feedName"`     // ex. "My email feed"
	Greetingname string              `json:"greetingName"` // ex. "peet"
	Email        string              `json:"email"`        // ex. peet@peet.ws
	Modules      []aggregator.Module `json:"modules"`      // collection of modules
}
