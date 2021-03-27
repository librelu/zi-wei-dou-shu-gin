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

	tianBoard, err := ziwei.NewTianBoard(birthday, gender)
	if err != nil {
		handleError(c, err)
		return
	}
	if tianBoard, err = tianBoard.CreateTianBoard(); err != nil {
		handleError(c, err)
		return
	}
	resp := mergeTianBoardAndTenBoard(tianBoard, tenBoard)

	c.JSON(http.StatusOK, resp)
}

func mergeTianBoardAndTenBoard(tianBoard *ziwei.TianBoard, tenYearBoard *ziwei.TenYearBoard) *GetYearBoardResponse {
	month := toChineseNums(int(tenYearBoard.LunaBirthday.Month))
	if tenYearBoard.LunaBirthday.IsLeap {
		month = "閏" + month
	}
	board := &GetYearBoardResponse{
		Birthday: fmt.Sprintf("%d年%d月%d日%d時", tianBoard.Birthday.Year(), tianBoard.Birthday.Month(), tianBoard.Birthday.Day(), tianBoard.Birthday.Hour()),
		LunaBirthDay: fmt.Sprintf("%s%s年%s月%s日%s時",
			tianBoard.LunaBirthday.Year.TianGan.String(),
			tianBoard.LunaBirthday.Year.DiZhi.String(),
			month,
			toChineseNums(int(tianBoard.LunaBirthday.Day)),
			tianBoard.LunaBirthday.Hour,
		),
		Blocks:           make([]*Block, defaultBoardBlock),
		Gender:           tianBoard.Gender.String(),
		MingJu:           tianBoard.MingJu.JuType.String(),
		MingJuValue:      int(tianBoard.MingJu.Number),
		ShenZhu:          tianBoard.ShenZhu,
		MingZhu:          tianBoard.MingZhu,
		ShenGongLocation: tianBoard.ShenGongLocation,
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
		for i, star := range yearBlock.Stars {
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
	for i, tianBlock := range tianBoard.Blocks {
		if board.Blocks[i] == nil {
			board.Blocks[i] = new(Block)
		}
		if len(board.Blocks[i].GongWei) == 0 {
			board.Blocks[i].GongWei = make([]*GongWei, 0)
		}
		board.Blocks[i].GongWei = append(board.Blocks[i].GongWei, &GongWei{
			Name: tianBlock.GongWeiName,
			Type: TypeTianBoard,
		})
		for _, star := range tianBlock.Stars {
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
				BoardType: TypeTianBoard,
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
