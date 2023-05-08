package main

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/jrudio/go-plex-client"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

const (
	errKeyNotFound        = "Key not found"
	errNoPlexToken        = "no plex auth token in datastore"
	errPleaseSignIn       = "use command 'signin' or 'link-app' to authorize us"
	errNoPlexServerInfo   = "no plex server in datastore"
	errPleaseChooseServer = "use command 'pick-server' to pick a plex server :)"
)

func initPlex(db store, checkForToken bool, checkForServerInfo bool) (*plex.Plex, error) {
	var plexConn *plex.Plex

	if !checkForToken && !checkForServerInfo {
		return plex.New("", "abc123")
	} else if checkForToken && !checkForServerInfo {
		plexToken, err := db.getPlexToken()

		if err != nil && err.Error() == errKeyNotFound {
			return plexConn, fmt.Errorf("%s\n%s", errNoPlexToken, errPleaseSignIn)
		} else if err != nil {
			return plexConn, fmt.Errorf("failed getting plex token: %v", err)
		}

		return plex.New("", plexToken)
	} else if !checkForToken && checkForServerInfo {
		// why would a caller use this?
		// we'll just capture the edge-case?
		return plexConn, fmt.Errorf("wait what")
	}

	plexToken, err := db.getPlexToken()

	if err != nil {
		return plexConn, fmt.Errorf("failed getting plex token: %v", err)
	}

	plexServer, err := db.getPlexServer()

	if err != nil && err.Error() == errKeyNotFound {
		return plexConn, fmt.Errorf("%s\n%s", errNoPlexServerInfo, errPleaseChooseServer)
	} else if err != nil {
		return plexConn, fmt.Errorf("failed getting plex server info from data store: %v", err)
	}

	return plex.New(plexServer.URL, plexToken)
}

func test(c *cli.Context) error {
	args := c.Args()

	// we need a url and an auth token
	if len(args) < 2 {
		return cli.NewExitError("a url and a token is required", 1)
	}

	var host string
	var token string

	// check if either argument is a url
	hostParsed, err := url.Parse(args[0])

	if err != nil {
		_host, err := url.Parse(args[1])

		if err != nil {
			return cli.NewExitError("a valid url is required", 1)
		}

		host = _host.String()
		token = args[0]
	} else {
		host = hostParsed.String()
		token = args[1]
	}

	plexConn, err := plex.New(host, token)

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	fmt.Println("testing connection to plex...")

	result, err := plexConn.Test()

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	if !result {
		fmt.Println("failed to connect to plex")
		return nil
	}

	fmt.Println("successfully connected to plex")

	return nil
}

func endTranscode(c *cli.Context) error {
	db, err := startDB()

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	defer db.Close()

	plexConn, err := initPlex(db, true, true)

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	sessionKey := c.Args().First()

	if sessionKey == "" {
		return cli.NewExitError("Missing required session key", 1)
	}

	result, err := plexConn.KillTranscodeSession(sessionKey)

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	fmt.Println(result)

	return nil
}

func getServersInfo(c *cli.Context) error {
	db, err := startDB()

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	defer db.Close()

	plexConn, err := initPlex(db, true, false)

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	info, err := plexConn.GetServersInfo()

	if err != nil {
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

func startDB() (store, error) {
	// create persistent key store in user home directory
	storeDirectory, err := homedir.Dir()

	if err != nil {
		return store{}, err
	}

	storeDirectory = filepath.Join(storeDirectory, homeFolderName)

	return initDataStore(storeDirectory)
}

func getSections(c *cli.Context) error {
	db, err := startDB()

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	defer db.Close()

	plexConn, err := initPlex(db, true, true)

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	// Grab machine id of the server we are connected to
	machineID, err := plexConn.GetMachineID()

	if err != nil {
		return fmt.Errorf("failed to retrieve machine id of plex server: %v", err)
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

func getSessions(c *cli.Context) error {
	db, err := startDB()

	if err != nil {
		return cli.NewExitError(fmt.Sprintf("start db failed: %v\n", err), 1)
	}

	defer db.Close()

	plexConn, err := initPlex(db, true, true)

	if err != nil {
		return cli.NewExitError("failed to initialize plex: "+err.Error(), 1)
	}

	// display sessions
	sessions, err := plexConn.GetSessions()

	if err != nil {
		return cli.NewExitError("failed to get sessions: "+err.Error(), 1)
	}

	if sessions.MediaContainer.Size == 0 {
		fmt.Println("no users in sessions")
		return nil
	}

	for _, session := range sessions.MediaContainer.Metadata {
		fmt.Print(session.User.Title)
		userIsWatching := "\t" + session.Session.ID + " (" + session.Type + ") "

		if session.GrandparentTitle != "" {
			userIsWatching += session.GrandparentTitle + " - " + session.ParentTitle
			userIsWatching += " - " + session.Title
		} else {
			userIsWatching += session.Title + " (" + string(session.Year) + ")"
		}

		fmt.Println(userIsWatching)
	}

	return nil
}

func authorizeApp(c *cli.Context) error {
	db, err := startDB()

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	defer db.Close()

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

	code := c.Args().First()
	codeLen := len(code)

	fmt.Println("code", code)

	if codeLen < 1 || codeLen > 4 {
		return errors.New("a 4 character code is required")
	}

	fmt.Println("attempting to link app with code " + code + "...")

	if err := plexConn.LinkAccount(code); err != nil {
		return err
	}

	fmt.Println("successfully linked app, enjoy!")

	return nil
}

// linkApp give user 4 character code that can be used authorized our app
func linkApp(c *cli.Context) error {
	db, err := startDB()

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	defer db.Close()

	// just need headers
	plexConn, err := initPlex(db, false, false)

	if err != nil {
		return cli.NewExitError(fmt.Sprintf("could not create headers: %v", err), 1)
	}

	info, err := plex.RequestPIN(plexConn.Headers)

	if err != nil {
		return cli.NewExitError("request plex pin failed: "+err.Error(), 1)
	}

	expireAtParsed, err := time.Parse(time.RFC3339, info.ExpiresAt)

	if err != nil {
		return cli.NewExitError(fmt.Sprintf("could not get expiration for plex pin: %v", err), 1)
	}

	expires := time.Until(expireAtParsed).String()

	fmt.Printf("your pin %s and expires in %s\n", info.Code, expires)

	var authToken string

	for {
		pinInformation, err := plex.CheckPIN(info.ID, plexConn.ClientIdentifier)

		if err != nil {
			fmt.Printf("\r%v", err)
		}

		// expiresAt := pinInformation.ExpiresAt

		// stop checking if time is expired
		// if time.Until(time.Unix(int64(expiresAt), 0)).Minutes() < 0 {
		// 	return errors.New("code has expired - please request another one")
		// }

		if pinInformation.AuthToken != "" {
			authToken = pinInformation.AuthToken
			break
		}

		time.Sleep(1 * time.Second)
	}

	fmt.Printf("\ryou have been successfully authorized!\nYour auth token is: %s\n", authToken)

	fmt.Println("saving plex token to disk...")

	if err := db.savePlexToken(authToken); err != nil {
		return cli.NewExitError(fmt.Sprintf("saving plex token failed: %v", err), 1)
	}

	fmt.Println("saved plex token!")

	return nil
}

func pickServer(c *cli.Context) error {
	db, err := startDB()

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	defer db.Close()

	plexConn, err := initPlex(db, true, false)

	if err != nil {
		return cli.NewExitError(err, 1)
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

// signIn displays the auth token on successful sign in
func signIn(c *cli.Context) error {
	db, err := startDB()

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	defer db.Close()

	if c.NArg() != 2 {
		return cli.NewExitError("signin requires 2 arguments - username and password", 1)
	}

	username := c.Args()[0]
	password := c.Args()[1]

	plexConn, err := plex.SignIn(username, password)

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	if plexConn.Token == "" {
		return cli.NewExitError("failed to receive a plex token", 1)
	}

	// fmt.Println("your auth token is:", plexConn.Token)
	fmt.Println("successfully signed in!")

	if isVerbose {
		fmt.Println("saving token locally...")
	}

	if err := db.savePlexToken(plexConn.Token); err != nil {
		return cli.NewExitError(err, 1)
	}

	return nil
}

func getLibraries(c *cli.Context) error {
	db, err := startDB()

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	defer db.Close()

	plexConn, err := initPlex(db, true, true)

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	fmt.Println("getting libraries...")

	libraries, err := plexConn.GetLibraries()

	if err != nil {
		return cli.NewExitError(fmt.Sprintf("failed fetching libraries: %v", err), 1)
	}

	for _, dir := range libraries.MediaContainer.Directory {
		fmt.Println(dir.Title)
	}

	return nil
}

func webhooks(c *cli.Context) error {
	db, err := startDB()

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	defer db.Close()

	plexConn, err := initPlex(db, true, false)

	if err != nil {
		return cli.NewExitError(err, 1)
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

func search(c *cli.Context) error {
	db, err := startDB()

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	defer db.Close()

	plexConn, err := initPlex(db, true, true)

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	title := strings.Join(c.Args(), " ")

	fmt.Println("searching plex server for " + title)

	results, err := plexConn.Search(title)

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	if len(results.MediaContainer.Metadata) == 0 {
		fmt.Println("could not find '" + title + "'")
		return nil
	}

	for _, searchResult := range results.MediaContainer.Metadata {
		fmt.Println(searchResult.Title)
	}

	return nil
}

func getEpisode(c *cli.Context) error {
	db, err := startDB()

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	defer db.Close()

	plexConn, err := initPlex(db, true, true)

	if err != nil {
		return err
	}

	if c.NArg() == 0 {
		return cli.NewExitError("episode id is required", 1)
	}

	result, err := plexConn.GetEpisode(c.Args().First())

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	episode := result.MediaContainer.Metadata

	if len(episode) == 0 {
		return cli.NewExitError("no episodes found", 1)
	}

	fmt.Println(episode[0].GrandparentTitle + ": " + episode[0].Title)

	return nil
}

func getOnDeck(c *cli.Context) error {
	db, err := startDB()

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	defer db.Close()

	plexConn, err := initPlex(db, true, true)

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	results, err := plexConn.GetOnDeck()

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	for _, result := range results.MediaContainer.Metadata {
		fmt.Println(result.Title, result.Rating)
	}

	return nil
}

func unlock(c *cli.Context) error {
	storeDirectory, err := homedir.Dir()

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	lockFilePath := filepath.Join(storeDirectory, homeFolderName, "LOCK")

	if err := os.Remove(lockFilePath); err != nil {
		return cli.NewExitError(fmt.Sprintf("failed to remove file: %v", err), 1)
	}

	fmt.Println("removed LOCK file")

	return nil
}

func stopPlayback(c *cli.Context) error {
	db, err := startDB()

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	defer db.Close()

	plexConn, err := initPlex(db, true, true)

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	sessions, err := plexConn.GetSessions()

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	sessionCount := sessions.MediaContainer.Size

	if sessionCount < 1 {
		fmt.Println("no users in session")

		return nil
	}

	fmt.Println("current sessions:")

	for i, session := range sessions.MediaContainer.Metadata {

		title := ""

		if session.GrandparentTitle != "" {
			title += session.GrandparentTitle + " - " + session.ParentTitle
			title += " - " + session.Title
		} else {
			title += session.Title + " (" + string(session.Year) + ")"
		}

		fmt.Printf("\t[%d] %s - %s\n", i, session.User.Title, title)
	}

	fmt.Println("choose a session to stop:")

	var sessionIndex int

	fmt.Scanln(&sessionIndex)

	// bound check user input
	if sessionIndex < 0 || sessionIndex > sessionCount-1 {
		return cli.NewExitError("invalid selection", 1)
	}

	selectedSession := sessions.MediaContainer.Metadata[sessionIndex]

	sessionID := selectedSession.Session.ID

	if err := plexConn.TerminateSession(sessionID, "stream stopped by github.com/jrudio/go-plex-client"); err != nil {
		return cli.NewExitError(err, 1)
	}

	fmt.Printf("sucessfully stopped %s from continuing %s\n", selectedSession.Title, selectedSession.ParentTitle)

	return nil
}

func getAccountInfo(c *cli.Context) error {
	db, err := startDB()

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	defer db.Close()

	plexConn, err := initPlex(db, true, false)

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	plexConn.HTTPClient.Timeout = time.Minute * 1

	account, err := plexConn.MyAccount()

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	// fmt.Println(account.Subscription.Features, account.Roles.Roles)
	// fmt.Println(account.Entitlements)

	fmt.Printf("%+v\n", account)

	return nil
}

func getMetadata(c *cli.Context) error {
	db, err := startDB()

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	defer db.Close()

	plexConn, err := initPlex(db, true, true)

	if err != nil {
		return err
	}

	if c.NArg() == 0 {
		return cli.NewExitError("episode id is required", 1)
	}

	result, err := plexConn.GetMetadata(c.Args().First())

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	fmt.Println(result)

	return nil
}

func downloadMedia(c *cli.Context) error {
	db, err := startDB()

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	defer db.Close()

	plexConn, err := initPlex(db, true, true)

	if err != nil {
		return err
	}

	if c.NArg() == 0 {
		return cli.NewExitError("search term is required", 1)
	}

	downloadPath := c.Args().Get(2)

	if downloadPath == "" {
		downloadPath = "."
	}

	createFolders := c.Bool("folders")

	skipIfExists := c.Bool("skip")

	// search for media
	results, err := plexConn.Search(c.Args().First())

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	if len(results.MediaContainer.Metadata) == 0 {
		return cli.NewExitError("no results found", 1)
	}

	// prompt user for media selection
	fmt.Println("results:")

	for i, result := range results.MediaContainer.Metadata {
		fmt.Printf("\t[%d] %s\n", i, result.Title)
	}

	// we use -1 to indicate that the user has not selected a media
	selection := -1

	fmt.Printf("choose media to download:")
	fmt.Scanln(&selection)

	// bound check user input
	if selection < 0 || selection > len(results.MediaContainer.Metadata)-1 {
		return cli.NewExitError("invalid selection", 1)
	}

	selectedMedia := results.MediaContainer.Metadata[selection]

	// download media
	if err := plexConn.Download(selectedMedia, downloadPath, createFolders, skipIfExists); err != nil {
		return cli.NewExitError(err, 1)
	}

	fmt.Printf("successfully downloaded %s\n", selectedMedia.Title)

	return nil
}

func getPlaylist(c *cli.Context) error {
	db, err := startDB()

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	defer db.Close()

	plexConn, err := initPlex(db, true, true)

	if err != nil {
		return err
	}

	if c.NArg() == 0 {
		return cli.NewExitError("playlist id is required", 1)
	}

	playlistID, err := strconv.ParseInt(c.Args().First(), 10, 64)

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	result, err := plexConn.GetPlaylist(int(playlistID))

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	fmt.Println(result)

	return nil
}

func deleteMedia(c *cli.Context) error {
	db, err := startDB()

	if err != nil {
		return cli.NewExitError(err, 1)
	}

	defer db.Close()

	plexConn, err := initPlex(db, true, true)

	if err != nil {
		return err
	}

	if c.NArg() == 0 {
		return cli.NewExitError("media id is required", 1)
	}

	mediaID := c.Args().First()

	if err := plexConn.DeleteMediaByID(mediaID); err != nil {
		return cli.NewExitError(err, 1)
	}

	fmt.Printf("successfully deleted media '%s'\n", mediaID)

	return nil
}
