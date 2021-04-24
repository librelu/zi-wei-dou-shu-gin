package boards

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/genders"
)

// GetMonthBoard get month board handler
func (h handler) GetDateBoard(c *gin.Context) {
	req := new(GetDateBoardRequest)
	if err := validateGetDateBoardRequest(c, req); err != nil {
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
	targetDate := time.Date(req.TargetYear, time.Month(req.TargetMonth), req.TargetDate, 0, 0, 0, 0, location)
	dateBoard, err := ziwei.NewDateBoard(birthday, targetDate, gender)
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
	resp := mergeBoards(tianBoard, dateBoard)

	c.JSON(http.StatusOK, resp)
}

func mergeBoards(tianBoard *ziwei.TianBoard, dateBoard *ziwei.DateBoard) *GetDateBoardResponse {
	month := toChineseNums(int(dateBoard.LunaBirthday.Month))
	if dateBoard.LunaBirthday.IsLeap {
		month = "閏" + month
	}
	board := &GetDateBoardResponse{
		Blocks: make([]*Block, defaultBoardBlock),
	}
	for i, dateBlock := range dateBoard.Blocks {
		if board.Blocks[i] == nil {
			board.Blocks[i] = new(Block)
		}
		if len(board.Blocks[i].GongWei) == 0 {
			board.Blocks[i].GongWei = make([]*GongWei, 0)
		}
		board.Blocks[i].Location = &Location{
			TianGan: dateBlock.Location.TianGan.String(),
			DiZhi:   dateBlock.Location.DiZhi.String(),
		}
		board.Blocks[i].GongWei = append(board.Blocks[i].GongWei, &GongWei{
			Name: dateBlock.GongWeiName,
			Type: TypeDateBoard,
		})
		board.Blocks[i].TenYearsRound = dateBlock.TenYearsRound
		for _, star := range dateBlock.Stars {
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
				BoardType: TypeDateBoard,
			})
		}
	}
	return board
}

func validateGetDateBoardRequest(c *gin.Context, req *GetDateBoardRequest) error {
	if err := c.ShouldBindQuery(req); err != nil {
		return err
	}
	if !(req.Gender == 0 || req.Gender == 1) {
		return fmt.Errorf("illegal gender, example gender: 0=>male, 1=>female, current: %d", req.Gender)
	}
	return nil
}
