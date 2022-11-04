package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/mail"
	"sync"
	"time"

	badger "github.com/dgraph-io/badger/v3"
)

const userDBPath = "./db/user.db"

var (
	userDB *badger.DB
	dbOnce sync.Once
)

func db() *badger.DB {
	dbOnce.Do(func() {
		if userDB != nil {
			db, err := badger.Open(badger.DefaultOptions(userDBPath))
			if err != nil {
				panic(fmt.Sprintf("open user badger db path (%v) error: %v", userDBPath, err))
			}
			userDB = db
		}
	})
	return userDB
}

// User definition
type User struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	Email     *mail.Address `json:"email"`
	ValidFrom time.Time     `json:"valid_from"`
	ValidTo   time.Time     `json:"valid_to"`
	Thumb     string        `json:"thumb"`
}

// Search if user exists
func Search(u *User) (found bool, s Status) {
	err := db().View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(u.Name))
		return err
	})

	found, s = false, Ok
	switch {
	case err == nil:
		found = true
	case errors.Is(err, badger.ErrKeyNotFound):
		found = false
	default:
		found = false
		s = InternalError
	}
	return
}

// Validate user login
func Validate(u *User) Status {
	s := Ok
	switch {
	case len(u.Name) == 0:
		s = UserNotFound
	case len(u.Thumb) == 0:
		s = PasswordIncorrect
	default:
		err := db().View(func(txn *badger.Txn) error {
			item, err := txn.Get([]byte(u.Name))
			if err != nil {
				return err
			}

			val, err := item.ValueCopy(nil)
			if err != nil {
				return err
			}

			pu := &User{}
			err = json.Unmarshal(val, pu)
			if err != nil {
				return err
			}
			if u.Thumb != pu.Thumb {
				s = PasswordIncorrect
			}
			return nil
		})
		if err != nil {
			s = InternalError
		}
	}
	return s
}

func Add(u *User) error {
	return db().Update(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(u.Name))
		if err != nil {
			return err
		}

		val, err := json.Marshal(u)
		if err != nil {
			return err
		}

		return txn.Set([]byte(u.Name), val)
	})
}
