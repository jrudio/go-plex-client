package main

import (
	"errors"
	"fmt"

	"github.com/jrudio/go-plex-client"
	"github.com/urfave/cli"
)

type commands struct{}

func (cmd *commands) test(c *cli.Context) error {
	initPlex(c)

	fmt.Println("Testing connection to Plex...")

	result, err := plexConn.Test()

	if err != nil {
		fmt.Println(err.Error())
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

func (cmd *commands) getSections(c *cli.Context) error {
	initPlex(c)

	// Grab machine id of the server we are connected to
	machineID, err := plexConn.GetMachineID()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	var sections []plex.ServerSections
	sections, err = plexConn.GetSections(machineID)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	for _, section := range sections {
		fmt.Println("Section title:", section.Title)
		fmt.Println("\tID:", section.ID)
		fmt.Println("\tKey:", section.Key)
		fmt.Println("\tType:", section.Type)
		fmt.Println("\t=========================")
	}

	return nil
}

func linkApp(c *cli.Context) error {
	token := c.String("token")
	tokenLen := len(token)

	fmt.Println("token", token)
	if token == "" || tokenLen <= 4 {
		return errors.New("a plex token is required")
	}

	code := c.Args().First()
	codeLen := len(code)

	fmt.Println("code", code)

	if codeLen < 1 || codeLen > 4 {
		return errors.New("A 4 character code is required")
	}

	fmt.Println("Attempting to link app with code " + code + "...")

	plexConn, err := plex.New("https://plex.tv", token)

	if err != nil {
		return err
	}

	if err := plexConn.LinkAccount(code); err != nil {
		return err
	}

	fmt.Println("Successfully linked app, enjoy!")

	return nil
}
