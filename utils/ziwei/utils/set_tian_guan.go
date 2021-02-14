package utils

import (
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/dizhi"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/stars"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/startype"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/tiangan"
)

// SetTianGuan 設定天官
func SetTianGuan(board *Board, tianGan tiangan.TianGan) *Board {
	tianGuanMap := []dizhi.DiZhi{
		dizhi.Wei,
		dizhi.Chen,
		dizhi.Si,
		dizhi.Yin,
		dizhi.Mao,
		dizhi.You,
		dizhi.Hai,
		dizhi.You,
		dizhi.Xu,
		dizhi.Wu,
	}
	tianGuanLocation := tianGuanMap[int(tianGan)]
	board.Blocks[tianGuanLocation].Stars = append(board.Blocks[tianGuanLocation].Stars, &Star{
		Name:     stars.TianGuan.String(),
		StarType: startype.NianGanXiZhuXing.String(),
	})
	return board
}
