package ziwei

import (
	"time"

	"github.com/pkg/errors"
	"github.com/zi-wei-dou-shu-gin/utils/lunacal"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/dizhi"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/genders"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/stars"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/startype"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/utils"
)

func NewYearsBoard(birthday time.Time, targetDate time.Time, gender genders.Gender) (*YearBoard, error) {
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
	currentLunaDate := lunacal.Solar2Lunar(targetDate)
	yearBoard := YearBoard(*tianBoard)
	index := targetDate.Year() - birthday.Year()
	yearBoard.rotateGongWeiNameByIndex(index)
	yearMingGong := getYearMingGong(index)
	currentTianGan := currentLunaDate.Year.TianGan
	b := utils.Board(yearBoard)
	utilsBoard := utils.SetTwelveGongs(&b, yearMingGong)
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
	yearBoard = YearBoard(*utilsBoard)
	return &yearBoard, nil
}

// rotateGongWeiNameByIndex one index is ten years. Basic on ages, this function is returning
// different Gong Wei location.
func (yb *YearBoard) rotateGongWeiNameByIndex(index int) {
	blocksLength := len(yb.Blocks)
	gongWeiNames := make([]string, blocksLength)
	yb.MingGongLocation = int(dizhi.DiZhi((yb.MingGongLocation + index) % 12))
	for i, block := range yb.Blocks {
		idx := (i + index) % 12
		if idx < 0 {
			idx = idx + blocksLength
		}
		gongWeiNames[idx] = block.GongWeiName
	}
	for i, name := range gongWeiNames {
		yb.Blocks[i].GongWeiName = name
	}
	return
}

func getYearMingGong(index int) dizhi.DiZhi {
	lunaDate := lunacal.Solar2Lunar(time.Now())
	mingGong := (lunaDate.Year.DiZhi + dizhi.DiZhi(index)) % 12
	return mingGong
}
