package boards

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/genders"
)

// GetYearBoard get board handler
func (h handler) SaveBoard(c *gin.Context) {
	req := new(SaveBoardRequest)
	if err := validateSaveBoardRequest(c, req); err != nil {
		handleError(c, err)
		return
	}
	gender := genders.Gender(req.Gender)
	h.dao.SaveBoard(gender.String(), req.Name, req.Birthday)
	c.JSON(http.StatusOK, map[string]string{})
}

func validateSaveBoardRequest(c *gin.Context, req *SaveBoardRequest) error {
	if err := c.ShouldBindQuery(req); err != nil {
		return err
	}
	if !(req.Gender == 0 || req.Gender == 1) {
		return fmt.Errorf("illegal gender, example gender: 0=>male, 1=>female, current: %d", req.Gender)
	}
	return nil
}
