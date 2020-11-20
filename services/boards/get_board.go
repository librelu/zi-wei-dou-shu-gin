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
	if err := validateGetBoardRequest(c, req); err != nil {
		handleError(c, err)
		return
	}
	birthday := time.Unix(req.Birthday, 0)
	gender := genders.Gender(req.Gender)
	board, err := ziwei.NewBoard(birthday, gender).CreateTianBoard()
	if err != nil {
		handleError(c, err)
		return
	}
	resp := convertBoardToGetBoardResponse(board, birthday)
	c.JSON(200, resp)
}

func convertBoardToGetBoardResponse(board *ziwei.Board, birthday time.Time) *GetBoardResponse {
	// convert blocks
	blocks := make([]*Block, len(board.Blocks))
	for i, b := range board.Blocks {
		blocks[i] = &Block{
			GongWeiName: b.GongWeiName,
			Stars:       b.Stars,
			Location: &Location{
				TianGan: b.Location.TianGan.String(),
				DiZhi:   b.Location.DiZhi.String(),
			},
			TenYearsRound: b.TenYearsRound,
		}
	}
	return &GetBoardResponse{
		Blocks:   blocks,
		BirthDay: fmt.Sprintf("%d年%d月%d日%d時", birthday.Year(), birthday.Month(), birthday.Day(), birthday.Hour()),
		LunaBirthDay: fmt.Sprintf("%s%s年%s月%s日%s時",
			board.LunaBirthday.Year.TianGan.String(),
			board.LunaBirthday.Year.DiZhi.String(),
			toChineseNums(int(board.LunaBirthday.Month)),
			toChineseNums(int(board.LunaBirthday.Day)),
			board.LunaBirthday.Hour,
		),
		Gender:           board.Gender.String(),
		MingJu:           board.MingJu.JuType.String(),
		MingJuValue:      int(board.MingJu.Number),
		ShenZhu:          board.ShenZhu,
		MingZhu:          board.MingZhu,
		ShenGongLocation: board.ShenGongLocation,
	}
}

func validateGetBoardRequest(c *gin.Context, req *GetBoardRequest) error {
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

func toChineseNums(number int) string {
	result := ""
	for number > 0 {
		i := number / 10
		if number <= 10 {
			return result + numberMap[number]
		}
		if i == 1 {
			result = result + numberMap[10]
		} else {
			result = result + numberMap[i]
		}
		number = number - (i * 10)
		if number == 0 && len(result) > 0 {
			result = result + numberMap[10]
		}
	}
	return result
}
