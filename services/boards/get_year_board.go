package boards

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/genders"
)

// GetYearBoard get board handler
func (h handler) GetYearBoard(c *gin.Context) {
	req := new(GetYearBoardRequest)
	if err := validateGetYearBoardRequest(c, req); err != nil {
		handleError(c, err)
		return
	}
	birthday := time.Unix(req.Birthday, 0)
	gender := genders.Gender(req.Gender)
	index := req.Index
	board, err := ziwei.NewTenYearsBoard(birthday, gender, index)
	if err != nil {
		handleError(c, err)
		return
	}
	resp := convertBoardToGetYearBoardResponse(board.Board, birthday)
	c.JSON(200, resp)
}

func validateGetYearBoardRequest(c *gin.Context, req *GetYearBoardRequest) error {
	if err := c.ShouldBindQuery(req); err != nil {
		return err
	}
	if !(req.Gender == 0 || req.Gender == 1) {
		return fmt.Errorf("illegal gender, example gender: 0=>male, 1=>female, current: %d", req.Gender)
	}
	return nil
}

func convertBoardToGetYearBoardResponse(board *ziwei.Board, birthday time.Time) *GetYearBoardResponse {
	getBoardResp := convertBoardToGetBoardResponse(board, birthday)
	return &GetYearBoardResponse{
		Blocks:           getBoardResp.Blocks,
		BirthDay:         getBoardResp.BirthDay,
		LunaBirthDay:     getBoardResp.LunaBirthDay,
		Gender:           getBoardResp.Gender,
		MingJu:           getBoardResp.MingJu,
		MingJuValue:      getBoardResp.MingJuValue,
		ShenZhu:          getBoardResp.ShenZhu,
		MingZhu:          getBoardResp.MingZhu,
		ShenGongLocation: getBoardResp.ShenGongLocation,
	}
}
