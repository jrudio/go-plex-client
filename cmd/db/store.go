package db

import (
	"fmt"

	"github.com/dgraph-io/badger/v3"
)

var (
	isVerbose bool
)

type DB struct {
	conn       *badger.DB
	isClosed bool
	keys     storeKeys
}

type DBOptions struct {
	IsVerbose bool
}

type storeKeys struct {
	authorizations []byte
	// appSecret  []byte
	// plexToken  []byte
	// plexServer []byte
}

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
	db.keys = storeKeys{
		authorizations: []byte(KeyAuthorizations),
	// 	appSecret:  []byte("app-secret"),
	// 	plexToken:  []byte("plex-token"),
	// 	plexServer: []byte("plex-server"),
	}

	if err = db.initKeys(); err != nil {
		return db, err
	}

	return db, nil
}

func (db *DB) initKeys() error {
	return db.AddKey(db.keys.authorizations)
}

func (db *DB) Close() {
	if db.isClosed {
		fmt.Println("data store already closed")
		return
	}

	if err := db.conn.Close(); err != nil {
		fmt.Printf("data store failed to closed: %v\n", err)
	}

	db.isClosed = true
}

func (db DB) AddKey(key []byte) error {
	return db.conn.Update(func (txn *badger.Txn) error {
		_, err := txn.Get(key)

		if err == badger.ErrKeyNotFound {
			// initialize key
			// fmt.Printf("key '%v' doesn't exist adding...\n", key)

			entry := badger.NewEntry(key, []byte{})

			return txn.SetEntry(entry)
		}

		return nil
	})
}

func (db DB) getData(key []byte) ([]byte, error) {
	data := []byte{}

	err := db.conn.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)

		if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {
			data = append([]byte{}, val...)

			return nil
		})
	})

	if err != nil {
		return data, err
	}

	return data, nil
}

// saveData provides a straightforward save to database method
// we should not encounter badger.ErrKeyNotFound if we properly create the keys in our init process
func (db DB) saveData(key []byte, data []byte) error {
	return db.conn.Update(func(txn *badger.Txn) error {
		entry := badger.NewEntry(key, data)

		return txn.SetEntry(entry)
	})
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
