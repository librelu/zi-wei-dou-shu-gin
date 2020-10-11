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
	blocks = setUpNainGanStars(&lunaDate.Year.TianGan, blocks)
	blocks = setXunKong(lunaDate.Year, blocks)
	blocks = setJieKong(&lunaDate.Year.TianGan, blocks)
	blocks = setNianZhiXiZhuXing(&lunaDate.Year.DiZhi, blocks)

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
		Name:     stars.ZiWei.String(),
		StarType: startype.FourteenMainStars,
	})
	tianFuIndex, err := setTianFuStarLocation(ziWeiStarIndex)
	if err == nil {
		blocks[tianFuIndex].Stars = append(blocks[tianFuIndex].Stars, &Star{
			Name:     stars.TianFu.String(),
			StarType: startype.FourteenMainStars,
		})
	} else {
		return nil, fmt.Errorf("tian fu star not found, error: %w", err)
	}

	blocks = setStarsBeggingWithZiWei(ziWeiStarIndex, blocks)
	blocks = setStarsBeggingWithTianFu(tianFuIndex, blocks)
	//TODO: logic missing in setSiHua
	blocks = setSiHua(blocks)
	return blocks, nil
}

//setSiHua TODO: 四化 should set after 安時諸星
func setSiHua(blocks []*Block) []*Block {
	return blocks
}

//setStarsBeggingWithZiWei 順時針一宮安天機星，跳隔一宮，安太陽星，順時針一宮安武曲星，順時針一宮安天同星，跳隔兩宮，安廉貞星
func setStarsBeggingWithZiWei(ziWeiStarIndex int, blocks []*Block) []*Block {
	tianJi := ziWeiStarIndex - 1
	if tianJi < 0 {
		tianJi = 12 + tianJi
	}
	blocks[tianJi].Stars = append(blocks[tianJi].Stars, &Star{
		Name:     stars.TianJi.String(),
		StarType: startype.FourteenMainStars,
	})
	taiYang := tianJi - 2
	if taiYang < 0 {
		taiYang = 12 + taiYang
	}
	blocks[taiYang].Stars = append(blocks[taiYang].Stars, &Star{
		Name:     stars.TaiYang.String(),
		StarType: startype.FourteenMainStars,
	})
	wuQu := taiYang - 1
	if wuQu < 0 {
		wuQu = 12 + wuQu
	}
	blocks[wuQu].Stars = append(blocks[wuQu].Stars, &Star{
		Name:     stars.WuQu.String(),
		StarType: startype.FourteenMainStars,
	})
	tianTong := wuQu - 1
	if tianTong < 0 {
		tianTong = 12 + tianTong
	}
	blocks[tianTong].Stars = append(blocks[tianTong].Stars, &Star{
		Name:     stars.TianTong.String(),
		StarType: startype.FourteenMainStars,
	})
	lianZhen := tianTong - 3
	if lianZhen < 0 {
		lianZhen = 12 + lianZhen
	}
	blocks[lianZhen].Stars = append(blocks[lianZhen].Stars, &Star{
		Name:     stars.LianZhen.String(),
		StarType: startype.FourteenMainStars,
	})
	return blocks
}

//setStarsBeggingWithTianFu 逆時針一宮安太陰星，逆時針一宮安貪狼星，逆時針一宮安巨門星，逆時針一宮安天相星，逆時針一宮安天梁星，逆時針一宮安七殺星，跳隔三宮，安破軍星
func setStarsBeggingWithTianFu(tianFuIndex int, blocks []*Block) []*Block {
	taiYin := tianFuIndex + 1
	if taiYin > 11 {
		taiYin = taiYin - 12
	}
	blocks[taiYin].Stars = append(blocks[taiYin].Stars, &Star{
		Name:     stars.TaiYin.String(),
		StarType: startype.FourteenMainStars,
	})
	tanLang := taiYin + 1
	if tanLang > 11 {
		tanLang = tanLang - 12
	}
	blocks[tanLang].Stars = append(blocks[tanLang].Stars, &Star{
		Name:     stars.TanLang.String(),
		StarType: startype.FourteenMainStars,
	})
	juMen := tanLang + 1
	if juMen > 11 {
		juMen = juMen - 12
	}
	blocks[juMen].Stars = append(blocks[juMen].Stars, &Star{
		Name:     stars.JuMen.String(),
		StarType: startype.FourteenMainStars,
	})
	tianXiang := juMen + 1
	if tianXiang > 11 {
		tianXiang = tianXiang - 12
	}
	blocks[tianXiang].Stars = append(blocks[tianXiang].Stars, &Star{
		Name:     stars.TianXiang.String(),
		StarType: startype.FourteenMainStars,
	})
	tianLiang := tianXiang + 1
	if tianLiang > 11 {
		tianLiang = tianLiang - 12
	}
	blocks[tianLiang].Stars = append(blocks[tianLiang].Stars, &Star{
		Name:     stars.TianLiang.String(),
		StarType: startype.FourteenMainStars,
	})
	qiSha := tianLiang + 1
	if qiSha > 11 {
		qiSha = qiSha - 12
	}
	blocks[qiSha].Stars = append(blocks[qiSha].Stars, &Star{
		Name:     stars.QiSha.String(),
		StarType: startype.FourteenMainStars,
	})
	poJun := qiSha + 4
	if poJun > 11 {
		poJun = poJun - 12
	}
	blocks[poJun].Stars = append(blocks[poJun].Stars, &Star{
		Name:     stars.PoJun.String(),
		StarType: startype.FourteenMainStars,
	})
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

// setUpNainGanStars　設定年干系諸星
func setUpNainGanStars(tainGan *tiangan.TianGan, blocks []*Block) []*Block {
	setLuCun(tainGan, blocks)
	setQingYang(tainGan, blocks)
	setTuoLuo(tainGan, blocks)
	setTainKui(tainGan, blocks)
	setTainYue(tainGan, blocks)
	setTainGuan(tainGan, blocks)
	setTainFu(tainGan, blocks)
	return blocks
}

// setLuCun 設定祿存
func setLuCun(tainGan *tiangan.TianGan, blocks []*Block) []*Block {
	luCunMap := []dizhi.DiZhi{
		dizhi.Yin,
		dizhi.Mao,
		dizhi.Si,
		dizhi.Wu,
		dizhi.Si,
		dizhi.Wu,
		dizhi.Shen,
		dizhi.You,
		dizhi.Hai,
		dizhi.Zi,
	}
	luCunLocation := luCunMap[int(*tainGan)]
	blocks[luCunLocation].Stars = append(blocks[luCunLocation].Stars, &Star{
		Name:     stars.LuCun.String(),
		StarType: startype.NianGanXiZhuXing,
	})
	return blocks
}

// setQingYang 設定擎羊
func setQingYang(tainGan *tiangan.TianGan, blocks []*Block) []*Block {
	qingYangMap := []dizhi.DiZhi{
		dizhi.Mao,
		dizhi.Chen,
		dizhi.Wu,
		dizhi.Wei,
		dizhi.Wu,
		dizhi.Wei,
		dizhi.You,
		dizhi.Xu,
		dizhi.Zi,
		dizhi.Chou,
	}
	qingYangLocation := qingYangMap[int(*tainGan)]
	blocks[qingYangLocation].Stars = append(blocks[qingYangLocation].Stars, &Star{
		Name:     stars.QingYang.String(),
		StarType: startype.NianGanXiZhuXing,
	})
	return blocks
}

// setTuoLuo 設定陀羅
func setTuoLuo(tainGan *tiangan.TianGan, blocks []*Block) []*Block {
	tuoLuoMap := []dizhi.DiZhi{
		dizhi.Chou,
		dizhi.Yin,
		dizhi.Chen,
		dizhi.Si,
		dizhi.Chen,
		dizhi.Si,
		dizhi.Wei,
		dizhi.Shen,
		dizhi.Xu,
		dizhi.Hai,
	}
	tuoLuoLocation := tuoLuoMap[int(*tainGan)]
	blocks[tuoLuoLocation].Stars = append(blocks[tuoLuoLocation].Stars, &Star{
		Name:     stars.TuoLuo.String(),
		StarType: startype.NianGanXiZhuXing,
	})
	return blocks
}

// setTainKui 設天魁
func setTainKui(tainGan *tiangan.TianGan, blocks []*Block) []*Block {
	tainKuiMap := []dizhi.DiZhi{
		dizhi.Chou,
		dizhi.Zi,
		dizhi.Hai,
		dizhi.Hai,
		dizhi.Chou,
		dizhi.Zi,
		dizhi.Yin,
		dizhi.Yin,
		dizhi.Mao,
		dizhi.Mao,
	}
	tainKuiLocation := tainKuiMap[int(*tainGan)]
	blocks[tainKuiLocation].Stars = append(blocks[tainKuiLocation].Stars, &Star{
		Name:     stars.TianKui.String(),
		StarType: startype.NianGanXiZhuXing,
	})
	return blocks
}

// setTainYue 設定天鉞
func setTainYue(tainGan *tiangan.TianGan, blocks []*Block) []*Block {
	tainYueMap := []dizhi.DiZhi{
		dizhi.Wei,
		dizhi.Shen,
		dizhi.You,
		dizhi.You,
		dizhi.Wei,
		dizhi.Shen,
		dizhi.Wu,
		dizhi.Wu,
		dizhi.Si,
		dizhi.Si,
	}
	tainYueLocation := tainYueMap[int(*tainGan)]
	blocks[tainYueLocation].Stars = append(blocks[tainYueLocation].Stars, &Star{
		Name:     stars.TianYue.String(),
		StarType: startype.NianGanXiZhuXing,
	})
	return blocks
}

// setTainGuan 設定天官
func setTainGuan(tainGan *tiangan.TianGan, blocks []*Block) []*Block {
	tainGuanMap := []dizhi.DiZhi{
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
	tainGuanLocation := tainGuanMap[int(*tainGan)]
	blocks[tainGuanLocation].Stars = append(blocks[tainGuanLocation].Stars, &Star{
		Name:     stars.TianGuan.String(),
		StarType: startype.NianGanXiZhuXing,
	})
	return blocks
}

// setTainGuan 設定天福
func setTainFu(tainGan *tiangan.TianGan, blocks []*Block) []*Block {
	tainFuMap := []dizhi.DiZhi{
		dizhi.You,
		dizhi.Shen,
		dizhi.Zi,
		dizhi.Hai,
		dizhi.Mao,
		dizhi.Yin,
		dizhi.Wu,
		dizhi.Si,
		dizhi.Wu,
		dizhi.Si,
	}
	tainFuLocation := tainFuMap[int(*tainGan)]
	blocks[tainFuLocation].Stars = append(blocks[tainFuLocation].Stars, &Star{
		Name:     stars.TianFu.String(),
		StarType: startype.NianGanXiZhuXing,
	})
	return blocks
}

// setXunKong 設定旬空星
func setXunKong(birthYear *lunacal.TianGanDiZhi, blocks []*Block) []*Block {
	xunKongMapXAxis := 6 - int(birthYear.DiZhi/2) - 1
	xunKongIndex := (xunKongMapXAxis*10 + int(birthYear.TianGan)) % 12
	blocks[xunKongIndex].Stars = append(blocks[xunKongIndex].Stars, &Star{
		Name:     stars.XunKong.String(),
		StarType: startype.NianGanXiZhuXing,
	})

	return blocks
}

// setJieKong 設定截空星
// TODO setting jie kong in two locations
func setJieKong(birthYear *tiangan.TianGan, blocks []*Block) []*Block {
	index := int(*birthYear) % 5
	if index == 0 {
		index = 5
	} else {
		index--
	}
	return blocks
}

// NianZhiXiZhuXing 設定年支系諸星
func setNianZhiXiZhuXing(birthYear *dizhi.DiZhi, blocks []*Block) []*Block {
	blocks = setTainKu(birthYear, blocks)
	blocks = setTainXu(birthYear, blocks)
	blocks = setLongChi(birthYear, blocks)
	blocks = setFengGe(birthYear, blocks)
	blocks = setHongLuan(birthYear, blocks)
	blocks = setTainXi(birthYear, blocks)
	blocks = setGuChen(birthYear, blocks)
	blocks = setGuaXiu(birthYear, blocks)
	blocks = setJieShen(birthYear, blocks)
	blocks = setPoSui(birthYear, blocks)
	return blocks
}

// setTainKu 設定天哭
func setTainKu(birthYear *dizhi.DiZhi, blocks []*Block) []*Block {
	index := (11 - int(*birthYear) + 7) % 12
	blocks[index].Stars = append(blocks[index].Stars, &Star{
		Name:     stars.TianKu.String(),
		StarType: startype.NianZhiXiZhuXing,
	})
	return blocks
}

// setTainXu 設定天虛
func setTainXu(birthYear *dizhi.DiZhi, blocks []*Block) []*Block {
	index := (int(*birthYear) + 6) % 12
	blocks[index].Stars = append(blocks[index].Stars, &Star{
		Name:     stars.TianXu.String(),
		StarType: startype.NianZhiXiZhuXing,
	})
	return blocks
}

// setLongChi 設定龍池
func setLongChi(birthYear *dizhi.DiZhi, blocks []*Block) []*Block {
	index := (int(*birthYear) + 4) % 12
	blocks[index].Stars = append(blocks[index].Stars, &Star{
		Name:     stars.LongChi.String(),
		StarType: startype.NianZhiXiZhuXing,
	})
	return blocks
}

// setFengGe 設定鳳閣
func setFengGe(birthYear *dizhi.DiZhi, blocks []*Block) []*Block {
	index := (11 - int(*birthYear) + 11) % 12
	blocks[index].Stars = append(blocks[index].Stars, &Star{
		Name:     stars.FengGe.String(),
		StarType: startype.NianZhiXiZhuXing,
	})
	return blocks
}

// setHongLuan 設定紅鸞
func setHongLuan(birthYear *dizhi.DiZhi, blocks []*Block) []*Block {
	index := (11 - int(*birthYear) + 4) % 12
	blocks[index].Stars = append(blocks[index].Stars, &Star{
		Name:     stars.HongLuan.String(),
		StarType: startype.NianZhiXiZhuXing,
	})
	return blocks
}

// setTainXi 設定天喜
func setTainXi(birthYear *dizhi.DiZhi, blocks []*Block) []*Block {
	index := (11 - int(*birthYear) + 10) % 12
	blocks[index].Stars = append(blocks[index].Stars, &Star{
		Name:     stars.TainXi.String(),
		StarType: startype.NianZhiXiZhuXing,
	})
	return blocks
}

// setGuChen 設定孤辰
func setGuChen(birthYear *dizhi.DiZhi, blocks []*Block) []*Block {
	locationIndex := (int(*birthYear) + 1) / 3 % 4
	locations := []dizhi.DiZhi{dizhi.Yin, dizhi.Si, dizhi.Shen, dizhi.Hai}
	index := locations[locationIndex]
	blocks[index].Stars = append(blocks[index].Stars, &Star{
		Name:     stars.GuChen.String(),
		StarType: startype.NianZhiXiZhuXing,
	})
	return blocks
}

// setGuaXiu　設定寡宿
func setGuaXiu(birthYear *dizhi.DiZhi, blocks []*Block) []*Block {
	locationIndex := (int(*birthYear) + 1) / 3 % 4
	locations := []dizhi.DiZhi{dizhi.Xu, dizhi.Chou, dizhi.Chen, dizhi.Wei}
	index := locations[locationIndex]
	blocks[index].Stars = append(blocks[index].Stars, &Star{
		Name:     stars.GuaXiu.String(),
		StarType: startype.NianZhiXiZhuXing,
	})
	return blocks
}

// setJieShen 設定解神
func setJieShen(birthYear *dizhi.DiZhi, blocks []*Block) []*Block {
	index := (11 - int(*birthYear) + 11) % 12
	blocks[index].Stars = append(blocks[index].Stars, &Star{
		Name:     stars.JieShen.String(),
		StarType: startype.NianZhiXiZhuXing,
	})
	return blocks
}

// setPoSui 設定破碎
func setPoSui(birthYear *dizhi.DiZhi, blocks []*Block) []*Block {
	locationIndex := int(int(*birthYear) % 3)
	locations := []dizhi.DiZhi{dizhi.Si, dizhi.Chou, dizhi.You}
	blocks[locations[locationIndex]].Stars = append(
		blocks[locations[locationIndex]].Stars, &Star{
			Name:     stars.PoSui.String(),
			StarType: startype.NianZhiXiZhuXing,
		})
	return blocks
}
