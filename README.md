# Plex.tv and Plex Media Server client written in Go

[![Go Reference](https://pkg.go.dev/badge/github.com/jrudio/go-plex-client.svg)](https://pkg.go.dev/github.com/jrudio/go-plex-client)

An unofficial client library for your [Plex Media Server](https://www.plex.tv/)

### Command-line tool

You have a few options for getting the command-line tool:

- Download the latest [release](https://github.com/jrudio/go-plex-client/releases)

- `go install github.com/jrudio/go-plex-client/cmd`

- You can also clone and build from source [here](./cmd)



#### Usage

Go get the library and use it in your application like below:

1. `go get -u github.com/jrudio/go-plex-client`

2. Import and authenticate

```Go
import (
	"github.com/jrudio/go-plex-client"
)

client, err := plex.New("http://192.168.1.2:32400", "myPlexToken")

// or


```

```Go
// Test your connection to your Plex server
result, err := client.Test()

// Search for media in your plex server
results, err := client.Search("The Walking Dead")

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
events.OnPlaying(func(n NotificationContainer) {
	mediaID := n.PlaySessionStateNotification[0].RatingKey
	sessionID := n.PlaySessionStateNotification[0].SessionKey
	var title

	sessions, err := client.GetSessions()

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

	metadata, err := client.GetMetadata(mediaID)

	if err != nil {
		fmt.Printf("failed to get metadata for key %s: %v\n", mediaID, err)
	} else {
		title = metadata.MediaContainer.Metadata[0].Title
	}

	fmt.Printf("user (id: %s) has started playing %s (id: %s) %s\n", username, userID, title, mediaID)
})

client.SubscribeToNotifications(events, ctrlC, onError)
```
