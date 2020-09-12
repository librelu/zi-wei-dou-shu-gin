package ziwei

import (
	"fmt"
	"math"
	"time"

	"github.com/zi-wei-dou-shu-gin/utils/lunacal"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/dizhi"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/gong"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/mingju"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/tiangan"
)

func NewBoard(birthday time.Time) *Board {
	lunaDate := lunacal.Solar2Lunar(birthday)
	blocks := setupDiZhi()
	blocks = setYinShou(&lunaDate.Year.TianGan, blocks)
	blocks = setupGongWei(lunaDate, blocks)
	mingGongLocation := getMingGong(lunaDate.Hour, lunaDate.Month)
	mingJu := getMingJu(mingGongLocation, lunaDate.Year.TianGan)
	blocks = settZiWeiStar(mingJu, lunaDate.Day, blocks)
	return &Board{
		Blocks: blocks,
		MingJu: mingJu,
	}
}

func setupDiZhi() []*Block {
	blocks := make([]*Block, defaultBoardBlock)
	for i := range blocks {
		blocks[i] = &Block{
			Location: &Location{
				DiZhi: dizhi.DiZhi(i),
			},
		}
	}
	return blocks
}

func setYinShou(birthYear *tiangan.TianGan, blocks []*Block) []*Block {
	yinShou := yinShouMap[*birthYear]
	for i := range blocks {
		index := i + int(yinShou) - 2
		if index < 0 {
			index = index + 2
		} else if index > 9 {
			index -= 10
		}
		blocks[i].Location.TianGan = tiangan.TianGan(index)
	}
	return blocks
}

func setupGongWei(lunaDate *lunacal.LunaDate, blocks []*Block) []*Block {
	hour := lunaDate.Hour
	month := lunaDate.Month
	mingGongLocation := getMingGong(hour, month)
	shengGongLocation := getShengGong(hour, month)
	blocks = setTwelveStars(mingGongLocation, blocks)
	_ = shengGongLocation

	return blocks
}

func setTwelveStars(mingGongLocation *dizhi.DiZhi, blocks []*Block) []*Block {
	for i := range blocks {
		index := int(*mingGongLocation) + i
		if index > 11 {
			index -= 12
		}
		blocks[index].GongWeiName = gong.Gong(i)
	}
	return blocks
}

func getMingGong(hour *dizhi.DiZhi, month uint) *dizhi.DiZhi {
	hourIndex := int(*hour) - 2
	mingGong := int(month-1) - int(hourIndex)
	if mingGong < 0 {
		mingGong += 12
	}
	mingGongLocation := dizhi.DiZhi(mingGong)
	return &mingGongLocation
}

func getShengGong(hour *dizhi.DiZhi, month uint) *dizhi.DiZhi {
	hourIndex := int(*hour) - 10
	shengGong := int(month-1) + int(hourIndex)
	if shengGong < 0 {
		shengGong += 12
	}
	shengGongLocation := dizhi.DiZhi(shengGong)
	return &shengGongLocation
}

func getMingJu(mingGongLocation *dizhi.DiZhi, lunaBirthYear tiangan.TianGan) *MingJu {
	totalTianGanIndex := 10 / 2
	tianGanIndex := 0
	for i := 0; i < totalTianGanIndex; i++ {
		idx := int(lunaBirthYear) - (i * 2)
		if idx == 0 || idx == 1 {
			tianGanIndex = i
			break
		}
	}
	totalDiZhiIndex := 12 / 2
	diZhiIndex := 0
	for i := 0; i < totalDiZhiIndex; i++ {
		idx := int(*mingGongLocation) - (i * 2)
		if idx == 0 || idx == 1 {
			diZhiIndex = i
			break
		}
	}
	mingJuIndex := tianGanIndex + (diZhiIndex)%3
	if mingJuIndex < 0 {
		mingJuIndex = 4 + mingJuIndex
	}

	juType := mingju.MingJuType(mingJuIndex)
	return &MingJu{
		JuType: juType,
		Number: mingju.JuShuMap[juType],
	}
}

func settZiWeiStar(mingJu *MingJu, birthdate uint, blocks []*Block) []*Block {
	n := mingJu.Number
	fmt.Println(birthdate)
	index := uint(math.Floor(float64(birthdate / n)))
	backStep := birthdate % n
	if backStep != 0 {
		location := int(index) - int(backStep)
		fmt.Println(birthdate)
		fmt.Println(n)
		fmt.Println(backStep)
		fmt.Println("--------")
		if location < 0 {
			location = 12 + location
		}
	}
	_ = append(blocks[index].Stars, &Star{
		Name:     "紫微",
		StarType: "zi_wei",
	})
	return blocks
}
