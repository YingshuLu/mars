package jwt

import (
	"bytes"
	"encoding/base64"
	"encoding/json"

	"github.com/yingshulu/mars/auth/sign"
)

type encoder struct {
	signer sign.Signer
	buf    bytes.Buffer
	err    error
}

func (e *encoder) encode(h *header, payload interface{}) []byte {
	if e.err != nil {
		return nil
	}
	e.write(e.base64Encode(e.jsonMarshal(h)))
	e.writeByte('.')
	e.write(e.base64Encode(e.jsonMarshal(payload)))

	if e.err != nil {
		return nil
	}

	var sig []byte
	sig, e.err = e.signer.Sign(e.bytes())
	e.writeByte('.')
	e.write(e.base64Encode(sig))
	return e.bytes()
}

func (e *encoder) write(b []byte) {
	if e.err != nil {
		return
	}
	_, e.err = e.buf.Write(b)
}

func (e *encoder) writeByte(c byte) {
	if e.err != nil {
		return
	}
	e.err = e.buf.WriteByte(c)
}

func (e *encoder) bytes() []byte {
	return e.buf.Bytes()
}

func (e *encoder) base64Encode(d []byte) []byte {
	if e.err != nil {
		return nil
	}
	return []byte(base64.StdEncoding.EncodeToString(d))
}

func (e *encoder) jsonMarshal(any interface{}) []byte {
	if e.err != nil {
		return nil
	}
	var d []byte
	d, e.err = json.Marshal(any)
	return d
}
