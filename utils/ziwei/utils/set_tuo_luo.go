package utils

import (
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/dizhi"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/miaoxian"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/stars"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/startype"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/tiangan"
)

// SetTuoLuo 設定陀羅
func SetTuoLuo(board *Board, tianGan tiangan.TianGan) *Board {
	tuoLuoMap := []dizhi.DiZhi{
		dizhi.Chou,
		dizhi.Yin,
		dizhi.Chen,
		dizhi.Si,
		dizhi.Chen,
		dizhi.Si,
		dizhi.Wei,
		dizhi.Shen,
		dizhi.Xu,
		dizhi.Hai,
	}
	tuoLuoLocation := tuoLuoMap[int(tianGan)]
	board.Blocks[tuoLuoLocation].Stars = append(board.Blocks[tuoLuoLocation].Stars, &Star{
		Name:     stars.TuoLuo.String(),
		StarType: startype.RightFuXing.String(),
		MiaoXian: getTuoLuoMiaoXian(int(tuoLuoLocation)).String(),
	})
	board.StarsMap[stars.TuoLuo] = int(tuoLuoLocation)
	return board
}

// getTuoLuoMiaoXian 得陀羅廟陷
func getTuoLuoMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.None,
		miaoxian.Miao,
		miaoxian.Xian,
		miaoxian.None,
		miaoxian.Miao,
		miaoxian.Xian,
		miaoxian.None,
		miaoxian.Miao,
		miaoxian.Xian,
		miaoxian.None,
		miaoxian.Miao,
		miaoxian.Xian,
	}
	return miaoXianMap[index]
}
