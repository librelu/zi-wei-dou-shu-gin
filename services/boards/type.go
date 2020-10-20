package boards

import (
	"github.com/gin-gonic/gin"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei"
)

type handler struct {
}

type Handler interface {
	GetBoard(c *gin.Context)
}

type GetBoardRequest struct {
	Birthday int64 `form:"birthday" binding:"required"`
	Gender   int   `form:"gender"`
}

type GetBoardResponse struct {
	Blocks       []*ziwei.Block
	BirthDay     string
	LunaBirthDay string
	Gender       string
	MingJu       string
	MingJuValue  int
	ShenZhu      string
	MingZhu      string
}
