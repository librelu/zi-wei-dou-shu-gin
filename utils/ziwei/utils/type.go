package utils

import (
	"time"

	"github.com/zi-wei-dou-shu-gin/utils/lunacal"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/dizhi"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/genders"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/mingju"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/stars"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/tiangan"
)

const (
	defaultBoardBlock = 12
)

type Location struct {
	TianGan tiangan.TianGan
	DiZhi   dizhi.DiZhi
}

type Block struct {
	GongWeiName   string
	Stars         []*Star
	Location      *Location
	TenYearsRound string
}

type Star struct {
	Name     string
	StarType string
	MiaoXian string
	FourStar string
}

type MingJu struct {
	JuType mingju.MingJuType
	Number uint
}

type Board struct {
	Birthday            time.Time
	LunaBirthday        *lunacal.LunaDate
	Blocks              []*Block
	StarsMap            map[stars.StarName]int
	Gender              genders.Gender
	MingJu              *MingJu
	ShenZhu             string
	MingZhu             string
	LocationTainGan     string
	ShenGongLocation    int
	MingGongLocation    int
	MainStarConnections []int
}

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
