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

	token := os.Getenv("BEARER_TOKEN")
	twt := twt.InitTwt(token)

	err := twt.Post(*postPtr)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Tweeted!")
}
