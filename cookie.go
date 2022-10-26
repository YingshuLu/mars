package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yingshulu/mars/auth"
	"github.com/yingshulu/mars/auth/jwt"
)

func cookieWare(c *gin.Context) {
	co, err := c.Request.Cookie("oidc")
	if err != nil {
		c.Redirect(403, "/static/html/login.html")
		return
	}

	refresh := false

	now := time.Now()
	if now.After(co.Expires) {
		c.Redirect(403, "/static/html/login.html"))
		return
	}

	if now.Sub(co.Expires) < time.Minute {
		co.Expires = now.Add(time.Minute * 30)
		refresh = true
	}

	user := &auth.User{}
	err = jwt.Decode("rsa_single", []byte(co.Value), user)
	if err != nil {
		
	}

	if now.After(user.ValidTo) {
		
		return
	}

	if now.Sub(user.ValidTo) < time.Minute {
		user.ValidFrom = now
		user.ValidTo = now.Add(time.Minute * 30)
		d, _ := jwt.Encode("rsa_single", user)
		co.Value = string(d)
		refresh = true
	}

	c.Next()

	if (refresh) {
		c.SetCookie(co.Name, co.Value, co.Expires)
	}
}
