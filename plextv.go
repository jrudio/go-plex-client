package plex

// I'll slowly migrate plex.tv related functions to this file

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"
)

type requestPinResponse struct {
	ID               int         `json:"id"`
	Code             string      `json:"code"`
	ClientIdentifier string      `json:"clientIdentifier"`
	ExpiresAt        time.Time   `json:"expiresAt"`
	AuthToken        interface{} `json:"authToken"`
}

// RequestPIN will retrieve a code (valid for 4 minutes) from plex.tv to link an app to your plex account
func RequestPIN() (string, error) {
	endpoint := "/api/v2/pins.json"

	// POST request and returns a 201 status code
	// response body returns json
	//
	// {
	// 		id: 123456,
	// 		code: "ABCD",
	// 		clientIdentifier: "goplexclient",
	// 		expiresAt: 15463757,
	// 		authToken: null
	// }

	resp, err := post(plexURL+endpoint, nil, defaultHeaders())

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return "", errors.New(resp.Status)
	}

	var pinResponse requestPinResponse

	if err := json.NewDecoder(resp.Body).Decode(&pinResponse); err != nil {
		return "", err
	}

	return pinResponse.Code, nil
}

// LinkAccount allows you to authorize an app via a 4 character pin. returns nil on success
func (p Plex) LinkAccount(code string) error {
	endpoint := "/api/v2/pins/link.json"

	body := url.Values{
		"code": []string{code},
	}

	headers := defaultHeaders()

	headers.ContentType = "application/x-www-form-urlencoded"

	// PUT request with 'code: <4-character-pin>' in the body
	resp, err := p.put(plexURL+endpoint, []byte(body.Encode()), headers)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// type linkAccountResponse struct {

	// }

	// var

	// json.NewDecoder(resp.Body).Decode()

	// should return 204 for success
	if resp.StatusCode != http.StatusNoContent {
		return errors.New("failed to link account: " + resp.Status)
	}

	return nil
}
