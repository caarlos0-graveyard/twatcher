package main

import (
	"os"

	"github.com/caarlos0/tvshows/feed"
	"github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "TVShows"
	app.Usage = "Let this binary running, and it will download your favorite tv shows to ~/Downloads"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "feed",
			Usage: "Feed URL to check for new shows",
		},
		cli.StringSliceFlag{
			Name:  "show",
			Usage: "Show you want to watch",
		},
	}
	app.Action = func(c *cli.Context) {
		if c.NumFlags() == 0 {
			cli.ShowAppHelp(c)
			return
		}
		feed.NewFeed(
			c.String("feed"),
			c.StringSlice("shows"),
		).Poll()
	}
	app.Run(os.Args)
}
