package auth

import (
	"net/http"
	"sinarmas-test/internal/domain/api"
	"sinarmas-test/internal/domain/auth"

	"github.com/gin-gonic/gin"
)

func (h *Auth) Login(c *gin.Context) {
	var (
		err error
		req auth.LoginReq
		ctx = c.Request.Context()
	)

	err = c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, api.Response{
			Status: "failed",
			Msg:    "invalid payload",
		})
		return
	}

	jwt, err := h.authSvc.Login(ctx, req)
	if err != nil {
		c.JSON(http.StatusOK, api.Response{
			Status: "failed",
			Msg:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.Response{
		Status: "ok",
		Data:   jwt,
	})
}
