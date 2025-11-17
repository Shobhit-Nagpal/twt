package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Post a tweet",
	Long:  "Post a tweet with the content you want to write",
	Args:  cobra.MinimumNArgs(1),
	Run:   postTweet,
}

func postTweet(cmd *cobra.Command, args []string) {
	description := strings.Join(args, " ")
	id, err := x.PostTweet(description)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	fmt.Printf("Tweeted! Check it out at https://x.com/%s/status/%s\n", username, id)
}
