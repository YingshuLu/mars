package jwt

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/yingshulu/mars/auth/sign"
)

type decoder struct {
	validator sign.Validator
	err       error
}

func (d *decoder) decode(token []byte, payload interface{}) error {
	if d.err != nil {
		return d.err
	}

	segments := bytes.Split(token, []byte("."))
	if len(segments) != 3 {
		d.err = fmt.Errorf("invalid jwt, segments size: %v", len(segments))
		return d.err
	}

	data := segments[1]
	sig := segments[2]
	d.jsonUnmarshal(d.base64Decode(data), payload)
	if d.err != nil {
		return d.err
	}

	if !d.validator.Check(token[:d.indexSig(token)-1], d.base64Decode(sig)) {
		d.err = fmt.Errorf("invalid jwt(%v), signature(%v) validation failure", string(token), string(sig))
	}

	return d.err
}

func (d *decoder) indexSig(token []byte) int {
	sum := 0
	for i, c := range token {
		if c == '.' {
			if sum != 0 {
				return i + 1
			} else {
				sum++
			}
		}
	}
	return -1
}

func (d *decoder) base64Decode(data []byte) []byte {
	var s []byte
	if d.err != nil {
		return s
	}
	s, d.err = base64.StdEncoding.DecodeString(string(data))
	return s
}

func (d *decoder) jsonUnmarshal(data []byte, payload interface{}) {
	if d.err != nil {
		return
	}
	d.err = json.Unmarshal(data, payload)
}
