package boards

import (
	"github.com/gin-gonic/gin"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei"
)

type handler struct {
}

type Handler interface {
	GetBoard(c *gin.Context)
}

type GetBoardRequest struct {
	Birthday int64 `form:"birthday" binding:"required"`
	Gender   int   `form:"gender"`
}

type GetBoardResponse struct {
	Blocks           []*ziwei.Block
	BirthDay         string
	LunaBirthDay     string
	Gender           string
	MingJu           string
	MingJuValue      int
	ShenZhu          string
	MingZhu          string
	ShenGongLocation int
}

var numberMap = []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}
