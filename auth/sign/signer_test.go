package sign

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rsaSigner(t *testing.T) {
	pk, err := rsa.GenerateKey(rand.Reader, 1024)
	assert.Nil(t, err)

	s := &rsaSigner{
		prk: pk,
	}

	data := []byte("test message to sign and check")
	sn, err := s.Sign(data)
	assert.Nil(t, err)
	assert.True(t, s.Check(data, sn))
	assert.False(t, s.Check([]byte("forge signure"), sn))
}
