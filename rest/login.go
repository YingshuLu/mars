package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yingshulu/mars/auth"
)

func Login(c *gin.Context) {
	var resp struct {
		Status  auth.Status `json:"status"`
		Message string      `json:"message"`
	}
	var err error
	defer func() {
		var code = http.StatusOK
		if err != nil {
			code = http.StatusUnauthorized
			resp.Message = err.Error()
		}
		c.JSON(code, resp)
	}()

	u := &auth.User{}
	err = c.ShouldBind(u)
	if err != nil {
		return
	}

	resp.Status = auth.Validate(u)
	if resp.Status != auth.Ok {
		err = fmt.Errorf("user profile can not validate")
		return
	}

}
