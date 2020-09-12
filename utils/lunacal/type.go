package lunacal

import (
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/dizhi"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/tiangan"
)

type TianGanDiZhi struct {
	TianGan tiangan.TianGan
	DiZhi   dizhi.DiZhi
}

type LunaDate struct {
	Year   *TianGanDiZhi
	Month  uint
	Day    uint
	Hour   *dizhi.DiZhi
	IsLeap bool
}
