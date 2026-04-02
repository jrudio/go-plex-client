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

For comprehensive examples, please check the [`example/`](./example) directory. It contains code for connecting to a Plex server, utilizing the webhook receiver, interacting with search hubs, and setting up WebSocket notifications.

**Basic Initialization:**

```Go
import "github.com/jrudio/go-plex-client/v2"

plexConnection, err := plex.New("http://192.168.1.2:32400", "myPlexToken")
if err != nil {
	panic(err)
}

// Test your connection to your Plex server
result, err := plexConnection.Test()

// Search for media in your plex server using the modern Hubs endpoint
results, err := plexConnection.HubSearch("The Walking Dead")

// ... Please checkout plex.go or the example folder for more methods
```
