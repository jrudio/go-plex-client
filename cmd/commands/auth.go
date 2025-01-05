package commands

import (
	"context"
	"fmt"
	"github.com/jrudio/go-plex-client"
	"github.com/urfave/cli/v3"
	"sync"
	"time"
)

func Login(ctx context.Context, c *cli.Command) error {
	timeout := 15 * time.Minute
	// timeout := 10 * time.Second
	interval := 1 * time.Second
	ticker	:=	time.NewTicker(interval)
	done := make(chan bool)
	wg := sync.WaitGroup{}

	headers := plex.PlexHeaders{}

	resp, err := plex.RequestPIN(headers)

	if err != nil {
		errMsg := fmt.Sprintf("requesting pin failed: %v", err)

		return cli.Exit(errMsg, 1)
	}

	fmt.Printf("To authorize plexctl, please navigate to https://plex.tv/link and enter this code: %s\n", resp.Pin())
	// fmt.Printf("To authorize plexctl, please navigate to https://plex.tv/link and enter this code: %s\n", "abc123")

	wg.Add(1)

	// timeout  go routine
	go (func() {
		time.Sleep(timeout)
		done <- true
		wg.Done()
	})()


	wg.Add(1)

	go (func() {
		for {
			select {
			case <- done:
				fmt.Println("stopped checking pin")
				ticker.Stop()
				wg.Done()
				wg.Done() // override the timeout's waitgroup
				return
			case <- ticker.C:
				fmt.Println("checking pin status")
				resp, err := plex.CheckPIN(resp.ID, resp.ClientIdentifier)

				if err != nil && err.Error() != plex.ErrorPINNotAuthorized {
					fmt.Println("pin is expired or doesn't exist, request new pin")

					done <- true
				} else if err != nil && err.Error() == plex.ErrorPINNotAuthorized {
					fmt.Println("pin is not authorized yet")

				}

				if resp.AuthToken != "" && err == nil {
					fmt.Println("plexctl is now authorized")

					// TODO: encrypt and save plex token in local database

					// TODO: successfully exit from goroutines to finish program
					done <- true
				}
			}
		}

	})()

	wg.Wait()

	fmt.Println("done")

	return nil
}