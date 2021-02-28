package boards

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/genders"
)

// GetTenYearBoard get ten years board handler
func (h handler) GetTenYearsBoard(c *gin.Context) {
	req := new(GetYearBoardRequest)
	if err := validateGetYearBoardRequest(c, req); err != nil {
		handleError(c, err)
		return
	}
	timezone := 0
	if req.TimeZone != 0 {
		timezone = req.TimeZone / 60
	}
	location := time.FixedZone(fmt.Sprintf("UTC %d", timezone), req.TimeZone)
	birthday := time.Date(req.BirthYear, time.Month(req.BirthMonth), req.BirthDate, req.BirthHour, 0, 0, 0, location)
	gender := genders.Gender(req.Gender)
	index := req.Index
	yearBoard, err := ziwei.NewYearsBoard(birthday, gender, index)
	if err != nil {
		handleError(c, err)
		return
	}

	tianBoard, err := ziwei.NewTianBoard(birthday, gender)
	if err != nil {
		handleError(c, err)
		return
	}
	if tianBoard, err = tianBoard.CreateTianBoard(); err != nil {
		handleError(c, err)
		return
	}
	resp := mergeTianBoardAndYearBoard(tianBoard, yearBoard)

	c.JSON(http.StatusOK, resp)
}
