package ziwei

import (
	"time"

	"github.com/pkg/errors"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/dizhi"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/genders"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/stars"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/startype"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/tiangan"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/utils"
)

func NewTenYearsBoard(birthday time.Time, gender genders.Gender, index int) (*TenYearBoard, error) {
	tianBoard, err := NewTianBoard(birthday, gender)
	if err != nil {
		return nil, errors.Wrap(err, "can't new tian board in year board process")
	}
	if tianBoard, err = tianBoard.CreateTianBoard(); err != nil {
		return nil, errors.Wrap(err, "can't create tian board in year board process")
	}
	// clean stars
	for i := range tianBoard.Blocks {
		tianBoard.Blocks[i].Stars = make([]*utils.Star, 0)
	}
	tenYearBoard := TenYearBoard(*tianBoard)
	tenYearBoard.rotateGongWeiNameByIndex(index)
	currentTianGan := tiangan.TianGan(int(tenYearBoard.LunaBirthday.Year.TianGan)+index) % 10
	b := utils.Board(tenYearBoard)
	utilsBoard := utils.SetTwelveGongs(&b, dizhi.DiZhi(tenYearBoard.MingGongLocation))
	utilsBoard = utils.SetLuCun(utilsBoard, currentTianGan)
	utilsBoard = utils.SetQingYang(utilsBoard, currentTianGan)
	utilsBoard = utils.SetTuoLuo(utilsBoard, currentTianGan)
	utilsBoard = utils.SetTianKui(utilsBoard, currentTianGan)
	utilsBoard = utils.SetTianGuan(utilsBoard, currentTianGan)
	if idx, err := utils.GetHuaJiLocation(utilsBoard, currentTianGan); err == nil {
		utilsBoard.Blocks[idx].Stars = append(utilsBoard.Blocks[idx].Stars, &utils.Star{
			Name:     stars.HuaLu.String(),
			StarType: startype.LiuNianGanXing.String(),
		})
	} else {
		return nil, err
	}
	if idx, err := utils.GetHuaKeLocation(utilsBoard, currentTianGan); err == nil {
		utilsBoard.Blocks[idx].Stars = append(utilsBoard.Blocks[idx].Stars, &utils.Star{
			Name:     stars.HuaLu.String(),
			StarType: startype.LiuNianGanXing.String(),
		})
	} else {
		return nil, err
	}
	if idx, err := utils.GetHuaQuanLocation(utilsBoard, currentTianGan); err == nil {
		utilsBoard.Blocks[idx].Stars = append(utilsBoard.Blocks[idx].Stars, &utils.Star{
			Name:     stars.HuaLu.String(),
			StarType: startype.LiuNianGanXing.String(),
		})
	} else {
		return nil, err
	}
	if idx, err := utils.GetHuaLuLocation(utilsBoard, &currentTianGan); err == nil {
		utilsBoard.Blocks[idx].Stars = append(utilsBoard.Blocks[idx].Stars, &utils.Star{
			Name:     stars.HuaQuan.String(),
			StarType: startype.LiuNianGanXing.String(),
		})
	} else {
		return nil, err
	}
	currentDiZhi := dizhi.DiZhi(int(utilsBoard.LunaBirthday.Year.DiZhi)+index) % 12
	utilsBoard = utils.SetHuo(utilsBoard, &currentDiZhi, utilsBoard.LunaBirthday.Hour)
	utilsBoard = utils.SetLing(utilsBoard, &currentDiZhi, utilsBoard.LunaBirthday.Hour)
	tenYearBoard = TenYearBoard(*utilsBoard)
	return &tenYearBoard, nil
}

// rotateGongWeiNameByIndex one index is ten years. Basic on ages, this function is returning
// different Gong Wei location.
func (yb *TenYearBoard) rotateGongWeiNameByIndex(index int) {
	blocks := make([]*utils.Block, len(yb.Blocks))
	if yb.Gender == genders.YangMale || yb.Gender == genders.YinFemale {
		mingGongLocation := int(dizhi.DiZhi((yb.MingGongLocation + index) % 12))
		yb.MingGongLocation = mingGongLocation
		// clockwise
		for i := range yb.Blocks {
			idx := (i + mingGongLocation) % 12
			blocks[idx] = yb.Blocks[idx]
		}
	}
	if yb.Gender == genders.YinMale || yb.Gender == genders.YangFemale {
		mingGongLocation := int(dizhi.DiZhi((yb.MingGongLocation - index) % 12))
		if mingGongLocation < 0 {
			mingGongLocation = 12 + mingGongLocation
		}
		yb.MingGongLocation = mingGongLocation
		for i := range yb.Blocks {
			// reverse clockwise
			idx := (mingGongLocation - i) % 12
			if idx < 0 {
				idx = 12 + idx
			}
			blocks[idx] = yb.Blocks[idx]
		}
	}
	yb.Blocks = blocks
	return
}
