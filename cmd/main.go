package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jrudio/go-plex-client/cmd/commands"
	"github.com/jrudio/go-plex-client/cmd/db"
	"github.com/urfave/cli/v3"
)

const (
	clientIdentifier = "plexctl"
)

var (
	dbConn *db.DB
	dbDir = "plexctl"
	cmds commands.CMDs
)

func init() {
	isDirNotExists := func(dir string) bool {
		_, err := os.Stat(dir)

		return os.IsNotExist(err)
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

	if isDirNotExists(dbDir) {
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

	dbConn = conn

	cmds = commands.New(conn, clientIdentifier)
}

func main() {
	if dbConn == nil {
		return
	}

	defer dbConn.Close()

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
						Action: cmds.Login,
					},
					{
						Name: "list",
						Action: cmds.ListAccounts,
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
