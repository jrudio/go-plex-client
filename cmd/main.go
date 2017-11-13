package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/jrudio/go-plex-client"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

var (
	cmd       = commands{}
	title     string
	baseURL   string
	token     string
	isVerbose bool
	plexConn  *plex.Plex
	// appSecret will used to seed encrypt()
	appSecret = []byte("iAmAseCReTuSEdTOENcrYp")
)

type server struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (s server) Serialize() ([]byte, error) {
	return json.Marshal(s)
}

func unserializeServer(serializedServer []byte) (server, error) {
	var s server

	err := json.Unmarshal(serializedServer, &s)

	return s, err
}

func main() {
	app := cli.NewApp()

	app.Name = "plex-cli"
	app.Usage = "Interact with your plex server and plex.tv from the command line"
	app.Version = "0.0.1"

	storeDirectory, err := homedir.Dir()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	storeDirectory = filepath.Join(storeDirectory, ".plex-cli")

	db, err := initDataStore(storeDirectory)

	if err != nil {
		fmt.Printf("data store initialization failed: %v\n", err)
		os.Exit(1)
	}

	// check if app secret exists in datastore
	if secret := db.getSecret(); len(secret) != 0 {
		db.secret = secret
	} else {
		// if not, create new appsecret and save
		// append a random number to make appsecret unique
		appSecret = append(appSecret, []byte(strconv.FormatInt(time.Now().Unix(), 10))...)

		if err := db.saveSecret(appSecret); err != nil {
			fmt.Printf("failed to save app secret: %v\n", err)
			db.Close()
			os.Exit(1)
		}

		db.secret = appSecret
	}

	shutdown := make(chan int, 1)

	wg := sync.WaitGroup{}

	cli.OsExiter = func(c int) {
		shutdown <- c
	}

	wg.Add(1)

	go func(shutdown chan int) {
		defer wg.Done()
		exitCode := <-shutdown
		// fmt.Printf("\nexiting with code %d...\n", exitCode)
		db.Close()
		os.Exit(exitCode)
	}(shutdown)

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
			Name:   "test",
			Usage:  "Test your connection to your Plex Media Server",
			Action: cmd.test,
		},
		{
			Name:   "end",
			Usage:  "End a transcode session",
			Action: cmd.endTranscode,
		},
		{
			Name:   "server-info",
			Usage:  "Print info about your servers - ip, machine id, access tokens, etc",
			Action: cmd.getServersInfo,
		},
		{
			Name:   "sections",
			Usage:  "Print info about your server's sections",
			Action: getSections(db),
		},
		{
			Name:  "link",
			Usage: "authorize an app (e.g. amazon fire app) with a 4 character `code`. REQUIRES a plex token",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "token",
					Value: "",
					Usage: "plex token required to link an app to your account. (e.g. `abc123`",
				},
			},
			Action: linkApp,
		},
		{
			Name:   "library",
			Usage:  "display your libraries",
			Action: getLibraries(db),
		},
		{
			Name:   "request-pin",
			Usage:  "request a pin (4 character code) from plex.tv to link account to an app. Use this to recieve an id to check for an auth token",
			Action: requestPIN,
		},
		{
			Name:  "check-pin",
			Usage: "check status of pin (4 character code) from plex.tv to link account to an app. Use this to recieve an auth token. REQUIRES an id",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "client-id",
					Value: "goplexclient",
					Usage: "plex token required to link an app to your account. (e.g. `abc123`",
				},
				cli.BoolFlag{
					Name:  "poll",
					Usage: "check every second if pin is authorized",
				},
			},
			Action: checkPIN,
		},
		{
			Name:   "signin",
			Usage:  "use your username and password to receive a plex auth token",
			Action: signIn(db),
		},
		{
			Name:   "sessions",
			Usage:  "display info on users currently consuming media",
			Action: getSessions(db),
		},
		{
			Name:   "pick-server",
			Usage:  "choose a server to interact with",
			Action: pickServer(db),
		},
		{
			Name:   "webhooks",
			Usage:  "display webhooks associated with your account",
			Action: webhooks(db),
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "add",
					Usage: "create a new webhook",
				},
				cli.BoolFlag{
					Name:  "delete",
					Usage: "delete a webhook",
				},
			},
		},
	}

	app.Run(os.Args)

	shutdown <- 0

	wg.Wait()
}

func initPlex(c *cli.Context) {
	var err error

	if plexConn, err = plex.New(c.GlobalString("url"), c.GlobalString("token")); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
