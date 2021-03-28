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

func NewDateBoard(birthday time.Time, gender genders.Gender, index int) (*DateBoard, error) {
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
	dateBoard := DateBoard(*tianBoard)
	yearMingGong := getYearMingGong(0)
	monthMingGong := getMonthMingGong(yearMingGong, tianBoard.LunaBirthday)
	dateBoard.MingGongLocation = int(monthMingGong)
	dateBoard.rotateGongWeiNameByIndex(index)
	birthday = birthday.AddDate(0, 0, index)
	date := lunacal.StemBranchDayWithTianGangDizhi(birthday.Year(), int(birthday.Month()), birthday.Day())
	currentTianGan := date.TianGan
	b := utils.Board(dateBoard)
	utilsBoard := utils.SetTwelveGongs(&b, dizhi.DiZhi(dateBoard.MingGongLocation))
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
	dateBoard = DateBoard(*utilsBoard)
	return &dateBoard, nil
}

// rotateGongWeiNameByIndex one index is ten years. Basic on ages, this function is returning
// different Gong Wei location.
func (db *DateBoard) rotateGongWeiNameByIndex(index int) {
	blocksLength := len(db.Blocks)
	gongWeiNames := make([]string, blocksLength)
	db.MingGongLocation = (db.MingGongLocation + index) % 12
	for i, block := range db.Blocks {
		idx := (i + index) % 12
		if idx < 0 {
			idx = idx + blocksLength
		}
		gongWeiNames[idx] = block.GongWeiName
	}
	for i, name := range gongWeiNames {
		db.Blocks[i].GongWeiName = name
	}
	return
}
