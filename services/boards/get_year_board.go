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
	yearBoard, err := ziwei.NewYearsBoard(birthday, gender, index)
	if err != nil {
		handleError(c, err)
		return
	}

	resp := convertToYearBoard(yearBoard)

	c.JSON(http.StatusOK, resp)
}

func convertToYearBoard(yearBoard *ziwei.YearBoard) *GetYearBoardResponse {
	month := toChineseNums(int(yearBoard.LunaBirthday.Month))
	if yearBoard.LunaBirthday.IsLeap {
		month = "閏" + month
	}
	board := &GetYearBoardResponse{
		Birthday: fmt.Sprintf("%d年%d月%d日%d時", yearBoard.Birthday.Year(), yearBoard.Birthday.Month(), yearBoard.Birthday.Day(), yearBoard.Birthday.Hour()),
		LunaBirthDay: fmt.Sprintf("%s%s年%s月%s日%s時",
			yearBoard.LunaBirthday.Year.TianGan.String(),
			yearBoard.LunaBirthday.Year.DiZhi.String(),
			month,
			toChineseNums(int(yearBoard.LunaBirthday.Day)),
			yearBoard.LunaBirthday.Hour,
		),
		Blocks:           make([]*Block, defaultBoardBlock),
		Gender:           yearBoard.Gender.String(),
		MingJu:           yearBoard.MingJu.JuType.String(),
		MingJuValue:      int(yearBoard.MingJu.Number),
		ShenZhu:          yearBoard.ShenZhu,
		MingZhu:          yearBoard.MingZhu,
		ShenGongLocation: yearBoard.ShenGongLocation,
	}
	for i, yearBlock := range yearBoard.Blocks {
		if board.Blocks[i] == nil {
			board.Blocks[i] = new(Block)
		}
		if len(board.Blocks[i].GongWei) == 0 {
			board.Blocks[i].GongWei = make([]*GongWei, 0)
		}
		board.Blocks[i].Location = &Location{
			TianGan: yearBlock.Location.TianGan.String(),
			DiZhi:   yearBlock.Location.DiZhi.String(),
		}
		board.Blocks[i].GongWei = append(board.Blocks[i].GongWei, &GongWei{
			Name: yearBlock.GongWeiName,
			Type: TypeYearBoard,
		})
		board.Blocks[i].TenYearsRound = yearBlock.TenYearsRound
		for _, star := range yearBlock.Stars {
			if board.Blocks[i] == nil {
				board.Blocks[i] = new(Block)
			}
			if len(board.Blocks[i].Stars) == 0 {
				board.Blocks[i].Stars = []*Star{}
			}
			board.Blocks[i].Stars = append(board.Blocks[i].Stars, &Star{
				Name:      star.Name,
				StarType:  star.StarType,
				MiaoXian:  star.MiaoXian,
				FourStar:  star.FourStar,
				BoardType: TypeYearBoard,
			})
		}
	}
	return board
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
