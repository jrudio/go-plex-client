# Plex.tv and Plex Media Server client written in Go

[![Go Reference](https://pkg.go.dev/badge/github.com/jrudio/go-plex-client/v2.svg)](https://pkg.go.dev/github.com/jrudio/go-plex-client/v2)

`go get -u github.com/jrudio/go-plex-client/v2`

## Version 2

Version 2 brings the library into alignment with the Plex **OpenAPI specification (PMS v1.43+)**. 

Major changes include:
- **`HubSearch`**: Switched from legacy `/search` to the modern `/hubs/search` endpoint. Use `plexConnection.HubSearch("Title")` to retrieve results categorized by hubs (Movies, Shows, etc). Legacy search points are still accessible via `Search()`.
- **Accurate Data Models**: Models now properly account for the dynamic and polymophic nature of Plex API payloads (e.g. `rating` values changing from `float` on searches to array of `PlexRating` evaluation objects on metadata requests).
- **Playlists & Watched Status**: Added native support for retrieving lists of custom playlists (`GetPlaylists`) and toggling watched status (`Scrobble` and `Unscrobble`) via the REST API.
- **Season Downloads**: Enabled season downloading in `plex-cli` by traversing child elements.

## Key Features

* **OpenAPI Compliant**: Structured Go models aligned with modern Plex Media Server JSON responses.
* **Real-time Notifications**: Native support for subscribing to Plex WebSocket event streams (playing, progress, preference changes, and transcode status).
* **Media & Library Control**: Comprehensive endpoints for content discovery, metadata retrieval, media deletion, and playback controls.
* **Watched Status & Playlists**: Native methods for scrobbling, unscrobbling, and playlist curation.

### CLI

You can tinker with this library using the command-line over [here](./cmd/plex-cli)

### Get a Plex Token

To interact with the Plex API, you often need an authentication token. You can obtain one by using the `plex.tv/link` flow, which allows a user to authorize your application using a 4-character code.

```go
import (
	"fmt"
	"time"

	"github.com/jrudio/go-plex-client/v2"
)

func getToken() {
	// 1. Get a PIN code from plex.tv
	// Note: It is recommended to provide your own Headers with a unique ClientIdentifier
	pin, err := plex.RequestPIN(plex.DefaultHeaders(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Go to https://plex.tv/link and enter the code: %s\n", pin.Code)

	// 2. Poll plex.tv to see if the user has authorized the code
	for {
		// Use the ID from the pin request and your client identifier
		auth, err := plex.CheckPIN(pin.ID, "your-client-id", nil)
		if err != nil {
			// Check if it's just waiting for authorization
			if err.Error() == plex.ErrorPINNotAuthorized {
				time.Sleep(2 * time.Second)
				continue
			}
			panic(err)
		}

		if auth.AuthToken != "" {
			fmt.Printf("Success! Your Plex token is: %s\n", auth.AuthToken)
			break
		}

		time.Sleep(2 * time.Second)
	}
}
```

### Usage

For comprehensive examples, please check the [`example/`](./example) directory. It contains code for connecting to a Plex server, utilizing the webhook receiver, interacting with search hubs, and setting up WebSocket notifications.

**Basic Operations:**

```go
import "github.com/jrudio/go-plex-client/v2"

plexConnection, err := plex.New("http://192.168.1.2:32400", "myPlexToken")
if err != nil {
	panic(err)
}

// Test your connection to your Plex server
result, err := plexConnection.Test()

// Search for media in your plex server using the modern Hubs endpoint
results, err := plexConnection.HubSearch("The Walking Dead")

// Mark a media item as watched (scrobble)
err = plexConnection.Scrobble("12345")
```

**Real-time WebSocket Events:**

```go
// Initialize events callbacks
events := plex.NewNotificationEvents()
events.OnPlaying(func(n plex.NotificationContainer) {
	if len(n.PlaySessionStateNotification) > 0 {
		fmt.Printf("Session updated: %s\n", n.PlaySessionStateNotification[0].SessionKey)
	}
})

// Subscribe to event stream
interrupt := make(chan os.Signal, 1)
plexConnection.SubscribeToNotifications(events, interrupt, func(err error) {
	fmt.Printf("WebSocket error: %v\n", err)
})
```

---

## Projects Using `go-plex-client`

* **[Synapse](https://github.com/jrudio/synapse-plex-mcp)**: An interactive, conversational Discord agent for Plex Media Server, Radarr, Sonarr, and SABnzbd, utilizing `go-plex-client` under the hood.
