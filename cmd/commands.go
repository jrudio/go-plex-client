package main

import (
	"fmt"
	"github.com/urfave/cli"
)

type commands struct{}

func (cmd *commands) test(c *cli.Context) error {
	initPlex(c)

	fmt.Println("Testing connection to Plex...")

	result, resErr := plexConn.Test()

	if resErr != nil {
		fmt.Println(resErr.Error())
		return nil
	}

	if !result {
		fmt.Println("Connect to Plex failed")
		return nil
	}

	fmt.Println("Connection to Plex successful")

	return nil
}

func (cmd *commands) endTranscode(c *cli.Context) error {
	initPlex(c)

	sessionKey := c.Args().First()

	if sessionKey == "" {
		fmt.Println("Missing required session key")
		return nil
	}

	result, err := plexConn.KillTranscodeSession(sessionKey)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Println(result)

	return nil
}

func (cmd *commands) getServersInfo(c *cli.Context) error {
	initPlex(c)

	info, err := plexConn.GetServersInfo()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Friendly name:", info.FriendlyName)
	fmt.Println("Machine id:", info.MachineIdentifier)
	fmt.Println("Size:", info.Size)
	fmt.Println("Servers:")

	for _, server := range info.Server {
		fmt.Println("\tName:", server.Name)
		fmt.Println("\tHost:", server.Host)
		fmt.Println("\tMachine id:", server.MachineIdentifier)
		fmt.Println("\tLocal address:", server.LocalAddresses)
		fmt.Println("\tScheme:", server.Scheme)
		fmt.Println("\tPort:", server.Port)
		fmt.Println("\tAddress:", server.Address)
		fmt.Println("\tCreated at:", server.CreatedAt)
		fmt.Println("\tUpdated at:", server.UpdatedAt)
		fmt.Println("\tVersion:", server.Version)
		fmt.Println("\tAccess token:", server.AccessToken)
		fmt.Println("\tOwned:", server.Owned)
		fmt.Println("\t=========================")
	}

	return nil
}
