package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_ConfigMarshal(t *testing.T) {
	c := &Config{
		Host: Host{
			Domain: "bulo.fun",
			Port:   80,
			UseSSL: 0,
		},
		Cookie: Cookie{
			MaxAge:   20,
			HttpOnly: 1,
		},
	}

	b, err := yaml.Marshal(c)
	assert.Nil(t, err)
	t.Logf("marshal config:\n%v", string(b))
}
