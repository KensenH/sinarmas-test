package items

import (
	"net/http"
	"sinarmas-test/internal/domain/api"

	"github.com/gin-gonic/gin"
)

func (h *Items) GetAllItems(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	items, err := h.itemSvc.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusOK, api.Response{
			Status: "failed",
			Msg:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.Response{
		Status: "ok",
		Data:   items,
	})
}
