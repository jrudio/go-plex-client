package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/jrudio/go-plex-client"
	"github.com/jrudio/go-plex-client/cmd/db"
	"github.com/urfave/cli/v3"
)

type CMDs struct {
	ClientIdentifier string
	dbClient         *db.DB
}

func New(dbConn *db.DB, clientIdentifier string) CMDs {
	return CMDs{
		ClientIdentifier: clientIdentifier,
		dbClient:         dbConn,
	}
}

func (cmd CMDs) Login(ctx context.Context, c *cli.Command) error {
	timeout := 15 * time.Minute
	interval := 1 * time.Second
	ticker := time.NewTicker(interval)
	done := make(chan bool, 1)

	headers := plex.PlexHeaders{ClientIdentifier: cmd.ClientIdentifier}

	resp, err := plex.RequestPIN(headers)

	if err != nil {
		errMsg := fmt.Sprintf("requesting pin failed: %v", err)

		return cli.Exit(errMsg, 1)
	}

	timeoutFn := func(done chan bool) {
		time.Sleep(timeout)
		done <- true
	}

	saveAuthorizationToDatabase := func(authToken string) error {
		client, err := plex.New("", authToken)

		if err != nil {
			return fmt.Errorf("creating plex client failed: %v", err)
		}

		account, err := client.MyAccount()

		if err != nil {
			return fmt.Errorf("fetching plex account information failed: %v", err)
		}

		if err := cmd.dbClient.SaveAuth(db.Authorization{
			Email:     account.Email,
			PlexToken: authToken,
		}); err != nil {
			return fmt.Errorf("saving authorization to database failed: %v", err)
		}

		return nil
	}

	authChecker := func(done chan bool, ticker *time.Ticker) {
		for {
			select {
			case <-done:
				fmt.Println("stopped checking pin")
				ticker.Stop()
				return
			case <-ticker.C:
				resp, err := plex.CheckPIN(resp.ID, resp.ClientIdentifier)

				if err != nil && err.Error() != plex.ErrorPINNotAuthorized {
					fmt.Println("pin is expired or doesn't exist, request new pin")

					done <- true

					continue
				}
				// else if err != nil && err.Error() == plex.ErrorPINNotAuthorized {
				// fmt.Println("pin is not authorized yet")

				// }

				if resp.AuthToken != "" && err == nil {

					fmt.Println("plexctl is now authorized, saving credentials...")

					if err := saveAuthorizationToDatabase(resp.AuthToken); err != nil {
						fmt.Printf("saving credentials failed: %v\n", err)
					} else {
						fmt.Println("successfully saved credentials")
					}

					done <- true
				}
			}
		}
	}

	fmt.Printf("To authorize plexctl, please navigate to https://plex.tv/link and enter this code: %s\n", resp.Pin())

	go timeoutFn(done)

	fmt.Println("checking pin status...")

	authChecker(done, ticker)

	fmt.Println("done")

	return nil
}

func (cmd CMDs) ListAccounts(ctx context.Context, c *cli.Command) error {
	auths, err := cmd.dbClient.GetAuthorizations()

	if err != nil {
		return cli.Exit(err, 1)
	}

	// print headers
	fmt.Printf("ACTIVE\tACCOUNT\n")

	// print accounts
	for i := 0; i < len(auths); i++ {
		auth := auths[i]

		isActive := ""

		if auth.IsActive {
			isActive = "*"
		}

		fmt.Printf("%s\t%s\n", isActive, auth.Email)
	}

	return nil
}

func (cmd CMDs) RevokeAccount(ctx context.Context, c *cli.Command) error {
	email := c.Args().Get(0)

	revokeClient := func(client *plex.Plex, account db.Authorization) error {
		devices, err := client.GetDevicesFromPlexTV()

		if err != nil {
			return fmt.Errorf("failed fetching plex devices: %v", err)
		}

		if len(devices) < 1 {
			return fmt.Errorf("no devices found")
		}

		isRevoked := false

		for _, device := range devices {
			if account.PlexToken == device.Token {
				if err := client.RevokeDevice(device.ID); err != nil {
					return err
				}

				isRevoked = true

				break
			}
		}

		if !isRevoked {
			return fmt.Errorf("did not find requested device to revoke")
		}

		return nil
	}

	removeAuthFromDB := func(email string) error {
		return cmd.dbClient.RemoveAuth(db.Authorization{
			Email: email,
		})
	}


	fmt.Printf("revoking '%s'...\n", email)

	authorizations, err := cmd.dbClient.GetAuthorizations()

	if err != nil {
		return cli.Exit(err, 1)
	}

	index := authorizations.FindIndexByEmail(email)
	account := authorizations[index]

	client, err := plex.New("", account.PlexToken)

	if err != nil {
		return cli.Exit(fmt.Errorf("failed creating plex client: %v", err), 1)
	}


	if err := revokeClient(client, account); err != nil {
		return cli.Exit(fmt.Errorf("revoking credentials failed: %v", err), 1)
	}

	fmt.Println("revoked credentials")

	if err := removeAuthFromDB(email); err != nil {
		return cli.Exit(fmt.Errorf("removing saved credentials failed: %v", err), 1)
	}

	return nil
}
