package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	aggregator "github.com/wwhtrbbtt/PersonalNewsletter/aggregator"
	sender "github.com/wwhtrbbtt/PersonalNewsletter/sender"
)

var template = "./template/email1.html"

// A config tells the program what data a program wants.
// Its similiar to "Feed", but doesn't yet contain data
type Config struct {
	Time         string         `json:"time"`         // ex. "10:00"
	Feedname     string         `json:"feedName"`     // ex. "My email feed"
	Greetingname string         `json:"greetingName"` // ex. "peet"
	Email        string         `json:"email"`        // ex. peet@peet.ws
	Modules      []ModuleConfig `json:"modules"`      // collection of modules
}

type ModuleConfig struct {
	Name     string                `json:"name" `    // ex. "rss-feed"
	Settings []ModuleConfigSetting `json:"settings"` // a bunch of settings for the module

}

type ModuleConfigSetting struct {
	Name  string `json:"name"`  // ex. feed-url
	Value string `json:"value"` // ex. https://github.com/wwhtrbbtt/PersonalNewsletter/commits.atom

}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found!")
		os.Exit(1)
	}
}

func main() {
	var c Config
	c.Email = "pe3et@protonmail.com"
	c.Feedname = "my cool feed"
	c.Greetingname = "peet"
	c.Time = "10:00"

	var m ModuleConfig
	m.Name = "rss-feed"
	AddSetting(&m, "URL", "https://github.com/wwhtrbbtt/PersonalNewsletter/commits.atom")
	AddSetting(&m, "count", "5")

	c.Modules = append(c.Modules, m)

	feed := ConfigToFeed(c)

	senderMail, exists := os.LookupEnv("SENDEREMAIL")
	if !exists {
		fmt.Println("Pleace set the SENDERMAIL")
		os.Exit(1)
	}

	password, exists := os.LookupEnv("PASSWORD")
	if !exists {
		fmt.Println("Pleace set the PASSWORD")
		os.Exit(1)
	}

	SMTPServer, exists := os.LookupEnv("SMTPSERVER")
	if !exists {
		fmt.Println("Pleace set the SMTPSERVER")
		os.Exit(1)
	}
	fmt.Println(feed)
	sender.SendEmail(feed, template, senderMail, password, SMTPServer, senderMail)
}

func ConfigToFeed(c Config) sender.Feed {
	var s sender.Feed
	// copy data that is the same
	s.Email = c.Email
	s.Feedname = c.Feedname
	s.Greetingname = c.Greetingname
	s.Time = c.Time

	for _, config := range c.Modules {
		for _, module := range GetAllModules().Modules { // module = demo module, all required data given
			if CheckEqualSettings(config, module) { // valid settings
				// fetch data for the module
				data, err := FetchData(config)
				if err != nil {
					fmt.Println(err)
					continue
				}
				// if no errors, append to modules
				s.Modules = append(s.Modules, data)
			}
		}
	}
	return s

}
func CheckEqualSettings(m1, m2 ModuleConfig) bool {
	// checks if a setting is valid
	// m1 = user input, m2 = valid

	// check name
	if m1.Name != m2.Name {
		return false
	}

	// check lenght
	if len(m1.Settings) != len(m2.Settings) {
		return false
	}

	// check all setting names
	for count, _ := range m2.Settings {
		if m1.Settings[count].Name != m2.Settings[count].Name {
			return false
		}
	}
	return true
}

func FetchData(c ModuleConfig) (aggregator.Module, error) {
	switch c.Name {
	case "rss-feed":
		// 0: URL, 1: Count
		i, err := strconv.Atoi(c.Settings[1].Value)
		if err != nil {
			fmt.Println(err)
			return aggregator.Module{}, errors.New("couldn't parse settings (FetchData)")
		}
		return aggregator.FetchRssFeed(c.Settings[0].Value, i)
	default:
		return aggregator.Module{}, errors.New("couldn't find module (FetchData)")
	}

}
