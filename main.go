package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	"github.com/urfave/cli"

	"./kusokora"
	"./twitter"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "dbname, d",
			Value: "./kusokora.db",
		},
		cli.StringFlag{
			Name:  "conf, c",
			Value: "./config.json",
		},
	}
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add to db",
			Action: func(c *cli.Context) error {
				// TODO default dbname from flag
				dbname := c.String("dbname")
				if dbname == "" {
					dbname = "./kusokora.db"
				}
				db, e := sql.Open("sqlite3", dbname)
				if e != nil {
					return e
				}
				defer db.Close()
				s := kusokora.NewKusokoraService(
					kusokora.NewKusokoraRepositoryOnSQLite(db),
				)
				mu, e := url.ParseRequestURI(c.Args().First())
				if e != nil {
					return e
				}
				e = s.AddKusokora(mu.String())
				if e != nil {
					return e
				}
				return nil
			},
		},
		{
			Name:    "crawl",
			Aliases: []string{"c"},
			Usage:   "crawl from twitter",
			Action: func(c *cli.Context) error {
				dbname := c.String("dbname")
				if dbname == "" {
					dbname = "./kusokora.db"
				}
				db, e := sql.Open("sqlite3", dbname)
				if e != nil {
					return e
				}
				defer db.Close()
				confpath := c.String("dbname")
				if confpath == "" {
					confpath = "./config.json"
				}

				conf, e := loadConfig(confpath)
				if e != nil {
					return e
				}

				r := kusokora.NewKusokoraRepositoryOnSQLite(db)
				ts, e := loadKusokoraFromTwitter(conf)
				if e != nil {
					return e
				}
				for _, t := range ts {
					e := r.Put(t)
					if e != nil {
						log.Println(e.Error())
					}
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}

func loadKusokoraFromTwitter(c *Config) ([]kusokora.Kusokora, error) {
	cli := twitter.NewClient(c.Twitter.ConsumerKey, c.Twitter.ConsumerSecret, c.Twitter.AccessToken, c.Twitter.AccessTokenSecret)
	ts, e := cli.Query("#papixクソコラグランプリ")
	if e != nil {
		return nil, e
	}
	ks := []kusokora.Kusokora{}
	for _, t := range ts {
		ks = append(ks, t.ToKusokora())
	}

	return ks, nil
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
