package db

import (
	"encoding/json"
	"fmt"
)

const (
	KeyAuthorizations = "authorizations"
)

type Authorization struct {
	Email     string `json:"emailAddress"`
	PlexToken string `json:"plexToken"`
	IsActive  bool   `json:"isActive"`
}

type Authorizations []Authorization

func (db DB) SaveAuth(auth Authorization) error {
	if auth.Email == "" {
		return fmt.Errorf("email must be provided")
	} else if auth.PlexToken == "" {
		return fmt.Errorf("plex token must be provided")
	}

	auths, err := db.GetAuthorizations()

	if err != nil {
		return fmt.Errorf("failed getting authorizations: %v", err)
	}

	if index := auths.FindIndexByEmail(auth.Email); index > -1 {
		// account should already be active, so just update plex token
		auths[index].PlexToken = auth.PlexToken
	} else {
		auths = append(auths, auth)
	}

	if err := auths.SetActive(auth.Email); err != nil {
		return fmt.Errorf("unable to set active account: %v", err)
	}

	data, err := auths.toJSON()

	if err != nil {
		return fmt.Errorf("converting authorization to json failed: %v", err)
	}

	if err = db.saveData(db.keys.authorizations, data); err != nil {
		return fmt.Errorf("saving authorization failed: %v", err)
	}

	return nil
}

func (db DB) RemoveAuth(auth Authorization) error {
	if auth.Email == "" {
		return fmt.Errorf("email must be provided")
	}

	auths, err := db.GetAuthorizations()

	if err != nil {
		return fmt.Errorf("failed getting authorizations: %v", err)
	}

	index := auths.FindIndexByEmail(auth.Email)

	if index < 0 {
		return fmt.Errorf("email not found")
	}

	newAuths := Authorizations{}

	setDefaultCredentials := func(auths Authorizations, newAuths Authorizations) (Authorizations, error) {
		if len(auths) == 0 {
			return newAuths, fmt.Errorf("no credentials to make active")
		}

		firstAuth := auths[0]

		if err := newAuths.SetActive(firstAuth.Email); err != nil {
			return newAuths, fmt.Errorf("unable to set new auth: %v", err)
		}

		return newAuths, nil
	}

	if authLen := len(auths); authLen > 1 {
		newAuths = append(newAuths, auths[:index]...)
		newAuths = append(newAuths, auths[index + 1: authLen]...)

		// if what we're removing is active, then set the first credential as active
		if auths[index].IsActive {
			newAuths, err = setDefaultCredentials(auths, newAuths)

			if err != nil {
				return fmt.Errorf("unable to set default credentials: %v", err)
			}

		}
	}

	data, err := newAuths.toJSON()

	if err != nil {
		return fmt.Errorf("failed converting authorization to json: %v", err)
	}

	if db.saveData([]byte(KeyAuthorizations), data); err != nil {
		return fmt.Errorf("failed saving authorizations: %v", err)
	}

	return nil
}

func (db DB) GetAuthorizations() (Authorizations, error) {
	auths := Authorizations{}

	data, err := db.getData(db.keys.authorizations)

	if err != nil {
		return auths, fmt.Errorf("failed fetching data from db: %v", err)
	}

	if len(data) > 0 {
		if err = auths.fromJSON(data); err != nil {
			return auths, fmt.Errorf("failed deserializing data: %v", err)
		}
	}

	return auths, nil
}

// SetActive will set the active credential by email
func (a *Authorizations) SetActive(email string) error {
	setActive := func(auths *Authorizations, email string) int {
		foundIndex := -1

		for i := 0; i < len(*auths); i++ {
			auth := (*auths)[i]

			if auth.Email == email {
				(*auths)[i].IsActive = true

				foundIndex = i
			}
		}

		return foundIndex
	}

	setInactive := func(auths *Authorizations, skipIndex int) {
		for i := 0; i < len(*auths); i++ {
			if skipIndex > -1 && i == skipIndex {
				continue
			}

			(*auths)[i].IsActive = false
		}
	}

	authIndex := setActive(a, email)

	if authIndex < 0 {
		return fmt.Errorf("email not found")
	}

	setInactive(a, authIndex)

	return nil
}

func (a *Authorizations) FindIndexByEmail(email string) int {
	for i := 0; i < len(*a); i++ {
		auth := (*a)[i]

		if auth.Email == email {
			return i
		}
	}

	return -1
}

func (a *Authorizations) IsExists(email string) bool {
	return a.FindIndexByEmail(email) > -1
}

func (a Authorizations) toJSON() ([]byte, error) {
	return json.Marshal(a)
}

// fromJSON fills struct with authorization data from the database by parsing raw bytes
func (a *Authorizations) fromJSON(data []byte) error {
	return json.Unmarshal(data, &a)
}
