package utils

import "github.com/zi-wei-dou-shu-gin/utils/ziwei/tiangan"

func GetYinShou(tianGan tiangan.TianGan) tiangan.TianGan {
	return yinShouMap[tianGan]
}
