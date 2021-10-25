package main

import (
	"fmt"
	"os"

	"github.com/dgraph-io/badger/v3"
)

type store struct {
	db       *badger.DB
	isClosed bool
	keys     storeKeys
}

type storeKeys struct {
	appSecret  []byte
	plexToken  []byte
	plexServer []byte
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

	options := badger.DefaultOptions(dirName)

	options = options.WithLoggingLevel(badger.WARNING)

	kvStore, err := badger.Open(options)

	if err != nil {
		return db, err
	}

	if isVerbose {
		fmt.Println("successfully opened data store")
	}

	db.db = kvStore
	db.keys = storeKeys{
		appSecret:  []byte("app-secret"),
		plexToken:  []byte("plex-token"),
		plexServer: []byte("plex-server"),
	}

	return db, nil
}

func (s store) Close() {
	if s.isClosed {
		fmt.Println("data store already closed")
		return
	}

	if err := s.db.Close(); err != nil {
		fmt.Printf("data store failed to closed: %v\n", err)
	}

	s.isClosed = true
}

func (s store) getPlexToken() (string, error) {
	var plexToken string

	if err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(s.keys.plexToken)

		if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {
			plexToken = string(append([]byte{}, val...))

			return nil
		})
	}); err != nil {
		return plexToken, err
	}

	if isVerbose {
		fmt.Printf("Your plex token is %s\n", plexToken)
	}

	return plexToken, nil
}

func (s store) savePlexToken(token string) error {
	if err := s.db.Update(func(txn *badger.Txn) error {
		return txn.Set(s.keys.plexToken, []byte(token))
	}); err != nil {
		return err
	}

	if isVerbose {
		fmt.Println("saved token to store")
	}

	return nil
}

func (s store) getPlexServer() (server, error) {
	var plexServer server

	err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(s.keys.plexServer)

		if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {
			serializedServer := append([]byte{}, val...)

			plexServer, err = unserializeServer(serializedServer)

			if err != nil {
				return err
			}

			return nil
		})
	})

	return plexServer, err
}

func (s store) savePlexServer(plexServer server) error {
	serializedServer, err := plexServer.Serialize()
	if err != nil {
		return err
	}

	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Set(s.keys.plexServer, serializedServer)
	})
}
