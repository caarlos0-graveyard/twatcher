package main

import (
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/caarlos0/twatcher/feed"
	"github.com/urfave/cli"
)

var version = "dev"

func main() {
	app := cli.NewApp()
	app.Name = "TorrentWatcher"
	app.Usage = "Watch a torrent feed for some expressions and downloads them to ~/Downloads"
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "feed",
			Usage: "Feed URL to watch for",
		},
		cli.StringSliceFlag{
			Name:  "name",
			Usage: "Name to look for",
		},
		cli.StringFlag{
			Name:  "filter",
			Usage: "A filter to use in all names, e.g.: 1080p",
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.NumFlags() == 0 {
			return cli.ShowAppHelp(c)
		}
		names := c.StringSlice("name")
		filter := c.String("filter")
		url := c.String("feed")
		log.WithField("url", url).
			WithField("names", strings.Join(names, ",")).
			WithField("filter", filter).
			Println("Looking for new torrents...")
		return feed.NewFeed(url, filter, names).Poll()
	}
	app.Run(os.Args)
}
