package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yingshulu/mars/auth"
	"github.com/yingshulu/mars/auth/jwt"
	"github.com/yingshulu/mars/config"
)

const (
	oidcKey = "OIDC"
	maxAge  = 3600
)

func cookie(c *gin.Context) (*http.Cookie, error) {
	return c.Request.Cookie(oidcKey)
}

func setCookie(u *auth.User, c *gin.Context) error {
	now := time.Now()
	u.Thumb = ""
	u.ValidFrom = now

	conf := config.Global()
	ma := conf.Cookie.MaxAge
	if ma <= 0 {
		ma = maxAge
	}

	u.ValidTo = now.Add(time.Duration(ma) * time.Second)
	token, err := jwt.Encode(conf.Security.Signer.Name, u)
	if err != nil {
		return err
	}

	httpOnly := conf.Cookie.HttpOnly != 0
	secure := conf.Cookie.Secure != 0
	c.SetCookie(oidcKey, string(token), ma, "/", conf.Host.Domain, secure, httpOnly)
	return nil
}
