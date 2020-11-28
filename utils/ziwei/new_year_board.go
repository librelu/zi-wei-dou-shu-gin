package ziwei

import (
	"time"

	"github.com/zi-wei-dou-shu-gin/utils/ziwei/dizhi"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/genders"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/stars"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/startype"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/tiangan"
)

func NewTenYearsBoard(birthday time.Time, gender genders.Gender, index int) (*YearBoard, error) {
	// TOOD: cache board if necessary
	b, err := NewBoard(birthday, gender)
	if err != nil {
		return nil, err
	}
	board, err := b.CreateTianBoard()
	if err != nil {
		return nil, err
	}
	// clean stars
	for i := range board.Blocks {
		board.Blocks[i].Stars = make([]*Star, 0)
	}
	yearBoard := &YearBoard{
		Board: board,
	}
	yearBoard.mingGongLocation = board.getMingGong(board.LunaBirthday.Hour, board.LunaBirthday.Month)
	yearBoard.rotateGongWeiNameByIndex(index)
	currentTianGan := tiangan.TianGan(int(board.LunaBirthday.Year.TianGan)+index) % 10
	board.setLuCun(&currentTianGan)
	board.setQingYang(&currentTianGan)
	board.setTuoLuo(&currentTianGan)
	board.setTianKui(&currentTianGan)
	board.setTianGuan(&currentTianGan)
	if idx, err := board.getHuaJiLocation(&currentTianGan); err == nil {
		board.Blocks[idx].Stars = append(board.Blocks[idx].Stars, &Star{
			Name:     stars.HuaLu.String(),
			StarType: startype.LiuNianGanXing.String(),
		})
	} else {
		return nil, err
	}
	if idx, err := board.getHuaKeLocation(&currentTianGan); err == nil {
		board.Blocks[idx].Stars = append(board.Blocks[idx].Stars, &Star{
			Name:     stars.HuaKe.String(),
			StarType: startype.LiuNianGanXing.String(),
		})
	} else {
		return nil, err
	}
	if idx, err := board.getHuaQuanLocation(&currentTianGan); err == nil {
		board.Blocks[idx].Stars = append(board.Blocks[idx].Stars, &Star{
			Name:     stars.HuaQuan.String(),
			StarType: startype.LiuNianGanXing.String(),
		})
	} else {
		return nil, err
	}
	if idx, err := board.getHuaLuLocation(&currentTianGan); err == nil {
		board.Blocks[idx].Stars = append(board.Blocks[idx].Stars, &Star{
			Name:     stars.HuaQuan.String(),
			StarType: startype.LiuNianGanXing.String(),
		})
	} else {
		return nil, err
	}
	currentDiZhi := dizhi.DiZhi(int(board.LunaBirthday.Year.DiZhi)+index) % 12
	board.setHuo(&currentDiZhi, board.LunaBirthday.Hour)
	board.setLing(&currentDiZhi, board.LunaBirthday.Hour)
	return yearBoard, nil
}

// rotateGongWeiNameByIndex one index is ten years. Basic on ages, this function is returning
// different Gong Wei location.
func (yb *YearBoard) rotateGongWeiNameByIndex(index int) {
	blocksLength := len(yb.Board.Blocks)
	gongWeiNames := make([]string, blocksLength)
	mingGongLocation := dizhi.DiZhi((int(*yb.mingGongLocation) + index) % 12)
	yb.mingGongLocation = &mingGongLocation
	for i, Block := range yb.Board.Blocks {
		idx := (i - index) % 12
		if idx < 0 {
			idx = idx + blocksLength
		}
		gongWeiNames[idx] = Block.GongWeiName
	}
	for i, name := range gongWeiNames {
		yb.Board.Blocks[i].GongWeiName = name
	}
	return
}
