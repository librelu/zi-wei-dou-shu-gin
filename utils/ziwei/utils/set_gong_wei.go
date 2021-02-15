package utils

import (
	"github.com/zi-wei-dou-shu-gin/utils/lunacal"
)

func SetupGongWei(board *Board, lunaDate *lunacal.LunaDate) *Board {
	hour := lunaDate.Hour
	month := lunaDate.Month
	mingGongLocation := GetMingGong(*hour, month)
	shengGongLocation := GetShengGong(*hour, month)
	board = SetTwelveGongs(board, mingGongLocation)
	board.ShenGongLocation = int(*shengGongLocation)
	return board
}
