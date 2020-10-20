package ziwei

import (
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/dizhi"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/genders"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/mingju"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/stars"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/tiangan"
)

type Block struct {
	GongWeiName string
	Stars       []*Star
	Location    *Location
}

type Star struct {
	Name     string
	StarType string
	MiaoXian string
}

type Location struct {
	TianGan tiangan.TianGan
	DiZhi   dizhi.DiZhi
}

const (
	defaultBoardBlock = 12
)

var yinShouMap = map[tiangan.TianGan]tiangan.TianGan{
	tiangan.Jia:  tiangan.Bing,
	tiangan.Ji:   tiangan.Bing,
	tiangan.Yi:   tiangan.Wu,
	tiangan.Geng: tiangan.Wu,
	tiangan.Bing: tiangan.Geng,
	tiangan.Xin:  tiangan.Geng,
	tiangan.Ding: tiangan.Ren,
	tiangan.Ren:  tiangan.Ren,
	tiangan.Wu:   tiangan.Jia,
	tiangan.Gui:  tiangan.Jia,
}

type MingJu struct {
	JuType mingju.MingJuType
	Number uint
}

type Board struct {
	Blocks   []*Block
	StarsMap map[stars.StarName]int
	Gender   genders.Gender
	MingJu   *MingJu
	ShenZhu  string
	MingZhu  string
}
