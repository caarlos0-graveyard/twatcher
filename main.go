package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/caarlos0/twatcher/feed"
	"github.com/urfave/cli"
)

const version = "dev"

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
	app.Action = func(c *cli.Context) {
		if c.NumFlags() == 0 {
			cli.ShowAppHelp(c)
			return
		}
		names := c.StringSlice("name")
		filter := c.String("filter")
		url := c.String("feed")
		fmt.Println(
			"Watching", url,
			"for", strings.Join(names, ", "),
			"with filter", filter,
			"Press CTRL+C to stop.",
		)
		feed.NewFeed(url, filter, names).Poll()
	}
	app.Run(os.Args)
}
