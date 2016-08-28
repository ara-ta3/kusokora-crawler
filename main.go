package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"./kusokora"
	"./twitter"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("args must be over 2")
		os.Exit(1)
	}
	db, e := sql.Open("sqlite3", "./kusokora.db")
	if e != nil {
		panic(e)
	}
	defer db.Close()

	s := kusokora.NewKusokoraService(
		kusokora.NewKusokoraRepositoryOnSQLite(db),
	)
	mu := os.Args[1]
	tu := os.Args[2]
	e = s.AddKusokora(mu, tu)
	if e != nil {
		panic(e)
	}
	//     c, e := loadConfig("./config.json")
	//     if e != nil {
	//         panic(e)
	//     }
	//     e = loadKusokoraFromTwitter(c)
	//     if e != nil {
	//         panic(e)
	//     }
}

func loadKusokoraFromTwitter(c *Config) error {
	cli := twitter.NewClient(c.Twitter.ConsumerKey, c.Twitter.ConsumerSecret, c.Twitter.AccessToken, c.Twitter.AccessTokenSecret)
	tweets, _ := cli.Query("#papixクソコラグランプリ")
	fmt.Println(tweets)
	return nil
}

func loadConfig(path string) (*Config, error) {
	f, e := ioutil.ReadFile(path)
	if e != nil {
		return nil, e
	}
	c := Config{}
	json.Unmarshal(f, &c)
	return &c, nil
}

type Config struct {
	Twitter TwitterKeys `json:"twitter"`
}

type TwitterKeys struct {
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
}
