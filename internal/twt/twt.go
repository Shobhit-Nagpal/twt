package twt

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/dghubble/oauth1"
	"github.com/g8rswimmer/go-twitter/v2"
)

type Twt struct {
	client *twitter.Client
}

type authorizer struct{}

func (a *authorizer) Add(req *http.Request) {}

func InitTwt(consumerToken, consumerSecret, accessToken, accessTokenSecret string) *Twt {
	config := oauth1.NewConfig(consumerToken, consumerSecret)
	httpClient := config.Client(oauth1.NoContext, &oauth1.Token{
		Token:       accessToken,
		TokenSecret: accessTokenSecret,
	})
	client := &twitter.Client{
		Authorizer: &authorizer{},
		Client:     httpClient,
		Host:       "https://api.twitter.com",
	}
	return &Twt{
		client: client,
	}
}

func (t *Twt) PostTweet(content string) (string, error) {
	if content == "" {
		return "", errors.New("Content is empty. Please write a tweet")
	}

	req := twitter.CreateTweetRequest{
		Text: content,
	}
	fmt.Println("Callout to create tweet callout")

	tweetResponse, err := t.client.CreateTweet(context.Background(), req)
	if err != nil {
		return "", err
	}

	return tweetResponse.Tweet.ID, nil
}
