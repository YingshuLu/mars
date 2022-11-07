package sign

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

var (
	signers    = map[string]Signer{}
	validators = map[string]Validator{}
)

var (
	singleRSAName          = "rsa_single"
	singleRSASignValidator = func() *rsaSigner {
		pk, _ := rsa.GenerateKey(rand.Reader, 1024)
		return &rsaSigner{
			prk: pk,
			pbk: &pk.PublicKey,
		}
	}()
)

func init() {
	RegisterSigner(singleRSAName, singleRSASignValidator)
	RegisterValidator(singleRSAName, singleRSASignValidator)
}

// FetchSigner get signer by name
func FetchSigner(name string) Signer {
	return signers[name]
}

// RegisterSigner register signer with name
func RegisterSigner(name string, s Signer) {
	signers[name] = s
}

// FetchValidator get validator by name
func FetchValidator(name string) Validator {
	return validators[name]
}

// RegisterValidator register validator with name
func RegisterValidator(name string, v Validator) {
	validators[name] = v
}

// Signer sign data
type Signer interface {
	Sign(data []byte) ([]byte, error)
}

// Validator validate sign of data
type Validator interface {
	Check(data []byte, sig []byte) bool
}

type rsaSigner struct {
	prk *rsa.PrivateKey
	pbk *rsa.PublicKey
}

func (s *rsaSigner) Sign(data []byte) ([]byte, error) {
	return s.prk.Sign(rand.Reader, s.hashSum(data), crypto.SHA256)
}

func (s *rsaSigner) Check(data []byte, sig []byte) bool {
	err := rsa.VerifyPKCS1v15(s.pbk, crypto.SHA256, s.hashSum(data), sig)
	return err == nil
}

func (s *rsaSigner) hashSum(data []byte) []byte {
	r := sha256.Sum256(data)
	return r[:]
}
