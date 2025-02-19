package items

import "github.com/gin-gonic/gin"

type Handler interface {
	GetAllItems(c *gin.Context)
	CreateItem(c *gin.Context)
}
