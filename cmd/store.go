package main

import (
	"fmt"
	"os"

	"github.com/dgraph-io/badger"
)

type store struct {
	db     *badger.DB
	keys   storeKeys
	secret []byte
}

type storeKeys struct {
	appSecret []byte
	plexToken []byte
}

func initDataStore(dirName string) (store, error) {
	var db store

	if isVerbose {
		fmt.Println("checking if our database exists in the home directory at:", dirName)
	}

	// create a directory for our database
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		if isVerbose {
			fmt.Println("creating directory because it doesn't exist")
		}

		if err := os.Mkdir(dirName, os.ModePerm); err != nil {
			return db, err
		}
	}

	options := badger.DefaultOptions

	options.Dir = dirName
	options.ValueDir = dirName

	kvStore, err := badger.Open(options)

	if err != nil {
		return db, err
	}

	if isVerbose {
		fmt.Println("successfully opened data store")
	}

	db.db = kvStore
	db.keys = storeKeys{
		appSecret: []byte("app-secret"),
		plexToken: []byte("plex-token"),
	}

	return db, nil
}
func (s store) Close() {
	if err := s.db.Close(); err != nil {
		fmt.Printf("data store failed to closed: %v\n", err)
	}
}

func (s store) getSecret() []byte {
	var secret []byte

	// an error is returned when the key is not found
	// so just return an empty secret
	s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(s.keys.appSecret)

		if err != nil {
			return err
		}

		_secret, err := item.Value()

		if err != nil {
			return err
		}

		secret = _secret

		return nil
	})

	return secret
}

func (s store) saveSecret(secret []byte) error {
	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Set(s.keys.appSecret, secret, 0)
	})
}

func (s store) getPlexToken() (string, error) {
	var plexToken string

	if err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(s.keys.plexToken)

		if err != nil {
			return err
		}

		tokenHash, err := item.Value()

		if err != nil {
			return err
		}

		_plexToken, err := decrypt(s.secret, string(tokenHash))

		if err != nil {
			if isVerbose {
				fmt.Println("token decryption failed")
			}
			return err
		}

		plexToken = _plexToken

		return nil
	}); err != nil {
		return plexToken, err
	}

	if isVerbose {
		fmt.Printf("Your plex token is %s\n", plexToken)
	}

	return plexToken, nil
}

func (s store) savePlexToken(token string) error {
	tokenHash, err := encrypt(s.secret, token)

	if err != nil {
		return err
	}

	if isVerbose {
		fmt.Printf("your plex token hash: %s\n", string(tokenHash))
	}

	if err := s.db.Update(func(txn *badger.Txn) error {
		return txn.Set(s.keys.plexToken, []byte(tokenHash), 0x00)
	}); err != nil {
		return err
	}

	if isVerbose {
		fmt.Println("saved token hash to store")
	}

	return nil
}
