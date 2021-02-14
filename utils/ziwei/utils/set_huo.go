package utils

import (
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/dizhi"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/miaoxian"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/stars"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/startype"
)

// SetHuo　設定火星
func SetHuo(board *Board, birthYear *dizhi.DiZhi, birthHour *dizhi.DiZhi) *Board {
	xaxis := getHuoLingGroupMap(birthYear)
	var index int
	switch xaxis {
	case 0:
		index = (int(*birthHour) + 1) % 12
	case 1:
		index = (int(*birthHour) + 2) % 12
	case 2:
		index = (int(*birthHour) + 3) % 12
	case 3:
		index = (int(*birthHour) + 9) % 12
	}
	board.Blocks[index].Stars = append(board.Blocks[index].Stars, &Star{
		Name:     stars.Huo.String(),
		StarType: startype.RightFuXing.String(),
		MiaoXian: getHuoMiaoXian(index).String(),
	})
	board.StarsMap[stars.Huo] = index
	return board
}

func getHuoLingGroupMap(birthYear *dizhi.DiZhi) int {
	birthYearGroup := map[dizhi.DiZhi]int{
		dizhi.Yin:  0,
		dizhi.Wu:   0,
		dizhi.Xu:   0,
		dizhi.Shen: 1,
		dizhi.Zi:   1,
		dizhi.Chen: 1,
		dizhi.Si:   2,
		dizhi.You:  2,
		dizhi.Chou: 2,
		dizhi.Hai:  3,
		dizhi.Mao:  3,
		dizhi.Wei:  3,
	}

	return birthYearGroup[*birthYear]
}

// getHuoMiaoXian 得火星廟陷
func getHuoMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Ping,
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Ping,
		miaoxian.Xian2,
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Xian2,
		miaoxian.Xian,
		miaoxian.Xian,
		miaoxian.Miao,
		miaoxian.Ping,
	}
	return miaoXianMap[index]
}
