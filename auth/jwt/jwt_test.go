package jwt

import (
	"net/mail"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_EncodeDecode(t *testing.T) {
	type User struct {
		ID        string        `json:"id"`
		Name      string        `json:"name"`
		Email     *mail.Address `json:"email"`
		ValidFrom time.Time     `json:"valid_from"`
		ValidTo   time.Time     `json:"valid_to"`
		Thumb     string        `json:"thumb"`
	}

	p := &User{
		Name:      "jack",
		ValidFrom: time.Now(),
		ValidTo:   time.Now().Add(12 * time.Hour),
	}

	signName := "rsa_single"
	token, err := Encode(signName, p)
	assert.Nil(t, err)

	t.Log("jwt token: ", string(token))

	u := &User{}
	err = Decode(signName, token, u)
	assert.Nil(t, err)
}
