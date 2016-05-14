package plex

// package main

import (
	"encoding/json"
	"errors"
	"net/url"
	"runtime"
)

var requestInfo = request{
	headers: headers{
		Platform:        runtime.GOOS,
		PlatformVersion: "0.0.0",
		Product:         "Go Plex Client",
		Version:         "0.0.1",
		Device:          runtime.GOOS + " " + runtime.GOARCH,
		ContainerSize:   "Plex-Container-Size=50",
		ContainerStart:  "X-Plex-Container-Start=0",
		Accept:          "application/json",
	},
}

// New Establish the required info to make requests to your Plex Media Server
func New(baseURL, token string) (*Plex, error) {
	if baseURL == "" {
		return &Plex{}, errors.New("url is required")
	}

	_, err := url.ParseRequestURI(baseURL)

	return &Plex{
		URL:   baseURL,
		token: token,
	}, err
}

// Search Searches your Plex Server for media
func (p *Plex) Search(title string) (SearchResults, error) {
	if title == "" {
		return SearchResults{}, errors.New("ERROR: A title is required")
	}

	title = url.QueryEscape(title)
	query := p.URL + "/search?query=" + title

	var results SearchResults

	resp, respErr := requestInfo.get(query)

	if respErr != nil {
		// resp.Body.Close()
		return SearchResults{}, respErr
	}

	defer resp.Body.Close()

	err := json.NewDecoder(resp.Body).Decode(&results)

	if err != nil {
		return SearchResults{}, err
	}

	return results, nil
}

// Test Checks if you can connect to your Plex Media Server
func (p *Plex) Test() (bool, error) {
	var isAvailable bool

	resp, respErr := requestInfo.get(p.URL)

	if respErr != nil {
		// resp.Body.Close()
		return isAvailable, respErr
	}

	var results plexResponse

	err := json.NewDecoder(resp.Body).Decode(&results)

	if err != nil {
		return isAvailable, err
	}

	err = resp.Body.Close()

	if err != nil {
		return isAvailable, err
	}

	if results.Version != "" {
		isAvailable = true
	}

	return isAvailable, nil
}

// KillTranscodeSession stops a transcode session with a session key
func (p *Plex) KillTranscodeSession(sessionKey string) (bool, error) {
	var result bool

	if sessionKey == "" {
		return result, errors.New("Missing sessionKey")
	}

	query := p.URL + "/video/:/transcode/universal/stop?session=" + sessionKey

	resp, respErr := requestInfo.get(query)

	if respErr != nil {
		return result, respErr
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		result = true
	}

	return result, nil
}

// GetTranscodeSessions retrieves a list of all active transcode sessions
func (p *Plex) GetTranscodeSessions() (transcodeSessionsResponse, error) {
	var result transcodeSessionsResponse

	query := p.URL + "/transcode/sessions"

	resp, respErr := requestInfo.get(query)

	if respErr != nil {
		return result, respErr
	}

	defer resp.Body.Close()

	return result, json.NewDecoder(resp.Body).Decode(&result)

}
