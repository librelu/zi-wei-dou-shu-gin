package utils

import (
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/dizhi"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/miaoxian"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/stars"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/startype"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/tiangan"
)

// SetLuCun 設定祿存
func SetLuCun(board *Board, tianGan tiangan.TianGan) *Board {
	luCunMap := []dizhi.DiZhi{
		dizhi.Yin,
		dizhi.Mao,
		dizhi.Si,
		dizhi.Wu,
		dizhi.Si,
		dizhi.Wu,
		dizhi.Shen,
		dizhi.You,
		dizhi.Hai,
		dizhi.Zi,
	}
	luCunLocation := luCunMap[int(tianGan)]
	board.Blocks[luCunLocation].Stars = append(board.Blocks[luCunLocation].Stars, &Star{
		Name:     stars.LuCun.String(),
		StarType: startype.ShiSiFuXing.String(),
		MiaoXian: getLuCunMiaoXian(int(luCunLocation)).String(),
	})
	board.StarsMap[stars.LuCun] = int(luCunLocation)
	return board
}

// getLuCunMiaoXian 得祿存廟陷
func getLuCunMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Wang,
		miaoxian.None,
		miaoxian.Miao,
		miaoxian.Wang,
		miaoxian.None,
		miaoxian.Miao,
		miaoxian.Wang,
		miaoxian.None,
		miaoxian.Miao,
		miaoxian.Wang,
		miaoxian.None,
		miaoxian.Miao,
	}
	return miaoXianMap[index]
}
