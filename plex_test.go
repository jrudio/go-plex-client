package plex

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	plexHost  string
	plexToken string
	plexConn  *Plex
)

func init() {
	plexHost = os.Getenv("PLEX_HOST")
	plexToken = os.Getenv("PLEX_TOKEN")

	if plexHost != "" {
		var err error
		if plexConn, err = New(plexHost, plexToken); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func newTestServer(code int, body string) (*httptest.Server, *Plex) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		w.Header().Set("Content-Type", applicationJson)
		fmt.Fprintln(w, body)
	}))

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	httpClient := http.Client{Transport: transport}
	plex := &Plex{URL: server.URL, Token: "", HTTPClient: httpClient}

	return server, plex
}

func TestSignIn(t *testing.T) {
	username := os.Getenv("PLEX_USERNAME")
	password := os.Getenv("PLEX_PASSWORD")

	if username == "" || password == "" {
		t.Skip("PLEX_USERNAME or PLEX_PASSWORD not set, skipping integration test")
	}

	plex, err := SignIn(username, password)

	if err != nil {
		t.Error(err.Error())
		return
	}

	if plex.Token == "" {
		t.Error("Received an empty token")
		return
	}
}

func TestGetSessions(t *testing.T) {
	testData := string(`
    {
      "MediaContainer": {
        "size": 2,
        "Metadata": [
          {
            "addedAt": 1461928551,
            "duration": 220497,
            "grandparentTitle": "Drake",
            "ratingKey": "910",
            "title": "Grammys (feat. Future)",
            "type": "track",
            "User": { "id": "1", "title": "jrudio" },
            "Player": { "state": "paused", "title": "note 5" }
          },
          {
            "addedAt": 1464015676,
            "duration": 2556288,
            "ratingKey": "1264",
            "title": "Shiva",
            "type": "episode",
            "User": { "id": "1", "title": "jrudio" },
            "Player": { "state": "paused", "title": "Plex Web (Chrome)" }
          }
        ]
      }
    }
  `)

	_, _plex := newTestServer(200, testData)

	_, err := _plex.GetSessions()

	if err != nil {
		t.Error(err.Error())
	}
}

func TestPlexTest(t *testing.T) {
	oldURL := plexURL
	defer func() { plexURL = oldURL }()

	_, _plex := newTestServer(200, "")
	plexURL = _plex.URL // override plexURL to local test server

	result, err := _plex.Test()

	if err != nil {
		t.Error(err.Error())
		return
	}

	if !result {
		t.Error(errors.New("the plex test returned false"))
		return
	}
}

func TestGetMetadata(t *testing.T) {
	query := "blahblah"

	testMetadataJSON := `
	{
		"MediaContainer": {
			"size": 1,
			"Metadata": [
				{
					"ratingKey": "123",
					"title": "Test Movie"
				}
			]
		}
	}`

	_, _plex := newTestServer(200, testMetadataJSON)

	_, err := _plex.GetMetadataChildren(query)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetServersInfo(t *testing.T) {
	if plexConn == nil {
		t.Skip("PLEX_HOST not set, skipping integration test")
	}

	info, err := plexConn.GetServersInfo()

	if err != nil {
		t.Error(err.Error())
		return
	}

	fmt.Println(info.Size)
}

func TestCheckUsernameOrEmailResponse(t *testing.T) {
	testData := []byte(`<?xml version="1.0" encoding="UTF-8"?>
		<Response code="0" status="Valid user"/>
	`)

	result := new(resultResponse)

	if err := xml.Unmarshal(testData, result); err != nil {
		t.Error(err.Error())
	}
}

func TestSectionIDResponse(t *testing.T) {
	testData := []byte(`
		{
			"MediaContainer": {
				"friendlyName": "myPlex",
				"size": 3,
				"Server": [
					{
						"name": "justin-server",
						"Section": [
							{"id": 2, "key": "2", "type": "movie", "title": "Movies"},
							{"id": 3, "key": "3", "type": "artist", "title": "Music"},
							{"id": 1, "key": "1", "type": "show", "title": "TV Shows"}
						]
					}
				]
			}
		}
	`)

	result := new(SectionIDResponse)

	if err := json.Unmarshal(testData, result); err != nil {
		t.Error(err.Error())
	}
}

func TestInviteFriendResponse(t *testing.T) {
	testData := []byte(`
		{
			"id": 1234,
			"name": "bob-server"
		}
	`)

	result := new(inviteFriendResponse)

	if err := json.Unmarshal(testData, result); err != nil {
		t.Error(err.Error())
	}
}

func TestPlex_GetInvitedFriends_Response(t *testing.T) {
	testData := []byte(`
	{
		"MediaContainer": {
			"friendlyName": "myPlex",
			"size": 3,
			"Invite": [
				{
					"id": "email1@gmail.com",
					"email": "email1@gmail.com",
					"friend": false,
					"home": false,
					"server": true,
					"Server": { "name": "Server123", "numLibraries": "3" }
				},
				{
					"id": "19661994",
					"username": "home-user",
					"email": "home-user@gmail.com",
					"friend": false,
					"home": true,
					"server": false
				},
				{
					"id": "22522496",
					"username": "existing-user",
					"email": "existing-user@umn.edu",
					"friend": true,
					"home": false,
					"server": true,
					"Server": { "name": "Server123", "numLibraries": "3" }
				}
			]
		}
	}
	`)

	result := new(invitedFriendsResponse)

	if err := json.Unmarshal(testData, result); err != nil {
		t.Error(err.Error())
	}
}

func TestPlex_RemoveInvitedFriend(t *testing.T) {
	if plexConn == nil {
		t.Skip("PLEX_HOST not set, skipping integration test")
	}
	success, err := plexConn.RemoveInvitedFriend("email-id-dne@gmail.com", false, true, false)
	if err.Error() != "404 Not Found" {
		// expect a 404
		t.Errorf("success: %v, error: %v", success, err)
	}
}

func TestGetLibraryLabels(t *testing.T) {
	testData := `{
		"MediaContainer": {
			"size": 1,
			"allowSync": true,
			"identifier": "com.plexapp.plugins.library",
			"mediaTagPrefix": "/system/bundle/media/providers/",
			"mediaTagVersion": 1421060938,
			"title1": "Library Name",
			"Directory": [
				{
					"key": "some_label",
					"title": "Some Label",
					"fastKey": "/library/sections/1/label/some_label"
				}
			]
		}
	}`

	server, _plex := newTestServer(200, testData)
	defer server.Close()

	labels, err := _plex.GetLibraryLabels("1", "1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	assert.Equal(t, "com.plexapp.plugins.library", labels.MediaContainer.Identifier)
	assert.True(t, labels.MediaContainer.AllowSync)
	assert.Equal(t, 1421060938, labels.MediaContainer.MediaTagVersion)
	if len(labels.MediaContainer.Directory) != 1 {
		t.Fatalf("expected 1 label, got %d", len(labels.MediaContainer.Directory))
	}
	assert.Equal(t, "some_label", labels.MediaContainer.Directory[0].Key)
	assert.Equal(t, "Some Label", labels.MediaContainer.Directory[0].Title)
}
