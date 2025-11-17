package twt

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
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

func InitTwt(config *oauth1.Config, token *oauth1.Token) *Twt {
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	return &Twt{
		client: client,
	}
}

func (t *Twt) Post(content string) error {
	if content == "" {
		return errors.New("Content is empty. Please write a tweet")
	}

	_, resp, err := t.client.Statuses.Update(content, nil)
	if err != nil {
		return err
	}

	fmt.Println(resp)
	return nil
}
