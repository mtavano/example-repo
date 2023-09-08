package storage

import "github.com/syndtr/goleveldb/leveldb/opt"

type dataStore struct {
	err error
}

func (ds *dataStore) Put(key, value []byte, wo *opt.WriteOptions) error {
	return ds.err
}
