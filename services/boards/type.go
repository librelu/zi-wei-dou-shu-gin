package boards

import "github.com/gin-gonic/gin"

type Handler interface {
	GetBoard(c *gin.Context)
	CreateBoard(c *gin.Context)
	UpdateBoard(c *gin.Context)
}
