package ziwei

import (
	"time"

	"github.com/zi-wei-dou-shu-gin/utils/ziwei/genders"
)

func NewTenYearsBoard(birthday time.Time, gender genders.Gender, index int) (*YearBoard, error) {
	// TOOD: cache board if necessary
	board, err := NewBoard(birthday, gender)
	if err != nil {
		return nil, err
	}
	yearBoard := &YearBoard{
		board,
	}
	yearBoard.rotateGongWeiNameByIndex(index)
	return yearBoard, nil
}

// rotateGongWeiNameByIndex one index is ten years. Basic on ages, this function is returning
// different Gong Wei location.
func (yb *YearBoard) rotateGongWeiNameByIndex(index int) {
	blocksLength := len(yb.Board.Blocks)
	gongWeiNames := make([]string, blocksLength)
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
