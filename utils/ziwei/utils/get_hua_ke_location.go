package utils

import (
	"fmt"

	"github.com/zi-wei-dou-shu-gin/utils/ziwei/stars"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/tiangan"
)

func GetHuaKeLocation(board *Board, birthYear tiangan.TianGan) (int, error) {
	starMap := []stars.StarName{
		stars.WuQu,
		stars.ZiWei,
		stars.WenChang,
		stars.TianJi,
		stars.TaiYang,
		stars.TianLiang,
		stars.TianFu,
		stars.WenQu,
		stars.TianFu,
		stars.TaiYin,
	}
	starName := starMap[birthYear]
	index, ok := board.StarsMap[starName]
	if !ok {
		return 0, fmt.Errorf("current star not found, current birth year: %d", birthYear)
	}
	return index, nil
}
