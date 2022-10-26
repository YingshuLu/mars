package jwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EncodeDecode(t *testing.T) {
	type Payload struct {
		Name string
		Age  int
	}
	p := &Payload{
		Name: "kamlu",
		Age:  29,
	}

	signName := "rsa_single"
	token, err := Encode(signName, p)
	assert.Nil(t, err)

	u := &Payload{}
	err = Decode(signName, token, u)
	assert.Nil(t, err)
	assert.Equal(t, p, u)
}
