package boards

import "github.com/gin-gonic/gin"

func BoardRegister(router *gin.RouterGroup, h Handler) {
	router.GET("/board", h.GetTianBoard)
	router.GET("/year-board", h.GetYearBoard)
	router.GET("/ten-board", h.GetTenYearsBoard)
	router.GET("/month-board", h.GetMonthBoard)
}
