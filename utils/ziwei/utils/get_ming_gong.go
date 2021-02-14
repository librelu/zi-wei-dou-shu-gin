package utils

import (
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/dizhi"
)

func GetMingGong(hour *dizhi.DiZhi, month uint) *dizhi.DiZhi {
	hourIndex := int(*hour) - 2
	mingGong := int(month-1) - int(hourIndex)
	if mingGong < 0 {
		mingGong += 12
	} else if mingGong > 11 {
		mingGong = mingGong - 12
	}
	mingGongLocation := dizhi.DiZhi(mingGong)
	return &mingGongLocation
}
