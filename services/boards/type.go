package boards

import "github.com/gin-gonic/gin"

type handler struct {
}

type Handler interface {
	GetBoard(c *gin.Context)
}

type GetBoardRequest struct {
	Birthday int64 `form:"birthday" binding:"required"`
	Gender   int   `form:"gender"`
}
