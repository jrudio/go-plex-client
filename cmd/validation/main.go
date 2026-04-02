package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/jrudio/go-plex-client/v2"
)

func main() {
	token := os.Getenv("PLEX_TOKEN")
	plexURL := os.Getenv("PLEX_URL")

	if plexURL == "" {
		log.Fatal("PLEX_URL environment variable is required")
	}

	fmt.Printf("Connecting to Plex at %s...\n", plexURL)

	p, err := plex.New(plexURL, token)
	if err != nil {
		log.Fatalf("Failed to create plex client: %v", err)
	}

	if token == "" {
		fmt.Println("No PLEX_TOKEN provided. Falling back to PIN auth...")
		info, err := plex.RequestPIN(p.Headers, nil)
		if err != nil {
			log.Fatalf("Failed to request PIN: %v", err)
		}

		fmt.Printf("Please go to https://plex.tv/link and enter your PIN: %s (Expires at: %s)\n", info.Code, info.ExpiresAt)

		var authToken string
		for {
			pinInfo, err := plex.CheckPIN(info.ID, p.ClientIdentifier, nil)
			if err != nil {
				if err.Error() == plex.ErrorPINNotAuthorized {
					// Expected, still waiting
				} else {
					fmt.Printf("\rChecking PIN status... %v", err)
				}
			} else if pinInfo.AuthToken != "" {
				authToken = pinInfo.AuthToken
				break
			}
			time.Sleep(2 * time.Second)
		}

		fmt.Println("\nSuccessfully authenticated via PIN!")

		// Re-initialize the client with the newly acquired token
		p, err = plex.New(plexURL, authToken)
		if err != nil {
			log.Fatalf("Failed to re-initialize plex client: %v", err)
		}
		
		// For the rest of the script that uses token directly
		token = authToken
	}

	// Test connection
	success, err := p.Test()
	if err != nil {
		fmt.Printf("❌ Connection test failed: %v\n", err)
	} else if !success {
		fmt.Printf("❌ Connection test returned false\n")
	} else {
		fmt.Printf("✅ Connection test passed\n")
	}

	// Test OnDeck
	fmt.Println("\nTesting GetOnDeck()...")
	onDeck, err := p.GetOnDeck()
	var searchTerm string = "Avatar" // Default

	if err != nil {
		fmt.Printf("❌ GetOnDeck failed: %v\n", err)
	} else {
		fmt.Printf("✅ GetOnDeck found %d items\n", onDeck.MediaContainer.Size)
		if len(onDeck.MediaContainer.Metadata) > 0 {
			searchTerm = onDeck.MediaContainer.Metadata[0].Title
			fmt.Printf("   Picked search term from OnDeck: '%s'\n", searchTerm)
		}
	}

	// Test Search
	fmt.Printf("\nTesting Search('%s')...\n", searchTerm)

	// Debug: Raw request to /hubs/search (New)
	fmt.Println("--- Debug: /hubs/search (New) ---")
	rawQueryHubs := fmt.Sprintf("%s/hubs/search?query=%s", plexURL, url.QueryEscape(searchTerm))
	reqHubs, _ := http.NewRequest("GET", rawQueryHubs, nil)
	reqHubs.Header.Set("X-Plex-Token", token)
	reqHubs.Header.Set("Accept", "application/json")

	client := &http.Client{}
	respHubs, err := client.Do(reqHubs)
	if err != nil {
		fmt.Printf("❌ /hubs/search request failed: %v\n", err)
	} else {
		defer respHubs.Body.Close()
		body, _ := io.ReadAll(respHubs.Body)
		fmt.Printf("🔍 Status: %s\n", respHubs.Status)
		bodyStr := string(body)
		if len(bodyStr) > 500 {
			fmt.Printf("🔍 Body (truncated): %s...\n", bodyStr[:500])
		} else {
			fmt.Printf("🔍 Body: %s\n", bodyStr)
		}
	}

	// Debug: Raw request to /search (Legacy)
	fmt.Println("\n--- Debug: /search (Old) ---")
	rawQueryLegacy := fmt.Sprintf("%s/search?query=%s", plexURL, url.QueryEscape(searchTerm))
	reqLegacy, _ := http.NewRequest("GET", rawQueryLegacy, nil)
	reqLegacy.Header.Set("X-Plex-Token", token)
	reqLegacy.Header.Set("Accept", "application/json")

	respLegacy, err := client.Do(reqLegacy)
	if err != nil {
		fmt.Printf("❌ /search request failed: %v\n", err)
	} else {
		defer respLegacy.Body.Close()
		body, _ := io.ReadAll(respLegacy.Body)
		fmt.Printf("🔍 Status: %s\n", respLegacy.Status)
		bodyStr := string(body)
		if len(bodyStr) > 500 {
			fmt.Printf("🔍 Body (truncated): %s...\n", bodyStr[:500])
		} else {
			fmt.Printf("🔍 Body: %s\n", bodyStr)
		}
	}
	fmt.Println("--------------------------------")

	results, err := p.HubSearch(searchTerm)
	if err != nil {
		fmt.Printf("❌ Search failed: %v\n", err)
	} else {
		// New response structure: MediaContainer -> Hub -> Metadata
		fmt.Printf("✅ Search response size: %d\n", results.MediaContainer.Size)

		processedCount := 0
		for _, hub := range results.MediaContainer.Hub {
			// fmt.Printf("   Hub: %s (Type: %s, Size: %d)\n", hub.Title, hub.Type, hub.Size)
			if len(hub.Metadata) > 0 {
				fmt.Printf("   Hub: %s (Type: %s, Size: %d)\n", hub.Title, hub.Type, hub.Size)
				first := hub.Metadata[0]
				fmt.Printf("   First item: %s (Type: %s, Key: %s)\n", first.Title, first.Type, first.RatingKey)
				processedCount++

				// Test GetMetadata using the first search result of the first hub
				if processedCount == 1 {
					if first.RatingKey != "" {
						fmt.Printf("\nTesting GetMetadata(%s)...\n", first.RatingKey)

						// Debug: Raw request to GetMetadata
						fmt.Println("--- Debug: GetMetadata ---")
						// Plex requires the key path, not just the ID.
						// If ratingKey is "198", the path is /library/metadata/198
						getMetaPath := fmt.Sprintf("/library/metadata/%s", first.RatingKey)
						rawQueryMeta := fmt.Sprintf("%s%s", plexURL, getMetaPath)

						reqMeta, _ := http.NewRequest("GET", rawQueryMeta, nil)
						reqMeta.Header.Set("X-Plex-Token", token)
						reqMeta.Header.Set("Accept", "application/json")

						respMeta, err := client.Do(reqMeta)
						if err != nil {
							fmt.Printf("❌ GetMetadata request failed: %v\n", err)
						} else {
							defer respMeta.Body.Close()
							body, _ := io.ReadAll(respMeta.Body)
							fmt.Printf("🔍 Status: %s\n", respMeta.Status)
							bodyStr := string(body)
							// Print enough to see the 'rating' field
							if len(bodyStr) > 2000 {
								fmt.Printf("🔍 Body (truncated): %s...\n", bodyStr[:2000])
							} else {
								fmt.Printf("🔍 Body: %s\n", bodyStr)
							}
						}

						meta, err := p.GetMetadata(first.RatingKey)
						if err != nil {
							fmt.Printf("❌ GetMetadata failed: %v\n", err)
						} else {
							if len(meta.MediaContainer.Metadata) > 0 {
								m := meta.MediaContainer.Metadata[0]
								fmt.Printf("✅ GetMetadata success: %s\n", m.Title)
								fmt.Printf("   Rating (float): %v\n", m.Rating)
								if len(m.Ratings) > 0 {
									fmt.Printf("   Ratings (slice): %v\n", m.Ratings)
								} else {
									fmt.Printf("   Ratings (slice): empty\n")
								}
							} else {
								fmt.Printf("❌ GetMetadata returned no metadata\n")
							}
						}
					} else {
						fmt.Println("⚠️ Skipping GetMetadata test because RatingKey is empty (this highlights a potential issue!)")
					}
				}
			}
		}
		if processedCount == 0 {
			fmt.Println("⚠️ No metadata found in any Hub")
		}
	}

	// Test GetTranscodeSessions
	fmt.Println("\nTesting GetTranscodeSessions()...")
	sessions, err := p.GetTranscodeSessions()
	if err != nil {
		fmt.Printf("❌ GetTranscodeSessions failed: %v\n", err)
	} else {
		fmt.Printf("✅ GetTranscodeSessions success (size: %d)\n", len(sessions.Children))
	}

	// Test GetFriends
	fmt.Println("\nTesting GetFriends()...")
	friends, err := p.GetFriends()
	if err != nil {
		fmt.Printf("❌ GetFriends failed: %v\n", err)
	} else {
		fmt.Printf("✅ GetFriends found %d friends\n", len(friends))
	}
}
