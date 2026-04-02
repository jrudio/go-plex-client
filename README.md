# Plex.tv and Plex Media Server client written in Go

[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/jrudio/go-plex-client)

`go get -u github.com/jrudio/go-plex-client/v2`

## Version 2

Version 2 brings the library into alignment with the Plex **OpenAPI specification (PMS v1.43+)**. 

Major changes include:
- **`HubSearch`**: Switched from legacy `/search` to the modern `/hubs/search` endpoint. Use `plexConnection.HubSearch("Title")` to retrieve results categorized by hubs (Movies, Shows, etc). Legacy search points are still accessible via `Search()`.
- **Accurate Data Models**: Models now properly account for the dynamic and polymophic nature of Plex API payloads (e.g. `rating` values changing from `float` on searches to array of `PlexRating` evaluation objects on metadata requests).

### CLI

You can tinker with this library using the command-line over [here](./cmd/plex-cli)

### Usage

```Go
import "github.com/jrudio/go-plex-client/v2"

plexConnection, err := plex.New("http://192.168.1.2:32400", "myPlexToken")

// Test your connection to your Plex server
result, err := plexConnection.Test()

// Search for media in your plex server using the modern Hubs endpoint
results, err := plexConnection.HubSearch("The Walking Dead")

// Webhook handler to easily handle events on your server
wh := plex.NewWebhook()

wh.OnPlay(func(w plex.Webhook) {
	fmt.Printf("%s is playing\n", w.Metadata.Title)
})

wh.OnPause(func(w plex.Webhook) {
	fmt.Printf("%s is paused\n", w.Metadata.Title)
})

wh.OnResume(func(w plex.Webhook) {
	fmt.Printf("%s has resumed\n", w.Metadata.Title)
})

wh.OnStop(func(w plex.Webhook) {
	fmt.Printf("%s has stopped\n", w.Metadata.Title)
})

http.HandleFunc("/", wh.Handler)

http.ListenAndServe("192.168.1.14:8080", nil)

// connect to your server via websockets to listen for events

ctrlC := make(chan os.Signal, 1)
onError := func(err error) {
	fmt.Println(err)
}

events := plex.NewNotificationEvents()
events.OnPlaying(func(n plex.NotificationContainer) {
	mediaID := n.PlaySessionStateNotification[0].RatingKey
	sessionID := n.PlaySessionStateNotification[0].SessionKey
	var title string
	var userID string
	var username string

	sessions, err := plexConnection.GetSessions()

	if err != nil {
		fmt.Printf("failed to fetch sessions on plex server: %v\n", err)
		return
	}

	for _, session := range sessions.MediaContainer.Video {
		if sessionID != session.SessionKey {
			continue
		}

		userID = session.User.ID
		username = session.User.Title

		break
	}

	metadata, err := plexConnection.GetMetadata(mediaID)

	if err != nil {
		fmt.Printf("failed to get metadata for key %s: %v\n", mediaID, err)
	} else {
		title = metadata.MediaContainer.Metadata[0].Title
	}

	fmt.Printf("user (id: %s) has started playing %s (id: %s)\n", username, userID, title, mediaID)
})

plexConnection.SubscribeToNotifications(events, ctrlC, onError)

// ... and more! Please checkout plex.go for more methods
```
