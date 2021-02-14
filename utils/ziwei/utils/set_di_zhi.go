package utils

import (
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/dizhi"
)

func SetupDiZhi(board *Board) *Board {
	board.Blocks = make([]*Block, defaultBoardBlock)
	for i := range board.Blocks {
		board.Blocks[i] = &Block{
			Location: &Location{
				DiZhi: dizhi.DiZhi(i),
			},
		}
	}
	return board
}
