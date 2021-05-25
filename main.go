package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
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
	secrets := GetSecrets()
	fs := Firestore{}
	fs.Collection = "NewLetters"
	fs.Connect(secrets.FireStoreKeyPath, secrets.FireStoreProjectID)

	// var c Config
	// c.Email = "pe3et@protonmail.com"
	// c.Feedname = "my cool feed"
	// c.Greetingname = "peet"
	// c.Time = "10:00"

	// var m ModuleConfig
	// m.Name = "rss-feed"
	// AddSetting(&m, "URL", "https://github.com/wwhtrbbtt/PersonalNewsletter/commits.atom")
	// AddSetting(&m, "count", "5")

	// c.Modules = append(c.Modules, m)
	// err := fs.SetDoc("test", c)
	// if err != nil {
	// 	panic(err)
	// }
	// // j, _ := json.MarshalIndent(c, "", "    ")
	// // fmt.Println(string(j))
	// os.Exit(1)

	// c := GetConfig("config.json")

	// feed := ConfigToFeed(c)

	for {
		// get current time stamp
		current := time.Now().Format("15:04") // gotta love golang...

		// get all letters that should be send in this minute
		docs, err := fs.WhereDoc("Time", "==", current)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Printf("[*] Checking at %s - %d Newsletters to send\n", current, len(docs))

		for _, letter := range docs {
			var c Config
			err := mapstructure.Decode(letter, &c)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			feed := ConfigToFeed(c)
			fmt.Printf("	[*] Sending newsletter to %s\n", feed.Email)
			sender.SendEmail(feed, template, secrets.SenderMail, secrets.Password, secrets.SMTPServer, secrets.SenderMail)
		}
		time.Sleep(time.Minute * 1)
	}

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

// func GetConfig(path string) Config {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	var config Config
// 	err = json.Unmarshal(data, &config)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return config
// }

type Secrets struct {
	SenderMail         string
	Password           string
	SMTPServer         string
	FireStoreProjectID string
	FireStoreKeyPath   string
}

func GetSecrets() Secrets {
	// get secrets from .env file
	senderMail, exists := os.LookupEnv("SENDEREMAIL")
	if !exists {
		fmt.Println("Pleace set the SENDERMAIL")
		os.Exit(1)
	}

	password, exists := os.LookupEnv("PASSWORD")
	if !exists {
		fmt.Println("Please set the PASSWORD")
		os.Exit(1)
	}

	SMTPServer, exists := os.LookupEnv("SMTPSERVER")
	if !exists {
		fmt.Println("Please set the SMTPSERVER")
		os.Exit(1)
	}

	fireStoreProjectID, exists := os.LookupEnv("FSPROJID")
	if !exists {
		fmt.Println("Please set the FSPROJID")
		os.Exit(1)
	}

	fireStoreKeyPath, exists := os.LookupEnv("FSKEYPATH")
	if !exists {
		fmt.Println("Please set the FSKEYPATH")
		os.Exit(1)
	}

	return Secrets{
		SenderMail:         senderMail,
		Password:           password,
		SMTPServer:         SMTPServer,
		FireStoreProjectID: fireStoreProjectID,
		FireStoreKeyPath:   fireStoreKeyPath,
	}
}
