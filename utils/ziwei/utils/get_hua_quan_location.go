package utils

import (
	"fmt"

	"github.com/zi-wei-dou-shu-gin/utils/ziwei/stars"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/tiangan"
)

func GetHuaQuanLocation(board *Board, birthYear tiangan.TianGan) (int, error) {
	starMap := []stars.StarName{
		stars.PoJun,
		stars.TianLiang,
		stars.TianJi,
		stars.TianTong,
		stars.TaiYin,
		stars.TanLang,
		stars.WuQu,
		stars.TaiYang,
		stars.ZiWei,
		stars.JuMen,
	}
	starName := starMap[birthYear]
	index, ok := board.StarsMap[starName]
	if !ok {
		return 0, fmt.Errorf("current star not found, current birth year: %d", birthYear)
	}
	return index, nil
}
