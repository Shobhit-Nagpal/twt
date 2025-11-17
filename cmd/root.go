package cmd

import (
	"log"
	"os"

	"github.com/Shobhit-Nagpal/twt/internal/twt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var x *twt.Twt

var username string

var rootCmd = &cobra.Command{
	Use:   "twt",
	Short: "Tweet from the terminal",
	Long:  `twt is a CLI tool for tweeting from the terminal.`,
}

func init() {
	// Add subcommands
	rootCmd.AddCommand(postCmd)

	// Setup deps
	setupDependencies()
}

func setupDependencies() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
		os.Exit(1)
	}

	consumerToken := os.Getenv("API_KEY")
	consumerTokenSecret := os.Getenv("API_KEY_SECRET")

	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	if consumerToken == "" || consumerTokenSecret == "" || accessToken == "" || accessTokenSecret == "" {
		log.Print("Missing API or Access token/secret")
		os.Exit(1)
	}

	x = twt.InitTwt(consumerToken, consumerTokenSecret, accessToken, accessTokenSecret)
	username = "shbhtngpl"
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
