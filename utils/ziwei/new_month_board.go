package ziwei

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/zi-wei-dou-shu-gin/utils/lunacal"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/dizhi"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/genders"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/stars"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/startype"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/utils"
)

func NewMonthBoard(birthday time.Time, targetDate time.Time, gender genders.Gender) (*MonthBoard, error) {
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
	monthBoard := MonthBoard(*tianBoard)
	yearMingGong := getYearMingGong(0)
	monthMingGong := getMonthMingGong(yearMingGong, tianBoard.LunaBirthday)
	monthBoard.MingGongLocation = int(monthMingGong)
	// TODO: fix month setting
	index := getIndex(targetDate, birthday)
	fmt.Println(index)
	monthBoard.rotateGongWeiNameByIndex(index)
	lunaDate := lunacal.Solar2Lunar(targetDate)
	currentTianGan := utils.GetYinShou(lunaDate.Year.TianGan)
	b := utils.Board(monthBoard)
	utilsBoard := utils.SetTwelveGongs(&b, dizhi.DiZhi(monthBoard.MingGongLocation))
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
	monthBoard = MonthBoard(*utilsBoard)
	return &monthBoard, nil
}

// rotateGongWeiNameByIndex one index is ten years. Basic on ages, this function is returning
// different Gong Wei location.
func (mb *MonthBoard) rotateGongWeiNameByIndex(index int) {
	blocksLength := len(mb.Blocks)
	gongWeiNames := make([]string, blocksLength)
	mb.MingGongLocation = (mb.MingGongLocation + index) % 12
	for i, block := range mb.Blocks {
		idx := (i + index) % 12
		if idx < 0 {
			idx = idx + blocksLength
		}
		gongWeiNames[idx] = block.GongWeiName
	}
	for i, name := range gongWeiNames {
		mb.Blocks[i].GongWeiName = name
	}
	return
}

func getIndex(targetDate time.Time, birthDate time.Time) int {
	if targetDate.Before(targetDate) {
		return 0
	}
	if targetDate.Year() == birthDate.Year() {
		return int(targetDate.Month()) - int(birthDate.Month())
	}
	yearDifference := targetDate.Year() - birthDate.Year()
	return yearDifference*12 + int(targetDate.Month())
}

func getMonthMingGong(yearMingGong dizhi.DiZhi, lunaBirthDay *lunacal.LunaDate) dizhi.DiZhi {
	month := lunaBirthDay.Month - 1
	startMonth := yearMingGong - dizhi.DiZhi(month)
	if int(startMonth) < 0 {
		startMonth = startMonth + 12
	}
	monthMingGong := (startMonth + *lunaBirthDay.Hour) % 12
	return monthMingGong
}
