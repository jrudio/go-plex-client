package main

import (
	"os"

	"fmt"
	"github.com/jrudio/go-plex-client"
	"github.com/urfave/cli"
)

var (
	cmd      = commands{}
	title    string
	baseURL  string
	token    string
	plexConn *plex.Plex
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "url, u",
			Usage:       "Plex url or ip",
			Destination: &baseURL,
		},
		cli.StringFlag{
			Name:        "token, tkn",
			Usage:       "abc123",
			Destination: &token,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   "Test your connection to your Plex Media Server",
			Action:  cmd.test,
		},
		{
			Name:   "end",
			Usage:  "End a transcode session",
			Action: cmd.endTranscode,
		},
		{
			Name:    "server-info",
			Aliases: []string{"si"},
			Usage:   "Print info about your servers - ip, machine id, access tokens, etc",
			Action:  cmd.getServersInfo,
		},
	}

	app.Run(os.Args)
}

func initPlex(c *cli.Context) {
	var err error

	if plexConn, err = plex.New(c.GlobalString("url"), c.GlobalString("token")); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
