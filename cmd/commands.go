package main

import (
	"fmt"
	"net/url"

	"github.com/codegangsta/cli"
	"github.com/jrudio/go-plex-client"
)

type commands struct{}

func (cmd *commands) test(c *cli.Context) error {
	_url := c.GlobalString("url")

	if _url == "" {
		fmt.Println("Missing required url")
		return nil
	}

	_, urlErr := url.ParseRequestURI(_url)

	// Make sure url is valid
	if urlErr != nil {
		fmt.Println(urlErr.Error())
		return nil
	}

	Plex, PlexErr := plex.New(_url, "")

	if PlexErr != nil {
		fmt.Println(PlexErr.Error())
		return nil
	}

	fmt.Println("Testing connection to Plex")

	result, resErr := Plex.Test()

	if resErr != nil {
		fmt.Println(resErr.Error())
		return nil
	}

	if result {
		fmt.Println("Connection to Plex successful")
	} else {
		fmt.Println("Connect to Plex failed")
	}

	return nil
}

func (cmd *commands) endTranscode(c *cli.Context) error {
	_url := c.GlobalString("url")
	sessionKey := c.Args().First()

	if _url == "" {
		fmt.Println("Missing required url")
		return nil
	}

	if sessionKey == "" {
		fmt.Println("Missing required session key")
		return nil
	}

	_, urlErr := url.ParseRequestURI(_url)

	// Make sure url is valid
	if urlErr != nil {
		fmt.Println(urlErr.Error())
		return nil
	}

	Plex, PlexErr := plex.New(_url, "")

	if PlexErr != nil {
		fmt.Println(PlexErr.Error())
		return nil
	}

	result, err := Plex.KillTranscodeSession(sessionKey)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Println(result)

	return nil
}
