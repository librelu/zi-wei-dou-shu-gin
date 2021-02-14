package boards

import (
	"github.com/gin-gonic/gin"
	"github.com/zi-wei-dou-shu-gin/db/dao"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/utils"
)

type handler struct {
	dao dao.DaoHandler
}

type Handler interface {
	GetTianBoard(c *gin.Context)
	GetYearBoard(c *gin.Context)
}

type GetBoardRequest struct {
	BirthYear  int `form:"birthYear" binding:"required"`
	BirthMonth int `form:"birthMonth" binding:"required"`
	BirthDate  int `form:"birthDate" binding:"required"`
	BirthHour  int `form:"birthHour"`
	TimeZone   int `form:"timezone"`
	Gender     int `form:"gender"`
}

type GetBoardResponse struct {
	Blocks              []*Block
	BirthDay            string
	LunaBirthDay        string
	Gender              string
	MingJu              string
	MingJuValue         int
	ShenZhu             string
	MingZhu             string
	ShenGongLocation    int
	MingGongLocation    int
	MainStarConnections []int
}

type GetYearBoardRequest struct {
	BirthYear  int `form:"birthYear" binding:"required"`
	BirthMonth int `form:"birthMonth" binding:"required"`
	BirthDate  int `form:"birthDate" binding:"required"`
	BirthHour  int `form:"birthHour"`
	TimeZone   int `form:"timezone"`
	Gender     int `form:"gender"`
	Index      int `form:"index"`
}

type GetYearBoardResponse struct {
	Blocks           []*Block
	BirthDay         string
	LunaBirthDay     string
	Gender           string
	MingJu           string
	MingJuValue      int
	ShenZhu          string
	MingZhu          string
	ShenGongLocation int
}

type SaveBoardRequest struct {
	Birthday int64  `form:"birthday" binding:"required"`
	Gender   int    `form:"gender"`
	Name     string `form:"name" binding:"required"`
}

type Block struct {
	GongWeiName   string
	Stars         []*utils.Star
	Location      *Location
	TenYearsRound string
}

type Location struct {
	TianGan string
	DiZhi   string
}

var numberMap = []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}
var tenDigitMap = []string{"零", "一", "廿", "三", "四", "五", "六", "七", "八", "九", "十"}
