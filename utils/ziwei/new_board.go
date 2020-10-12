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
	shenGongLocation := getShengGong(lunaDate.Hour, lunaDate.Month)
	blocks = setNianZhiXiZhuXing(&lunaDate.Year.DiZhi, mingGongLocation, shenGongLocation, blocks)
	blocks = setYueXiXing(int(lunaDate.Month), blocks)
	blocks = setShiXiZhuXing(lunaDate.Hour, blocks)

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
		// fmt.Println(index)
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
	} else if mingGong > 12 {
		mingGong = mingGong - 12
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
		index = index - 12
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
func setUpNainGanStars(tianGan *tiangan.TianGan, blocks []*Block) []*Block {
	setLuCun(tianGan, blocks)
	setQingYang(tianGan, blocks)
	setTuoLuo(tianGan, blocks)
	setTianKui(tianGan, blocks)
	setTianYue(tianGan, blocks)
	setTianGuan(tianGan, blocks)
	setTianFu(tianGan, blocks)
	return blocks
}

// setLuCun 設定祿存
func setLuCun(tianGan *tiangan.TianGan, blocks []*Block) []*Block {
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
	luCunLocation := luCunMap[int(*tianGan)]
	blocks[luCunLocation].Stars = append(blocks[luCunLocation].Stars, &Star{
		Name:     stars.LuCun.String(),
		StarType: startype.NianGanXiZhuXing,
	})
	return blocks
}

// setQingYang 設定擎羊
func setQingYang(tianGan *tiangan.TianGan, blocks []*Block) []*Block {
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
	qingYangLocation := qingYangMap[int(*tianGan)]
	blocks[qingYangLocation].Stars = append(blocks[qingYangLocation].Stars, &Star{
		Name:     stars.QingYang.String(),
		StarType: startype.NianGanXiZhuXing,
	})
	return blocks
}

// setTuoLuo 設定陀羅
func setTuoLuo(tianGan *tiangan.TianGan, blocks []*Block) []*Block {
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
	tuoLuoLocation := tuoLuoMap[int(*tianGan)]
	blocks[tuoLuoLocation].Stars = append(blocks[tuoLuoLocation].Stars, &Star{
		Name:     stars.TuoLuo.String(),
		StarType: startype.NianGanXiZhuXing,
	})
	return blocks
}

// setTianKui 設天魁
func setTianKui(tianGan *tiangan.TianGan, blocks []*Block) []*Block {
	tianKuiMap := []dizhi.DiZhi{
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
	tianKuiLocation := tianKuiMap[int(*tianGan)]
	blocks[tianKuiLocation].Stars = append(blocks[tianKuiLocation].Stars, &Star{
		Name:     stars.TianKui.String(),
		StarType: startype.NianGanXiZhuXing,
	})
	return blocks
}

// setTianYue 設定天鉞
func setTianYue(tianGan *tiangan.TianGan, blocks []*Block) []*Block {
	tianYueMap := []dizhi.DiZhi{
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
	tianYueLocation := tianYueMap[int(*tianGan)]
	blocks[tianYueLocation].Stars = append(blocks[tianYueLocation].Stars, &Star{
		Name:     stars.TianYue.String(),
		StarType: startype.NianGanXiZhuXing,
	})
	return blocks
}

// setTianGuan 設定天官
func setTianGuan(tianGan *tiangan.TianGan, blocks []*Block) []*Block {
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
	tianGuanLocation := tianGuanMap[int(*tianGan)]
	blocks[tianGuanLocation].Stars = append(blocks[tianGuanLocation].Stars, &Star{
		Name:     stars.TianGuan.String(),
		StarType: startype.NianGanXiZhuXing,
	})
	return blocks
}

// setTianGuan 設定天福
func setTianFu(tianGan *tiangan.TianGan, blocks []*Block) []*Block {
	tianFuMap := []dizhi.DiZhi{
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
	tianFuLocation := tianFuMap[int(*tianGan)]
	blocks[tianFuLocation].Stars = append(blocks[tianFuLocation].Stars, &Star{
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
func setJieKong(birthYear *tiangan.TianGan, blocks []*Block) []*Block {
	index := ((4 - int(*birthYear)%5 + 1) * 2) - 1
	blocks[index].Stars = append(blocks[index].Stars, &Star{
		Name:     stars.JieKong.String(),
		StarType: startype.NianGanXiZhuXing,
	})
	blocks[index-1].Stars = append(blocks[index-1].Stars, &Star{
		Name:     stars.JieKong.String(),
		StarType: startype.NianGanXiZhuXing,
	})
	return blocks
}

// NianZhiXiZhuXing 設定年支系諸星
func setNianZhiXiZhuXing(birthYear *dizhi.DiZhi, mingGongLocation *dizhi.DiZhi, shengGongLocation *dizhi.DiZhi, blocks []*Block) []*Block {
	blocks = setTianKu(birthYear, blocks)
	blocks = setTianXu(birthYear, blocks)
	blocks = setLongChi(birthYear, blocks)
	blocks = setFengGe(birthYear, blocks)
	blocks = setHongLuan(birthYear, blocks)
	blocks = setTianXi(birthYear, blocks)
	blocks = setGuChen(birthYear, blocks)
	blocks = setGuaXiu(birthYear, blocks)
	blocks = setJieShen(birthYear, blocks)
	blocks = setPoSui(birthYear, blocks)
	blocks = setTianMa(birthYear, blocks)
	blocks = setDaHao(birthYear, blocks)
	blocks = setTianDe(birthYear, blocks)
	blocks = setJieSha(birthYear, blocks)
	blocks = setHuaGai(birthYear, blocks)
	blocks = setXianChi(birthYear, blocks)
	blocks = setTianCai(birthYear, mingGongLocation, blocks)
	blocks = setTianShou(birthYear, shengGongLocation, blocks)
	blocks = setTianKong(birthYear, blocks)
	return blocks
}

// setTianKu 設定天哭
func setTianKu(birthYear *dizhi.DiZhi, blocks []*Block) []*Block {
	index := (11 - int(*birthYear) + 7) % 12
	blocks[index].Stars = append(blocks[index].Stars, &Star{
		Name:     stars.TianKu.String(),
		StarType: startype.NianZhiXiZhuXing,
	})
	return blocks
}

// setTianXu 設定天虛
func setTianXu(birthYear *dizhi.DiZhi, blocks []*Block) []*Block {
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

// setTianXi 設定天喜
func setTianXi(birthYear *dizhi.DiZhi, blocks []*Block) []*Block {
	index := (11 - int(*birthYear) + 10) % 12
	blocks[index].Stars = append(blocks[index].Stars, &Star{
		Name:     stars.TianXi.String(),
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

// setPoSui 設定天馬
func setTianMa(birthYear *dizhi.DiZhi, blocks []*Block) []*Block {
	locationIndex := int((int(*birthYear) + 1) % 4)
	locations := []dizhi.DiZhi{dizhi.Si, dizhi.Yin, dizhi.Hai, dizhi.Shen}
	blocks[locations[locationIndex]].Stars = append(
		blocks[locations[locationIndex]].Stars, &Star{
			Name:     stars.TianMa.String(),
			StarType: startype.NianZhiXiZhuXing,
		})
	return blocks
}

// setDaHao 設定大耗
func setDaHao(birthYear *dizhi.DiZhi, blocks []*Block) []*Block {
	locationMap := []dizhi.DiZhi{
		dizhi.Wei, dizhi.Wu, dizhi.You, dizhi.Shen, dizhi.Hai, dizhi.Xu, dizhi.Chou, dizhi.Zi, dizhi.Mao, dizhi.Yin, dizhi.Si, dizhi.Chen,
	}
	index := locationMap[*birthYear]
	blocks[index].Stars = append(
		blocks[index].Stars, &Star{
			Name:     stars.DaHao.String(),
			StarType: startype.NianZhiXiZhuXing,
		})
	return blocks
}

// setTianDe 設定天德
func setTianDe(birthYear *dizhi.DiZhi, blocks []*Block) []*Block {
	index := (int(*birthYear) + 9) % 12
	blocks[index].Stars = append(blocks[index].Stars, &Star{
		Name:     stars.TianDe.String(),
		StarType: startype.NianZhiXiZhuXing,
	})
	return blocks
}

// setJieSha 設定劫殺
func setJieSha(birthYear *dizhi.DiZhi, blocks []*Block) []*Block {
	locationIndex := int(int(*birthYear) % 4)
	locations := []dizhi.DiZhi{dizhi.Si, dizhi.Yin, dizhi.Hai, dizhi.Shen}
	blocks[locations[locationIndex]].Stars = append(
		blocks[locations[locationIndex]].Stars, &Star{
			Name:     stars.JieSha.String(),
			StarType: startype.NianZhiXiZhuXing,
		})
	return blocks
}

// setHuaGai 設定華蓋
func setHuaGai(birthYear *dizhi.DiZhi, blocks []*Block) []*Block {
	locationIndex := int(int(*birthYear) % 4)
	locations := []dizhi.DiZhi{dizhi.Chen, dizhi.Chou, dizhi.Xu, dizhi.Wei}
	blocks[locations[locationIndex]].Stars = append(
		blocks[locations[locationIndex]].Stars, &Star{
			Name:     stars.HuaGai.String(),
			StarType: startype.NianZhiXiZhuXing,
		})
	return blocks
}

// setXianChi 設定咸池
func setXianChi(birthYear *dizhi.DiZhi, blocks []*Block) []*Block {
	locationIndex := int(int(*birthYear) % 4)
	locations := []dizhi.DiZhi{dizhi.You, dizhi.Wu, dizhi.Mao, dizhi.Zi}
	blocks[locations[locationIndex]].Stars = append(
		blocks[locations[locationIndex]].Stars, &Star{
			Name:     stars.XianChi.String(),
			StarType: startype.NianZhiXiZhuXing,
		})
	return blocks
}

// setTianCai 設定天才
func setTianCai(birthYear *dizhi.DiZhi, mingGongLocation *dizhi.DiZhi, blocks []*Block) []*Block {
	index := (*mingGongLocation + *birthYear) % 12
	blocks[index].Stars = append(
		blocks[index].Stars, &Star{
			Name:     stars.TianCai.String(),
			StarType: startype.NianZhiXiZhuXing,
		})
	return blocks
}

// setTianShou 設定天壽
func setTianShou(birthYear *dizhi.DiZhi, shengGongLocation *dizhi.DiZhi, blocks []*Block) []*Block {
	index := (*shengGongLocation + *birthYear) % 12
	blocks[index].Stars = append(
		blocks[index].Stars, &Star{
			Name:     stars.TianShou.String(),
			StarType: startype.NianZhiXiZhuXing,
		})
	return blocks
}

// setTianKong 設定天空
func setTianKong(birthYear *dizhi.DiZhi, blocks []*Block) []*Block {
	index := (int(*birthYear) + 1) % 12
	blocks[index].Stars = append(blocks[index].Stars, &Star{
		Name:     stars.TianKong.String(),
		StarType: startype.NianZhiXiZhuXing,
	})
	return blocks
}

// setYueXiXing 安月系星
func setYueXiXing(birthMonth int, blocks []*Block) []*Block {
	blocks = setZuoFu(birthMonth, blocks)
	blocks = setYouBi(birthMonth, blocks)
	blocks = setTianXing(birthMonth, blocks)
	blocks = setTianYao(birthMonth, blocks)
	blocks = setTianWu(birthMonth, blocks)
	blocks = setYueXiXingTianYue(birthMonth, blocks)
	blocks = setYinSha(birthMonth, blocks)
	return blocks
}

// setZuoFu 設定左輔
func setZuoFu(birthMonth int, blocks []*Block) []*Block {
	index := (birthMonth + 3) % 12
	blocks[index].Stars = append(blocks[index].Stars, &Star{
		Name:     stars.ZuoFu.String(),
		StarType: startype.YueXiXing,
	})
	return blocks
}

// setYouBi 設定右弼
func setYouBi(birthMonth int, blocks []*Block) []*Block {
	index := ((13 - birthMonth) + 10) % 12
	blocks[index].Stars = append(blocks[index].Stars, &Star{
		Name:     stars.YouBi.String(),
		StarType: startype.YueXiXing,
	})
	return blocks
}

// setTianXing 設定天刑
func setTianXing(birthMonth int, blocks []*Block) []*Block {
	index := (birthMonth + 8) % 12
	blocks[index].Stars = append(blocks[index].Stars, &Star{
		Name:     stars.TianXing.String(),
		StarType: startype.YueXiXing,
	})
	return blocks
}

// setTianYao 設定天姚
func setTianYao(birthMonth int, blocks []*Block) []*Block {
	index := (birthMonth) % 12
	blocks[index].Stars = append(blocks[index].Stars, &Star{
		Name:     stars.TianYao.String(),
		StarType: startype.YueXiXing,
	})
	return blocks
}

// setTianWu 設定天巫
func setTianWu(birthMonth int, blocks []*Block) []*Block {
	locationIndex := (birthMonth - 1) % 4
	locations := []dizhi.DiZhi{dizhi.Si, dizhi.Shen, dizhi.Yin, dizhi.Hai}
	blocks[locations[locationIndex]].Stars = append(
		blocks[locations[locationIndex]].Stars, &Star{
			Name:     stars.TianWu.String(),
			StarType: startype.YueXiXing,
		})
	return blocks
}

// setYueXiXingTianYue 設定天月
// TODO: refactor Yue Xi Xing to a object
func setYueXiXingTianYue(birthMonth int, blocks []*Block) []*Block {
	locationMap := []dizhi.DiZhi{
		dizhi.Xu, dizhi.Si, dizhi.Chen, dizhi.Yin, dizhi.Wei, dizhi.Mao, dizhi.Hai, dizhi.Wei, dizhi.Yin, dizhi.Wu, dizhi.Xu, dizhi.Yin,
	}
	index := locationMap[birthMonth-1]
	blocks[index].Stars = append(
		blocks[index].Stars, &Star{
			Name:     stars.YueXiXingTianYue.String(),
			StarType: startype.YueXiXing,
		})
	return blocks
}

// setYinSha 設定陰煞
func setYinSha(birthMonth int, blocks []*Block) []*Block {
	locationIndex := (birthMonth + 2) % 6
	locations := []dizhi.DiZhi{dizhi.Shen, dizhi.Wu, dizhi.Chen, dizhi.Yin, dizhi.Zi, dizhi.Xu}
	blocks[locations[locationIndex]].Stars = append(
		blocks[locations[locationIndex]].Stars, &Star{
			Name:     stars.YinSha.String(),
			StarType: startype.YueXiXing,
		})
	return blocks
}

// setShiXiZhuXing 安時系諸星
func setShiXiZhuXing(birthHour *dizhi.DiZhi, blocks []*Block) []*Block {
	blocks = setWenChang(birthHour, blocks)
	blocks = setWenQu(birthHour, blocks)
	blocks = setDiJie(birthHour, blocks)
	blocks = setDiKong(birthHour, blocks)
	blocks = setTaiFu(birthHour, blocks)
	blocks = setFengGao(birthHour, blocks)
	return blocks
}

// setWenChang 設定文昌
func setWenChang(birthHour *dizhi.DiZhi, blocks []*Block) []*Block {
	index := (11 - int(*birthHour) + 11) % 12
	blocks[index].Stars = append(blocks[index].Stars, &Star{
		Name:     stars.WenChang.String(),
		StarType: startype.ShiXiZhuXing,
	})
	return blocks
}

// setWen 設定文曲
func setWenQu(birthHour *dizhi.DiZhi, blocks []*Block) []*Block {
	index := (int(*birthHour) + 4) % 12
	blocks[index].Stars = append(blocks[index].Stars, &Star{
		Name:     stars.WenQu.String(),
		StarType: startype.ShiXiZhuXing,
	})
	return blocks
}

// setDiJie 設定地劫
func setDiJie(birthHour *dizhi.DiZhi, blocks []*Block) []*Block {
	index := (int(*birthHour) + 11) % 12
	blocks[index].Stars = append(blocks[index].Stars, &Star{
		Name:     stars.DiJie.String(),
		StarType: startype.ShiXiZhuXing,
	})
	return blocks
}

// setDiKong 設定地空
func setDiKong(birthHour *dizhi.DiZhi, blocks []*Block) []*Block {
	index := (11 - int(*birthHour) + 12) % 12
	blocks[index].Stars = append(blocks[index].Stars, &Star{
		Name:     stars.DiKong.String(),
		StarType: startype.ShiXiZhuXing,
	})
	return blocks
}

// setTaiFu 設定台輔
func setTaiFu(birthHour *dizhi.DiZhi, blocks []*Block) []*Block {
	index := (int(*birthHour) + 6) % 12
	blocks[index].Stars = append(blocks[index].Stars, &Star{
		Name:     stars.TaiFu.String(),
		StarType: startype.ShiXiZhuXing,
	})
	return blocks
}

// setFengGao 設定封誥
func setFengGao(birthHour *dizhi.DiZhi, blocks []*Block) []*Block {
	index := (int(*birthHour) + 2) % 12
	blocks[index].Stars = append(blocks[index].Stars, &Star{
		Name:     stars.FengGao.String(),
		StarType: startype.ShiXiZhuXing,
	})
	return blocks
}
