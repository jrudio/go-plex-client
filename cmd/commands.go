package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"

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

func getSections(db store) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		plexToken, err := db.getPlexToken()

		if err != nil {
			return fmt.Errorf("failed getting plex token: %v", err)
		}

		plexServer, err := db.getPlexServer()

		if err != nil {
			return fmt.Errorf("failed getting plex server info from data store: %v", err)
		}

		plexConn, err := plex.New(plexServer.URL, plexToken)

		if err != nil {
			return fmt.Errorf("failed to create plex instance: %v", err)
		}

		// Grab machine id of the server we are connected to
		machineID, err := plexConn.GetMachineID()

		if err != nil {
			return fmt.Errorf("failed to retrieve machine id of plex server (%s): %v", plexServer.Name, err)
		}

		sections, err := plexConn.GetSections(machineID)

		if err != nil {
			return fmt.Errorf("failed to retrieve sections: %v", err)
		}

		fmt.Println("section count:", len(sections))

		if len(sections) < 1 {
			return errors.New("sections not found")
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
}

func getSessions(db store) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		plexToken, err := db.getPlexToken()

		if err != nil {
			return cli.NewExitError("failed to get plex token from datastore: "+err.Error(), 1)
		}

		plexServer, err := db.getPlexServer()

		if err != nil {
			return cli.NewExitError("failed to get plex server info from datastore: "+err.Error(), 1)
		}

		plexConn, err := plex.New(plexServer.URL, plexToken)

		if err != nil {
			return cli.NewExitError("failed to initialize plex: "+err.Error(), 1)
		}

		// display sessions
		sessions, err := plexConn.GetSessions()

		if err != nil {
			return cli.NewExitError("failed to get sessions: "+err.Error(), 1)
		}

		if len(sessions.MediaContainer.Video) == 0 && len(sessions.MediaContainer.Track) == 0 {
			fmt.Println("no users in sessions")
			return nil
		}

		for _, session := range sessions.MediaContainer.Video {
			fmt.Print(session.User.Title)
			userIsWatching := "\t(" + session.Type + ") "

			if session.GrandparentTitle != "" {
				userIsWatching += session.GrandparentTitle + " - " + session.ParentTitle
				userIsWatching += " - " + session.Title
			} else {
				userIsWatching += session.Title + " (" + session.Year + ")"
			}

			fmt.Println(userIsWatching)
		}

		for _, session := range sessions.MediaContainer.Track {
			fmt.Print(session.User.Title)
			fmt.Println("\t", "("+session.Type+")", session.Title)
		}

		return nil
	}
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

// requestPIN is good for just receiving the pin and you manually going to plex.tv/link to link the code
func requestPIN(c *cli.Context) error {
	info, err := plex.RequestPIN()

	if err != nil {
		return errors.New("request plex pin failed: " + err.Error())
	}

	expires := time.Until(time.Unix(info.ExpiresAt, 0)).String()

	fmt.Printf("your pin %s (%d) expires in %s", info.Code, info.ID, expires)

	return nil
}

// checkPIN will check the status of a pin/code via the id given in requestPIN. It will display the auth token when authorized
func checkPIN(c *cli.Context) error {
	idArg := c.Args().First()

	id, err := strconv.ParseInt(idArg, 0, 64)

	if err != nil {
		return errors.New("failed to parse id: " + err.Error())
	}

	clientID := c.String("client-id")

	if clientID == "" {
		return errors.New("client-id is required")
	}

	var authToken string

	for {
		pinInformation, err := plex.CheckPIN(int(id), clientID)

		if err != nil {
			fmt.Printf("\r%v", err)
		}

		expiresAt := pinInformation.ExpiresAt

		// stop checking if time is expired
		if time.Until(time.Unix(expiresAt, 0)).Minutes() < 0 {
			return errors.New("code has expired - please request another one")
		}

		if pinInformation.AuthToken != "" {
			authToken = pinInformation.AuthToken
			break
		}

		// just check once
		if !c.Bool("poll") {
			// still not authorized
			return nil
		}

		time.Sleep(1 * time.Second)
	}

	fmt.Printf("\ryou have been successfully authorized!\nYour auth token is: %s\n", authToken)

	return nil
}

func pickServer(db store) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		// look up servers - hopefully a token is already in store
		plexToken, err := db.getPlexToken()

		if err != nil {
			return fmt.Errorf("failed to retreive plex token: %v", err)
		}

		plexConn, err := plex.New("", plexToken)

		if err != nil {
			return err
		}

		// load list of servers
		servers, err := plexConn.GetServers()

		if err != nil {
			return fmt.Errorf("failed getting plex servers: %v", err)
		}

		fmt.Println("Server list:")

		for i, server := range servers {
			fmt.Printf("[%d] - %s\n", i, server.Name)
		}

		fmt.Print("\nSelect a server: ")

		var serverIndex int
		fmt.Scanln(&serverIndex)

		// bound check input
		if serverIndex < 0 || serverIndex > (len(servers)-1) {
			return errors.New("invalid selection")
		}

		selectedServer := servers[serverIndex]

		// choose to connect via local or remote
		fmt.Printf("\nshowing local and remote addresses for %s:\n", selectedServer.Name)

		for i, conn := range selectedServer.Connection {
			fmt.Printf("\t[%d] uri: %s, is local: %t\n", i, conn.Address, conn.Local == 1)
		}

		fmt.Print("\nPick the appropriate address: ")

		var urlIndex int
		fmt.Scanln(&urlIndex)

		// bound check again
		if urlIndex < 0 || urlIndex > (len(selectedServer.Connection)-1) {
			return errors.New("invalid selection")
		}

		// persist selection to disk
		fmt.Printf("\nsetting %s as the default server using url %s...\n", selectedServer.Name, selectedServer.Connection[urlIndex].URI)

		if err := db.savePlexServer(server{
			Name: selectedServer.Name,
			URL:  selectedServer.Connection[urlIndex].URI,
		}); err != nil {
			return fmt.Errorf("failed to save server info: %v", err)
		}

		fmt.Println("success!")

		return nil
	}
}

// signIn displays the auth token on successful sign in
func signIn(db store) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		if c.NArg() != 2 {
			return errors.New("signin requires 2 arguments - username and password")
		}

		username := c.Args()[0]
		password := c.Args()[1]

		plexConn, err := plex.SignIn(username, password)

		if err != nil {
			return err
		}

		if plexConn.Token == "" {
			return errors.New("failed to receive a plex token")
		}

		// fmt.Println("your auth token is:", plexConn.Token)
		fmt.Println("successfully signed in!")

		if isVerbose {
			fmt.Println("saving token locally...")
		}

		if err := db.savePlexToken(plexConn.Token); err != nil {
			return err
		}

		return nil
	}
}

func getLibraries(db store) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		plexToken, err := db.getPlexToken()

		if err != nil {
			return fmt.Errorf("failed to retreive plex token: %v", err)
		}

		plexServer, err := db.getPlexServer()

		if err != nil {
			return fmt.Errorf("failed to retreive plex server url: %v", err)
		}

		plexConn, err := plex.New(plexServer.URL, plexToken)

		if err != nil {
			return err
		}

		fmt.Println("getting libraries...")

		libraries, err := plexConn.GetLibraries()

		if err != nil {
			return err
		}

		for _, dir := range libraries.MediaContainer.Directory {
			fmt.Println(dir.Title)
		}

		return nil
	}
}

func webhooks(db store) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		plexToken, err := db.getPlexToken()

		if err != nil {
			return cli.NewExitError(fmt.Sprintf("failed to retreive plex token: %v", err), 1)
		}

		plexConn, err := plex.New("", plexToken)

		if err != nil {
			return err
		}

		// create webhook
		newWebhook := c.String("add")

		if newWebhook != "" {
			fmt.Println("adding new webhook:", newWebhook)

			if err := plexConn.AddWebhook(newWebhook); err != nil {
				return cli.NewExitError(fmt.Sprintf("adding webhook failed: %v", err), 1)
			}

			fmt.Println("success!")

			return nil
		}

		fmt.Println("displaying webhooks...")

		hooks, err := plexConn.GetWebhooks()

		if err != nil {
			return cli.NewExitError(err, 1)
		}

		for i, hook := range hooks {
			fmt.Printf("\t[%d] %s\n", i, hook)
		}

		// delete webhook
		if c.Bool("delete") {
			var index int

			fmt.Print("enter a number to delete that webhook: ")
			fmt.Scan(&index)

			bounds := len(hooks) - 1

			if index < bounds || index > bounds {
				return cli.NewExitError("invalid input", 1)
			}

			fmt.Printf("deleting webhook %s at index %d...\n", hooks[index], index)

			hooks = append(hooks[:index], hooks[index+1:]...)

			if err := plexConn.SetWebhooks(hooks); err != nil {
				return cli.NewExitError(fmt.Sprintf("failed to set webhooks: %v", err), 1)
			}

			fmt.Println("success!")
		}

		return nil
	}
}
