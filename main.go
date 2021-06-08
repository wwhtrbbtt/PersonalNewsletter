package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
	sender "github.com/wwhtrbbtt/PersonalNewsletter/sender"
)

var DEBUG = true
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
	Name  string      `json:"name"`  // ex. feed-url
	Value interface{} `json:"value"` // ex. https://github.com/wwhtrbbtt/PersonalNewsletter/commits.atom
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found!")
		os.Exit(1)
	}
}

func main() {

	// fmt.Println(GetAllModules())
	// os.Exit(0)
	secrets := GetSecrets()
	fs := Firestore{}
	fs.Collection = "NewLetters"
	fs.Connect(secrets.FireStoreKeyPath, secrets.FireStoreProjectID)

	if len(os.Args) > 1 {
		if os.Args[1] == "sync" {
			modules := getAllModules()
			fs.Collection = "misc"

			err := fs.SetDoc("newModules", modules)
			if err != nil {
				panic(err)
			}
			fmt.Println("Successfully synced all modules")
			os.Exit(0)
		}
	}

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
			if DEBUG {
				fmt.Println("Raw config: ")
				fmt.Println(letter)
				fmt.Println("Raw -> Config object:")
				fmt.Println(c)
			}
			// fmt.Println(letter)
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
		data, err := FetchData(config)
		if err != nil {
			fmt.Println(err)
			continue
		}
		s.Modules = append(s.Modules, data)
	}
	return s

}

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
