package utils

import (
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/dizhi"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/miaoxian"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/stars"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/startype"
)

// SetLing　設定鈴星
func SetLing(board *Board, birthYear *dizhi.DiZhi, birthHour *dizhi.DiZhi) *Board {
	xaxis := getHuoLingGroupMap(birthYear)
	var index int
	switch xaxis {
	case 0:
		index = (int(*birthHour) + 3) % 12
	default:
		index = (int(*birthHour) + 10) % 12
	}
	board.Blocks[index].Stars = append(board.Blocks[index].Stars, &Star{
		Name:     stars.Ling.String(),
		StarType: startype.RightFuXing.String(),
		MiaoXian: getLingMiaoXian(index).String(),
	})
	board.StarsMap[stars.Ling] = index
	return board
}

// getLingMiaoXian 得鈴星廟陷
func getLingMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Xian,
		miaoxian.Xian,
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Wang,
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Wang,
		miaoxian.Wang,
		miaoxian.Xian,
		miaoxian.Miao,
		miaoxian.Miao,
	}
	return miaoXianMap[index]
}
