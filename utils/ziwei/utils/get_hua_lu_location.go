package utils

import (
	"fmt"

	"github.com/zi-wei-dou-shu-gin/utils/ziwei/stars"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/tiangan"
)

func GetHuaLuLocation(board *Board, birthYear *tiangan.TianGan) (int, error) {
	starMap := []stars.StarName{
		stars.LianZhen,
		stars.TianJi,
		stars.TianTong,
		stars.TaiYin,
		stars.TanLang,
		stars.WuQu,
		stars.TaiYang,
		stars.JuMen,
		stars.TianLiang,
		stars.PoJun,
	}
	starName := starMap[*birthYear]
	index, ok := board.StarsMap[starName]
	if !ok {
		return 0, fmt.Errorf("current star not found, current birth year: %d", birthYear)
	}
	return index, nil
}
