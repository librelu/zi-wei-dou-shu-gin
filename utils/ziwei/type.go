package ziwei

import (
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/dizhi"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/gong"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/miaoxian"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/mingju"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/tiangan"
)

type Block struct {
	GongWei     tiangan.TianGan
	GongWeiName gong.Gong
	XiaoXian    bool
	DaXian      string
	Stars       []*Star
	Location    *Location
}

type Star struct {
	Name     string
	StarType string
	Location tiangan.TianGan
	MiaoXian *miaoxian.MiaoXian
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
	Blocks []*Block
	MingJu *MingJu
}