# Plex.tv and Plex Media Server client written in Go

[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/jrudio/go-plex-client)

`go get -u github.com/jrudio/go-plex-client`

### Usage

```Go
Plex, err := plex.New("http://192.168.1.2:32400", "myPlexToken")

// Test your connection to your Plex server
result, pErr := Plex.Test()

// Search for media in your plex server
results, pErr := Plex.Search("The Walking Dead")

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


// ... and more! Please checkout plex.go for more methods
```
