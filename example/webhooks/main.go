package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jrudio/go-plex-client/v2"
)

func main() {
	plexConnection, err := plex.New("http://192.168.1.2:32400", "myPlexToken")
	if err != nil {
		log.Fatalf("failed to connect to Plex: %v", err)
	}

	// 1. Webhook handler to easily handle events on your server HTTP
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

	go func() {
		fmt.Println("Listening for Webhooks on 192.168.1.14:8080")
		if err := http.ListenAndServe("192.168.1.14:8080", nil); err != nil {
			log.Fatalf("HTTP Server Error: %v", err)
		}
	}()

	// 2. Connect to your server via websockets to listen for events
	ctrlC := make(chan os.Signal, 1)
	onError := func(err error) {
		fmt.Println("Websocket Error:", err)
	}

	events := plex.NewNotificationEvents()
	events.OnPlaying(func(n plex.NotificationContainer) {
		if len(n.PlaySessionStateNotification) == 0 {
			return
		}

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

		for _, metadataObj := range sessions.MediaContainer.Metadata {
			if sessionID != metadataObj.SessionKey {
				continue
			}
			userID = metadataObj.User.ID
			username = metadataObj.User.Title
			break
		}

		metadata, err := plexConnection.GetMetadata(mediaID)
		if err != nil {
			fmt.Printf("failed to get metadata for key %s: %v\n", mediaID, err)
		} else {
			if len(metadata.MediaContainer.Metadata) > 0 {
				title = metadata.MediaContainer.Metadata[0].Title
			}
		}

		fmt.Printf("user (id: %s) has started playing %s (id: %s)\n", username, userID, title, mediaID)
	})

	fmt.Println("Listening to Websocket notifications...")
	plexConnection.SubscribeToNotifications(events, ctrlC, onError)
}
