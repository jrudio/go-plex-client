package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jrudio/go-plex-client/cmd/commands"
	"github.com/jrudio/go-plex-client/cmd/db"
	"github.com/urfave/cli/v3"
)

var (
	dbDir = "plexctl"
)

var (
	dbConn *db.DB
)

func init() {
	isDirExists := func(dir string) bool {
		_, err := os.Stat(dir)

		return os.IsExist(err)
	}

	createDir := func(dir string) error {
		if err := os.Mkdir(dir, 0700); err != nil {
			return err
		}

		return nil
	}

	homeDir, err := os.UserHomeDir()

	if err != nil {
		fmt.Printf("could not find user home directory when creating database: %v\n", err)

		return
	}

	dbDir = fmt.Sprintf("%s/%s/%s", homeDir, ".config", dbDir)

	if !isDirExists(dbDir) {
		if err := createDir(dbDir); err != nil {
			fmt.Printf("create directory for database failed: %v\n", err)
			return
		}
	}

  conn, err	:= db.New(dbDir)

	if err != nil {
		fmt.Printf("failed to create database: %v\n", err)
		return
	}

	defer conn.Close()

	dbConn = conn
}

func main() {
	if dbConn == nil {
		return
	}

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
