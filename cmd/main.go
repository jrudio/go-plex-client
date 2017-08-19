package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/jrudio/go-plex-client"
	"github.com/urfave/cli"
)

var (
	cmd       = commands{}
	title     string
	baseURL   string
	token     string
	isVerbose bool
	plexConn  *plex.Plex
	appSecret = "iAmAseCReTuSEdTOENcrYptTHangs123"
)

func main() {
	app := cli.NewApp()

	app.Name = "plex-cli"
	app.Usage = "Interact with your plex server and plex.tv from the command line"
	app.Version = "0.0.1"

	// check if app secret already in datastore

	// if not, save new appsecret
	appSecret = appSecret + string(time.Now().Unix())

	// store, err := initDataStore("plex-cli")

	// if err != nil {
	// 	fmt.Printf("data store initialization failed: %v\n", err)
	// 	os.Exit(2)
	// }

	// tokenKey := "token"

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
			Name: "test",
			// Aliases: []string{"t"},
			Usage:  "Test your connection to your Plex Media Server",
			Action: cmd.test,
		},
		{
			Name:   "end",
			Usage:  "End a transcode session",
			Action: cmd.endTranscode,
		},
		{
			Name: "server-info",
			// Aliases: []string{"si"},
			Usage:  "Print info about your servers - ip, machine id, access tokens, etc",
			Action: cmd.getServersInfo,
		},
		{
			Name:   "sections",
			Usage:  "Print info about your server's sections",
			Action: cmd.getSections,
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
	}

	app.Run(os.Args)
}

func getPlexToken(store *badger.KV, secret, tokenKey string) (string, error) {
	var tokenHash badger.KVItem

	if err := store.Get([]byte(tokenKey), &tokenHash); err != nil {
		if isVerbose {
			fmt.Println("could not retrieve plex token from datastore")
		}

		return "", err
	}

	token, err := decrypt([]byte(secret), string(tokenHash.Value()))

	if err != nil {
		if isVerbose {
			fmt.Println("token decryption failed")
		}
		return "", err
	}

	if isVerbose {
		fmt.Printf("Your plex token is %s\n", token)
	}

	return token, nil
}

func savePlexToken(store *badger.KV, secret, tokenKey, token string) error {
	tokenHash, err := encrypt([]byte(secret), token)

	if err != nil {
		return err
	}

	if isVerbose {
		fmt.Printf("your plex token hash: %s\n", string(tokenHash))
	}

	if err := store.Set([]byte(tokenKey), []byte(tokenHash), 0x00); err != nil {
		return err
	}

	if isVerbose {
		fmt.Println("saved token hash to store")
	}

	return nil
}

func initDataStore(dirName string) (*badger.KV, error) {
	options := badger.DefaultOptions
	dir, err := ioutil.TempDir("", dirName)

	if err != nil {
		return &badger.KV{}, err
	}

	options.Dir = dir
	options.ValueDir = dir

	return badger.NewKV(&options)
}

func initPlex(c *cli.Context) {
	var err error

	if plexConn, err = plex.New(c.GlobalString("url"), c.GlobalString("token")); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
