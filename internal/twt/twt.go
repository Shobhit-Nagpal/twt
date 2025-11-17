package twt

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
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

func InitTwt(config *clientcredentials.Config) *Twt {
	httpClient := config.Client(oauth2.NoContext)
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
