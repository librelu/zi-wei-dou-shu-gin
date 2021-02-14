package ziwei

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/zi-wei-dou-shu-gin/utils/lunacal"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/dizhi"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/genders"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/miaoxian"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/mingju"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/stars"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/startype"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/tiangan"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/utils"
)

func NewTianBoard(birthday time.Time, gender genders.Gender) (*TianBoard, error) {
	board := new(TianBoard)
	board.Birthday = birthday
	board.Gender = gender
	lunaDate := lunacal.Solar2Lunar(birthday)
	board.LunaBirthday = lunaDate
	err := board.setGender(&board.LunaBirthday.Year.TianGan, board.Gender)
	if err != nil {
		return nil, fmt.Errorf("failed in set gender: %w", err)
	}
	board.StarsMap = make(map[stars.StarName]int)
	b := utils.Board(*board)
	utilBoard := utils.SetupDiZhi(&b)
	utilBoard = utils.SetupGongWei(utilBoard, lunaDate)
	utilBoard = utils.SetYinShouAndTianGanLocation(utilBoard, &board.LunaBirthday.Year.TianGan)
	tainBoard := TianBoard(*utilBoard)
	return &tainBoard, nil
}

func (b *TianBoard) CreateTianBoard() (*TianBoard, error) {
	mingGongLocation := b.getMingGong(b.LunaBirthday.Hour, b.LunaBirthday.Month)
	b.MingGongLocation = int(*mingGongLocation)
	b.setMingJu(mingGongLocation)
	err := b.setFourteenMainStars(mingGongLocation, b.MingJu, b.LunaBirthday.Day)
	if err != nil {
		return nil, fmt.Errorf("failed in set fourteen main stars, error: %w", err)
	}
	b.setupMainStarsConnections(mingGongLocation)
	b = setUpNainGanStars(b, b.LunaBirthday.Year.TianGan)
	b.setXunKong(b.LunaBirthday.Year)
	b.setJieKong(&b.LunaBirthday.Year.TianGan)
	shenGongLocation := b.getShengGong(b.LunaBirthday.Hour, b.LunaBirthday.Month)
	b.setNianZhiXiZhuXing(&b.LunaBirthday.Year.DiZhi, mingGongLocation, shenGongLocation)
	b.setYueXiXing(int(b.LunaBirthday.Month))
	b = setShiXiZhuXing(b, &b.LunaBirthday.Year.DiZhi, int(b.LunaBirthday.Month), int(b.LunaBirthday.Day), b.LunaBirthday.Hour)
	err = b.setMingZhu(&b.LunaBirthday.Year.DiZhi)
	if err != nil {
		return nil, fmt.Errorf("failed in set ming zhu, error: %w", err)
	}
	err = b.setShenZhu(&b.LunaBirthday.Year.DiZhi)
	if err != nil {
		return nil, fmt.Errorf("failed in set shen zhu, error: %w", err)
	}

	luCunLocation, ok := b.StarsMap[stars.LuCun]
	if !ok {
		return nil, fmt.Errorf("failed to get lu Cun location")
	}
	b.setBoShiTwelveStars(luCunLocation)
	todaysLunaDate := lunacal.Solar2Lunar(time.Now())
	b.setLiuNianSuiQianZhuXing(&todaysLunaDate.Year.DiZhi)
	err = b.setSiHua(&b.LunaBirthday.Year.TianGan)
	if err != nil {
		return nil, fmt.Errorf("failed in set si hua: %w", err)
	}
	b.setTenYearsRound(mingGongLocation)
	return b, nil
}

func (b *TianBoard) setupMainStarsConnections(minGongLocation *dizhi.DiZhi) {
	mingGongIndex := int(*minGongLocation)
	b.MainStarConnections = make([]int, 3)
	b.MainStarConnections[0] = (mingGongIndex + 4) % 12
	b.MainStarConnections[1] = (mingGongIndex + 6) % 12
	b.MainStarConnections[2] = (mingGongIndex + 8) % 12
}

func (b *TianBoard) setTenYearsRound(mingGongLocation *dizhi.DiZhi) {
	number := int(b.MingJu.Number)
	tenYearsRoundFormat := "%d-%d"
	if b.Gender == genders.YangMale || b.Gender == genders.YinFemale {
		// clockwise
		for i := range b.Blocks {
			idx := (i + int(*mingGongLocation)) % 12
			startYear := number + i*10
			endYear := number + 9 + i*10
			b.Blocks[idx].TenYearsRound = fmt.Sprintf(tenYearsRoundFormat, startYear, endYear)
		}
	}
	if b.Gender == genders.YinMale || b.Gender == genders.YangFemale {
		for i := range b.Blocks {
			// reverse clockwise
			idx := (int(*mingGongLocation) - i) % 12
			if idx < 0 {
				idx = 12 + idx
			}
			startYear := number + i*10
			endYear := number + 9 + i*10
			b.Blocks[idx].TenYearsRound = fmt.Sprintf(tenYearsRoundFormat, startYear, endYear)
		}
	}

}

func (b *TianBoard) getMingGong(hour *dizhi.DiZhi, month uint) *dizhi.DiZhi {
	hourIndex := int(*hour) - 2
	mingGong := int(month-1) - int(hourIndex)
	if mingGong < 0 {
		mingGong += 12
	} else if mingGong > 11 {
		mingGong = mingGong - 12
	}
	mingGongLocation := dizhi.DiZhi(mingGong)
	return &mingGongLocation
}

func (b *TianBoard) getShengGong(hour *dizhi.DiZhi, month uint) *dizhi.DiZhi {
	hourIndex := int(*hour) - 10
	shengGong := int(month-1) + int(hourIndex)
	if shengGong < 0 {
		shengGong += 12
	}
	shengGongLocation := dizhi.DiZhi(shengGong)
	return &shengGongLocation
}

func (b *TianBoard) setMingJu(mingGongLocation *dizhi.DiZhi) {
	minGongTainGan := b.Blocks[*mingGongLocation].Location.TianGan
	tianGanIndex := int((minGongTainGan)/2) % 5
	diZhiIndex := int((*mingGongLocation)/2) % 3
	mingJuIndex := tianGanIndex + diZhiIndex
	if mingJuIndex > 4 {
		mingJuIndex = mingJuIndex - 5
	}

	juType := mingju.MingJuType(mingJuIndex)
	b.MingJu = &utils.MingJu{
		JuType: juType,
		Number: mingju.JuShuMap[juType],
	}
}

func (b *TianBoard) setFourteenMainStars(mingGongLocation *dizhi.DiZhi, mingJu *utils.MingJu, birthdate uint) error {
	// 設定紫微
	ziWeiStarIndex := getZiWeiStarLocation(mingGongLocation, mingJu, birthdate)
	b.Blocks[ziWeiStarIndex].Stars = append(b.Blocks[ziWeiStarIndex].Stars, &utils.Star{
		Name:     stars.ZiWei.String(),
		StarType: startype.FourteenMainStars.String(),
		MiaoXian: getZiWeiMiaoXian(ziWeiStarIndex).String(),
	})
	b.StarsMap[stars.ZiWei] = ziWeiStarIndex

	// 設定天府
	tianFuIndex, err := b.setTianFuStarLocation(ziWeiStarIndex)
	if err == nil {
		b.Blocks[tianFuIndex].Stars = append(b.Blocks[tianFuIndex].Stars, &utils.Star{
			Name:     stars.TianFu.String(),
			StarType: startype.FourteenMainStars.String(),
			MiaoXian: getTianFuMiaoXian(tianFuIndex).String(),
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
func (b *TianBoard) setSiHua(birthYear *tiangan.TianGan) error {
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
func (b *TianBoard) setHuaLu(birthYear *tiangan.TianGan) error {
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
	starName := starMap[*birthYear]
	index, ok := b.StarsMap[starName]
	if !ok {
		return fmt.Errorf("current star not found, current birth year: %d", birthYear)
	}
	hasFound := false
	for i, s := range b.Blocks[index].Stars {
		if s.Name == starName.String() {
			b.Blocks[index].Stars[i].FourStar = stars.HuaLu.String()
			hasFound = true
		}
	}
	if !hasFound {
		return fmt.Errorf("stars %s doesn't located in %d block", starName, index)
	}
	return nil
}

// setHuaQuan 設定化權
func (b *TianBoard) setHuaQuan(birthYear *tiangan.TianGan) error {
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
	starName := starMap[*birthYear]
	index, ok := b.StarsMap[starName]
	if !ok {
		return fmt.Errorf("current star not found, current birth year: %d", birthYear)
	}
	hasFound := false
	for i, s := range b.Blocks[index].Stars {
		if s.Name == starName.String() {
			b.Blocks[index].Stars[i].FourStar = stars.HuaQuan.String()
			hasFound = true
		}
	}
	if !hasFound {
		return fmt.Errorf("stars %s doesn't located in %d block", starName, index)
	}
	return nil
}

// setHuaKe 設定化科
func (b *TianBoard) setHuaKe(birthYear *tiangan.TianGan) error {
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
	starName := starMap[*birthYear]
	index, ok := b.StarsMap[starName]
	if !ok {
		return fmt.Errorf("current star not found, current birth year: %d", birthYear)
	}
	hasFound := false
	for i, s := range b.Blocks[index].Stars {
		if s.Name == starName.String() {
			b.Blocks[index].Stars[i].FourStar = stars.HuaKe.String()
			hasFound = true
		}
	}
	if !hasFound {
		return fmt.Errorf("stars %s doesn't located in %d block", starName, index)
	}
	return nil
}

// setHuaJi 設定化忌
func (b *TianBoard) setHuaJi(birthYear *tiangan.TianGan) error {
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
	starName := starMap[*birthYear]
	index, ok := b.StarsMap[starName]
	if !ok {
		return fmt.Errorf("current star not found, current birth year: %d", birthYear)
	}
	hasFound := false
	for i, s := range b.Blocks[index].Stars {
		if s.Name == starName.String() {
			b.Blocks[index].Stars[i].FourStar = stars.HuaJi.String()
			hasFound = true
		}
	}
	if !hasFound {
		return fmt.Errorf("stars %s doesn't located in %d block", starName, index)
	}
	return nil
}

//setStarsBeggingWithZiWei 順時針一宮安天機星，跳隔一宮，安太陽星，順時針一宮安武曲星，順時針一宮安天同星，跳隔兩宮，安廉貞星
func (b *TianBoard) setStarsBeggingWithZiWei(ziWeiStarIndex int) {
	// 設定天機
	tianJi := ziWeiStarIndex - 1
	if tianJi < 0 {
		tianJi = 12 + tianJi
	}
	b.Blocks[tianJi].Stars = append(b.Blocks[tianJi].Stars, &utils.Star{
		Name:     stars.TianJi.String(),
		StarType: startype.FourteenMainStars.String(),
		MiaoXian: getTianJiMiaoXian(tianJi).String(),
	})
	b.StarsMap[stars.TianJi] = tianJi

	//　設定太陽
	taiYang := tianJi - 2
	if taiYang < 0 {
		taiYang = 12 + taiYang
	}
	b.Blocks[taiYang].Stars = append(b.Blocks[taiYang].Stars, &utils.Star{
		Name:     stars.TaiYang.String(),
		StarType: startype.FourteenMainStars.String(),
		MiaoXian: getTaiYangMiaoXian(taiYang).String(),
	})
	b.StarsMap[stars.TaiYang] = taiYang

	//　設定武曲
	wuQu := taiYang - 1
	if wuQu < 0 {
		wuQu = 12 + wuQu
	}
	b.Blocks[wuQu].Stars = append(b.Blocks[wuQu].Stars, &utils.Star{
		Name:     stars.WuQu.String(),
		StarType: startype.FourteenMainStars.String(),
		MiaoXian: getWuQuMiaoXian(wuQu).String(),
	})
	b.StarsMap[stars.WuQu] = wuQu

	// 設定天同
	tianTong := wuQu - 1
	if tianTong < 0 {
		tianTong = 12 + tianTong
	}
	b.Blocks[tianTong].Stars = append(b.Blocks[tianTong].Stars, &utils.Star{
		Name:     stars.TianTong.String(),
		StarType: startype.FourteenMainStars.String(),
		MiaoXian: getTianTongMiaoXian(tianTong).String(),
	})
	b.StarsMap[stars.TianTong] = tianTong

	// 設定廉貞
	lianZhen := tianTong - 3
	if lianZhen < 0 {
		lianZhen = 12 + lianZhen
	}
	b.Blocks[lianZhen].Stars = append(b.Blocks[lianZhen].Stars, &utils.Star{
		Name:     stars.LianZhen.String(),
		StarType: startype.FourteenMainStars.String(),
		MiaoXian: getLianZhenMiaoXian(lianZhen).String(),
	})
	b.StarsMap[stars.LianZhen] = lianZhen

	return
}

//setStarsBeggingWithTianFu 逆時針一宮安太陰星，逆時針一宮安貪狼星，逆時針一宮安巨門星，逆時針一宮安天相星，逆時針一宮安天梁星，逆時針一宮安七殺星，跳隔三宮，安破軍星
func (b *TianBoard) setStarsBeggingWithTianFu(tianFuIndex int) {
	// 設定太陰
	taiYin := tianFuIndex + 1
	if taiYin > 11 {
		taiYin = taiYin - 12
	}
	b.Blocks[taiYin].Stars = append(b.Blocks[taiYin].Stars, &utils.Star{
		Name:     stars.TaiYin.String(),
		StarType: startype.FourteenMainStars.String(),
		MiaoXian: getTaiYinMiaoXian(taiYin).String(),
	})
	b.StarsMap[stars.TaiYin] = taiYin

	// 設定貪狼
	tanLang := taiYin + 1
	if tanLang > 11 {
		tanLang = tanLang - 12
	}
	b.Blocks[tanLang].Stars = append(b.Blocks[tanLang].Stars, &utils.Star{
		Name:     stars.TanLang.String(),
		StarType: startype.FourteenMainStars.String(),
		MiaoXian: getTanLangMiaoXian(tanLang).String(),
	})
	b.StarsMap[stars.TanLang] = tanLang

	//　設定巨門
	juMen := tanLang + 1
	if juMen > 11 {
		juMen = juMen - 12
	}
	b.Blocks[juMen].Stars = append(b.Blocks[juMen].Stars, &utils.Star{
		Name:     stars.JuMen.String(),
		StarType: startype.FourteenMainStars.String(),
		MiaoXian: getJuMenMiaoXian(juMen).String(),
	})
	b.StarsMap[stars.JuMen] = juMen

	// 設定天相
	tianXiang := juMen + 1
	if tianXiang > 11 {
		tianXiang = tianXiang - 12
	}
	b.Blocks[tianXiang].Stars = append(b.Blocks[tianXiang].Stars, &utils.Star{
		Name:     stars.TianXiang.String(),
		StarType: startype.FourteenMainStars.String(),
		MiaoXian: getTianXiangXian(tianXiang).String(),
	})
	b.StarsMap[stars.TianXiang] = tianXiang

	// 設定天梁
	tianLiang := tianXiang + 1
	if tianLiang > 11 {
		tianLiang = tianLiang - 12
	}
	b.Blocks[tianLiang].Stars = append(b.Blocks[tianLiang].Stars, &utils.Star{
		Name:     stars.TianLiang.String(),
		StarType: startype.FourteenMainStars.String(),
		MiaoXian: getTianLiangXian(tianLiang).String(),
	})
	b.StarsMap[stars.TianLiang] = tianLiang

	// 設定七殺
	qiSha := tianLiang + 1
	if qiSha > 11 {
		qiSha = qiSha - 12
	}
	b.Blocks[qiSha].Stars = append(b.Blocks[qiSha].Stars, &utils.Star{
		Name:     stars.QiSha.String(),
		StarType: startype.FourteenMainStars.String(),
		MiaoXian: getQiShaMiaoXian(qiSha).String(),
	})
	b.StarsMap[stars.QiSha] = qiSha

	// 設定破軍
	poJun := qiSha + 4
	if poJun > 11 {
		poJun = poJun - 12
	}
	b.Blocks[poJun].Stars = append(b.Blocks[poJun].Stars, &utils.Star{
		Name:     stars.PoJun.String(),
		StarType: startype.FourteenMainStars.String(),
		MiaoXian: getPoJunMiaoXian(poJun).String(),
	})
	b.StarsMap[stars.PoJun] = poJun
	return
}

func getZiWeiStarLocation(mingGongLocation *dizhi.DiZhi, mingJu *utils.MingJu, birthdate uint) int {
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

func (b *TianBoard) setTianFuStarLocation(ziWeiStarIndex int) (int, error) {
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
func setUpNainGanStars(b *TianBoard, tianGan tiangan.TianGan) *TianBoard {
	board := utils.Board(*b)
	utilBoard := utils.SetLuCun(&board, tianGan)
	utilBoard = utils.SetQingYang(utilBoard, tianGan)
	utilBoard = utils.SetTuoLuo(utilBoard, tianGan)
	utilBoard = utils.SetTianKui(utilBoard, tianGan)
	utilBoard = utils.SetTianGuan(utilBoard, tianGan)
	tainBoard := TianBoard(*utilBoard)
	tb := &tainBoard
	tb = setTianYue(tb, tianGan)
	tb = setTianFu(tb, tianGan)
	return tb
}

// setTianYue 設定天鉞
func setTianYue(b *TianBoard, tianGan tiangan.TianGan) *TianBoard {
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
	tianYueLocation := tianYueMap[int(tianGan)]
	b.Blocks[tianYueLocation].Stars = append(b.Blocks[tianYueLocation].Stars, &utils.Star{
		Name:     stars.TianYue.String(),
		StarType: startype.LeftFuXing.String(),
		MiaoXian: getTianYueMiaoXian(int(tianYueLocation)).String(),
	})
	b.StarsMap[stars.TianYue] = int(tianYueLocation)
	return b
}

// setTianGuan 設定天福
func setTianFu(b *TianBoard, tianGan tiangan.TianGan) *TianBoard {
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
	tianFuLocation := tianFuMap[int(tianGan)]
	b.Blocks[tianFuLocation].Stars = append(b.Blocks[tianFuLocation].Stars, &utils.Star{
		Name:     stars.NainGanTianFu.String(),
		StarType: startype.NianGanXiZhuXing.String(),
	})
	return b
}

// setXunKong 設定旬空星
func (b *TianBoard) setXunKong(birthYear *lunacal.TianGanDiZhi) {
	dizhi := int(birthYear.DiZhi)
	if int(birthYear.TianGan) > dizhi {
		dizhi += 11
	}
	idx := dizhi - int(birthYear.TianGan) - 1
	if idx < 0 {
		idx += 11
	}

	if (idx % 2) == 0 {
		b.Blocks[idx].Stars = append(b.Blocks[idx].Stars, &utils.Star{
			Name:     stars.XunKong.String(),
			StarType: startype.NianGanXiZhuXing.String(),
		})

		idx++
		b.Blocks[idx].Stars = append(b.Blocks[idx].Stars, &utils.Star{
			Name:     stars.XunKong.String(),
			StarType: startype.NianGanXiZhuXing.String(),
		})
	} else {
		b.Blocks[idx].Stars = append(b.Blocks[idx].Stars, &utils.Star{
			Name:     stars.XunKong.String(),
			StarType: startype.NianGanXiZhuXing.String(),
		})

		idx--
		b.Blocks[idx].Stars = append(b.Blocks[idx].Stars, &utils.Star{
			Name:     stars.XunKong.String(),
			StarType: startype.NianGanXiZhuXing.String(),
		})
	}

}

// setJieKong 設定截空星
func (b *TianBoard) setJieKong(birthYear *tiangan.TianGan) {
	index := ((4 - int(*birthYear)%5 + 1) * 2) - 1
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &utils.Star{
		Name:     stars.JieKong.String(),
		StarType: startype.NianGanXiZhuXing.String(),
	})
	b.Blocks[index-1].Stars = append(b.Blocks[index-1].Stars, &utils.Star{
		Name:     stars.JieKong.String(),
		StarType: startype.NianGanXiZhuXing.String(),
	})
	return
}

// NianZhiXiZhuXing 設定年支系諸星
func (b *TianBoard) setNianZhiXiZhuXing(birthYear *dizhi.DiZhi, mingGongLocation *dizhi.DiZhi, shengGongLocation *dizhi.DiZhi) {
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
func (b *TianBoard) setTianKu(birthYear *dizhi.DiZhi) {
	index := (11 - int(*birthYear) + 7) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &utils.Star{
		Name:     stars.TianKu.String(),
		StarType: startype.NianZhiXiZhuXing.String(),
		MiaoXian: getTianKuMiaoXian(index).String(),
	})
	b.StarsMap[stars.TianKu] = index
	return
}

// setTianXu 設定天虛
func (b *TianBoard) setTianXu(birthYear *dizhi.DiZhi) {
	index := (int(*birthYear) + 6) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &utils.Star{
		Name:     stars.TianXu.String(),
		StarType: startype.NianZhiXiZhuXing.String(),
		MiaoXian: getTianXuMiaoXian(index).String(),
	})
	b.StarsMap[stars.TianXu] = index
	return
}

// setLongChi 設定龍池
func (b *TianBoard) setLongChi(birthYear *dizhi.DiZhi) {
	index := (int(*birthYear) + 4) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &utils.Star{
		Name:     stars.LongChi.String(),
		StarType: startype.NianZhiXiZhuXing.String(),
	})
	return
}

// setFengGe 設定鳳閣
func (b *TianBoard) setFengGe(birthYear *dizhi.DiZhi) {
	index := (11 - int(*birthYear) + 11) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &utils.Star{
		Name:     stars.FengGe.String(),
		StarType: startype.NianZhiXiZhuXing.String(),
	})
	return
}

// setHongLuan 設定紅鸞
func (b *TianBoard) setHongLuan(birthYear *dizhi.DiZhi) {
	index := (11 - int(*birthYear) + 4) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &utils.Star{
		Name:     stars.HongLuan.String(),
		StarType: startype.NianZhiXiZhuXing.String(),
		MiaoXian: getHongLuanMiaoXian(index).String(),
	})
	b.StarsMap[stars.HongLuan] = index
	return
}

// setTianXi 設定天喜
func (b *TianBoard) setTianXi(birthYear *dizhi.DiZhi) {
	index := (11 - int(*birthYear) + 10) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &utils.Star{
		Name:     stars.TianXi.String(),
		StarType: startype.NianZhiXiZhuXing.String(),
		MiaoXian: getTianXiMiaoXian(index).String(),
	})
	b.StarsMap[stars.TianXi] = index
	return
}

// setGuChen 設定孤辰
func (b *TianBoard) setGuChen(birthYear *dizhi.DiZhi) {
	locationIndex := (int(*birthYear) + 1) / 3 % 4
	locations := []dizhi.DiZhi{dizhi.Yin, dizhi.Si, dizhi.Shen, dizhi.Hai}
	index := locations[locationIndex]
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &utils.Star{
		Name:     stars.GuChen.String(),
		StarType: startype.NianZhiXiZhuXing.String(),
	})
	return
}

// setGuaXiu　設定寡宿
func (b *TianBoard) setGuaXiu(birthYear *dizhi.DiZhi) {
	locationIndex := (int(*birthYear) + 1) / 3 % 4
	locations := []dizhi.DiZhi{dizhi.Xu, dizhi.Chou, dizhi.Chen, dizhi.Wei}
	index := locations[locationIndex]
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &utils.Star{
		Name:     stars.GuaXiu.String(),
		StarType: startype.NianZhiXiZhuXing.String(),
	})
	return
}

// setJieShen 設定解神
func (b *TianBoard) setJieShen(birthYear *dizhi.DiZhi) {
	index := (11 - int(*birthYear) + 11) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &utils.Star{
		Name:     stars.JieShen.String(),
		StarType: startype.NianZhiXiZhuXing.String(),
	})
	return
}

// setPoSui 設定破碎
func (b *TianBoard) setPoSui(birthYear *dizhi.DiZhi) {
	locationIndex := int(int(*birthYear) % 3)
	locations := []dizhi.DiZhi{dizhi.Si, dizhi.Chou, dizhi.You}
	b.Blocks[locations[locationIndex]].Stars = append(
		b.Blocks[locations[locationIndex]].Stars, &utils.Star{
			Name:     stars.PoSui.String(),
			StarType: startype.NianZhiXiZhuXing.String(),
		})
	return
}

// setPoSui 設定天馬
func (b *TianBoard) setTianMa(birthYear *dizhi.DiZhi) {
	locationIndex := int((int(*birthYear) + 1) % 4)
	locations := []dizhi.DiZhi{dizhi.Si, dizhi.Yin, dizhi.Hai, dizhi.Shen}
	blockIndex := locations[locationIndex]
	b.Blocks[blockIndex].Stars = append(
		b.Blocks[blockIndex].Stars, &utils.Star{
			Name:     stars.TianMa.String(),
			StarType: startype.ShiSiFuXing.String(),
			MiaoXian: getTianMaMiaoXian(int(blockIndex)).String(),
		})
	b.StarsMap[stars.TianMa] = int(blockIndex)
	return
}

// setDaHao 設定大耗
func (b *TianBoard) setDaHao(birthYear *dizhi.DiZhi) {
	locationMap := []dizhi.DiZhi{
		dizhi.Wei, dizhi.Wu, dizhi.You, dizhi.Shen, dizhi.Hai, dizhi.Xu, dizhi.Chou, dizhi.Zi, dizhi.Mao, dizhi.Yin, dizhi.Si, dizhi.Chen,
	}
	index := locationMap[*birthYear]
	b.Blocks[index].Stars = append(
		b.Blocks[index].Stars, &utils.Star{
			Name:     stars.DaHao.String(),
			StarType: startype.NianZhiXiZhuXing.String(),
		})
	return
}

// setTianDe 設定天德
func (b *TianBoard) setTianDe(birthYear *dizhi.DiZhi) {
	index := (int(*birthYear) + 9) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &utils.Star{
		Name:     stars.TianDe.String(),
		StarType: startype.NianZhiXiZhuXing.String(),
	})
	return
}

// setJieSha 設定劫殺
func (b *TianBoard) setJieSha(birthYear *dizhi.DiZhi) {
	locationIndex := int(int(*birthYear) % 4)
	locations := []dizhi.DiZhi{dizhi.Si, dizhi.Yin, dizhi.Hai, dizhi.Shen}
	b.Blocks[locations[locationIndex]].Stars = append(
		b.Blocks[locations[locationIndex]].Stars, &utils.Star{
			Name:     stars.JieSha.String(),
			StarType: startype.NianZhiXiZhuXing.String(),
		})
	return
}

// setHuaGai 設定華蓋
func (b *TianBoard) setHuaGai(birthYear *dizhi.DiZhi) {
	locationIndex := int(int(*birthYear) % 4)
	locations := []dizhi.DiZhi{dizhi.Chen, dizhi.Chou, dizhi.Xu, dizhi.Wei}
	b.Blocks[locations[locationIndex]].Stars = append(
		b.Blocks[locations[locationIndex]].Stars, &utils.Star{
			Name:     stars.HuaGai.String(),
			StarType: startype.NianZhiXiZhuXing.String(),
		})
	return
}

// setXianChi 設定咸池
func (b *TianBoard) setXianChi(birthYear *dizhi.DiZhi) {
	locationIndex := int(int(*birthYear) % 4)
	locations := []dizhi.DiZhi{dizhi.You, dizhi.Wu, dizhi.Mao, dizhi.Zi}
	b.Blocks[locations[locationIndex]].Stars = append(
		b.Blocks[locations[locationIndex]].Stars, &utils.Star{
			Name:     stars.XianChi.String(),
			StarType: startype.NianZhiXiZhuXing.String(),
		})
	return
}

// setTianCai 設定天才
func (b *TianBoard) setTianCai(birthYear *dizhi.DiZhi, mingGongLocation *dizhi.DiZhi) {
	index := (*mingGongLocation + *birthYear) % 12
	b.Blocks[index].Stars = append(
		b.Blocks[index].Stars, &utils.Star{
			Name:     stars.TianCai.String(),
			StarType: startype.NianZhiXiZhuXing.String(),
		})
	return
}

// setTianShou 設定天壽
func (b *TianBoard) setTianShou(birthYear *dizhi.DiZhi, shengGongLocation *dizhi.DiZhi) {
	index := (*shengGongLocation + *birthYear) % 12
	b.Blocks[index].Stars = append(
		b.Blocks[index].Stars, &utils.Star{
			Name:     stars.TianShou.String(),
			StarType: startype.NianZhiXiZhuXing.String(),
		})
	return
}

// setTianKong 設定天空
func (b *TianBoard) setTianKong(birthYear *dizhi.DiZhi) {
	index := (int(*birthYear) + 1) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &utils.Star{
		Name:     stars.TianKong.String(),
		StarType: startype.NianZhiXiZhuXing.String(),
	})
	return
}

// setYueXiXing 安月系星
func (b *TianBoard) setYueXiXing(birthMonth int) {
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
func (b *TianBoard) setZuoFu(birthMonth int) {
	index := getZuoFuLocation(birthMonth)
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &utils.Star{
		Name:     stars.ZuoFu.String(),
		StarType: startype.LeftFuXing.String(),
		MiaoXian: getZuoFuMiaoXian(index).String(),
	})
	b.StarsMap[stars.ZuoFu] = index
	return
}

func getZuoFuLocation(birthMonth int) int {
	return (birthMonth + 3) % 12
}

// setYouBi 設定右弼
func (b *TianBoard) setYouBi(birthMonth int) {
	index := getYouBiLocation(birthMonth)
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &utils.Star{
		Name:     stars.YouBi.String(),
		StarType: startype.LeftFuXing.String(),
		MiaoXian: getYouBiMiaoXian(index).String(),
	})
	b.StarsMap[stars.YouBi] = index
	return
}

func getYouBiLocation(birthMonth int) int {
	return ((13 - birthMonth) + 10) % 12
}

// setTianXing 設定天刑
func (b *TianBoard) setTianXing(birthMonth int) {
	index := (birthMonth + 8) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &utils.Star{
		Name:     stars.TianXing.String(),
		StarType: startype.YueXiXing.String(),
		MiaoXian: getTianXingMiaoXian(index).String(),
	})
	b.StarsMap[stars.TianXing] = index
	return
}

// setTianYao 設定天姚
func (b *TianBoard) setTianYao(birthMonth int) {
	index := (birthMonth) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &utils.Star{
		Name:     stars.TianYao.String(),
		StarType: startype.YueXiXing.String(),
		MiaoXian: getTianYaoMiaoXian(index).String(),
	})
	b.StarsMap[stars.TianYao] = index
	return
}

// setTianWu 設定天巫
func (b *TianBoard) setTianWu(birthMonth int) {
	locationIndex := (birthMonth - 1) % 4
	locations := []dizhi.DiZhi{dizhi.Si, dizhi.Shen, dizhi.Yin, dizhi.Hai}
	b.Blocks[locations[locationIndex]].Stars = append(
		b.Blocks[locations[locationIndex]].Stars, &utils.Star{
			Name:     stars.TianWu.String(),
			StarType: startype.YueXiXing.String(),
		})
	return
}

// setYueXiXingTianYue 設定天月
// TODO: refactor Yue Xi Xing to a object
func (b *TianBoard) setYueXiXingTianYue(birthMonth int) {
	locationMap := []dizhi.DiZhi{
		dizhi.Xu, dizhi.Si, dizhi.Chen, dizhi.Yin, dizhi.Wei, dizhi.Mao, dizhi.Hai, dizhi.Wei, dizhi.Yin, dizhi.Wu, dizhi.Xu, dizhi.Yin,
	}
	index := locationMap[birthMonth-1]
	b.Blocks[index].Stars = append(
		b.Blocks[index].Stars, &utils.Star{
			Name:     stars.YueXiXingTianYue.String(),
			StarType: startype.YueXiXing.String(),
		})
	return
}

// setYinSha 設定陰煞
func (b *TianBoard) setYinSha(birthMonth int) {
	locationIndex := (birthMonth + 2) % 6
	locations := []dizhi.DiZhi{dizhi.Shen, dizhi.Wu, dizhi.Chen, dizhi.Yin, dizhi.Zi, dizhi.Xu}
	b.Blocks[locations[locationIndex]].Stars = append(
		b.Blocks[locations[locationIndex]].Stars, &utils.Star{
			Name:     stars.YinSha.String(),
			StarType: startype.YueXiXing.String(),
		})
	return
}

// setShiXiZhuXing 安時系諸星
func setShiXiZhuXing(b *TianBoard, birthYear *dizhi.DiZhi, birthMonth int, birthDate int, birthHour *dizhi.DiZhi) *TianBoard {
	b.setWenChang(birthHour)
	b.setWenQu(birthHour)
	b.setDiJie(birthHour)
	b.setDiKong(birthHour)
	b.setTaiFu(birthHour)
	b.setFengGao(birthHour)
	utilBoard := utils.Board(*b)
	ub := &utilBoard
	ub = utils.SetHuo(ub, birthYear, birthHour)
	ub = utils.SetLing(ub, birthYear, birthHour)
	*b = TianBoard(*ub)
	zuoFuLocation := getZuoFuLocation(birthMonth)
	b.setSanTai(zuoFuLocation, birthDate)
	youBiLocation := getYouBiLocation(birthMonth)
	b.setBaZuo(youBiLocation, birthDate)
	wenChangLocation := getWenChangLocation(birthHour)
	b.setEnGuang(wenChangLocation, birthDate)
	wenQuLocation := getWenQuLocation(birthHour)
	b.setTianGui(wenQuLocation, birthDate)
	return b
}

// setWenChang 設定文昌
func (b *TianBoard) setWenChang(birthHour *dizhi.DiZhi) {
	index := getWenChangLocation(birthHour)
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &utils.Star{
		Name:     stars.WenChang.String(),
		StarType: startype.LeftFuXing.String(),
		MiaoXian: getWenChangMiaoXian(index).String(),
	})
	b.StarsMap[stars.WenChang] = index
	return
}

func getWenChangLocation(birthHour *dizhi.DiZhi) int {
	return (11 - int(*birthHour) + 11) % 12
}

// setWen 設定文曲
func (b *TianBoard) setWenQu(birthHour *dizhi.DiZhi) {
	index := getWenQuLocation(birthHour)
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &utils.Star{
		Name:     stars.WenQu.String(),
		StarType: startype.LeftFuXing.String(),
		MiaoXian: getWenQuMiaoXian(index).String(),
	})
	b.StarsMap[stars.WenQu] = index
	return
}

func getWenQuLocation(birthHour *dizhi.DiZhi) int {
	return (int(*birthHour) + 4) % 12
}

// setDiJie 設定地劫
func (b *TianBoard) setDiJie(birthHour *dizhi.DiZhi) {
	index := (int(*birthHour) + 11) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &utils.Star{
		Name:     stars.DiJie.String(),
		StarType: startype.RightFuXing.String(),
		MiaoXian: getDiJieMiaoXian(index).String(),
	})
	b.StarsMap[stars.DiJie] = index
	return
}

// setDiKong 設定地空
func (b *TianBoard) setDiKong(birthHour *dizhi.DiZhi) {
	index := (11 - int(*birthHour) + 12) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &utils.Star{
		Name:     stars.DiKong.String(),
		StarType: startype.RightFuXing.String(),
		MiaoXian: getDiKongMiaoXian(index).String(),
	})
	b.StarsMap[stars.DiKong] = index
	return
}

// setTaiFu 設定台輔
func (b *TianBoard) setTaiFu(birthHour *dizhi.DiZhi) {
	index := (int(*birthHour) + 6) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &utils.Star{
		Name:     stars.TaiFu.String(),
		StarType: startype.ShiXiZhuXing.String(),
	})
	return
}

// setFengGao 設定封誥
func (b *TianBoard) setFengGao(birthHour *dizhi.DiZhi) {
	index := (int(*birthHour) + 2) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &utils.Star{
		Name:     stars.FengGao.String(),
		StarType: startype.ShiXiZhuXing.String(),
	})
	return
}

// setSanTai 設定三台
func (b *TianBoard) setSanTai(zuoFuLocation int, birthDate int) {
	index := (zuoFuLocation + birthDate - 1) % 12
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &utils.Star{
		Name:     stars.SanTai.String(),
		StarType: startype.ShiXiZhuXing.String(),
	})
	return
}

// setBaZuo 設定八座
func (b *TianBoard) setBaZuo(youBiLocation int, birthDate int) {
	index := (youBiLocation - (birthDate - 1)) % 12
	if index < 0 {
		index += 12
	}
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &utils.Star{
		Name:     stars.BaZuo.String(),
		StarType: startype.ShiXiZhuXing.String(),
	})
	return
}

// setEnGuang 設定恩光
func (b *TianBoard) setEnGuang(wenChangLocation int, birthDate int) {
	index := (wenChangLocation + birthDate - 2) % 12
	if index < 0 {
		index += 11
	}
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &utils.Star{
		Name:     stars.EnGuang.String(),
		StarType: startype.ShiXiZhuXing.String(),
	})
	return
}

// setTianGui 設定天貴
func (b *TianBoard) setTianGui(wenQuLocation int, birthDate int) {
	index := (wenQuLocation + birthDate - 2) % 12
	if index < 0 {
		index += 11
	}
	b.Blocks[index].Stars = append(b.Blocks[index].Stars, &utils.Star{
		Name:     stars.TianGui.String(),
		StarType: startype.ShiXiZhuXing.String(),
	})
	return
}

// setMingZhu 設定命主
func (b *TianBoard) setMingZhu(birthYear *dizhi.DiZhi) error {
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
	_, ok := b.StarsMap[starName]
	if !ok {
		return fmt.Errorf("the star name not found in birth year = %d", birthYear)
	}
	b.MingZhu = starMap[*birthYear].String()
	return nil
}

// setShenZhu 設定身主
func (b *TianBoard) setShenZhu(birthYear *dizhi.DiZhi) error {
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
	_, ok := b.StarsMap[starName]
	if !ok {
		return fmt.Errorf("the star name not found in birth year = %d", birthYear)
	}
	b.ShenZhu = starMap[*birthYear].String()
	return nil
}

// setGetGenders
func (b *TianBoard) setGender(birthYear *tiangan.TianGan, gender genders.Gender) error {
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
func (b *TianBoard) setBoShiTwelveStars(LuCunLocation int) {
	starsMap := []*utils.Star{
		{
			Name:     stars.BoShi.String(),
			StarType: startype.BoShiTwelveStars.String(),
		},
		{
			Name:     stars.LiShi.String(),
			StarType: startype.BoShiTwelveStars.String(),
		},
		{
			Name:     stars.QingLong.String(),
			StarType: startype.BoShiTwelveStars.String(),
		},
		{
			Name:     stars.XiaoHao.String(),
			StarType: startype.BoShiTwelveStars.String(),
		},
		{
			Name:     stars.JiangJun.String(),
			StarType: startype.BoShiTwelveStars.String(),
		},
		{
			Name:     stars.ZouShu.String(),
			StarType: startype.BoShiTwelveStars.String(),
		},
		{
			Name:     stars.FeiLian.String(),
			StarType: startype.BoShiTwelveStars.String(),
		},
		{
			Name:     stars.XiShen.String(),
			StarType: startype.BoShiTwelveStars.String(),
		},
		{
			Name:     stars.BingFu.String(),
			StarType: startype.BoShiTwelveStars.String(),
		},
		{
			Name:     stars.BoShiDaHao.String(),
			StarType: startype.BoShiTwelveStars.String(),
		},
		{
			Name:     stars.FuBing.String(),
			StarType: startype.BoShiTwelveStars.String(),
		},
		{
			Name:     stars.GuanFu.String(),
			StarType: startype.BoShiTwelveStars.String(),
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
func (b *TianBoard) setLiuNianSuiQianZhuXing(currentDiZhiYear *dizhi.DiZhi) {
	starsMap := []*utils.Star{
		{
			Name:     stars.SuiJian.String(),
			StarType: startype.LiuNianSuiQianZhuXing.String(),
		},
		{
			Name:     stars.HuiQi.String(),
			StarType: startype.LiuNianSuiQianZhuXing.String(),
		},
		{
			Name:     stars.SangMen.String(),
			StarType: startype.LiuNianSuiQianZhuXing.String(),
		},
		{
			Name:     stars.GuanSuo.String(),
			StarType: startype.LiuNianSuiQianZhuXing.String(),
		},
		{
			Name:     stars.GuanFu.String(),
			StarType: startype.LiuNianSuiQianZhuXing.String(),
		},
		{
			Name:     stars.LiuNianXiaoHao.String(),
			StarType: startype.LiuNianSuiQianZhuXing.String(),
		},
		{
			Name:     stars.LiuNianDaHao.String(),
			StarType: startype.LiuNianSuiQianZhuXing.String(),
		},
		{
			Name:     stars.LongDe.String(),
			StarType: startype.LiuNianSuiQianZhuXing.String(),
		},
		{
			Name:     stars.BaiHu.String(),
			StarType: startype.LiuNianSuiQianZhuXing.String(),
		},
		{
			Name:     stars.TianDe.String(),
			StarType: startype.LiuNianSuiQianZhuXing.String(),
		},
		{
			Name:     stars.DiKe.String(),
			StarType: startype.LiuNianSuiQianZhuXing.String(),
		},
		{
			Name:     stars.LiuNianBingFu.String(),
			StarType: startype.LiuNianSuiQianZhuXing.String(),
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

// getQiShaMiaoXian 得七殺廟陷
func getQiShaMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Xian,
		miaoxian.Wang,
		miaoxian.Ping,
		miaoxian.Wang,
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Xian2,
		miaoxian.Miao,
		miaoxian.Ping,
	}
	return miaoXianMap[index]
}

// getPoJunMiaoXian 得破軍廟陷
func getPoJunMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Miao,
		miaoxian.Wang,
		miaoxian.Xian,
		miaoxian.Wang,
		miaoxian.Wang,
		miaoxian.Xian2,
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Xian,
		miaoxian.Xian,
		miaoxian.Wang,
		miaoxian.Ping,
	}
	return miaoXianMap[index]
}

// getTianMaMiaoXian 得天馬廟陷
func getTianMaMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.None,
		miaoxian.None,
		miaoxian.Wang,
		miaoxian.None,
		miaoxian.None,
		miaoxian.Ping,
		miaoxian.None,
		miaoxian.None,
		miaoxian.Wang,
		miaoxian.None,
		miaoxian.None,
		miaoxian.Ping,
	}
	return miaoXianMap[index]
}

// getDiKongMiaoXian 得地空廟陷
func getDiKongMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Ping,
		miaoxian.Xian,
		miaoxian.Xian,
		miaoxian.Ping,
		miaoxian.Xian,
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Ping,
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Xian,
		miaoxian.Xian,
	}
	return miaoXianMap[index]
}

// getDiJieMiaoXian 得地劫廟陷
func getDiJieMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Xian,
		miaoxian.Xian,
		miaoxian.Ping,
		miaoxian.Ping,
		miaoxian.Xian,
		miaoxian.Xian2,
		miaoxian.Miao,
		miaoxian.Ping,
		miaoxian.Miao,
		miaoxian.Ping,
		miaoxian.Ping,
		miaoxian.Wang,
	}
	return miaoXianMap[index]
}

// getTianYueMiaoXian 得天鉞廟陷
func getTianYueMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.None,
		miaoxian.None,
		miaoxian.Wang,
		miaoxian.None,
		miaoxian.None,
		miaoxian.Wang,
		miaoxian.None,
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.None,
		miaoxian.None,
	}
	return miaoXianMap[index]
}

// getZuoFuMiaoXian 得左輔廟陷
func getZuoFuMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Xian,
		miaoxian.Miao,
		miaoxian.Ping,
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Ping,
		miaoxian.Xian,
		miaoxian.Miao,
		miaoxian.Xian2,
	}
	return miaoXianMap[index]
}

// getYouBiMiaoXian 得右弼廟陷
func getYouBiMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Wang,
		miaoxian.Xian,
		miaoxian.Miao,
		miaoxian.Ping,
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Xian2,
		miaoxian.Xian,
		miaoxian.Miao,
		miaoxian.Ping,
	}
	return miaoXianMap[index]
}

// getWenChangMiaoXian 得文昌廟陷
func getWenChangMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Xian,
		miaoxian.Ping,
		miaoxian.Wang,
		miaoxian.Xian,
		miaoxian.Miao,
		miaoxian.Ping,
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Xian,
		miaoxian.Wang,
	}
	return miaoXianMap[index]
}

// getWenQuMiaoXian 得文曲廟陷
func getWenQuMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Ping,
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Xian,
		miaoxian.Wang,
		miaoxian.Ping,
		miaoxian.Miao,
		miaoxian.Xian,
		miaoxian.Wang,
	}
	return miaoXianMap[index]
}

// getHongLuanMiaoXian 得紅鸞廟陷
func getHongLuanMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Miao,
		miaoxian.Xian,
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Wang,
		miaoxian.Wang,
		miaoxian.Xian,
		miaoxian.Miao,
		miaoxian.Wang,
		miaoxian.Xian,
		miaoxian.Miao,
	}
	return miaoXianMap[index]
}

// getTianXiMiaoXian 得天喜廟陷
func getTianXiMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Wang,
		miaoxian.Xian,
		miaoxian.Miao,
		miaoxian.Wang,
		miaoxian.Xian,
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Xian,
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Xian,
		miaoxian.Wang,
	}
	return miaoXianMap[index]
}

// getTianXingMiaoXian 得天刑廟陷
func getTianXingMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Ping,
		miaoxian.Xian,
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Ping,
		miaoxian.Xian,
		miaoxian.Ping,
		miaoxian.Xian,
		miaoxian.Xian,
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Xian,
	}
	return miaoXianMap[index]
}

// getTianYaoMiaoXian 得天姚廟陷
func getTianYaoMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Xian,
		miaoxian.Ping,
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Xian,
		miaoxian.Ping,
		miaoxian.Ping,
		miaoxian.Wang,
		miaoxian.Xian,
		miaoxian.Miao,
		miaoxian.Miao,
		miaoxian.Xian,
	}
	return miaoXianMap[index]
}

// getTianKuMiaoXian 得天哭廟陷
func getTianKuMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Ping,
		miaoxian.Miao,
		miaoxian.Ping,
		miaoxian.Miao,
		miaoxian.Ping,
		miaoxian.Di,
		miaoxian.Xian,
		miaoxian.Ping,
		miaoxian.Miao,
		miaoxian.Xian2,
		miaoxian.Ping,
		miaoxian.Ping,
	}
	return miaoXianMap[index]
}

// getTianXuMiaoXian 得天虛廟陷
func getTianXuMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Xian,
		miaoxian.Miao,
		miaoxian.Wang,
		miaoxian.Miao,
		miaoxian.Xian,
		miaoxian.Wang,
		miaoxian.Ping,
		miaoxian.Xian,
		miaoxian.Miao,
		miaoxian.Wang,
		miaoxian.Xian,
		miaoxian.Ping,
	}
	return miaoXianMap[index]
}
