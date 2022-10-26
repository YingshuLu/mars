package jwt

import "github.com/yingshulu/mars/auth/sign"

type header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

// Encode payload to jwt bytes
func Encode(signerName string, payload interface{}) ([]byte, error) {
	e := &encoder{
		signer: sign.FetchSigner(signerName),
	}

	h := &header{
		Alg: "RS256",
		Typ: "JWT",
	}

	token := e.encode(h, payload)
	return token, e.err
}

// Decode jwt bytes to payload and validate signature
func Decode(validatorName string, token []byte, payload interface{}) error {
	d := &decoder{
		validator: sign.FetchValidator(validatorName),
	}

	return d.decode(token, payload)
}
