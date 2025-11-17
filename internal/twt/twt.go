package twt

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	twitter "github.com/g8rswimmer/go-twitter/v2"
)

type Twt struct {
	client *twitter.Client
}

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

func InitTwt(token string) *Twt {
	client := &twitter.Client{
		Authorizer: authorize{
			Token: token,
		},
		Client: http.DefaultClient,
		Host:   "https://api.twitter.com",
	}

	return &Twt{
		client: client,
	}
}

func (t *Twt) Post(content string) error {
	if content == "" {
		return errors.New("Content is empty. Please write a tweet")
	}

	tweetReq := twitter.CreateTweetRequest{
		Text: content,
	}
	resp, err := t.client.CreateTweet(context.Background(), tweetReq)
	if err != nil {
		return err
	}

	enc, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(string(enc))
	return nil
}
