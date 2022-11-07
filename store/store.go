package store

import (
	"sync"
	"time"

	badger "github.com/dgraph-io/badger/v3"
	"github.com/yingshulu/mars/config"
)

var (
	storageTable = map[string]Storage{}
	storageMutex sync.RWMutex
)

type Storage interface {
	Set(k string, v interface{}, d time.Duration) error
	Get(k string, v interface{}) error
}

func New(path string) (Storage, error) {
	storageMutex.RLock()
	s, ok := storageTable[path]
	if ok {
		storageMutex.RUnlock()
		return s, nil
	}

	storageMutex.Lock()
	defer storageMutex.Unlock()
	s, ok = storageTable[path]
	if ok {
		return s, nil
	}

	s, err := open(path)
	if err != nil {
		return nil, err
	}
	storageTable[path] = s
	return s, nil
}

func open(path string) (*storage, error) {
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return nil, err
	}
	s := &storage{
		db:         db,
		Serializer: serializerMap[config.Global().Storage.Serializer],
	}
	return s, nil
}

type storage struct {
	db *badger.DB
	Serializer
}

func (s *storage) Set(k string, v interface{}, d time.Duration) error {
	payload, err := s.Marshal(v)
	if err != nil {
		return err
	}

	if d <= 0 {
		return s.set(k, payload)
	}
	return s.setTTL(k, payload, d)
}

func (s *storage) Get(k string, v interface{}) error {
	return s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(k))
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			return s.Unmarshal(val, v)
		})
		return err
	})
}

func (s *storage) set(k string, payload []byte) error {
	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(k), payload)
	})
}

func (s *storage) setTTL(k string, payload []byte, d time.Duration) error {
	return s.db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(k), payload).WithTTL(d)
		return txn.SetEntry(e)
	})
}
