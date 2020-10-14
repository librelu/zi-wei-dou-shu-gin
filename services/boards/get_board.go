package boards

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/genders"
)

// GetBoard get board handler
func (h handler) GetBoard(c *gin.Context) {
	req := new(GetBoardRequest)
	if err := validate(c, req); err != nil {
		handleError(c, err)
		return
	}
	birthday := time.Unix(req.Birthday, 0)
	gender := genders.Gender(req.Gender)
	board, err := ziwei.NewBoard(birthday, gender)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(200, board)
}

func validate(c *gin.Context, req *GetBoardRequest) error {
	if err := c.ShouldBindQuery(req); err != nil {
		return err
	}
	if !(req.Gender == 0 || req.Gender == 1) {
		return fmt.Errorf("illegal gender, example gender: 0=>male, 1=>female, current: %d", req.Gender)
	}
	return nil
}

func handleError(c *gin.Context, err error) {
	c.String(http.StatusBadRequest, err.Error())
	c.Abort()
}
