package utils

import "github.com/zi-wei-dou-shu-gin/utils/ziwei/tiangan"

func SetYinShouAndTianGanLocation(board *Board, birthYear *tiangan.TianGan) *Board {
	yinShou := yinShouMap[*birthYear]
	for i := 0; i < 12; i++ {
		blockIndex := (i + 2) % 12
		tainGanName := (int(yinShou) + i) % 10
		board.Blocks[blockIndex].Location.TianGan = tiangan.TianGan(tainGanName)
	}
	return board
}
