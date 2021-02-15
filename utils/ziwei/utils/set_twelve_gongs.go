package utils

import (
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/dizhi"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/gong"
)

func SetTwelveGongs(board *Board, mingGongLocation dizhi.DiZhi) *Board {
	for i := range board.Blocks {
		index := int(mingGongLocation) + i
		if index > 11 {
			index -= 12
		}
		board.Blocks[index].GongWeiName = gong.Gong(i).String()
	}
	return board
}
