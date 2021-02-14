package boards

import (
	"fmt"
	"net/http"
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
	timezone := 0
	if req.TimeZone != 0 {
		timezone = req.TimeZone / 60
	}
	location := time.FixedZone(fmt.Sprintf("UTC %d", timezone), req.TimeZone)
	birthday := time.Date(req.BirthYear, time.Month(req.BirthMonth), req.BirthDate, req.BirthHour, 0, 0, 0, location)
	gender := genders.Gender(req.Gender)
	index := req.Index
	board, err := ziwei.NewTenYearsBoard(birthday, gender, index)
	if err != nil {
		handleError(c, err)
		return
	}

	resp := convertBoardToGetYearBoardResponse(board, birthday)
	c.JSON(http.StatusOK, resp)
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

func convertBoardToGetYearBoardResponse(board *ziwei.YearBoard, birthday time.Time) *GetYearBoardResponse {
	getBoardResp := convertYearBoardToGetBoardResponse(board, birthday)
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

func convertYearBoardToGetBoardResponse(board *ziwei.YearBoard, birthday time.Time) *GetBoardResponse {
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
	month := toChineseNums(int(board.LunaBirthday.Month))
	if board.LunaBirthday.IsLeap {
		month = "閏" + month
	}
	return &GetBoardResponse{
		Blocks:   blocks,
		BirthDay: fmt.Sprintf("%d年%d月%d日%d時", birthday.Year(), birthday.Month(), birthday.Day(), birthday.Hour()),
		LunaBirthDay: fmt.Sprintf("%s%s年%s月%s日%s時",
			board.LunaBirthday.Year.TianGan.String(),
			board.LunaBirthday.Year.DiZhi.String(),
			month,
			toChineseNums(int(board.LunaBirthday.Day)),
			board.LunaBirthday.Hour,
		),
		Gender:              board.Gender.String(),
		MingJu:              board.MingJu.JuType.String(),
		MingJuValue:         int(board.MingJu.Number),
		ShenZhu:             board.ShenZhu,
		MingZhu:             board.MingZhu,
		ShenGongLocation:    board.ShenGongLocation,
		MingGongLocation:    board.MingGongLocation,
		MainStarConnections: board.MainStarConnections,
	}
}
