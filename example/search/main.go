package main

import (
	"fmt"
	"log"

	"github.com/jrudio/go-plex-client/v2"
)

func main() {
	// Initialize the Plex connection
	plexConnection, err := plex.New("http://192.168.1.2:32400", "myPlexToken")
	if err != nil {
		log.Fatalf("failed to connect to Plex: %v", err)
	}

	// Test your connection to your Plex server
	result, err := plexConnection.Test()
	if err != nil || !result {
		log.Fatalf("connection test failed: %v", err)
	}
	fmt.Println("Successfully connected to Plex!")

	// Search for media in your plex server using the modern Hubs endpoint
	results, err := plexConnection.HubSearch("The Walking Dead")
	if err != nil {
		log.Fatalf("search failed: %v", err)
	}

	for _, hub := range results.MediaContainer.Hub {
		fmt.Printf("Found %d items in %s hub\n", hub.Size, hub.Title)
		for _, metadata := range hub.Metadata {
			fmt.Printf("- %s (%d)\n", metadata.Title, metadata.Year)
		}
	}
}
