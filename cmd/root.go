package cmd

import (
	"log"
	"os"

	"github.com/Shobhit-Nagpal/twt/internal/twt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

	// Setup config
	setupConfig()

	// Setup deps
	setupDependencies()
}

func setupConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("$HOME/.config/twt")

	viper.ReadInConfig()
}

func setupDependencies() {
	consumerToken := viper.GetString("API_KEY")
	consumerTokenSecret := viper.GetString("API_KEY_SECRET")

	accessToken := viper.GetString("ACCESS_TOKEN")
	accessTokenSecret := viper.GetString("ACCESS_TOKEN_SECRET")

	if consumerToken == "" || consumerTokenSecret == "" || accessToken == "" || accessTokenSecret == "" {
		log.Print("Missing API or Access token/secret in config")
		os.Exit(1)
	}

	username = viper.GetString("USERNAME")
	if consumerToken == "" || consumerTokenSecret == "" || accessToken == "" || accessTokenSecret == "" {
		log.Print("Missing username in config")
		os.Exit(1)
	}

	x = twt.InitTwt(consumerToken, consumerTokenSecret, accessToken, accessTokenSecret)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
