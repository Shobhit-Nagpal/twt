package main

import (
	"flag"
	"log"
	"os"

	"github.com/Shobhit-Nagpal/twt/internal/twt"
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

	consumerToken := os.Getenv("API_KEY")
	consumerTokenSecret := os.Getenv("API_KEY_SECRET")

	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	if consumerToken == "" || consumerTokenSecret == "" || accessToken == "" || accessTokenSecret == "" {
		log.Fatal("Bearer Token required")
	}

	twt := twt.InitTwt(consumerToken, consumerTokenSecret, accessToken, accessTokenSecret)

	err := twt.Post(*postPtr)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Tweeted!")
}
