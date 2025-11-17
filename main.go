package main

import (
	"flag"
	"log"
	"os"

	"github.com/Shobhit-Nagpal/twt/internal/twt"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2/clientcredentials"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	postPtr := flag.String("post", "", "tweet to post")
	flag.Parse()

	consumerKey := os.Getenv("CLIENT_ID")
	consumerSecret := os.Getenv("CLIENT_SECRET")

	config := &clientcredentials.Config{
		ClientID:     consumerKey,
		ClientSecret: consumerSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}

	twt := twt.InitTwt(config)

	err := twt.Post(*postPtr)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Tweeted!")
}
