package items

import (
	"net/http"
	"sinarmas-test/internal/domain/api"
	"sinarmas-test/internal/domain/items"

	"github.com/gin-gonic/gin"
)

func (h *Items) CreateItem(c *gin.Context) {
	var (
		err error
		req items.CreateReq
		ctx = c.Request.Context()
	)

	// userPass := c.GetHeader("Authentication")

	err = c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, api.Response{
			Status: "failed",
			Msg:    "invalid payload",
		})
		return
	}

	err = h.itemSvc.CreateItem(ctx, req)
	if err != nil {
		c.JSON(http.StatusOK, api.Response{
			Status: "failed",
			Msg:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.Response{
		Status: "ok",
	})
}

// func (h *Items) CheckAuth(userPass string) bool {

// 	userPassSplit := strings.Split(userPass, ":")

// 	h.authSvc.Login(ctx, auth.LoginReq{Username: userPassSplit[0], Password: userPassSplit[1]})
// }
