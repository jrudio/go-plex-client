package plex

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// RefreshSection triggers a library scan for a specific library section ID.
func (p *Plex) RefreshSection(sectionID string) error {
	if sectionID == "" {
		return errors.New("sectionID is required")
	}

	query := fmt.Sprintf("%s/library/sections/%s/refresh", p.URL, sectionID)
	resp, err := p.get(query, p.Headers)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	return nil
}

// GetCollections lists all collections within a specific library section.
func (p *Plex) GetCollections(sectionID string) (MediaMetadata, error) {
	var results MediaMetadata
	if sectionID == "" {
		return results, errors.New("sectionID is required")
	}

	query := fmt.Sprintf("%s/library/sections/%s/collections", p.URL, sectionID)
	resp, err := p.get(query, p.Headers)
	if err != nil {
		return results, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return results, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return results, err
	}

	return results, nil
}

// CreateCollection creates a new manual collection in a specific library section.
func (p *Plex) CreateCollection(sectionID, title string) (MediaMetadata, error) {
	var results MediaMetadata
	if sectionID == "" || title == "" {
		return results, errors.New("sectionID and title are required")
	}

	// Endpoint: POST /library/collections?type=18&sectionId={sectionId}&title={collectionName}
	query := fmt.Sprintf("%s/library/collections?type=18&sectionId=%s&title=%s", p.URL, sectionID, url.QueryEscape(title))
	resp, err := p.post(query, nil, p.Headers)
	if err != nil {
		return results, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return results, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return results, err
	}

	return results, nil
}

// AddCollectionItem adds a media item (by ratingKey) to a manual collection.
func (p *Plex) AddCollectionItem(collectionID, ratingKey string) error {
	if collectionID == "" || ratingKey == "" {
		return errors.New("collectionID and ratingKey are required")
	}

	machineID, err := p.GetMachineIDLocal()
	if err != nil {
		return fmt.Errorf("failed to fetch machine identifier: %w", err)
	}

	// URI: server://{machineIdentifier}/com.plexapp.plugins.library/library/metadata/{ratingKey}
	uri := fmt.Sprintf("server://%s/com.plexapp.plugins.library/library/metadata/%s", machineID, ratingKey)

	// Endpoint: PUT /library/collections/{collectionId}/items?uri={uri}
	query := fmt.Sprintf("%s/library/collections/%s/items?uri=%s", p.URL, collectionID, url.QueryEscape(uri))
	resp, err := p.put(query, nil, p.Headers)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	return nil
}

// RemoveCollectionItem removes a media item (by ratingKey) from a manual collection.
func (p *Plex) RemoveCollectionItem(collectionID, ratingKey string) error {
	if collectionID == "" || ratingKey == "" {
		return errors.New("collectionID and ratingKey are required")
	}

	// Endpoint: DELETE /library/collections/{collectionID}/items/{ratingKey}
	query := fmt.Sprintf("%s/library/collections/%s/items/%s", p.URL, collectionID, ratingKey)
	resp, err := p.delete(query, p.Headers)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	return nil
}

// GetServerInfo retrieves base API information from the Plex Media Server.
func (p *Plex) GetServerInfo() (BaseAPIResponse, error) {
	var result BaseAPIResponse
	resp, err := p.get(p.URL, p.Headers)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}

// GetMachineIDLocal fetches the local Plex server's machine identifier directly from the server root.
func (p *Plex) GetMachineIDLocal() (string, error) {
	info, err := p.GetServerInfo()
	if err != nil {
		return "", err
	}
	if info.MediaContainer.MachineIdentifier == "" {
		return "", errors.New("machineIdentifier not found in response")
	}
	return info.MediaContainer.MachineIdentifier, nil
}

// PausePlayback sends the 'pause' command to a player client.
func (p *Plex) PausePlayback(machineID string) error {
	query := p.URL + "/player/playback/pause"
	newHeaders := p.Headers
	newHeaders.Accept = applicationXml
	newHeaders.TargetClientIdentifier = machineID

	resp, err := p.get(query, newHeaders)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

// ResumePlayback sends the 'play' command to a player client to resume playback.
func (p *Plex) ResumePlayback(machineID string) error {
	query := p.URL + "/player/playback/play"
	newHeaders := p.Headers
	newHeaders.Accept = applicationXml
	newHeaders.TargetClientIdentifier = machineID

	resp, err := p.get(query, newHeaders)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

// SkipNextPlayback sends the 'skipNext' command to skip to the next item.
func (p *Plex) SkipNextPlayback(machineID string) error {
	query := p.URL + "/player/playback/skipNext"
	newHeaders := p.Headers
	newHeaders.Accept = applicationXml
	newHeaders.TargetClientIdentifier = machineID

	resp, err := p.get(query, newHeaders)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

// SkipPreviousPlayback sends the 'skipPrevious' command to skip to the previous item.
func (p *Plex) SkipPreviousPlayback(machineID string) error {
	query := p.URL + "/player/playback/skipPrevious"
	newHeaders := p.Headers
	newHeaders.Accept = applicationXml
	newHeaders.TargetClientIdentifier = machineID

	resp, err := p.get(query, newHeaders)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

// StepForwardPlayback sends the 'stepForward' command to skip forward.
func (p *Plex) StepForwardPlayback(machineID string) error {
	query := p.URL + "/player/playback/stepForward"
	newHeaders := p.Headers
	newHeaders.Accept = applicationXml
	newHeaders.TargetClientIdentifier = machineID

	resp, err := p.get(query, newHeaders)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

// StepBackPlayback sends the 'stepBack' command to skip backward.
func (p *Plex) StepBackPlayback(machineID string) error {
	query := p.URL + "/player/playback/stepBack"
	newHeaders := p.Headers
	newHeaders.Accept = applicationXml
	newHeaders.TargetClientIdentifier = machineID

	resp, err := p.get(query, newHeaders)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

// SeekToPlayback sends the 'seekTo' command to jump to a specific offset in milliseconds.
func (p *Plex) SeekToPlayback(machineID string, offsetMs int) error {
	query := fmt.Sprintf("%s/player/playback/seekTo?offset=%d", p.URL, offsetMs)
	newHeaders := p.Headers
	newHeaders.Accept = applicationXml
	newHeaders.TargetClientIdentifier = machineID

	resp, err := p.get(query, newHeaders)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

// SetVolumePlayback sends the 'setParameters' command with a volume percentage (0-100).
func (p *Plex) SetVolumePlayback(machineID string, volume int) error {
	if volume < 0 {
		volume = 0
	} else if volume > 100 {
		volume = 100
	}
	query := fmt.Sprintf("%s/player/playback/setParameters?volume=%d", p.URL, volume)
	newHeaders := p.Headers
	newHeaders.Accept = applicationXml
	newHeaders.TargetClientIdentifier = machineID

	resp, err := p.get(query, newHeaders)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

// PlaylistMetadata represents Plex's playlists container JSON structure
type PlaylistMetadata struct {
	MediaContainer struct {
		Size     int        `json:"size"`
		Metadata []Metadata `json:"Metadata"`
	} `json:"MediaContainer"`
}

// GetPlaylists retrieves custom playlists from the Plex server, optionally filtered by playlistType
func (p *Plex) GetPlaylists(playlistType string) (PlaylistMetadata, error) {
	var results PlaylistMetadata
	plType := url.QueryEscape(playlistType)
	var query string
	if plType != "" {
		query = fmt.Sprintf("%s/playlists?playlistType=%s", p.URL, plType)
	} else {
		query = p.URL + "/playlists"
	}
	resp, err := p.get(query, p.Headers)
	if err != nil {
		return results, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return results, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return results, err
	}
	return results, nil
}

// CreatePlaylist creates a new custom playlist with an initial media item
func (p *Plex) CreatePlaylist(title, ratingKey string) (PlaylistMetadata, error) {
	var results PlaylistMetadata
	machineID, err := p.GetMachineIDLocal()
	if err != nil {
		return results, err
	}

	uri := fmt.Sprintf("server://%s/com.plexapp.plugins.library/library/metadata/%s", machineID, ratingKey)
	query := fmt.Sprintf("%s/playlists?type=video&title=%s&smart=0&uri=%s", p.URL, url.QueryEscape(title), url.QueryEscape(uri))

	resp, err := p.post(query, nil, p.Headers)
	if err != nil {
		return results, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return results, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return results, err
	}
	return results, nil
}

// AddPlaylistItems adds a media item to an existing playlist
func (p *Plex) AddPlaylistItems(playlistID, ratingKey string) error {
	machineID, err := p.GetMachineIDLocal()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("server://%s/com.plexapp.plugins.library/library/metadata/%s", machineID, ratingKey)
	query := fmt.Sprintf("%s/playlists/%s/items?uri=%s", p.URL, playlistID, url.QueryEscape(uri))

	resp, err := p.put(query, nil, p.Headers)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

// RemovePlaylistItems resolves a ratingKey within a playlist and deletes the corresponding item
func (p *Plex) RemovePlaylistItems(playlistID, ratingKey string) error {
	// 1. Fetch current items in the playlist to locate the playlistItemID (id) for the target ratingKey
	queryItems := fmt.Sprintf("%s/playlists/%s/items", p.URL, playlistID)
	respItems, err := p.get(queryItems, p.Headers)
	if err != nil {
		return err
	}
	defer respItems.Body.Close()

	if respItems.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch playlist items: %s", respItems.Status)
	}

	var playlistContent struct {
		MediaContainer struct {
			Metadata []Metadata `json:"Metadata"`
		} `json:"MediaContainer"`
	}
	if err := json.NewDecoder(respItems.Body).Decode(&playlistContent); err != nil {
		return err
	}

	var playlistItemID string
	for _, item := range playlistContent.MediaContainer.Metadata {
		if item.RatingKey == ratingKey {
			playlistItemID = item.PlaylistItemID
			break
		}
	}

	if playlistItemID == "" {
		return fmt.Errorf("media rating key '%s' not found inside playlist ID '%s'", ratingKey, playlistID)
	}

	// 2. Delete the item using its playlistItemID
	queryDelete := fmt.Sprintf("%s/playlists/%s/items/%s", p.URL, playlistID, playlistItemID)
	respDelete, err := p.delete(queryDelete, p.Headers)
	if err != nil {
		return err
	}
	defer respDelete.Body.Close()

	if respDelete.StatusCode != http.StatusOK && respDelete.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code on delete: %d", respDelete.StatusCode)
	}
	return nil
}

// OptimizeMedia starts a background media optimization job for a specific profile (e.g. mobile, tv)
func (p *Plex) OptimizeMedia(ratingKey, profile string) error {
	if profile == "" {
		profile = "mobile" // Default profile
	}
	// Endpoint: POST /library/metadata/{ratingKey}/optimize?profile={profile}&locationId=-1
	query := fmt.Sprintf("%s/library/metadata/%s/optimize?profile=%s&locationId=-1", p.URL, ratingKey, url.QueryEscape(profile))
	resp, err := p.post(query, nil, p.Headers)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}
	return nil
}

// BackgroundProcessingResponse holds details of running background transcoder/optimization tasks
type BackgroundProcessingResponse struct {
	MediaContainer struct {
		Size int `json:"size"`
		// Plex returns background processing info
	} `json:"MediaContainer"`
}

// GetBackgroundProcessing retrieves running background transcoding or optimization processes
func (p *Plex) GetBackgroundProcessing() (string, error) {
	query := p.URL + "/library/background"
	resp, err := p.get(query, p.Headers)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status: %s", resp.Status)
	}

	// Read raw body to pass to MCP as JSON text
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
