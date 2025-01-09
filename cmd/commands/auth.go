package commands

import (
	"context"
	"fmt"
	// "github.com/jrudio/go-plex-client"
	"github.com/urfave/cli/v3"
	"sync"
	"time"
)

// TODO: delete, this is only used for temporary testing implementation
func doneInSeconds(seconds int, done chan bool) {
	time.Sleep(time.Duration(seconds) * time.Second)

	done <- true
}

func Login(ctx context.Context, c *cli.Command) error {
	// timeout := 15 * time.Minute
	timeout := 5 * time.Second
	shortCircuitSeconds := 10
	interval := 1 * time.Second
	ticker	:=	time.NewTicker(interval)
	done := make(chan bool)
	wg := sync.WaitGroup{}

	// headers := plex.PlexHeaders{}

	// resp, err := plex.RequestPIN(headers)

	// if err != nil {
	// 	errMsg := fmt.Sprintf("requesting pin failed: %v", err)

	// 	return cli.Exit(errMsg, 1)
	// }

	// fmt.Printf("To authorize plexctl, please navigate to https://plex.tv/link and enter this code: %s\n", resp.Pin())
	fmt.Printf("To authorize plexctl, please navigate to https://plex.tv/link and enter this code: %s\n", "abc123")

	timeoutFn := func(done chan bool) {
		time.Sleep(timeout)
		done <- true
	}

	authChecker := func() {
		for {
			select {
			case <- done:
				fmt.Println("stopped checking pin")
				ticker.Stop()
				wg.Done()
				return
			case <- ticker.C:
				fmt.Println("checking pin status")
				// resp, err := plex.CheckPIN(resp.ID, resp.ClientIdentifier)

				// if err != nil && err.Error() != plex.ErrorPINNotAuthorized {
				// 	fmt.Println("pin is expired or doesn't exist, request new pin")

				// 	done <- true
				// } else if err != nil && err.Error() == plex.ErrorPINNotAuthorized {
				// 	fmt.Println("pin is not authorized yet")

				// }

				// if resp.AuthToken != "" && err == nil {
				// 	fmt.Println("plexctl is now authorized")

				// 	// TODO: encrypt and save plex token in local database

				// 	// TODO: successfully exit from goroutines to finish program
				// 	done <- true
				// }
			}
		}
	}

	// timeout goroutine
	go timeoutFn(done)

	// TEST: short-circuit the main goroutine and finish gracefuly; delete after successful test
	go doneInSeconds(shortCircuitSeconds, done)

	// main goroutine
	wg.Add(1)
	go authChecker()

	wg.Wait()

	fmt.Println("done")

	return nil
}