package sign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rsaSigner(t *testing.T) {
	s := FetchSigner("rsa_single")
	d := FetchValidator("rsa_single")

	data := []byte("test message to sign and check")
	sn, err := s.Sign(data)
	assert.Nil(t, err)
	assert.True(t, d.Check(data, sn))
	assert.False(t, d.Check([]byte("forge signure"), sn))
}
