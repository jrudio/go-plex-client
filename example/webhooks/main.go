package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jrudio/go-plex-client/v2"
)

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	plexHost := os.Getenv("PLEX_HOST")
	plexToken := os.Getenv("PLEX_TOKEN")

	if host == "" {
		host = "localhost"
	}

	if port == "" {
		port = "8080"
	}

	if plexHost == "" {
		fmt.Println("PLEX_HOST not set")
		return
	}

	if plexToken == "" {
		headers := plex.DefaultHeaders()

		pin, err := plex.RequestPIN(headers, nil)
		if err != nil {
			log.Fatalf("failed to request PIN: %v", err)
		}

		fmt.Println("Please visit https://plex.tv/link to link your account")
		fmt.Println("Code: ", pin.Code)

		timeout := time.Now().Add(5 * time.Minute)

		for time.Now().Before(timeout) {
			if timeout.Before(time.Now()) {
				log.Fatalf("failed to check PIN: %v", err)
			}

			pinResponse, err := plex.CheckPIN(pin.ID, headers.ClientIdentifier, nil)

			if err != nil || pinResponse.AuthToken == "" {
				time.Sleep(3 * time.Second)
				continue
			}

			plexToken = pinResponse.AuthToken
			fmt.Println("Successfully linked to Plex!")

			// Save base64 encoded token to file
			if err := os.WriteFile("token", []byte(base64.StdEncoding.EncodeToString([]byte(plexToken))), 0644); err != nil {
				log.Fatalf("failed to save token: %v", err)
			}

			break
		}
	}

	plexConnection, err := plex.New(plexHost, plexToken)
	if err != nil {
		log.Fatalf("failed to connect to Plex: %v", err)
	}

	fmt.Println("setting up webhooks...")

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
		fmt.Printf("Listening for Webhooks on %s\n", plexHost)
		if err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil); err != nil {
			log.Fatalf("HTTP Server Error: %v", err)
		}
	}()

	// 2. Connect to your server via websockets to listen for events
	ctrlC := make(chan os.Signal, 1)
	signal.Notify(ctrlC, os.Interrupt, syscall.SIGTERM)
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

		fmt.Printf("user %s (id: %s) has started playing %s (id: %s)\n", username, userID, title, mediaID)
	})

	fmt.Println("Listening to Websocket notifications...")
	plexConnection.SubscribeToNotifications(events, ctrlC, onError)
}
