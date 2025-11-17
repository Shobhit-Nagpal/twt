package main

import (
	"flag"
	"log"
	"os"

	"github.com/Shobhit-Nagpal/twt/internal/twt"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
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

	consumerKey := os.Getenv("API_KEY")
	consumerSecret := os.Getenv("API_KEY_SECRET")

	accessToken := os.Getenv("ACCESS_TOKEN")
	accessSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	twt := twt.InitTwt(config, token)

	err := twt.Post(*postPtr)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Tweeted!")
}
