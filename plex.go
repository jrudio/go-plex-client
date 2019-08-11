package plex

// plex is a Plex Media Server and Plex.tv client

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const plexURL = "https://plex.tv"

func defaultHeaders() headers {
	version := "0.0.1"

	return headers{
		Platform:         runtime.GOOS,
		PlatformVersion:  "0.0.0",
		Product:          "Go Plex Client",
		Version:          version,
		Device:           runtime.GOOS + " " + runtime.GOARCH,
		ClientIdentifier: "go-plex-client-v" + version,
		ContainerSize:    "Plex-Container-Size=50",
		ContainerStart:   "X-Plex-Container-Start=0",
		Accept:           "application/json",
		ContentType:      "application/json",
	}
}

// New creates a new plex instance that is required to
// to make requests to your Plex Media Server
func New(baseURL, token string) (*Plex, error) {
	var p Plex

	// allow empty url so caller can use GetServers() to set the server url later

	if baseURL == "" && token == "" {
		return &p, errors.New("url or a token is required")
	}

	p.HTTPClient = http.Client{
		Timeout: 3 * time.Second,
	}

	// id, err := uuid.NewRandom()

	// if err != nil {
	// 	return &p, err
	// }

	// p.ClientIdentifier = id.String()
	p.ClientIdentifier = p.Headers.ClientIdentifier
	p.Headers.ClientIdentifier = p.ClientIdentifier

	// has url and token
	if baseURL != "" && token != "" {
		_, err := url.ParseRequestURI(baseURL)

		p.URL = baseURL
		p.Token = token

		return &p, err
	}

	// just has token
	if baseURL == "" && token != "" {
		p.Token = token

		return &p, nil
	}

	// just url
	p.URL = baseURL

	return &p, nil
}

// SignIn creates a plex instance using a user name and password instead of an auth
// token.
func SignIn(username, password string) (*Plex, error) {
	id, err := uuid.NewRandom()

	if err != nil {
		return &Plex{}, err
	}

	p := Plex{
		ClientIdentifier: id.String(),
		HTTPClient: http.Client{
			Timeout: 3 * time.Second,
		},
	}

	query := plexURL + "/users/sign_in.json"

	// Encode login in the specific format they require
	body := url.Values{}
	body.Add("user[login]", username)
	body.Add("user[password]", password)

	newHeaders := p.Headers
	// Doesn't like having a content type, even form-data
	newHeaders.ContentType = ""
	resp, err := p.post(query, []byte(body.Encode()), newHeaders)

	if err != nil {
		return &Plex{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return &Plex{}, errors.New(resp.Status)
	}

	var signInResponse SignInResponse

	if err := json.NewDecoder(resp.Body).Decode(&signInResponse); err != nil {
		return &Plex{}, err
	}

	p.Token = signInResponse.User.AuthToken

	return &p, err
}

// Search your Plex Server for media
func (p *Plex) Search(title string) (SearchResults, error) {
	if title == "" {
		return SearchResults{}, errors.New("ERROR: A title is required")
	}

	title = url.QueryEscape(title)
	query := p.URL + "/search?query=" + title

	var results SearchResults

	resp, err := p.get(query, p.Headers)

	if err != nil {
		return SearchResults{}, err
	}

	// Unauthorized
	if resp.StatusCode == 401 {
		return SearchResults{}, errors.New("You are not authorized to access that server")
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return SearchResults{}, err
	}

	return results, nil
}

// GetMetadata can get some media info
func (p *Plex) GetMetadata(key string) (MediaMetadata, error) {
	if key == "" {
		return MediaMetadata{}, errors.New("ERROR: A key is required")
	}

	var results MediaMetadata

	query := fmt.Sprintf("%s/library/metadata/%s", p.URL, key)

	newHeaders := p.Headers

	resp, err := p.get(query, newHeaders)

	if err != nil {
		return results, err
	}

	if resp.StatusCode != http.StatusOK {
		return results, errors.New("server error: " + resp.Status)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return results, err
	}

	return results, nil
}

// GetMetadataChildren can get a show's season titles. My use-case would be getting the season titles after using Search()
func (p *Plex) GetMetadataChildren(key string) (MetadataChildren, error) {
	if key == "" {
		return MetadataChildren{}, errors.New("ERROR: A key is required")
	}

	query := fmt.Sprintf("%s/library/metadata/%s/children", p.URL, key)

	newHeaders := p.Headers

	resp, err := p.get(query, newHeaders)

	if err != nil {
		return MetadataChildren{}, err
	}

	// Unauthorized
	if resp.StatusCode == 401 {
		return MetadataChildren{}, errors.New("You are not authorized to access that server")
	}

	defer resp.Body.Close()

	var results MetadataChildren

	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return MetadataChildren{}, err
	}

	return results, nil
}

// GetEpisodes returns episodes of a season of a show
func (p *Plex) GetEpisodes(key string) (SearchResultsEpisode, error) {
	if key == "" {
		return SearchResultsEpisode{}, errors.New("Key is required")
	}

	query := fmt.Sprintf("%s/library/metadata/%s/children", p.URL, key)

	resp, err := p.get(query, p.Headers)

	if err != nil {
		return SearchResultsEpisode{}, err
	}

	// Unauthorized
	if resp.StatusCode == 401 {
		return SearchResultsEpisode{}, errors.New("You are not authorized to access that server")
	}

	defer resp.Body.Close()

	var results SearchResultsEpisode

	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return SearchResultsEpisode{}, err
	}

	return results, nil
}

// GetEpisode returns a single episode of a show.
func (p *Plex) GetEpisode(key string) (SearchResultsEpisode, error) {
	if key == "" {
		return SearchResultsEpisode{}, errors.New("Key is required")
	}

	query := fmt.Sprintf("%s/library/metadata/%s", p.URL, key)

	resp, err := p.get(query, p.Headers)

	if err != nil {
		return SearchResultsEpisode{}, err
	}

	// Unauthorized
	if resp.StatusCode == 401 {
		return SearchResultsEpisode{}, errors.New("You are not authorized to access that server")
	}

	defer resp.Body.Close()

	var results SearchResultsEpisode

	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return SearchResultsEpisode{}, err
	}

	return results, nil
}

// GetOnDeck gets the on-deck videos.
func (p *Plex) GetOnDeck() (SearchResultsEpisode, error) {
	query := fmt.Sprintf("%s/library/onDeck", p.URL)

	resp, err := p.get(query, p.Headers)

	if err != nil {
		return SearchResultsEpisode{}, err
	}

	// Unauthorized
	if resp.StatusCode == 401 {
		return SearchResultsEpisode{}, errors.New("You are not authorized to access that server")
	}

	defer resp.Body.Close()

	var results SearchResultsEpisode

	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return SearchResultsEpisode{}, err
	}

	return results, nil
}

// GetPlaylist gets all videos in a playlist.
func (p *Plex) GetPlaylist(key int) (SearchResultsEpisode, error) {
	query := fmt.Sprintf("%s/playlists/%d/items", p.URL, key)

	resp, err := p.get(query, p.Headers)

	if err != nil {
		return SearchResultsEpisode{}, err
	}

	// Unauthorized
	if resp.StatusCode == 401 {
		return SearchResultsEpisode{}, errors.New("You are not authorized to access that server")
	}

	defer resp.Body.Close()

	var results SearchResultsEpisode

	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return SearchResultsEpisode{}, err
	}

	return results, nil
}

// GetThumbnail returns the response of a request to pms thumbnail
// My ideal use case would be to proxy a request to pms without exposing the plex token
func (p *Plex) GetThumbnail(key, thumbnailID string) (*http.Response, error) {
	query := fmt.Sprintf("%s/library/metadata/%s/thumb/%s", p.URL, key, thumbnailID)

	return p.get(query, p.Headers)
}

// Test your connection to your Plex Media Server
func (p *Plex) Test() (bool, error) {
	resp, err := p.get(plexURL+"/api/servers", p.Headers)

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		return false, errors.New("You are not authorized to access this server")
	} else if resp.StatusCode != 200 {
		statusCode := strconv.Itoa(resp.StatusCode)
		return false, errors.New("Server replied with " + statusCode + " status code")
	}

	return true, nil
}

// KillTranscodeSession stops a transcode session
func (p *Plex) KillTranscodeSession(sessionKey string) (bool, error) {

	if sessionKey == "" {
		return false, errors.New("Missing sessionKey")
	}

	query := p.URL + "/video/:/transcode/universal/stop?session=" + sessionKey

	resp, err := p.get(query, p.Headers)

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		return false, errors.New("You are not authorized to access this server")
	} else if resp.StatusCode != 200 {
		statusCode := strconv.Itoa(resp.StatusCode)
		return false, errors.New("Server replied with " + statusCode + " status code")
	}

	return true, nil
}

// GetTranscodeSessions retrieves a list of all active transcode sessions
func (p *Plex) GetTranscodeSessions() (TranscodeSessionsResponse, error) {
	var result TranscodeSessionsResponse

	query := p.URL + "/transcode/sessions"

	resp, err := p.get(query, p.Headers)

	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		return result, errors.New("You are not authorized to access this server")
	} else if resp.StatusCode != 200 {
		statusCode := strconv.Itoa(resp.StatusCode)
		return result, errors.New("Server replied with " + statusCode + " status code")
	}

	return result, json.NewDecoder(resp.Body).Decode(&result)

}

// GetPlexTokens not sure if it works
func (p *Plex) GetPlexTokens(token string) (DevicesResponse, error) {
	var result DevicesResponse

	query := plexURL + "/devices.json"

	resp, err := p.get(query, p.Headers)

	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		return result, errors.New("You are not authorized to access this server")
	} else if resp.StatusCode != 200 {
		statusCode := strconv.Itoa(resp.StatusCode)
		return result, errors.New("Server replied with " + statusCode + " status code")
	}

	return result, json.NewDecoder(resp.Body).Decode(&result)
}

// DeletePlexToken is currently not tested
func (p *Plex) DeletePlexToken(token string) (bool, error) {
	var result bool

	query := plexURL + "/devices/" + token + ".json"

	resp, err := p.get(query, p.Headers)

	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		return result, errors.New("You are not authorized to access this server")
	} else if resp.StatusCode != 200 {
		statusCode := strconv.Itoa(resp.StatusCode)
		return result, errors.New("Server replied with " + statusCode + " status code")
	}

	return result, json.NewDecoder(resp.Body).Decode(&result)
}

// GetFriends returns all of your plex friends
func (p *Plex) GetFriends() ([]Friends, error) {

	var plexFriendsResp friendsResponse

	query := plexURL + "/api/users"

	newHeaders := p.Headers

	newHeaders.Accept = "application/xml"

	resp, err := p.get(query, newHeaders)

	if err != nil {
		return []Friends{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		return []Friends{}, errors.New("You are not authorized to access this server")
	} else if resp.StatusCode != 200 {
		statusCode := strconv.Itoa(resp.StatusCode)
		return []Friends{}, errors.New("Server replied with " + statusCode + " status code")
	}

	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return []Friends{}, err
	}

	err = xml.Unmarshal(respBytes, &plexFriendsResp)

	if err != nil {
		return []Friends{}, err
	}

	friendCount := plexFriendsResp.Size

	plexFriends := make([]Friends, friendCount)

	for ii, f := range plexFriendsResp.User {
		plexFriends[ii] = f
	}

	return plexFriends, nil
}

// RemoveFriend from your friend's list which stops access to your Plex server
func (p *Plex) RemoveFriend(id string) (bool, error) {

	query := plexURL + "/api/friends/" + id

	resp, err := p.delete(query, p.Headers)

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 400 {
		return false, errors.New(resp.Status)
	}

	result := new(resultResponse)

	if err := xml.NewDecoder(resp.Body).Decode(result); err != nil {
		return false, err
	}

	return result.Response.Code == 0, nil
}

// InviteFriend to access your Plex server. Add restrictions to media or give them full access.
func (p *Plex) InviteFriend(params InviteFriendParams) (int, error) {
	usernameOrEmail := url.QueryEscape(params.UsernameOrEmail)

	label := url.QueryEscape(params.Label)

	query := fmt.Sprintf("%s/api/servers/%s/shared_servers", plexURL, params.MachineID)

	var restrictions inviteFriendBody

	restrictions.ServerID = params.MachineID
	restrictions.SharedServer = inviteFriendSharedServer{
		InvitedEmail:      params.UsernameOrEmail,
		LibrarySectionIDs: params.LibraryIDs,
	}

	settings := inviteFriendSharingSettings{}

	if label != "" {
		settings.FilterMovies = fmt.Sprintf("label=%s", label)
		settings.FilterTelevision = fmt.Sprintf("label=%s", label)
	}

	restrictions.SharingSettings = settings

	jsonBody, jsonErr := json.Marshal(restrictions)

	if jsonErr != nil {
		return 0, jsonErr
	}

	resp, err := p.post(query, jsonBody, p.Headers)

	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	result := new(inviteFriendResponse)

	if err := xml.NewDecoder(resp.Body).Decode(result); err != nil {
		return 0, err
	}

	sharedServer := result.SharedServer

	if resp.StatusCode != 200 {
		return 0, errors.New("non-200 response code")
	}

	if sharedServer.Username != usernameOrEmail && sharedServer.Email != usernameOrEmail {
		return 0, errors.New("username or email does not match expected output")
	}

	return sharedServer.UserID, nil
}

// UpdateFriendAccess limit your friends access to your plex server
func (p *Plex) UpdateFriendAccess(userID string, params UpdateFriendParams) (bool, error) {
	// Fix any defaults to statisfy what plex expects
	if params.AllowSync == "" {
		params.AllowSync = "0"
	}

	if params.AllowCameraUpload == "" {
		params.AllowCameraUpload = "0"
	}

	if params.AllowChannels == "" {
		params.AllowChannels = "0"
	}

	query := fmt.Sprintf("%s/api/friends/%s", plexURL, userID)

	parsedQuery, parseErr := url.Parse(query)

	if parseErr != nil {
		return false, parseErr
	}

	vals := parsedQuery.Query()

	vals.Add("allowSync", params.AllowSync)
	vals.Add("allowCameraUpload", params.AllowCameraUpload)
	vals.Add("allowChannels", params.AllowChannels)
	vals.Add("filterMovies", params.FilterMovies)
	vals.Add("filterMusic", params.FilterMusic)
	vals.Add("filterTelevision", params.FilterTelevision)
	vals.Add("filterPhotos", params.FilterPhotos)

	parsedQuery.RawQuery = vals.Encode()

	query = parsedQuery.String()

	resp, err := p.put(query, nil, p.Headers)

	if err != nil {
		return false, err
	}

	resp.Body.Close()

	if resp.StatusCode != 200 {
		return false, errors.New(resp.Status)
	}

	return true, nil
}

// RemoveFriendAccessToLibrary you can individually revoke access to a library on your server. Such as movies, tv shows, music, etc
func (p *Plex) RemoveFriendAccessToLibrary(userID, machineID, serverID string) (bool, error) {
	query := fmt.Sprintf("%s/api/servers/%s/shared_servers/%s", plexURL, machineID, serverID)

	resp, err := p.delete(query, p.Headers)

	if err != nil {
		return false, err
	}

	resp.Body.Close()

	if resp.StatusCode != 200 {
		return false, errors.New(resp.Status)
	}

	return true, nil
}

// CheckUsernameOrEmail will check if the username is a Plex user or will verify an email is valid
func (p *Plex) CheckUsernameOrEmail(usernameOrEmail string) (bool, error) {

	usernameOrEmail = url.QueryEscape(usernameOrEmail)

	query := fmt.Sprintf("%s/api/users/validate?invited_email=%s", plexURL, usernameOrEmail)

	resp, err := p.post(query, nil, p.Headers)

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 400 {
		return false, errors.New(resp.Status)
	}

	result := new(resultResponse)

	if err := xml.NewDecoder(resp.Body).Decode(result); err != nil {
		return false, err
	}

	return result.Response.Code == 0, nil
}

// StopPlayback acts as a remote controller and sends the 'stop' command
func (p *Plex) StopPlayback(machineID string) error {
	query := p.URL + "/player/playback/stop"

	newHeaders := p.Headers

	newHeaders.Accept = "application/xml"
	newHeaders.TargetClientIdentifier = machineID

	resp, err := p.get(query, newHeaders)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(resp.Status)
	}

	return nil
}

// GetServers returns a list of your Plex servers
func (p *Plex) GetServers() ([]PMSDevices, error) {

	query := plexURL + "/pms/resources.xml?includeHttps=1"

	resp, err := p.get(query, p.Headers)

	if err != nil {
		return []PMSDevices{}, err
	}

	defer resp.Body.Close()

	result := new(resourcesResponse)

	if err := xml.NewDecoder(resp.Body).Decode(result); err != nil {
		fmt.Println(err.Error())

		return []PMSDevices{}, err
	}

	var servers []PMSDevices

	for _, r := range result.Device {
		if r.Provides != "server" {
			continue
		}

		servers = append(servers, r)
	}

	return servers, nil
}

// GetServersInfo returns info about your Plex servers
func (p *Plex) GetServersInfo() (ServerInfo, error) {
	query := plexURL + "/api/servers"

	resp, err := p.get(query, p.Headers)

	if err != nil {
		return ServerInfo{}, err
	}

	defer resp.Body.Close()

	result := ServerInfo{}

	if err := xml.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println(err.Error())

		return ServerInfo{}, err
	}

	return result, nil
}

// GetMachineID returns the machine id of the currently connected server
func (p *Plex) GetMachineID() (string, error) {
	resp, err := p.get(p.URL, p.Headers)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var result BaseAPIResponse

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println(err.Error())

		return "", err
	}

	return result.MediaContainer.MachineIdentifier, nil
}

// GetSections of your plex server. This is useful when inviting a user
// as you can restrict the invited user to a library (i.e. Movie's, TV Shows)
func (p *Plex) GetSections(machineID string) ([]ServerSections, error) {
	query := fmt.Sprintf("%s/api/servers/%s", plexURL, machineID)

	newHeaders := p.Headers

	newHeaders.Accept = "application/xml"

	resp, err := p.get(query, newHeaders)

	if err != nil {
		return []ServerSections{}, err
	}

	defer resp.Body.Close()

	var result SectionIDResponse

	if err := xml.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println(err.Error())

		return []ServerSections{}, err
	}

	// Look for our server via the machine id
	for _, server := range result.Server {
		if server.MachineIdentifier != machineID {
			continue
		}

		return server.Section, nil
	}

	return []ServerSections{}, nil
}

// GetLibraries of your Plex server. My ideal use-case would be
// to get library count to determine label index
func (p *Plex) GetLibraries() (LibrarySections, error) {

	query := fmt.Sprintf("%s/library/sections", p.URL)

	resp, err := p.get(query, p.Headers)

	if err != nil {
		return LibrarySections{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LibrarySections{}, errors.New(resp.Status)
	}

	var result LibrarySections

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println(err.Error())

		return LibrarySections{}, err
	}

	return result, nil
}

// GetLibraryContent retrieve the content inside a library
func (p *Plex) GetLibraryContent(sectionKey string, filter string) (SearchResults, error) {
	query := fmt.Sprintf("%s/library/sections/%s/all%s", p.URL, sectionKey, filter)

	resp, err := p.get(query, p.Headers)

	if err != nil {
		return SearchResults{}, err
	}

	if resp.Status == ErrorInvalidToken {
		return SearchResults{}, errors.New("invalid token")
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return SearchResults{}, errors.New("You are not authorized to access that server")
	}

	if resp.StatusCode == http.StatusBadRequest {
		return SearchResults{}, errors.New("There was an error in the request")
	}

	defer resp.Body.Close()

	var results SearchResults

	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return SearchResults{}, err
	}

	return results, nil
}

// CreateLibrary will create a new library on your Plex server
func (p *Plex) CreateLibrary(params CreateLibraryParams) error {
	// all params are required
	if params.Name == "" {
		return errors.New("name is required")
	}

	if params.Location == "" {
		return errors.New("location is required")
	}

	if params.LibraryType == "" {
		return errors.New("libraryType is required")
	}

	if params.Agent == "" {
		return errors.New("agent is required")
	}

	if params.Scanner == "" {
		return errors.New("scanner is required")
	}

	if params.Language == "" {
		params.Language = "en"
	}

	query := p.URL + "/library/sections"

	parsedQuery, err := url.Parse(query)

	if err != nil {
		return err
	}

	queryValues := parsedQuery.Query()

	queryValues.Add("name", params.Name)
	queryValues.Add("location", params.Location)
	queryValues.Add("language", params.Language)
	queryValues.Add("type", params.LibraryType)
	queryValues.Add("agent", params.Agent)
	queryValues.Add("scanner", params.Scanner)

	parsedQuery.RawQuery = queryValues.Encode()

	query = parsedQuery.String()

	resp, err := p.post(query, nil, p.Headers)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return errors.New(resp.Status)
	}

	return nil
}

// DeleteLibrary removes the library from your Plex server via library key (or id)
func (p *Plex) DeleteLibrary(key string) error {
	query := fmt.Sprintf("%s/library/sections/%s", p.URL, key)

	resp, err := p.delete(query, p.Headers)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	return nil
}

// GetLibraryLabels of your plex server
func (p *Plex) GetLibraryLabels(sectionKey, sectionIndex string) (LibraryLabels, error) {

	if sectionIndex == "" {
		sectionIndex = "1"
	}

	query := fmt.Sprintf("%s/library/sections/%s/labels?type=%s", p.URL, sectionKey, sectionIndex)

	resp, err := p.get(query, p.Headers)

	if err != nil {
		return LibraryLabels{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LibraryLabels{}, errors.New(resp.Status)
	}

	var result LibraryLabels

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println(err.Error())

		return LibraryLabels{}, err
	}

	return result, nil
}

// AddLabelToMedia restrict access to certain media. Requires a Plex Pass.
// mediaType is the media type (1), id is the ratingKey or media id, label is your label, locked is unknown
// 1. A reference to the plex media types: https://github.com/Arcanemagus/plex-api/wiki/MediaTypes
// XXX: Currently plex is capitalizing the first letter
func (p *Plex) AddLabelToMedia(mediaType, sectionID, id, label, locked string) (bool, error) {

	query := fmt.Sprintf("%s/library/sections/%s/all", p.URL, sectionID)

	parsedQuery, parseErr := url.Parse(query)

	if parseErr != nil {
		return false, parseErr
	}

	vals := parsedQuery.Query()

	vals.Add("type", mediaType)
	vals.Add("id", id)
	vals.Add("label[0].tag.tag", label)
	// vals.Add("label.locked", locked)

	parsedQuery.RawQuery = vals.Encode()

	query = parsedQuery.String()

	resp, err := p.put(query, nil, p.Headers)

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	return resp.StatusCode == 200, nil
}

// RemoveLabelFromMedia to remove a label from a piece of media Requires a Plex Pass.
func (p *Plex) RemoveLabelFromMedia(mediaType, sectionID, id, label, locked string) (bool, error) {

	query := fmt.Sprintf("%s/library/sections/%s/all", p.URL, sectionID)

	parsedQuery, parseErr := url.Parse(query)

	if parseErr != nil {
		return false, parseErr
	}

	vals := parsedQuery.Query()

	vals.Add("type", mediaType)
	vals.Add("id", id)
	vals.Add("label[].tag.tag-", label)
	vals.Add("label.locked", locked)

	parsedQuery.RawQuery = vals.Encode()

	query = parsedQuery.String()

	resp, err := p.put(query, nil, p.Headers)

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	return resp.StatusCode == 200, nil
}

// GetSessions of devices currently consuming media
func (p *Plex) GetSessions() (CurrentSessions, error) {
	newHeaders := p.Headers

	query := fmt.Sprintf("%s/status/sessions", p.URL)

	resp, err := p.get(query, newHeaders)

	if err != nil {
		return CurrentSessions{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return CurrentSessions{}, errors.New(resp.Status)
	}

	var result CurrentSessions

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return CurrentSessions{}, err
	}

	return result, nil
}

// TerminateSession will end a streaming session - plex pass feature
func (p *Plex) TerminateSession(sessionID string, reason string) error {
	if reason == "" {
		reason = "The server owner has ended the stream"
	}

	sessionID = url.QueryEscape(sessionID)
	reason = url.QueryEscape(reason)

	query := fmt.Sprintf("%s/status/sessions/terminate?sessionId=%s&reason=%s", p.URL, sessionID, reason)

	newHeaders := p.Headers
	newHeaders.Accept = "application/xml"

	resp, err := p.get(query, newHeaders)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", resp.Status)
	}

	return nil
}
