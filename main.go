package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	aggregator "github.com/wwhtrbbtt/PersonalNewsletter/aggregator"
	sender "github.com/wwhtrbbtt/PersonalNewsletter/sender"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found!")
		os.Exit(1)
	}
}

func main() {

	var feed sender.Feed

	feed.Time = "00:00"
	feed.Email = "pe3et@protonmail.com"
	feed.Feedname = "Test feed"
	feed.Greetingname = "Test"

	// RSS feed
	rss, _ := aggregator.FetchRssFeed("https://github.com/wwhtrbbtt/PersonalNewsletter/commits.atom", 10)
	feed.Modules = append(feed.Modules, rss)

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

	fmt.Println(senderMail, password, SMTPServer)
	sender.SendEmail(feed, senderMail, password, SMTPServer)
}
