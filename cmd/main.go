package main

import (
	"os"

	"github.com/codegangsta/cli"
)

var (
	cmd     = commands{}
	title   string
	baseURL string
)

func main() {

	// flag.StringVar(&title, "title", "", "Enter the title of media to search within Plex")

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "url, u",
			Usage:       "Plex url or ip",
			Destination: &baseURL,
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
	}

	app.Run(os.Args)

	// Display the search results
	// results, resultErr := Plex.Search(title)

	// if resultErr != nil {
	// 	log.Println(resultErr.Error())
	// 	return
	// }

	// Last 4 results are not relevant
	// itemCount := len(results.Children) - 4

	// log.Println("Found")
	// log.Println(itemCount)
	// log.Println("items")

	// for ii, r := range results.Children {
	// 	if ii >= itemCount {
	// 		break
	// 	}

	// 	log.Println(r.Title)
	// }
}
