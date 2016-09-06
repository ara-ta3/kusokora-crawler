package twitter

import (
	"../kusokora"
	"github.com/ChimeraCoder/anaconda"
)

type ClientOnTwitter struct {
	Api *anaconda.TwitterApi
}

type Client interface {
	Query(q string) ([]Tweet, error)
}

func NewClient(consumerKey, consumerSecret, accessToken, accessTokenSecret string) ClientOnTwitter {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	return ClientOnTwitter{
		Api: api,
	}
}

func (cli *ClientOnTwitter) Query(q string) ([]Tweet, error) {
	sr, e := cli.Api.GetSearch(q, nil)
	if e != nil {
		return nil, e
	}
	tweets := []Tweet{}
	for _, tweet := range sr.Statuses {
		for _, m := range tweet.Entities.Media {
			tweets = append(tweets, Tweet{
				MediaURL:   m.Media_url_https,
				TwitterURL: m.Url,
			})
		}
	}

	return tweets, nil
}

type Tweet struct {
	MediaURL   string
	TwitterURL string
}

func (t *Tweet) ToKusokora() kusokora.Kusokora {
	return kusokora.Kusokora{
		PictureURL: t.MediaURL,
	}
}
