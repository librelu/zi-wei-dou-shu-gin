package boards

import "github.com/gin-gonic/gin"

func BoardRegiester(router *gin.RouterGroup, h Handler) {
	router.GET("/board", h.GetBoard)
	router.POST("/board", h.CreateBoard)
	router.PUT("/board", h.UpdateBoard)
}
