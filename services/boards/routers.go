package boards

import "github.com/gin-gonic/gin"

func BoardRegister(router *gin.RouterGroup, h Handler) {
	router.GET("/board", h.GetBoard)
}
