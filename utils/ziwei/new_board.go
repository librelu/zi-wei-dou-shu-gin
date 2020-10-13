package ziwei

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/zi-wei-dou-shu-gin/utils/lunacal"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/dizhi"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/genders"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/gong"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/miaoxian"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/mingju"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/stars"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/startype"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/tiangan"
)

func NewBoard(birthday time.Time, gender genders.Gender) (*Board, error) {
	board := new(Board)
	board.StarsMap = make(map[stars.StarName]int)
	lunaDate := lunacal.Solar2Lunar(birthday)
	board.setupDiZhi()
	board.setYinShou(&lunaDate.Year.TianGan)
	board.setupGongWei(lunaDate)
	mingGongLocation := getMingGong(lunaDate.Hour, lunaDate.Month)
	board.MingJu = getMingJu(mingGongLocation, lunaDate.Year.TianGan)
	err := board.setFourteenMainStars(mingGongLocation, board.MingJu, lunaDate.Day)
	if err != nil {
		return nil, fmt.Errorf("failed in set fourteen main stars, error: %w", err)
	}
	board.setUpNainGanStars(&lunaDate.Year.TianGan)
	board.setXunKong(lunaDate.Year)
	board.setJieKong(&lunaDate.Year.TianGan)
	shenGongLocation := getShengGong(lunaDate.Hour, lunaDate.Month)
	board.setNianZhiXiZhuXing(&lunaDate.Year.DiZhi, mingGongLocation, shenGongLocation)
	board.setYueXiXing(int(lunaDate.Month))
	board.setShiXiZhuXing(&lunaDate.Year.DiZhi, int(lunaDate.Month), int(lunaDate.Day), lunaDate.Hour)
	err = board.setMingZhu(&lunaDate.Year.DiZhi)
	if err != nil {
		return nil, fmt.Errorf("failed in set ming zhu, error: %w", err)
	}
	err = board.setShenZhu(&lunaDate.Year.DiZhi)
	if err != nil {
		return nil, fmt.Errorf("failed in set shen zhu, error: %w", err)
	}
	err = board.setGender(&lunaDate.Year.TianGan, gender)
	if err != nil {
		return nil, fmt.Errorf("failed in set gender: %w", err)
	}

	luCunLocation, ok := board.StarsMap[stars.LuCun]
	if !ok {
		return nil, fmt.Errorf("failed to get lu Cun location")
	}
	board.setBoShiTwelveStars(luCunLocation)
	currentLunaDate := lunacal.Solar2Lunar(time.Now())
	board.setLiuNianSuiQianZhuXing(&currentLunaDate.Year.DiZhi)
	err = board.setSiHua(&lunaDate.Year.TianGan)
	if err != nil {
		return nil, fmt.Errorf("failed in set si hua: %w", err)
	}
	return board, nil
}

func (b *Board) setupDiZhi() {
	b.Blocks = make([]*Block, defaultBoardBlock)
	for i := range b.Blocks {
		b.Blocks[i] = &Block{
			Location: &Location{
				DiZhi: dizhi.DiZhi(i),
			},
		}
	}
	return
}

func (b *Board) setYinShou(birthYear *tiangan.TianGan) {
	yinShou := yinShouMap[*birthYear]
	for i := range b.Blocks {
		index := i + int(yinShou) - 2
		if index < 0 {
			index = index + 2
		} else if index > 9 {
			index -= 10
		}
		b.Blocks[i].Location.TianGan = tiangan.TianGan(index)
	}
	return
}

func (b *Board) setupGongWei(lunaDate *lunacal.LunaDate) {
	hour := lunaDate.Hour
	month := lunaDate.Month
	mingGongLocation := getMingGong(hour, month)
	shengGongLocation := getShengGong(hour, month)
	b.setTwelveGongs(mingGongLocation)
	_ = shengGongLocation

	return
}

func (b *Board) setTwelveGongs(mingGongLocation *dizhi.DiZhi) {
	for i := range b.Blocks {
		index := int(*mingGongLocation) + i
		if index > 11 {
			index -= 12
		}
		b.Blocks[index].GongWeiName = gong.Gong(i)
	}
	return
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

func (b *Board) setFourteenMainStars(mingGongLocation *dizhi.DiZhi, mingJu *MingJu, birthdate uint) error {
	// 設定紫微
	ziWeiStarIndex := getZiWeiStarLocation(mingGongLocation, mingJu, birthdate)
	b.Blocks[ziWeiStarIndex].Stars = append(b.Blocks[ziWeiStarIndex].Stars, &Star{
		Name:     stars.ZiWei.String(),
		StarType: startype.FourteenMainStars,
		MiaoXian: getZiWeiMiaoXian(ziWeiStarIndex),
	})
	b.StarsMap[stars.ZiWei] = ziWeiStarIndex

	// 設定天府
	tianFuIndex, err := b.setTianFuStarLocation(ziWeiStarIndex)
	if err == nil {
		b.Blocks[tianFuIndex].Stars = append(b.Blocks[tianFuIndex].Stars, &Star{
			Name:     stars.TianFu.String(),
			StarType: startype.FourteenMainStars,
			MiaoXian: getTianFuMiaoXian(tianFuIndex),
		})
		b.StarsMap[stars.TianFu] = tianFuIndex
	} else {
		return fmt.Errorf("tian fu star not found, error: %w", err)
	}

	b.setStarsBeggingWithZiWei(ziWeiStarIndex)
	b.setStarsBeggingWithTianFu(tianFuIndex)
	return nil
}

// setSiHua 安四化
func (b *Board) setSiHua(birthYear *tiangan.TianGan) error {
	if err := b.setHuaLu(birthYear); err != nil {
		return err
	}
	if err := b.setHuaQuan(birthYear); err != nil {
		return err
	}
	if err := b.setHuaKe(birthYear); err != nil {
		return err
	}
	if err := b.setHuaJi(birthYear); err != nil {
		return err
	}
	return nil
}

// setHuaLu 設定化祿
func (b *Board) setHuaLu(birthYear *tiangan.TianGan) error {
	starMap := []stars.StarName{
		stars.LianZhen,
		stars.TianJi,
		stars.TianTong,
		stars.TaiYin,
		stars.TanLang,
		stars.WuQu,
		stars.TaiYang,
		stars.JuMen,
		stars.TianLiang,
		stars.PoJun,
	}
	index, ok := b.StarsMap[starMap[*birthYear]]
	if !ok {
		return fmt.Errorf("current star not found, current birth year: %d", birthYear)
	}
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.HuaLu.String(),
		StarType: startype.SiHua,
	})
	return nil
}

// setHuaQuan 設定化權
func (b *Board) setHuaQuan(birthYear *tiangan.TianGan) error {
	starMap := []stars.StarName{
		stars.PoJun,
		stars.TianLiang,
		stars.TianJi,
		stars.TianTong,
		stars.TaiYin,
		stars.TanLang,
		stars.WuQu,
		stars.TaiYang,
		stars.ZiWei,
		stars.JuMen,
	}
	index, ok := b.StarsMap[starMap[*birthYear]]
	if !ok {
		return fmt.Errorf("current star not found, current birth year: %d", birthYear)
	}
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.HuaQuan.String(),
		StarType: startype.SiHua,
	})
	return nil
}

// setHuaKe 設定化科
func (b *Board) setHuaKe(birthYear *tiangan.TianGan) error {
	starMap := []stars.StarName{
		stars.WuQu,
		stars.ZiWei,
		stars.WenChang,
		stars.TianJi,
		stars.TaiYang,
		stars.TianLiang,
		stars.TianFu,
		stars.WenQu,
		stars.TianFu,
		stars.TaiYin,
	}
	index, ok := b.StarsMap[starMap[*birthYear]]
	if !ok {
		return fmt.Errorf("current star not found, current birth year: %d", birthYear)
	}
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.HuaKe.String(),
		StarType: startype.SiHua,
	})
	return nil
}

// setHuaJi 設定化忌
func (b *Board) setHuaJi(birthYear *tiangan.TianGan) error {
	starMap := []stars.StarName{
		stars.TaiYang,
		stars.TaiYin,
		stars.LianZhen,
		stars.JuMen,
		stars.TianJi,
		stars.WenQu,
		stars.TianTong,
		stars.WenChang,
		stars.WuQu,
		stars.TanLang,
	}
	index, ok := b.StarsMap[starMap[*birthYear]]
	if !ok {
		return fmt.Errorf("current star not found, current birth year: %d", birthYear)
	}
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.HuaJi.String(),
		StarType: startype.SiHua,
	})
	return nil
}

//setStarsBeggingWithZiWei 順時針一宮安天機星，跳隔一宮，安太陽星，順時針一宮安武曲星，順時針一宮安天同星，跳隔兩宮，安廉貞星
func (b *Board) setStarsBeggingWithZiWei(ziWeiStarIndex int) {
	// 設定天機
	tianJi := ziWeiStarIndex - 1
	if tianJi < 0 {
		tianJi = 12 + tianJi
	}
	b.Blocks[tianJi].Stars = append(b.Blocks[tianJi].Stars, &Star{
		Name:     stars.TianJi.String(),
		StarType: startype.FourteenMainStars,
		MiaoXian: getTianJiMiaoXian(tianJi),
	})
	b.StarsMap[stars.TianJi] = tianJi

	//　設定太陽
	taiYang := tianJi - 2
	if taiYang < 0 {
		taiYang = 12 + taiYang
	}
	b.Blocks[taiYang].Stars = append(b.Blocks[taiYang].Stars, &Star{
		Name:     stars.TaiYang.String(),
		StarType: startype.FourteenMainStars,
		MiaoXian: getTaiYangMiaoXian(taiYang),
	})
	b.StarsMap[stars.TaiYang] = taiYang

	//　設定武曲
	wuQu := taiYang - 1
	if wuQu < 0 {
		wuQu = 12 + wuQu
	}
	b.Blocks[wuQu].Stars = append(b.Blocks[wuQu].Stars, &Star{
		Name:     stars.WuQu.String(),
		StarType: startype.FourteenMainStars,
		MiaoXian: getWuQuMiaoXian(wuQu),
	})
	b.StarsMap[stars.WuQu] = wuQu

	// 設定天同
	tianTong := wuQu - 1
	if tianTong < 0 {
		tianTong = 12 + tianTong
	}
	b.Blocks[tianTong].Stars = append(b.Blocks[tianTong].Stars, &Star{
		Name:     stars.TianTong.String(),
		StarType: startype.FourteenMainStars,
		MiaoXian: getTianTongMiaoXian(tianTong),
	})
	b.StarsMap[stars.TianTong] = tianTong

	// 設定廉貞
	lianZhen := tianTong - 3
	if lianZhen < 0 {
		lianZhen = 12 + lianZhen
	}
	b.Blocks[lianZhen].Stars = append(b.Blocks[lianZhen].Stars, &Star{
		Name:     stars.LianZhen.String(),
		StarType: startype.FourteenMainStars,
		MiaoXian: getLianZhenMiaoXian(lianZhen),
	})
	b.StarsMap[stars.LianZhen] = lianZhen

	return
}

//setStarsBeggingWithTianFu 逆時針一宮安太陰星，逆時針一宮安貪狼星，逆時針一宮安巨門星，逆時針一宮安天相星，逆時針一宮安天梁星，逆時針一宮安七殺星，跳隔三宮，安破軍星
func (b *Board) setStarsBeggingWithTianFu(tianFuIndex int) {
	// 設定太陰
	taiYin := tianFuIndex + 1
	if taiYin > 11 {
		taiYin = taiYin - 12
	}
	b.Blocks[taiYin].Stars = append(b.Blocks[taiYin].Stars, &Star{
		Name:     stars.TaiYin.String(),
		StarType: startype.FourteenMainStars,
		MiaoXian: getTaiYinMiaoXian(taiYin),
	})
	b.StarsMap[stars.TaiYin] = taiYin

	// 設定貪狼
	tanLang := taiYin + 1
	if tanLang > 11 {
		tanLang = tanLang - 12
	}
	b.Blocks[tanLang].Stars = append(b.Blocks[tanLang].Stars, &Star{
		Name:     stars.TanLang.String(),
		StarType: startype.FourteenMainStars,
		MiaoXian: getTanLangMiaoXian(tanLang),
	})
	b.StarsMap[stars.TanLang] = tanLang

	//　設定巨門
	juMen := tanLang + 1
	if juMen > 11 {
		juMen = juMen - 12
	}
	b.Blocks[juMen].Stars = append(b.Blocks[juMen].Stars, &Star{
		Name:     stars.JuMen.String(),
		StarType: startype.FourteenMainStars,
		MiaoXian: getJuMenMiaoXian(juMen),
	})
	b.StarsMap[stars.JuMen] = juMen

	// 設定天相
	tianXiang := juMen + 1
	if tianXiang > 11 {
		tianXiang = tianXiang - 12
	}
	b.Blocks[tianXiang].Stars = append(b.Blocks[tianXiang].Stars, &Star{
		Name:     stars.TianXiang.String(),
		StarType: startype.FourteenMainStars,
		MiaoXian: getTianXiangXian(tianXiang),
	})
	b.StarsMap[stars.TianXiang] = tianXiang

	// 設定天梁
	tianLiang := tianXiang + 1
	if tianLiang > 11 {
		tianLiang = tianLiang - 12
	}
	b.Blocks[tianLiang].Stars = append(b.Blocks[tianLiang].Stars, &Star{
		Name:     stars.TianLiang.String(),
		StarType: startype.FourteenMainStars,
		MiaoXian: getTianLiangXian(tianLiang),
	})
	b.StarsMap[stars.TianLiang] = tianLiang

	// 設定七殺
	qiSha := tianLiang + 1
	if qiSha > 11 {
		qiSha = qiSha - 12
	}
	b.Blocks[qiSha].Stars = append(b.Blocks[qiSha].Stars, &Star{
		Name:     stars.QiSha.String(),
		StarType: startype.FourteenMainStars,
	})
	b.StarsMap[stars.QiSha] = qiSha

	// 設定破軍
	poJun := qiSha + 4
	if poJun > 11 {
		poJun = poJun - 12
	}
	b.Blocks[poJun].Stars = append(b.Blocks[poJun].Stars, &Star{
		Name:     stars.PoJun.String(),
		StarType: startype.FourteenMainStars,
	})
	b.StarsMap[stars.PoJun] = poJun
	return
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

func (b *Board) setTianFuStarLocation(ziWeiStarIndex int) (int, error) {
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
func (b *Board) setUpNainGanStars(tianGan *tiangan.TianGan) {
	b.setLuCun(tianGan)
	b.setQingYang(tianGan)
	b.setTuoLuo(tianGan)
	b.setTianKui(tianGan)
	b.setTianYue(tianGan)
	b.setTianGuan(tianGan)
	b.setTianFu(tianGan)
}

// setLuCun 設定祿存
func (b *Board) setLuCun(tianGan *tiangan.TianGan) {
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
	b.Blocks[luCunLocation].Stars = append(b.Blocks[luCunLocation].Stars, &Star{
		Name:     stars.LuCun.String(),
		StarType: startype.NianGanXiZhuXing,
	})
	b.StarsMap[stars.LuCun] = int(luCunLocation)
	return
}

// setQingYang 設定擎羊
func (b *Board) setQingYang(tianGan *tiangan.TianGan) {
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
	b.Blocks[qingYangLocation].Stars = append(b.Blocks[qingYangLocation].Stars, &Star{
		Name:     stars.QingYang.String(),
		StarType: startype.NianGanXiZhuXing,
	})
	b.StarsMap[stars.QingYang] = int(qingYangLocation)
	return
}

// setTuoLuo 設定陀羅
func (b *Board) setTuoLuo(tianGan *tiangan.TianGan) {
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
	b.Blocks[tuoLuoLocation].Stars = append(b.Blocks[tuoLuoLocation].Stars, &Star{
		Name:     stars.TuoLuo.String(),
		StarType: startype.NianGanXiZhuXing,
	})
	b.StarsMap[stars.TuoLuo] = int(tuoLuoLocation)
	return
}

// setTianKui 設天魁
func (b *Board) setTianKui(tianGan *tiangan.TianGan) {
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
	b.Blocks[tianKuiLocation].Stars = append(b.Blocks[tianKuiLocation].Stars, &Star{
		Name:     stars.TianKui.String(),
		StarType: startype.NianGanXiZhuXing,
	})
	b.StarsMap[stars.TianKui] = int(tianKuiLocation)
	return
}

// setTianYue 設定天鉞
func (b *Board) setTianYue(tianGan *tiangan.TianGan) {
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
	b.Blocks[tianYueLocation].Stars = append(b.Blocks[tianYueLocation].Stars, &Star{
		Name:     stars.TianYue.String(),
		StarType: startype.NianGanXiZhuXing,
	})
	b.StarsMap[stars.TianYue] = int(tianYueLocation)
	return
}

// setTianGuan 設定天官
func (b *Board) setTianGuan(tianGan *tiangan.TianGan) {
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
	b.Blocks[tianGuanLocation].Stars = append(b.Blocks[tianGuanLocation].Stars, &Star{
		Name:     stars.TianGuan.String(),
		StarType: startype.NianGanXiZhuXing,
	})
	return
}

// setTianGuan 設定天福
func (b *Board) setTianFu(tianGan *tiangan.TianGan) {
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
	b.Blocks[tianFuLocation].Stars = append(b.Blocks[tianFuLocation].Stars, &Star{
		Name:     stars.NainGanTianFu.String(),
		StarType: startype.NianGanXiZhuXing,
	})
	return
}

// setXunKong 設定旬空星
func (b *Board) setXunKong(birthYear *lunacal.TianGanDiZhi) {
	xunKongMapXAxis := 6 - int(birthYear.DiZhi/2) - 1
	xunKongIndex := (xunKongMapXAxis*10 + int(birthYear.TianGan)) % 12
	b.Blocks[xunKongIndex].Stars = append(b.Blocks[xunKongIndex].Stars, &Star{
		Name:     stars.XunKong.String(),
		StarType: startype.NianGanXiZhuXing,
	})

	return
}

// setJieKong 設定截空星
func (b *Board) setJieKong(birthYear *tiangan.TianGan) {
	index := ((4 - int(*birthYear)%5 + 1) * 2) - 1
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.JieKong.String(),
		StarType: startype.NianGanXiZhuXing,
	})
	b.Blocks[index-1].Stars = append(b.Blocks[index-1].Stars, &Star{
		Name:     stars.JieKong.String(),
		StarType: startype.NianGanXiZhuXing,
	})
	return
}

// NianZhiXiZhuXing 設定年支系諸星
func (b *Board) setNianZhiXiZhuXing(birthYear *dizhi.DiZhi, mingGongLocation *dizhi.DiZhi, shengGongLocation *dizhi.DiZhi) {
	b.setTianKu(birthYear)
	b.setTianXu(birthYear)
	b.setLongChi(birthYear)
	b.setFengGe(birthYear)
	b.setHongLuan(birthYear)
	b.setTianXi(birthYear)
	b.setGuChen(birthYear)
	b.setGuaXiu(birthYear)
	b.setJieShen(birthYear)
	b.setPoSui(birthYear)
	b.setTianMa(birthYear)
	b.setDaHao(birthYear)
	b.setTianDe(birthYear)
	b.setJieSha(birthYear)
	b.setHuaGai(birthYear)
	b.setXianChi(birthYear)
	b.setTianCai(birthYear, mingGongLocation)
	b.setTianShou(birthYear, shengGongLocation)
	b.setTianKong(birthYear)
	return
}

// setTianKu 設定天哭
func (b *Board) setTianKu(birthYear *dizhi.DiZhi) {
	index := (11 - int(*birthYear) + 7) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.TianKu.String(),
		StarType: startype.NianZhiXiZhuXing,
	})
	b.StarsMap[stars.TianKu] = index
	return
}

// setTianXu 設定天虛
func (b *Board) setTianXu(birthYear *dizhi.DiZhi) {
	index := (int(*birthYear) + 6) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.TianXu.String(),
		StarType: startype.NianZhiXiZhuXing,
	})
	b.StarsMap[stars.TianXu] = index
	return
}

// setLongChi 設定龍池
func (b *Board) setLongChi(birthYear *dizhi.DiZhi) {
	index := (int(*birthYear) + 4) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.LongChi.String(),
		StarType: startype.NianZhiXiZhuXing,
	})
	return
}

// setFengGe 設定鳳閣
func (b *Board) setFengGe(birthYear *dizhi.DiZhi) {
	index := (11 - int(*birthYear) + 11) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.FengGe.String(),
		StarType: startype.NianZhiXiZhuXing,
	})
	return
}

// setHongLuan 設定紅鸞
func (b *Board) setHongLuan(birthYear *dizhi.DiZhi) {
	index := (11 - int(*birthYear) + 4) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.HongLuan.String(),
		StarType: startype.NianZhiXiZhuXing,
	})
	b.StarsMap[stars.HongLuan] = index
	return
}

// setTianXi 設定天喜
func (b *Board) setTianXi(birthYear *dizhi.DiZhi) {
	index := (11 - int(*birthYear) + 10) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.TianXi.String(),
		StarType: startype.NianZhiXiZhuXing,
	})
	b.StarsMap[stars.TianXi] = index
	return
}

// setGuChen 設定孤辰
func (b *Board) setGuChen(birthYear *dizhi.DiZhi) {
	locationIndex := (int(*birthYear) + 1) / 3 % 4
	locations := []dizhi.DiZhi{dizhi.Yin, dizhi.Si, dizhi.Shen, dizhi.Hai}
	index := locations[locationIndex]
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.GuChen.String(),
		StarType: startype.NianZhiXiZhuXing,
	})
	return
}

// setGuaXiu　設定寡宿
func (b *Board) setGuaXiu(birthYear *dizhi.DiZhi) {
	locationIndex := (int(*birthYear) + 1) / 3 % 4
	locations := []dizhi.DiZhi{dizhi.Xu, dizhi.Chou, dizhi.Chen, dizhi.Wei}
	index := locations[locationIndex]
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.GuaXiu.String(),
		StarType: startype.NianZhiXiZhuXing,
	})
	return
}

// setJieShen 設定解神
func (b *Board) setJieShen(birthYear *dizhi.DiZhi) {
	index := (11 - int(*birthYear) + 11) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.JieShen.String(),
		StarType: startype.NianZhiXiZhuXing,
	})
	return
}

// setPoSui 設定破碎
func (b *Board) setPoSui(birthYear *dizhi.DiZhi) {
	locationIndex := int(int(*birthYear) % 3)
	locations := []dizhi.DiZhi{dizhi.Si, dizhi.Chou, dizhi.You}
	b.Blocks[locations[locationIndex]].Stars = append(
		b.Blocks[locations[locationIndex]].Stars, &Star{
			Name:     stars.PoSui.String(),
			StarType: startype.NianZhiXiZhuXing,
		})
	return
}

// setPoSui 設定天馬
func (b *Board) setTianMa(birthYear *dizhi.DiZhi) {
	locationIndex := int((int(*birthYear) + 1) % 4)
	locations := []dizhi.DiZhi{dizhi.Si, dizhi.Yin, dizhi.Hai, dizhi.Shen}
	blockIndex := locations[locationIndex]
	b.Blocks[blockIndex].Stars = append(
		b.Blocks[blockIndex].Stars, &Star{
			Name:     stars.TianMa.String(),
			StarType: startype.NianZhiXiZhuXing,
		})
	b.StarsMap[stars.TianMa] = int(blockIndex)
	return
}

// setDaHao 設定大耗
func (b *Board) setDaHao(birthYear *dizhi.DiZhi) {
	locationMap := []dizhi.DiZhi{
		dizhi.Wei, dizhi.Wu, dizhi.You, dizhi.Shen, dizhi.Hai, dizhi.Xu, dizhi.Chou, dizhi.Zi, dizhi.Mao, dizhi.Yin, dizhi.Si, dizhi.Chen,
	}
	index := locationMap[*birthYear]
	b.Blocks[index].Stars = append(
		b.Blocks[index].Stars, &Star{
			Name:     stars.DaHao.String(),
			StarType: startype.NianZhiXiZhuXing,
		})
	return
}

// setTianDe 設定天德
func (b *Board) setTianDe(birthYear *dizhi.DiZhi) {
	index := (int(*birthYear) + 9) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.TianDe.String(),
		StarType: startype.NianZhiXiZhuXing,
	})
	return
}

// setJieSha 設定劫殺
func (b *Board) setJieSha(birthYear *dizhi.DiZhi) {
	locationIndex := int(int(*birthYear) % 4)
	locations := []dizhi.DiZhi{dizhi.Si, dizhi.Yin, dizhi.Hai, dizhi.Shen}
	b.Blocks[locations[locationIndex]].Stars = append(
		b.Blocks[locations[locationIndex]].Stars, &Star{
			Name:     stars.JieSha.String(),
			StarType: startype.NianZhiXiZhuXing,
		})
	return
}

// setHuaGai 設定華蓋
func (b *Board) setHuaGai(birthYear *dizhi.DiZhi) {
	locationIndex := int(int(*birthYear) % 4)
	locations := []dizhi.DiZhi{dizhi.Chen, dizhi.Chou, dizhi.Xu, dizhi.Wei}
	b.Blocks[locations[locationIndex]].Stars = append(
		b.Blocks[locations[locationIndex]].Stars, &Star{
			Name:     stars.HuaGai.String(),
			StarType: startype.NianZhiXiZhuXing,
		})
	return
}

// setXianChi 設定咸池
func (b *Board) setXianChi(birthYear *dizhi.DiZhi) {
	locationIndex := int(int(*birthYear) % 4)
	locations := []dizhi.DiZhi{dizhi.You, dizhi.Wu, dizhi.Mao, dizhi.Zi}
	b.Blocks[locations[locationIndex]].Stars = append(
		b.Blocks[locations[locationIndex]].Stars, &Star{
			Name:     stars.XianChi.String(),
			StarType: startype.NianZhiXiZhuXing,
		})
	return
}

// setTianCai 設定天才
func (b *Board) setTianCai(birthYear *dizhi.DiZhi, mingGongLocation *dizhi.DiZhi) {
	index := (*mingGongLocation + *birthYear) % 12
	b.Blocks[index].Stars = append(
		b.Blocks[index].Stars, &Star{
			Name:     stars.TianCai.String(),
			StarType: startype.NianZhiXiZhuXing,
		})
	return
}

// setTianShou 設定天壽
func (b *Board) setTianShou(birthYear *dizhi.DiZhi, shengGongLocation *dizhi.DiZhi) {
	index := (*shengGongLocation + *birthYear) % 12
	b.Blocks[index].Stars = append(
		b.Blocks[index].Stars, &Star{
			Name:     stars.TianShou.String(),
			StarType: startype.NianZhiXiZhuXing,
		})
	return
}

// setTianKong 設定天空
func (b *Board) setTianKong(birthYear *dizhi.DiZhi) {
	index := (int(*birthYear) + 1) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.TianKong.String(),
		StarType: startype.NianZhiXiZhuXing,
	})
	return
}

// setYueXiXing 安月系星
func (b *Board) setYueXiXing(birthMonth int) {
	b.setZuoFu(birthMonth)
	b.setYouBi(birthMonth)
	b.setTianXing(birthMonth)
	b.setTianYao(birthMonth)
	b.setTianWu(birthMonth)
	b.setYueXiXingTianYue(birthMonth)
	b.setYinSha(birthMonth)
	return
}

// setZuoFu 設定左輔
func (b *Board) setZuoFu(birthMonth int) {
	index := getZuoFuLocation(birthMonth)
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.ZuoFu.String(),
		StarType: startype.YueXiXing,
	})
	b.StarsMap[stars.ZuoFu] = index
	return
}

func getZuoFuLocation(birthMonth int) int {
	return (birthMonth + 3) % 12
}

// setYouBi 設定右弼
func (b *Board) setYouBi(birthMonth int) {
	index := getYouBiLocation(birthMonth)
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.YouBi.String(),
		StarType: startype.YueXiXing,
	})
	b.StarsMap[stars.YouBi] = index
	return
}

func getYouBiLocation(birthMonth int) int {
	return ((13 - birthMonth) + 10) % 12
}

// setTianXing 設定天刑
func (b *Board) setTianXing(birthMonth int) {
	index := (birthMonth + 8) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.TianXing.String(),
		StarType: startype.YueXiXing,
	})
	b.StarsMap[stars.TianXing] = index
	return
}

// setTianYao 設定天姚
func (b *Board) setTianYao(birthMonth int) {
	index := (birthMonth) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.TianYao.String(),
		StarType: startype.YueXiXing,
	})
	b.StarsMap[stars.TianYao] = index
	return
}

// setTianWu 設定天巫
func (b *Board) setTianWu(birthMonth int) {
	locationIndex := (birthMonth - 1) % 4
	locations := []dizhi.DiZhi{dizhi.Si, dizhi.Shen, dizhi.Yin, dizhi.Hai}
	b.Blocks[locations[locationIndex]].Stars = append(
		b.Blocks[locations[locationIndex]].Stars, &Star{
			Name:     stars.TianWu.String(),
			StarType: startype.YueXiXing,
		})
	return
}

// setYueXiXingTianYue 設定天月
// TODO: refactor Yue Xi Xing to a object
func (b *Board) setYueXiXingTianYue(birthMonth int) {
	locationMap := []dizhi.DiZhi{
		dizhi.Xu, dizhi.Si, dizhi.Chen, dizhi.Yin, dizhi.Wei, dizhi.Mao, dizhi.Hai, dizhi.Wei, dizhi.Yin, dizhi.Wu, dizhi.Xu, dizhi.Yin,
	}
	index := locationMap[birthMonth-1]
	b.Blocks[index].Stars = append(
		b.Blocks[index].Stars, &Star{
			Name:     stars.YueXiXingTianYue.String(),
			StarType: startype.YueXiXing,
		})
	return
}

// setYinSha 設定陰煞
func (b *Board) setYinSha(birthMonth int) {
	locationIndex := (birthMonth + 2) % 6
	locations := []dizhi.DiZhi{dizhi.Shen, dizhi.Wu, dizhi.Chen, dizhi.Yin, dizhi.Zi, dizhi.Xu}
	b.Blocks[locations[locationIndex]].Stars = append(
		b.Blocks[locations[locationIndex]].Stars, &Star{
			Name:     stars.YinSha.String(),
			StarType: startype.YueXiXing,
		})
	return
}

// setShiXiZhuXing 安時系諸星
func (b *Board) setShiXiZhuXing(birthYear *dizhi.DiZhi, birthMonth int, birthDate int, birthHour *dizhi.DiZhi) {
	b.setWenChang(birthHour)
	b.setWenQu(birthHour)
	b.setDiJie(birthHour)
	b.setDiKong(birthHour)
	b.setTaiFu(birthHour)
	b.setFengGao(birthHour)
	b.setHuo(birthYear, birthHour)
	b.setLing(birthYear, birthHour)
	zuoFuLocation := getZuoFuLocation(birthMonth)
	b.setSanTai(zuoFuLocation, birthDate)
	youBiLocation := getYouBiLocation(birthMonth)
	b.setBaZuo(youBiLocation, birthDate)
	wenChangLocation := getWenChangLocation(birthHour)
	b.setEnGuang(wenChangLocation, birthDate)
	wenQuLocation := getWenQuLocation(birthHour)
	b.setTianGui(wenQuLocation, birthDate)
	return
}

// setWenChang 設定文昌
func (b *Board) setWenChang(birthHour *dizhi.DiZhi) {
	index := getWenChangLocation(birthHour)
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.WenChang.String(),
		StarType: startype.ShiXiZhuXing,
	})
	b.StarsMap[stars.WenChang] = index
	return
}

func getWenChangLocation(birthHour *dizhi.DiZhi) int {
	return (11 - int(*birthHour) + 11) % 12
}

// setWen 設定文曲
func (b *Board) setWenQu(birthHour *dizhi.DiZhi) {
	index := getWenQuLocation(birthHour)
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.WenQu.String(),
		StarType: startype.ShiXiZhuXing,
	})
	b.StarsMap[stars.WenQu] = index
	return
}

func getWenQuLocation(birthHour *dizhi.DiZhi) int {
	return (int(*birthHour) + 4) % 12
}

// setDiJie 設定地劫
func (b *Board) setDiJie(birthHour *dizhi.DiZhi) {
	index := (int(*birthHour) + 11) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.DiJie.String(),
		StarType: startype.ShiXiZhuXing,
	})
	b.StarsMap[stars.DiJie] = index
	return
}

// setDiKong 設定地空
func (b *Board) setDiKong(birthHour *dizhi.DiZhi) {
	index := (11 - int(*birthHour) + 12) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.DiKong.String(),
		StarType: startype.ShiXiZhuXing,
	})
	b.StarsMap[stars.DiKong] = index
	return
}

// setTaiFu 設定台輔
func (b *Board) setTaiFu(birthHour *dizhi.DiZhi) {
	index := (int(*birthHour) + 6) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.TaiFu.String(),
		StarType: startype.ShiXiZhuXing,
	})
	return
}

// setFengGao 設定封誥
func (b *Board) setFengGao(birthHour *dizhi.DiZhi) {
	index := (int(*birthHour) + 2) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.FengGao.String(),
		StarType: startype.ShiXiZhuXing,
	})
	return
}

// setHuo　設定火星
func (b *Board) setHuo(birthYear *dizhi.DiZhi, birthHour *dizhi.DiZhi) {
	xaxis := getHuoLingGroupMap(birthYear)
	var index int
	switch xaxis {
	case 0:
		index = (int(*birthHour) + 1) % 12
	case 1:
		index = (int(*birthHour) + 2) % 12
	case 2:
		index = (int(*birthHour) + 3) % 12
	case 3:
		index = (int(*birthHour) + 9) % 12
	}
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.Huo.String(),
		StarType: startype.ShiXiZhuXing,
	})
	b.StarsMap[stars.Huo] = index
	return
}

// setHuo　設定鈴星
func (b *Board) setLing(birthYear *dizhi.DiZhi, birthHour *dizhi.DiZhi) {
	xaxis := getHuoLingGroupMap(birthYear)
	var index int
	switch xaxis {
	case 0:
		index = (int(*birthHour) + 3) % 12
	default:
		index = (int(*birthHour) + 10) % 12
	}
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.Ling.String(),
		StarType: startype.ShiXiZhuXing,
	})
	b.StarsMap[stars.Ling] = index
	return
}

func getHuoLingGroupMap(birthYear *dizhi.DiZhi) int {
	birthYearGroup := map[dizhi.DiZhi]int{
		dizhi.Yin:  0,
		dizhi.Wu:   0,
		dizhi.Xu:   0,
		dizhi.Shen: 1,
		dizhi.Zi:   1,
		dizhi.Chen: 1,
		dizhi.Si:   2,
		dizhi.You:  2,
		dizhi.Chou: 2,
		dizhi.Hai:  3,
		dizhi.Mao:  3,
		dizhi.Wei:  3,
	}

	return birthYearGroup[*birthYear]
}

// setSanTai 設定三台
func (b *Board) setSanTai(zuoFuLocation int, birthDate int) {
	index := (zuoFuLocation + birthDate - 1) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.SanTai.String(),
		StarType: startype.ShiXiZhuXing,
	})
	return
}

// setBaZuo 設定八座
func (b *Board) setBaZuo(youBiLocation int, birthDate int) {
	index := (youBiLocation - birthDate + 1) % 12
	if index < 0 {
		index += 11
	}
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.BaZuo.String(),
		StarType: startype.ShiXiZhuXing,
	})
	return
}

// setEnGuang 設定恩光
func (b *Board) setEnGuang(wenChangLocation int, birthDate int) {
	index := (wenChangLocation + birthDate - 2) % 12
	if index < 0 {
		index += 11
	}
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.EnGuang.String(),
		StarType: startype.ShiXiZhuXing,
	})
	return
}

// setTianGui 設定天貴
func (b *Board) setTianGui(wenQuLocation int, birthDate int) {
	index := (wenQuLocation + birthDate - 2) % 12
	if index < 0 {
		index += 11
	}
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.TianGui.String(),
		StarType: startype.ShiXiZhuXing,
	})
	return
}

// setMingZhu 設定命主
func (b *Board) setMingZhu(birthYear *dizhi.DiZhi) error {
	starMap := []stars.StarName{
		stars.TanLang,
		stars.JuMen,
		stars.LuCun,
		stars.WenQu,
		stars.LianZhen,
		stars.WuQu,
		stars.PoJun,
		stars.WuQu,
		stars.LianZhen,
		stars.WenQu,
		stars.LuCun,
		stars.JuMen,
	}
	starName := starMap[*birthYear]
	index, ok := b.StarsMap[starName]
	if !ok {
		return fmt.Errorf("the star name not found in birth year = %d", birthYear)
	}
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.MingZhu.String(),
		StarType: startype.ShenMing,
	})
	return nil
}

// setShenZhu 設定身主
func (b *Board) setShenZhu(birthYear *dizhi.DiZhi) error {
	starMap := []stars.StarName{
		stars.Ling,
		stars.TianXiang,
		stars.TianLiang,
		stars.TianTong,
		stars.WenChang,
		stars.TianJi,
		stars.Huo,
		stars.TianXiang,
		stars.TianLiang,
		stars.TianTong,
		stars.WenChang,
		stars.TianJi,
	}
	starName := starMap[*birthYear]
	index, ok := b.StarsMap[starName]
	if !ok {
		return fmt.Errorf("the star name not found in birth year = %d", birthYear)
	}
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &Star{
		Name:     stars.ShenZhu.String(),
		StarType: startype.ShenMing,
	})
	return nil
}

// setGetGenders
func (b *Board) setGender(birthYear *tiangan.TianGan, gender genders.Gender) error {
	genderIndex := int(*birthYear % 2)
	switch gender {
	case genders.Male:
		b.Gender = genders.Gender(genderIndex + 2)
	case genders.Female:
		b.Gender = genders.Gender(genderIndex + 4)
	default:
		return fmt.Errorf("please insert male or female index, example: 0 or 1, current=%d", gender)
	}
	return nil
}

// setBoShiTwelveStars 安博士十二星
func (b *Board) setBoShiTwelveStars(LuCunLocation int) {
	starsMap := []*Star{
		{
			Name:     stars.BoShi.String(),
			StarType: startype.BoShiTwelveStars,
		},
		{
			Name:     stars.LiShi.String(),
			StarType: startype.BoShiTwelveStars,
		},
		{
			Name:     stars.QingLong.String(),
			StarType: startype.BoShiTwelveStars,
		},
		{
			Name:     stars.XiaoHao.String(),
			StarType: startype.BoShiTwelveStars,
		},
		{
			Name:     stars.JiangJun.String(),
			StarType: startype.BoShiTwelveStars,
		},
		{
			Name:     stars.ZouShu.String(),
			StarType: startype.BoShiTwelveStars,
		},
		{
			Name:     stars.FeiLian.String(),
			StarType: startype.BoShiTwelveStars,
		},
		{
			Name:     stars.XiShen.String(),
			StarType: startype.BoShiTwelveStars,
		},
		{
			Name:     stars.BingFu.String(),
			StarType: startype.BoShiTwelveStars,
		},
		{
			Name:     stars.BoShiDaHao.String(),
			StarType: startype.BoShiTwelveStars,
		},
		{
			Name:     stars.FuBing.String(),
			StarType: startype.BoShiTwelveStars,
		},
		{
			Name:     stars.GuanFu.String(),
			StarType: startype.BoShiTwelveStars,
		},
	}
	for i := 0; i < 12; i++ {
		index := LuCunLocation
		switch b.Gender {
		case genders.YangMale:
			index += i
		case genders.YinMale:
			index -= i
		case genders.YangFemale:
			index -= i
		case genders.YinFemale:
			index += i
		}
		if index > 11 {
			index -= 12
		} else if index < 0 {
			index += 12
		}
		b.Blocks[index].Stars = append(b.Blocks[index].Stars, starsMap[i])
	}

	return
}

// setLiuNianSuiQianZhuXing 安流年歲前諸星
func (b *Board) setLiuNianSuiQianZhuXing(currentDiZhiYear *dizhi.DiZhi) {
	starsMap := []*Star{
		{
			Name:     stars.SuiJian.String(),
			StarType: startype.LiuNianSuiQianZhuXing,
		},
		{
			Name:     stars.HuiQi.String(),
			StarType: startype.LiuNianSuiQianZhuXing,
		},
		{
			Name:     stars.SangMen.String(),
			StarType: startype.LiuNianSuiQianZhuXing,
		},
		{
			Name:     stars.GuanSuo.String(),
			StarType: startype.LiuNianSuiQianZhuXing,
		},
		{
			Name:     stars.GuanFu.String(),
			StarType: startype.LiuNianSuiQianZhuXing,
		},
		{
			Name:     stars.LiuNianXiaoHao.String(),
			StarType: startype.LiuNianSuiQianZhuXing,
		},
		{
			Name:     stars.LiuNianDaHao.String(),
			StarType: startype.LiuNianSuiQianZhuXing,
		},
		{
			Name:     stars.LongDe.String(),
			StarType: startype.LiuNianSuiQianZhuXing,
		},
		{
			Name:     stars.BaiHu.String(),
			StarType: startype.LiuNianSuiQianZhuXing,
		},
		{
			Name:     stars.TianDe.String(),
			StarType: startype.LiuNianSuiQianZhuXing,
		},
		{
			Name:     stars.DiKe.String(),
			StarType: startype.LiuNianSuiQianZhuXing,
		},
		{
			Name:     stars.LiuNianBingFu.String(),
			StarType: startype.LiuNianSuiQianZhuXing,
		},
	}
	for i := 0; i < 12; i++ {
		index := int(*currentDiZhiYear)
		index += i
		if index > 11 {
			index -= 12
		} else if index < 0 {
			index += 12
		}
		b.Blocks[index].Stars = append(b.Blocks[index].Stars, starsMap[i])
	}
	return
}

// getZiWeiMiaoXian 得紫微廟陷
func getZiWeiMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Ping,
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Wang,
		miaoxian.Xian,
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Wang,
		miaoxian.Ping,
		miaoxian.Xian2,
		miaoxian.Wang,
	}
	return miaoXianMap[index]
}

// getTianJiMiaoXian 得天機廟陷
func getTianJiMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Miao,
		miaoxian.Xian,
		miaoxian.Wang,
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Ping,
		miaoxian.Miao,
		miaoxian.Xian,
		miaoxian.Ping,
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Ping,
	}
	return miaoXianMap[index]
}

// getTaiYangMiaoXian 得太陽廟陷
func getTaiYangMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Xian,
		miaoxian.Xian,
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Ping,
		miaoxian.Miao,
		miaoxian.Ping,
		miaoxian.Xian2,
		miaoxian.Xian2,
		miaoxian.Xian,
		miaoxian.Xian,
	}
	return miaoXianMap[index]
}

// getWuQuMiaoXian 得武曲廟陷
func getWuQuMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Xian2,
		miaoxian.Xian,
		miaoxian.Miao,
		miaoxian.Ping,
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Ping,
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Ping,
	}
	return miaoXianMap[index]
}

// getTianTongMiaoXian 得天同廟陷
func getTianTongMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Wang,
		miaoxian.Xian,
		miaoxian.Xian2,
		miaoxian.Miao,
		miaoxian.Ping,
		miaoxian.Miao,
		miaoxian.Xian,
		miaoxian.Xian,
		miaoxian.Wang,
		miaoxian.Ping,
		miaoxian.Ping,
		miaoxian.Miao,
	}
	return miaoXianMap[index]
}

// getLianZhenMiaoXian 得廉貞廟陷
func getLianZhenMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Ping,
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Xian2,
		miaoxian.Wang,
		miaoxian.Xian,
		miaoxian.Ping,
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Ping,
		miaoxian.Wang,
		miaoxian.Xian,
	}
	return miaoXianMap[index]
}

// getTianFuMiaoXian 得天府廟陷
func getTianFuMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Ping,
		miaoxian.Miao,
		miaoxian.Ping,
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Ping,
		miaoxian.Xian,
		miaoxian.Miao,
		miaoxian.Wang,
	}
	return miaoXianMap[index]
}

// getTaiYinMiaoXian 得太陰廟陷
func getTaiYinMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Xian2,
		miaoxian.Xian,
		miaoxian.Xian2,
		miaoxian.Xian,
		miaoxian.Xian,
		miaoxian.Ping,
		miaoxian.Ping,
		miaoxian.Wang,
		miaoxian.Wang,
		miaoxian.Miao,
	}
	return miaoXianMap[index]
}

// getTanLangMiaoXian 得貪狼廟陷
func getTanLangMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Ping,
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Xian,
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Ping,
		miaoxian.Miao,
		miaoxian.Xian,
	}
	return miaoXianMap[index]
}

// getJuMenMiaoXian 得巨門廟陷
func getJuMenMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Wang,
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Ping,
		miaoxian.Ping,
		miaoxian.Wang,
		miaoxian.Xian,
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Wang,
		miaoxian.Wang,
	}
	return miaoXianMap[index]
}

// getTianXiangMiaoXian 得天相廟陷
func getTianXiangXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Xian,
		miaoxian.Wang,
		miaoxian.Ping,
		miaoxian.Wang,
		miaoxian.Xian2,
		miaoxian.Miao,
		miaoxian.Xian,
		miaoxian.Xian2,
		miaoxian.Ping,
	}
	return miaoXianMap[index]
}

// getTianLiangMiaoXian 得天梁廟陷
func getTianLiangXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Miao,
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Wang,
		miaoxian.Xian,
		miaoxian.Miao,
		miaoxian.Wang,
		miaoxian.Xian,
		miaoxian.Di,
		miaoxian.Wang,
		miaoxian.Xian,
	}
	return miaoXianMap[index]
}
