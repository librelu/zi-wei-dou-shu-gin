package boards

import (
	"github.com/gin-gonic/gin"
	"github.com/zi-wei-dou-shu-gin/db/dao"
)

type handler struct {
	dao dao.DaoHandler
}

type Handler interface {
	GetTianBoard(c *gin.Context)
	GetYearBoard(c *gin.Context)
	GetTenYearsBoard(c *gin.Context)
	GetMonthBoard(c *gin.Context)
	GetDateBoard(c *gin.Context)
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
	Blocks              []*Block `json:"blocks"`
	Birthday            string   `json:"birthday"`
	LunaBirthDay        string   `json:"luna_birthday"`
	Gender              string   `json:"gender"`
	MingJu              string   `json:"ming_ju"`
	MingJuValue         int      `json:"ming_ju_value"`
	ShenZhu             string   `json:"shen_zhu"`
	MingZhu             string   `json:"ming_zhu"`
	ShenGongLocation    int      `json:"shen_gong_location"`
	MingGongLocation    int      `json:"ming_gong_location"`
	MainStarConnections []int    `json:"main_star_connection"`
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
	Blocks           []*Block `json:"blocks"`
	Birthday         string   `json:"birthday"`
	LunaBirthDay     string   `json:"luna_birthday"`
	Gender           string   `json:"gender"`
	MingJu           string   `json:"ming_ju"`
	MingJuValue      int      `json:"ming_ju_value"`
	ShenZhu          string   `json:"shen_zhu"`
	MingZhu          string   `json:"ming_zhu"`
	ShenGongLocation int      `json:"shen_gong_location"`
}

type GetTenBoardRequest struct {
	BirthYear  int `form:"birthYear" binding:"required"`
	BirthMonth int `form:"birthMonth" binding:"required"`
	BirthDate  int `form:"birthDate" binding:"required"`
	BirthHour  int `form:"birthHour"`
	TimeZone   int `form:"timezone"`
	Gender     int `form:"gender"`
	Index      int `form:"index"`
}

type GetTenBoardResponse struct {
	Blocks           []*Block `json:"blocks"`
	Birthday         string   `json:"birthday"`
	LunaBirthDay     string   `json:"luna_birthday"`
	Gender           string   `json:"gender"`
	MingJu           string   `json:"ming_ju"`
	MingJuValue      int      `json:"ming_ju_value"`
	ShenZhu          string   `json:"shen_zhu"`
	MingZhu          string   `json:"ming_zhu"`
	ShenGongLocation int      `json:"shen_gong_location"`
}

type SaveBoardRequest struct {
	Birthday int64  `form:"birthday" binding:"required"`
	Gender   int    `form:"gender"`
	Name     string `form:"name" binding:"required"`
}

type Block struct {
	GongWei       []*GongWei `json:"gong_wei"`
	Stars         []*Star    `json:"stars"`
	Location      *Location  `json:"location"`
	TenYearsRound string     `json:"ten_years_round"`
}

type Star struct {
	Name      string `json:"name"`
	StarType  string `json:"star_type"`
	MiaoXian  string `json:"miao_xian"`
	FourStar  string `json:"four_star"`
	BoardType string `json:"board_type"`
}

type GongWei struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Location struct {
	TianGan string `json:"tian_gan"`
	DiZhi   string `json:"dizhi"`
}

var numberMap = []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}
var tenDigitMap = []string{"零", "一", "廿", "三", "四", "五", "六", "七", "八", "九", "十"}

const (
	TypeTenBoard      = "ten_years_board"
	TypeYearBoard     = "year_board"
	TypeTianBoard     = "tian_board"
	TypeMonthBoard    = "month_board"
	TypeDateBoard     = "date_board"
	defaultBoardBlock = 12
)

type GetMonthBoardRequest struct {
	BirthYear  int `form:"birthYear" binding:"required"`
	BirthMonth int `form:"birthMonth" binding:"required"`
	BirthDate  int `form:"birthDate" binding:"required"`
	BirthHour  int `form:"birthHour"`
	TimeZone   int `form:"timezone"`
	Gender     int `form:"gender"`
	Index      int `form:"index"`
}

type GetMonthBoardResponse struct {
	Blocks           []*Block `json:"blocks"`
	Birthday         string   `json:"birthday"`
	LunaBirthDay     string   `json:"luna_birthday"`
	Gender           string   `json:"gender"`
	MingJu           string   `json:"ming_ju"`
	MingJuValue      int      `json:"ming_ju_value"`
	ShenZhu          string   `json:"shen_zhu"`
	MingZhu          string   `json:"ming_zhu"`
	ShenGongLocation int      `json:"shen_gong_location"`
}

type GetDateBoardRequest struct {
	BirthYear  int `form:"birthYear" binding:"required"`
	BirthMonth int `form:"birthMonth" binding:"required"`
	BirthDate  int `form:"birthDate" binding:"required"`
	BirthHour  int `form:"birthHour"`
	TimeZone   int `form:"timezone"`
	Gender     int `form:"gender"`
	Index      int `form:"index"`
}

type GetDateBoardResponse struct {
	Blocks           []*Block `json:"blocks"`
	Birthday         string   `json:"birthday"`
	LunaBirthDay     string   `json:"luna_birthday"`
	Gender           string   `json:"gender"`
	MingJu           string   `json:"ming_ju"`
	MingJuValue      int      `json:"ming_ju_value"`
	ShenZhu          string   `json:"shen_zhu"`
	MingZhu          string   `json:"ming_zhu"`
	ShenGongLocation int      `json:"shen_gong_location"`
}
