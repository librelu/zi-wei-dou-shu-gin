package boards

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/genders"
)

// GetTianBoard get board handler
func (h handler) GetTianBoard(c *gin.Context) {
	req := new(GetBoardRequest)
	if err := validateGetBoardRequest(c, req); err != nil {
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
	b, err := ziwei.NewTianBoard(birthday, gender)
	board, err := b.CreateTianBoard()
	if err != nil {
		handleError(c, err)
		return
	}
	resp := convertTianBoardToGetBoardResponse(board, birthday)
	c.JSON(http.StatusOK, resp)
}

func convertTianBoardToGetBoardResponse(board *ziwei.TianBoard, birthday time.Time) *GetBoardResponse {
	// convert blocks
	blocks := make([]*Block, len(board.Blocks))
	for i, b := range board.Blocks {
		blocks[i] = &Block{
			GongWei: []*GongWei{
				{
					Name: b.GongWeiName,
					Type: TypeTianBoard,
				},
			},
			Location: &Location{
				TianGan: b.Location.TianGan.String(),
				DiZhi:   b.Location.DiZhi.String(),
			},
			TenYearsRound: b.TenYearsRound,
		}
		stars := []*Star{}
		for _, star := range b.Stars {
			stars = append(stars, &Star{
				Name:      star.Name,
				StarType:  star.StarType,
				MiaoXian:  star.MiaoXian,
				FourStar:  star.FourStar,
				BoardType: TypeTianBoard,
			})
		}
		blocks[i].Stars = stars
	}
	month := toChineseNums(int(board.LunaBirthday.Month))
	if board.LunaBirthday.IsLeap {
		month = "閏" + month
	}
	return &GetBoardResponse{
		Blocks:   blocks,
		Birthday: fmt.Sprintf("%d年%d月%d日%d時", birthday.Year(), birthday.Month(), birthday.Day(), birthday.Hour()),
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
			result = result + tenDigitMap[10]
		} else {
			result = result + tenDigitMap[i]
		}
		number = number - (i * 10)
		if number == 0 && len(result) > 0 {
			result = result + numberMap[10]
		}
	}
	return result
}
