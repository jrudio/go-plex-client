package plex

// package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"runtime"
	"strconv"
)

const (
	plexURL = "https://plex.tv"
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

// New creates a new plex instance that is required to
// to make requests to your Plex Media Server
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

// Search your Plex Server for media
func (p *Plex) Search(title string) (SearchResults, error) {
	if title == "" {
		return SearchResults{}, errors.New("ERROR: A title is required")
	}

	requestInfo.headers.Token = p.token

	title = url.QueryEscape(title)
	query := p.URL + "/search?query=" + title

	var results SearchResults

	resp, respErr := requestInfo.get(query)

	if respErr != nil {
		return SearchResults{}, respErr
	}

	// Unauthorized
	if resp.StatusCode == 401 {
		return SearchResults{}, errors.New("You are not authorized to access that server")
	}

	defer resp.Body.Close()

	err := json.NewDecoder(resp.Body).Decode(&results)

	if err != nil {
		return SearchResults{}, err
	}

	return results, nil
}

// Test your connection to your Plex Media Server
func (p *Plex) Test() (bool, error) {
	requestInfo.headers.Token = p.token

	resp, respErr := requestInfo.get(p.URL)

	if respErr != nil {
		return false, respErr
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
	requestInfo.headers.Token = p.token

	if sessionKey == "" {
		return false, errors.New("Missing sessionKey")
	}

	query := p.URL + "/video/:/transcode/universal/stop?session=" + sessionKey

	resp, respErr := requestInfo.get(query)

	if respErr != nil {
		return false, respErr
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
func (p *Plex) GetTranscodeSessions() (transcodeSessionsResponse, error) {
	var result transcodeSessionsResponse

	requestInfo.headers.Token = p.token

	query := p.URL + "/transcode/sessions"

	resp, respErr := requestInfo.get(query)

	if respErr != nil {
		return result, respErr
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
func (p *Plex) GetPlexTokens(token string) (devicesResponse, error) {
	var result devicesResponse

	requestInfo.headers.Token = p.token

	query := plexURL + "/devices.json"

	resp, respErr := requestInfo.get(query)

	if respErr != nil {
		return result, respErr
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

// DeletePlexToken is currently not test
func (p *Plex) DeletePlexToken(token string) (bool, error) {
	var result bool

	requestInfo.headers.Token = p.token

	query := plexURL + "/devices/" + token + ".json"

	resp, respErr := requestInfo.get(query)

	if respErr != nil {
		return result, respErr
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
func (p *Plex) GetFriends() ([]friends, error) {
	requestInfo.headers.Token = p.token

	var plexFriendsResp friendsResponse

	query := plexURL + "/api/users"

	resp, respErr := requestInfo.get(query)

	if respErr != nil {
		return []friends{}, respErr
	}

	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		return []friends{}, errors.New("You are not authorized to access this server")
	} else if resp.StatusCode != 200 {
		statusCode := strconv.Itoa(resp.StatusCode)
		return []friends{}, errors.New("Server replied with " + statusCode + " status code")
	}

	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return []friends{}, err
	}

	err = xml.Unmarshal(respBytes, &plexFriendsResp)

	if err != nil {
		return []friends{}, err
	}

	friendCount := plexFriendsResp.Size

	plexFriends := make([]friends, friendCount)

	for ii, f := range plexFriendsResp.User {
		plexFriends[ii] = f
	}

	return plexFriends, nil
}

// RemoveFriend from your friend's list which stops access to your Plex server
func (p *Plex) RemoveFriend(id string) (bool, error) {
	requestInfo.headers.Token = p.token

	query := plexURL + "/api/friends/" + id

	resp, respErr := requestInfo.delete(query)

	if respErr != nil {
		return false, respErr
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 400 {
		return false, errors.New(resp.Status)
	}

	result := new(resultResponse)

	if err := xml.NewDecoder(resp.Body).Decode(result); err != nil {
		return false, err
	}

	return result.Code == 0, nil
}

// InviteFriend to access your Plex server. Add restrictions to media or give them full access.
func (p *Plex) InviteFriend(params InviteFriendParams) (bool, error) {
	requestInfo.headers.Token = p.token

	usernameOrEmail := url.QueryEscape(params.UsernameOrEmail)
	label := url.QueryEscape(params.Label)

	query := fmt.Sprintf("%s/api/servers/%s/shared_servers", plexURL, params.MachineID)

	var restrictions inviteFriendBody

	restrictions.ServerID = params.MachineID
	restrictions.SharedServer = inviteFriendSharedServer{
		InvitedEmail:      params.UsernameOrEmail,
		LibrarySectionIDs: params.LibraryIDs,
	}

	settings := inviteFriendSharingSettings{
		FilterMovies:     fmt.Sprintf("label=%s", label),
		FilterTelevision: fmt.Sprintf("label=%s", label),
	}

	restrictions.SharingSettings = settings

	jsonBody, jsonErr := json.Marshal(restrictions)

	if jsonErr != nil {
		return false, jsonErr
	}

	resp, respErr := requestInfo.post(query, jsonBody)

	if respErr != nil {
		return false, respErr
	}

	defer resp.Body.Close()

	result := new(inviteFriendResponse)

	if err := xml.NewDecoder(resp.Body).Decode(result); err != nil {
		return false, err
	}

	if resp.StatusCode != 200 {
		return false, nil
	}

	sharedServer := result.SharedServer

	if sharedServer.Username != usernameOrEmail && sharedServer.Email != usernameOrEmail {
		return false, nil
	}

	return true, nil
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

	resp, respErr := requestInfo.put(query)

	if respErr != nil {
		return false, respErr
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

	resp, respErr := requestInfo.delete(query)

	if respErr != nil {
		return false, respErr
	}

	resp.Body.Close()

	if resp.StatusCode != 200 {
		return false, errors.New(resp.Status)
	}

	return true, nil
}

// CheckUsernameOrEmail will check if the username is a Plex user or will verify an email is valid
func (p *Plex) CheckUsernameOrEmail(usernameOrEmail string) (bool, error) {
	requestInfo.headers.Token = p.token

	usernameOrEmail = url.QueryEscape(usernameOrEmail)

	query := fmt.Sprintf("%s/api/users/validate?invited_email=%s", plexURL, usernameOrEmail)

	resp, respErr := requestInfo.post(query, nil)

	if respErr != nil {
		return false, respErr
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 400 {
		return false, errors.New(resp.Status)
	}

	result := new(resultResponse)

	if err := xml.NewDecoder(resp.Body).Decode(result); err != nil {
		return false, err
	}

	return result.Code == 0, nil
}

// GetServers A list of your Plex servers
func (p *Plex) GetServers() ([]pmsDevices, error) {
	requestInfo.headers.Token = p.token

	query := plexURL + "/pms/resources.xml?includeHttps=1"

	resp, respErr := requestInfo.get(query)

	if respErr != nil {
		return []pmsDevices{}, respErr
	}

	defer resp.Body.Close()

	result := new(resourcesResponse)

	if err := xml.NewDecoder(resp.Body).Decode(result); err != nil {
		fmt.Println(err.Error())

		return []pmsDevices{}, err
	}

	var servers []pmsDevices

	for _, r := range result.Device {
		if r.Provides != "server" {
			continue
		}

		servers = append(servers, r)
	}

	return servers, nil
}

// GetSectionIDs of your plex server. This is useful when inviting a user
// as you can restrict the invited user to a library (i.e. Movie's, TV Shows)
func (p *Plex) GetSectionIDs(machineID string) (sectionIDResponse, error) {
	requestInfo.headers.Token = p.token

	query := fmt.Sprintf("%s/api/servers/%s", plexURL, machineID)

	resp, respErr := requestInfo.get(query)

	if respErr != nil {
		return sectionIDResponse{}, respErr
	}

	defer resp.Body.Close()

	var result sectionIDResponse

	if err := xml.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println(err.Error())

		return sectionIDResponse{}, err
	}

	return result, nil
}

// GetLibraries of your Plex server. My ideal use-case would be
// to get library count to determine label index
func (p *Plex) GetLibraries() (librarySections, error) {
	requestInfo.headers.Token = p.token

	query := fmt.Sprintf("%s/library/sections", p.URL)

	resp, respErr := requestInfo.get(query)

	if respErr != nil {
		return librarySections{}, respErr
	}

	defer resp.Body.Close()

	var result librarySections

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		fmt.Println(err.Error())

		return librarySections{}, err
	}

	return result, nil
}

// GetLibraryLabels of your plex server
func (p *Plex) GetLibraryLabels(sectionKey, sectionIndex string) (libraryLabels, error) {
	requestInfo.headers.Token = p.token

	if sectionIndex == "" {
		sectionIndex = "1"
	}

	query := fmt.Sprintf("%s/library/sections/%s/labels?type=%s", p.URL, sectionKey, sectionIndex)

	resp, respErr := requestInfo.get(query)

	if respErr != nil {
		return libraryLabels{}, respErr
	}

	defer resp.Body.Close()

	var result libraryLabels

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		fmt.Println(err.Error())

		return libraryLabels{}, err
	}

	return result, nil
}

// AddLabelToMedia restrict access to certain media. Requires a Plex Pass.
// XXX: Currently plex is capitalizing the first letter
func (p *Plex) AddLabelToMedia(_type, id, label, locked string) (bool, error) {
	requestInfo.headers.Token = p.token

	query := fmt.Sprintf("%s/library/sections/3/all", p.URL)

	parsedQuery, parseErr := url.Parse(query)

	if parseErr != nil {
		return false, parseErr
	}

	vals := parsedQuery.Query()

	vals.Add("type", _type)
	vals.Add("id", id)
	vals.Add("label[0].tag.tag", label)
	vals.Add("label.locked", locked)

	parsedQuery.RawQuery = vals.Encode()

	query = parsedQuery.String()

	resp, respErr := requestInfo.put(query)

	if respErr != nil {
		return false, respErr
	}

	defer resp.Body.Close()

	return resp.StatusCode == 200, nil
}

// RemoveLabelFromMedia to remove a label from a piece of media Requires a Plex Pass.
func (p *Plex) RemoveLabelFromMedia(_type, id, label, locked string) (bool, error) {
	requestInfo.headers.Token = p.token

	query := fmt.Sprintf("%s/library/sections/3/all", p.URL)

	parsedQuery, parseErr := url.Parse(query)

	if parseErr != nil {
		return false, parseErr
	}

	vals := parsedQuery.Query()

	vals.Add("type", _type)
	vals.Add("id", id)
	vals.Add("label[].tag.tag-", label)
	vals.Add("label.locked", locked)

	parsedQuery.RawQuery = vals.Encode()

	query = parsedQuery.String()

	resp, respErr := requestInfo.put(query)

	if respErr != nil {
		return false, respErr
	}

	defer resp.Body.Close()

	return resp.StatusCode == 200, nil
}

// GetSessions of devices currently consuming media
func (p *Plex) GetSessions() (currentSessions, error) {
	acceptType := requestInfo.headers.Accept

	requestInfo.headers.Token = p.token
	// Don't request json
	requestInfo.headers.Accept = ""

	query := fmt.Sprintf("%s/status/sessions", p.URL)

	resp, respErr := requestInfo.get(query)

	if respErr != nil {
		return currentSessions{}, respErr
	}

	// Return value to what it was before we touched it
	requestInfo.headers.Accept = acceptType

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return currentSessions{}, errors.New(resp.Status)
	}

	var result currentSessions

	if err := xml.NewDecoder(resp.Body).Decode(&result); err != nil {
		return currentSessions{}, err
	}

	return result, nil
}
