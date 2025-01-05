package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jrudio/go-plex-client/cmd/commands"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:   "test",
				Usage:  "Test your connection to your Plex server",
				Action: testConnection,
			},
			{
				Name:   "auth",
				Usage:  "Authenticate your plexctl client",
				Commands: []*cli.Command{
					{
						Name: "login",
						Action: commands.Login,
					},
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Printf("failed to start: %v\n", err)
		os.Exit(1)
	}
}
