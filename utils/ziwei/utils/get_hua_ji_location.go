package utils

import (
	"fmt"

	"github.com/zi-wei-dou-shu-gin/utils/ziwei/stars"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/tiangan"
)

func GetHuaJiLocation(board *Board, birthYear tiangan.TianGan) (int, error) {
	starMap := []stars.StarName{
		stars.TaiYang,
		stars.TaiYin,
		stars.LianZhen,
		stars.JuMen,
		stars.TianJi,
		stars.WenQu,
		stars.TianTong,
		stars.WenChang,
		stars.WuQu,
		stars.TanLang,
	}
	starName := starMap[birthYear]
	index, ok := board.StarsMap[starName]
	if !ok {
		return 0, fmt.Errorf("current star not found, current birth year: %d", birthYear)
	}
	return index, nil
}
