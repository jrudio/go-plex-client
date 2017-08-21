package plex

// I'll slowly migrate plex.tv related functions to this file

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

// ErrorResponse contains a code and an error message
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// PinResponse holds information to successfully check a pin when linking an account
type PinResponse struct {
	ID               int             `json:"id"`
	Code             string          `json:"code"`
	ClientIdentifier string          `json:"clientIdentifier"`
	ExpiresAt        int64           `json:"expiresAt"`
	AuthToken        string          `json:"authToken"`
	Errors           []ErrorResponse `json:"errors"`
}

// RequestPIN will retrieve a code (valid for 15 minutes) from plex.tv to link an app to your plex account
func RequestPIN() (PinResponse, error) {
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
	var pinInformation PinResponse

	resp, err := post(plexURL+endpoint, nil, defaultHeaders())

	if err != nil {
		return pinInformation, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return pinInformation, errors.New(resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&pinInformation); err != nil {
		return pinInformation, err
	}

	return pinInformation, nil
}

// CheckPIN will return information related to the pin such as the auth token if your code has been approved.
// will return an error if code expired or still not linked
// clientIdentifier must be the same when requesting a pin
func CheckPIN(id int, clientIdentifier string) (PinResponse, error) {
	endpoint := "/api/v2/pins/"

	endpoint = endpoint + strconv.Itoa(id) + ".json"

	headers := defaultHeaders()

	resp, err := get(plexURL+endpoint, headers)

	if err != nil {
		return PinResponse{}, err
	}

	defer resp.Body.Close()

	var pinInformation PinResponse

	if err := json.NewDecoder(resp.Body).Decode(&pinInformation); err != nil {
		return pinInformation, err
	}

	// code doesn't exist or expired
	if len(pinInformation.Errors) > 0 {
		return pinInformation, errors.New(pinInformation.Errors[0].Message)
	}

	// we are not authorized yet
	if pinInformation.AuthToken == "" {
		return pinInformation, errors.New("pin is not authorized yet")
	}

	// we are authorized! Yay!
	return pinInformation, nil
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
