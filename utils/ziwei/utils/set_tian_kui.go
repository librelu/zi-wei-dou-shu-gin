package utils

import (
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/dizhi"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/miaoxian"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/stars"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/startype"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/tiangan"
)

// SetTianKui 設定天𣁽
func SetTianKui(board *Board, tianGan tiangan.TianGan) *Board {
	tianKuiMap := []dizhi.DiZhi{
		dizhi.Chou,
		dizhi.Zi,
		dizhi.Hai,
		dizhi.Hai,
		dizhi.Chou,
		dizhi.Zi,
		dizhi.Yin,
		dizhi.Yin,
		dizhi.Mao,
		dizhi.Mao,
	}
	tianKuiLocation := tianKuiMap[int(tianGan)]
	board.Blocks[tianKuiLocation].Stars = append(board.Blocks[tianKuiLocation].Stars, &Star{
		Name:     stars.TianKui.String(),
		StarType: startype.LeftFuXing.String(),
		MiaoXian: getTianKuiMiaoXian(int(tianKuiLocation)).String(),
	})
	board.StarsMap[stars.TianKui] = int(tianKuiLocation)
	return board
}

// getTianKuiMiaoXian 得天𣁽廟陷
func getTianKuiMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Wang,
		miaoxian.Wang,
		miaoxian.None,
		miaoxian.Miao,
		miaoxian.None,
		miaoxian.None,
		miaoxian.Miao,
		miaoxian.None,
		miaoxian.None,
		miaoxian.None,
		miaoxian.None,
		miaoxian.Wang,
	}
	return miaoXianMap[index]
}
