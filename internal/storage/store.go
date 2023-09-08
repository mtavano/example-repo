package storage

import (
	"encoding/json"
	"fmt"

	"github.com/mtavano/devkit/errors"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type Store struct {
	// db *leveldb.DB
	db DataStore
}

func NewStore(dbName string) (*Store, error) {
	pathToDB := fmt.Sprintf("data/%s", dbName)
	db, err := leveldb.OpenFile(pathToDB, nil)
	if err != nil {
		return nil, err
	}
	return &Store{db: db}, nil
}

func (st *Store) Put(key string, value interface{}) error {
	inputKey := []byte(key)
	inputData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = st.db.Put(inputKey, inputData, nil)
	if err != nil {
		return errors.Wrap(err, "storage: Store.Put st.db.Put error")
	}

	return nil
}

func (st *Store) Get(key string) ([]byte, error) {
	return nil, nil
}

type DataStore interface {
	Put(key, value []byte, wo *opt.WriteOptions) error
}
