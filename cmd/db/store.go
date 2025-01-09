package db

import (
	// "encoding/json"
	"fmt"

	"github.com/dgraph-io/badger/v3"
)

var (
	isVerbose bool
)

// type server struct {
// 	Name string `json:"name"`
// 	URL  string `json:"url"`
// }

// func (s server) Serialize() ([]byte, error) {
// 	return json.Marshal(s)
// }

// func unserializeServer(serializedServer []byte) (server, error) {
// 	var s server

// 	err := json.Unmarshal(serializedServer, &s)

// 	return s, err
// }

type DB struct {
	conn       *badger.DB
	isClosed bool
	// keys     storeKeys
}

type DBOptions struct {
	IsVerbose bool
}

// type storeKeys struct {
// 	appSecret  []byte
// 	plexToken  []byte
// 	plexServer []byte
// }

func New(dir string) (*DB, error) {
	db := new(DB)

	if isVerbose {
		fmt.Println("checking if our database exists in the home directory at:", dir)
	}

	options := badger.DefaultOptions(dir)

	options = options.WithLoggingLevel(badger.WARNING)

	conn, err := badger.Open(options)

	if err != nil {
		return db, err
	}

	if isVerbose {
		fmt.Println("successfully opened data store")
	}

	db.conn = conn
	// db.keys = storeKeys{
	// 	appSecret:  []byte("app-secret"),
	// 	plexToken:  []byte("plex-token"),
	// 	plexServer: []byte("plex-server"),
	// }

	return db, nil
}

func (d DB) Close() {
	if d.isClosed {
		fmt.Println("data store already closed")
		return
	}

	if err := d.conn.Close(); err != nil {
		fmt.Printf("data store failed to closed: %v\n", err)
	}

	d.isClosed = true
}

func (d DB) AddKey() error {
	tx := d.conn.NewTransaction(true)

	tx.SetEntry()
}

// func (s DB) getPlexToken() (string, error) {
// 	var plexToken string

// 	if err := s.db.View(func(txn *badger.Txn) error {
// 		item, err := txn.Get(s.keys.plexToken)

// 		if err != nil {
// 			return err
// 		}

// 		return item.Value(func(val []byte) error {
// 			plexToken = string(append([]byte{}, val...))

// 			return nil
// 		})
// 	}); err != nil {
// 		return plexToken, err
// 	}

// 	if isVerbose {
// 		fmt.Printf("Your plex token is %s\n", plexToken)
// 	}

// 	return plexToken, nil
// }

// func (s DB) savePlexToken(token string) error {
// 	if err := s.db.Update(func(txn *badger.Txn) error {
// 		return txn.Set(s.keys.plexToken, []byte(token))
// 	}); err != nil {
// 		return err
// 	}

// 	if isVerbose {
// 		fmt.Println("saved token to store")
// 	}

// 	return nil
// }

// func (s DB) getPlexServer() (server, error) {
// 	var plexServer server

// 	err := s.db.View(func(txn *badger.Txn) error {
// 		item, err := txn.Get(s.keys.plexServer)

// 		if err != nil {
// 			return err
// 		}

// 		return item.Value(func(val []byte) error {
// 			serializedServer := append([]byte{}, val...)

// 			plexServer, err = unserializeServer(serializedServer)

// 			if err != nil {
// 				return err
// 			}

// 			return nil
// 		})
// 	})

// 	return plexServer, err
// }

// func (s DB) savePlexServer(plexServer server) error {
// 	serializedServer, err := plexServer.Serialize()
// 	if err != nil {
// 		return err
// 	}

// 	return s.db.Update(func(txn *badger.Txn) error {
// 		return txn.Set(s.keys.plexServer, serializedServer)
// 	})
// }
