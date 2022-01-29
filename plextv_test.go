package plex

import (
	"testing"
)

func TestMyAccount(t *testing.T) {
	// plex token is required to test
	plexConn, err := SignIn(username, password)

	if err != nil {
		t.Errorf("valid username and password is required to test: %v", err)

		return
	}

	user, err := plexConn.MyAccount()

	if err != nil {
		t.Error(err)
		return
	}

	// username is from env var PLEX_USERNAME
	// if we got to this point somehow our previous checks failed
	if user.Username != username {
		t.Errorf("plex.tv username does not given username; have: '%s', need: '%s'", user.Username, username)
	}
}
