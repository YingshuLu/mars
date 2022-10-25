package sign

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

var signers = map[string]Signer{}

func Fetch(name string) Signer {
	return signers[name]
}

func Register(name string, s Signer) {
	signers[name] = s
}

type Signer interface {
	Sign(data []byte) ([]byte, error)
	Check(data []byte, sn []byte) bool
}

type rsaSigner struct {
	prk *rsa.PrivateKey
}

func (s *rsaSigner) Sign(data []byte) ([]byte, error) {
	return s.prk.Sign(rand.Reader, s.hashSum(data), crypto.SHA256)
}

func (s *rsaSigner) Check(data []byte, sn []byte) bool {
	err := rsa.VerifyPKCS1v15(&s.prk.PublicKey, crypto.SHA256, s.hashSum(data), sn)
	return err == nil
}

func (s *rsaSigner) hashSum(data []byte) []byte {
	r := sha256.Sum256(data)
	return r[:]
}
