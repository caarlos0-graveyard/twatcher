package main

import (
	"fmt"
	"os"
	"strings"

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
		fmt.Println(
			"Watching",
			strings.Join(c.StringSlice("show"), ", ")+".",
			"Press CTRL+C to stop.",
		)
		feed.NewFeed(
			c.String("feed"),
			c.StringSlice("show"),
		).Poll()
	}
	app.Run(os.Args)
}
