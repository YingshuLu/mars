package rest

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yingshulu/mars/auth"
	"github.com/yingshulu/mars/auth/jwt"
	"github.com/yingshulu/mars/config"
)

func Auth(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			c.JSON(http.StatusUnauthorized, struct {
				Error string `json:"error"`
			}{
				Error: err.Error(),
			})
			c.Abort()
			return
		}
	}()

	co, err := cookie(c)
	if err != nil {
		return
	}

	token, err := url.QueryUnescape(co.Value)
	if err != nil {
		return
	}

	conf := config.Global()
	u := &auth.User{}
	err = jwt.Decode(conf.Security.Validator.Name, []byte(token), u)
	if err != nil {
		return
	}

	now := time.Now()
	if u.ValidTo.Before(now) {
		err = fmt.Errorf("user %v cookie expires", u.Name)
		return
	}

	if now.Sub(u.ValidTo) < time.Minute {
		err = setCookie(u, c)
	}

	if err != nil {
		return
	}

	auth.SetUser(c, u)
	c.Next()
}
