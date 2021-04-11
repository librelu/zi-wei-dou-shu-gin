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
	req := new(GetTenBoardRequest)
	if err := validateGetTenBoardRequest(c, req); err != nil {
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
	tenBoard, err := ziwei.NewTenYearsBoard(birthday, gender, index)
	if err != nil {
		handleError(c, err)
		return
	}

	resp := convertToTenYearBoardResponse(tenBoard)

	c.JSON(http.StatusOK, resp)
}

func convertToTenYearBoardResponse(tenYearBoard *ziwei.TenYearBoard) *GetYearBoardResponse {
	month := toChineseNums(int(tenYearBoard.LunaBirthday.Month))
	if tenYearBoard.LunaBirthday.IsLeap {
		month = "閏" + month
	}
	board := &GetYearBoardResponse{
		Blocks: make([]*Block, defaultBoardBlock),
	}
	for i, yearBlock := range tenYearBoard.Blocks {
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
			Type: TypeTenBoard,
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
				BoardType: TypeTenBoard,
			})
		}
	}
	return board
}

func validateGetTenBoardRequest(c *gin.Context, req *GetTenBoardRequest) error {
	if err := c.ShouldBindQuery(req); err != nil {
		return err
	}
	if !(req.Gender == 0 || req.Gender == 1) {
		return fmt.Errorf("illegal gender, example gender: 0=>male, 1=>female, current: %d", req.Gender)
	}
	return nil
}
