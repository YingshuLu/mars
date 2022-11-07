package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yingshulu/mars/auth"
)

func Register(c *gin.Context) {
	var resp struct {
		Message string      `json:"error"`
		Status  auth.Status `json:"status"`
	}
	var err error
	defer func() {
		code := http.StatusOK
		if err != nil {
			code = http.StatusForbidden
			resp.Message = err.Error()
		}
		c.JSON(code, resp)
	}()

	u := &auth.User{}
	err = c.ShouldBind(u)
	if err != nil {
		resp.Status = auth.InternalError
		return
	}

	err = auth.Add(u)
	if err != nil {
		resp.Status = auth.InternalError
		return
	}
	err = setCookie(u, c)
}
