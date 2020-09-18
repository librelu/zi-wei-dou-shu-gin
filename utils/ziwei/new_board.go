package ziwei

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/zi-wei-dou-shu-gin/utils/lunacal"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/dizhi"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/gong"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/mingju"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/stars"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/startype"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/tiangan"
)

func NewBoard(birthday time.Time) (*Board, error) {
	var err error
	lunaDate := lunacal.Solar2Lunar(birthday)
	blocks := setupDiZhi()
	blocks = setYinShou(&lunaDate.Year.TianGan, blocks)
	blocks = setupGongWei(lunaDate, blocks)
	mingGongLocation := getMingGong(lunaDate.Hour, lunaDate.Month)
	mingJu := getMingJu(mingGongLocation, lunaDate.Year.TianGan)
	blocks, err = setFourteenMainStars(mingGongLocation, mingJu, lunaDate.Day, blocks)
	if err != nil {
		return nil, fmt.Errorf("failed in set fourteen main stars, error: %w", err)
	}

	return &Board{
		Blocks: blocks,
		MingJu: mingJu,
	}, nil
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
	blocks = setTwelveGongs(mingGongLocation, blocks)
	_ = shengGongLocation

	return blocks
}

func setTwelveGongs(mingGongLocation *dizhi.DiZhi, blocks []*Block) []*Block {
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
	totalTianGanIndex := 5
	tianGanIndex := 0
	for i := 0; i < totalTianGanIndex; i++ {
		idx := int(lunaBirthYear) - (i * 2)
		if idx == 0 || idx == 1 {
			tianGanIndex = i
			break
		}
	}
	totalDiZhiIndex := 6
	diZhiIndex := 0
	for i := 0; i < totalDiZhiIndex; i++ {
		idx := int(*mingGongLocation) - (i * 2)
		if idx == 0 || idx == 1 {
			diZhiIndex = i % 3
			break
		}
	}

	mingJuIndex := tianGanIndex + diZhiIndex
	if mingJuIndex > 4 {
		mingJuIndex = mingJuIndex - 5
	}

	juType := mingju.MingJuType(mingJuIndex)
	return &MingJu{
		JuType: juType,
		Number: mingju.JuShuMap[juType],
	}
}

func setFourteenMainStars(mingGongLocation *dizhi.DiZhi, mingJu *MingJu, birthdate uint, blocks []*Block) ([]*Block, error) {
	ziWeiStarIndex := getZiWeiStarLocation(mingGongLocation, mingJu, birthdate)
	blocks[ziWeiStarIndex].Stars = append(blocks[ziWeiStarIndex].Stars, &Star{
		Name:     stars.ZiWei,
		StarType: startype.FourteenMainStars,
	})
	tianFuIndex, err := setTianFuStarLocation(ziWeiStarIndex)
	if err == nil {
		blocks[tianFuIndex].Stars = append(blocks[tianFuIndex].Stars, &Star{
			Name:     stars.TianFu,
			StarType: startype.FourteenMainStars,
		})
	} else {
		return nil, fmt.Errorf("tian fu star not found, error: %w", err)
	}

	blocks = setStarsBeggingWithZiWei(ziWeiStarIndex, blocks)
	blocks = setStarsBeggingWithTianFu(tianFuIndex, blocks)
	return blocks, nil
}

//setStarsBeggingWithZiWei 逆時針一宮安天機星，跳隔一宮，安太陽星，逆時針一宮安武曲星，逆時針一宮安天同星，跳隔兩宮，安廉貞星
func setStarsBeggingWithZiWei(ziWeiStarIndex int, blocks []*Block) []*Block {
	tianJi := ziWeiStarIndex - 1
	if tianJi < 0 {
		tianJi = 12 + tianJi
	}
	blocks[tianJi].Stars = append(blocks[tianJi].Stars, &Star{
		Name:     stars.TianJi,
		StarType: startype.FourteenMainStars,
	})
	taiYang := tianJi - 2
	if taiYang < 0 {
		taiYang = 12 + taiYang
	}
	blocks[taiYang].Stars = append(blocks[taiYang].Stars, &Star{
		Name:     stars.TaiYang,
		StarType: startype.FourteenMainStars,
	})
	wuQu := taiYang - 1
	if wuQu < 0 {
		wuQu = 12 + wuQu
	}
	blocks[wuQu].Stars = append(blocks[wuQu].Stars, &Star{
		Name:     stars.WuQu,
		StarType: startype.FourteenMainStars,
	})
	tianTong := wuQu - 1
	if tianTong < 0 {
		tianTong = 12 + tianTong
	}
	blocks[tianTong].Stars = append(blocks[tianTong].Stars, &Star{
		Name:     stars.TianTong,
		StarType: startype.FourteenMainStars,
	})
	lianZhen := tianTong - 3
	if lianZhen < 0 {
		lianZhen = 12 + lianZhen
	}
	blocks[lianZhen].Stars = append(blocks[lianZhen].Stars, &Star{
		Name:     stars.LianZhen,
		StarType: startype.FourteenMainStars,
	})
	return blocks
}

//setStarsBeggingWithTianFu 順時針一宮安太陰星，順時針一宮安貪狼星，順時針一宮安巨門星，順時針一宮安天相星，順時針一宮安天梁星，順時針一宮安七殺星，跳隔三宮，安破軍星
func setStarsBeggingWithTianFu(tianFuIndex int, blocks []*Block) []*Block {
	return blocks
}

func getZiWeiStarLocation(mingGongLocation *dizhi.DiZhi, mingJu *MingJu, birthdate uint) int {
	steps := int(math.Floor(float64(birthdate / mingJu.Number)))
	if birthdate%mingJu.Number > 0 {
		steps = steps + 1
	}

	stepsRemainder := steps*int(mingJu.Number) - int(birthdate)
	if (stepsRemainder % 2) == 0 {
		steps = steps + int(stepsRemainder)
	} else {
		steps = steps - int(stepsRemainder)
	}

	index := steps + 1
	if index > 11 {
		index = 12 - index
	} else if index < 0 {
		index = 12 + index
	}

	return index
}

func setTianFuStarLocation(ziWeiStarIndex int) (int, error) {
	index := 0
	if ziWeiStarIndex < 8 {
		index = 8 + (8 - ziWeiStarIndex)
	} else if ziWeiStarIndex > 8 {
		index = 8 - (ziWeiStarIndex - 8)
	} else if ziWeiStarIndex == 2 || ziWeiStarIndex == 8 {
		index = ziWeiStarIndex
	} else {
		return 0, errors.New("no index found")
	}

	if index > 11 {
		index = index - 12
	}

	return index, nil
}
