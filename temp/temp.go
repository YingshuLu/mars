package temp

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/url"
	"time"

	"github.com/yingshulu/mars/auth"
	"github.com/yingshulu/mars/store"
)

type Temp interface {
	Gen(u *auth.User, path string, duration time.Duration) (*url.URL, error)
	Check(*url.URL) error
}

type Session struct {
	User      *auth.User `yaml:"user"`
	Path      string     `yaml:"path"`
	ExpiredAt time.Time  `yaml:"expired_at"`
}

type tempUrl struct {
	store.Storage
}

func (t *tempUrl) Gen(u *auth.User, path string, d time.Duration) (*url.URL, error) {
	r := rand.Reader
	var buf [64]byte
	n, err := r.Read(buf[:])
	if err != nil {
		return nil, err
	}

	hash := md5.Sum(buf[:n])
	key := hex.EncodeToString(hash[:])

	s := &Session{
		User:      u,
		Path:      path,
		ExpiredAt: time.Now().Add(d),
	}

	err = t.Set(key, s, d)
	if err != nil {
		return nil, err
	}

	tu := &url.URL{
		Path: path,
	}
	tu.Query().Add("token", key)

	return tu, nil
}

func (t *tempUrl) Check(u *url.URL) error {
	key := u.Query().Get("token")

	s := &Session{}
	err := t.Get(key, s)
	if err != nil {
		return err
	}

	if time.Now().Before(s.ExpiredAt) {
		return errors.New("access expired")
	}
	return nil
}
