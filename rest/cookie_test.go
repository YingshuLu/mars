package rest

import (
	"errors"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/gin-gonic/gin"
	"github.com/yingshulu/mars/auth"
	"github.com/yingshulu/mars/auth/jwt"
	"github.com/yingshulu/mars/config"
)

func Test_Cookie(t *testing.T) {
	var p = gomonkey.ApplyFunc(config.Global, func() *config.Config {
		return &config.Config{}
	})
	defer p.Reset()

	var p0 = gomonkey.ApplyFunc(domain, func(c *gin.Context) string {
		return "bulo.fun"
	})
	defer p0.Reset()

	c := &gin.Context{}
	p1 := gomonkey.ApplyMethod(reflect.TypeOf(c), "SetCookie", func(_ *gin.Context, name string,
		value string, maxAge int, path string, domain string, secure bool, httpOnly bool) {
		return
	})
	defer p1.Reset()

	tests := []struct {
		name string
		data []byte
		err  error

		wantErr bool
	}{
		{
			name: "no_error",

			data: []byte("test"),
			err:  nil,

			wantErr: false,
		},
		{
			name: "error",

			data: nil,
			err:  errors.New("mock error"),

			wantErr: true,
		},
	}

	for _, te := range tests {
		var p2 = gomonkey.ApplyFunc(jwt.Encode, func(signerName string, payload interface{}) ([]byte, error) {
			return te.data, te.err
		})

		err := setCookie(&auth.User{}, c)
		if te.wantErr != (err != nil) {
			t.Errorf("case: %s, error: %v", te.name, err)
		}
		p2.Reset()
	}

}
