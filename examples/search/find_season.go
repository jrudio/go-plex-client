package main

import (
	"flag"
	"fmt"
	"strconv"

	plex "github.com/jrudio/go-plex-client"
)

const ERRUSAGE = "use -h or --help to see command usage"

func usage() {
	fmt.Println(ERRUSAGE)
}

func main() {
	// get required token and url to plex media server
	pmsURL := flag.String("url", "http://localhost:32400", "url to your plex media server")
	pmsToken := flag.String("token", "abc123", "auth token to your plex media server")

	flag.Parse()

	// url and token are required to connect to plex media server
	if *pmsURL == "" {
		fmt.Println("url is required")
		usage()
		return
	}

	if *pmsToken == "" {
		fmt.Println("token is required")
		usage()
		return
	}

	// get first argument
	if flag.NArg() < 2 {
		fmt.Println("title and season is required")
		usage()
		return
	}

	title := flag.Arg(0)
	targetSeason := flag.Arg(1)

	plexClient, err := plex.New(*pmsURL, *pmsToken)

	if err != nil {
		fmt.Println(err)
		return
	}

	searchResults, err := plexClient.Search(title)

	if err != nil {
		fmt.Printf("search failed: %v\n", err)
		return
	}

	results := searchResults.MediaContainer.Metadata

	if len(results) == 0 {
		fmt.Printf("could not find show '%s'\n", title)
		return
	}

	for _, result := range results {
		if result.Title == title && result.Type == "show" {
			// the children of a show are seasons, the ratingKey is the show's id
			tvShowMetadata, err := plexClient.GetMetadataChildren(result.RatingKey)

			if err != nil {
				fmt.Printf("failed to get seasons for '%s': %v\n", result.Title, err)
				continue
			}

			seasons := tvShowMetadata.MediaContainer.Metadata

			if seasonExists(targetSeason, seasons) {
				fmt.Printf("season %s of '%s' exists\n", targetSeason, result.Title)
			} else {
				fmt.Printf("could not find season %s of '%s'\n", targetSeason, result.Title)
			}

			// found the show, no need to continue loop
			break
		}
	}

}

func seasonExists(targetSeason string, seasons []plex.Metadata) bool {
	targetSeasonInt, err := strconv.Atoi(targetSeason)

	if err != nil || targetSeasonInt < 1 {
		return false
	}

	for _, season := range seasons {
		if season.Type != "season" {
			continue
		}

		if int(season.Index) == targetSeasonInt {
			return true
		}
	}

	return false
}
