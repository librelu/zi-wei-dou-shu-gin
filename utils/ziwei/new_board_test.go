package ziwei_test

import (
	"time"

	"bou.ke/monkey"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/dizhi"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/genders"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/gong"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/mingju"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/stars"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/startype"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/tiangan"
)

var _ = Describe("ziwei", func() {
	Describe("NewBoard", func() {
		var (
			board    *ziwei.Board
			birthday time.Time
			gender   genders.Gender
			err      error
		)
		BeforeEach(func() {
			gender = genders.Male
		})
		JustBeforeEach(func() {
			board, err = ziwei.NewBoard(birthday, gender)
		})
		When("init board", func() {
			BeforeEach(func() {
				// luna date: 丙寅年 7 / 3 子時
				birthday = time.Date(1984, 8, 8, 0, 4, 0, 0, time.Local)
			})
			It("should returns location", func() {
				Expect(err).Should(BeNil())
				Expect(len(board.Blocks)).Should(Equal(12))
				for i, block := range board.Blocks {
					Expect(block.Location.DiZhi.String()).Should(Equal(dizhi.DiZhi(i).String()))
				}
			})
		})
		Context("testing Yin Shou", func() {
			When("given a birthday with Jia year", func() {
				BeforeEach(func() {
					// luna date: 丙寅年 7 / 3 子時
					birthday = time.Date(1984, 8, 8, 0, 4, 0, 0, time.Local)
				})
				It("shouldn't returns any errors", func() {
					Expect(err).Should(BeNil())
				})
				It("should display Yin Shou by birthday", func() {
					Expect(board.Blocks[0].Location.TianGan).Should(Equal(tiangan.Jia))
					Expect(board.Blocks[1].Location.TianGan).Should(Equal(tiangan.Yi))
					Expect(board.Blocks[2].Location.TianGan).Should(Equal(tiangan.Bing))
					Expect(board.Blocks[3].Location.TianGan).Should(Equal(tiangan.Ding))
					Expect(board.Blocks[4].Location.TianGan).Should(Equal(tiangan.Wu))
					Expect(board.Blocks[5].Location.TianGan).Should(Equal(tiangan.Ji))
					Expect(board.Blocks[6].Location.TianGan).Should(Equal(tiangan.Geng))
					Expect(board.Blocks[7].Location.TianGan).Should(Equal(tiangan.Xin))
					Expect(board.Blocks[8].Location.TianGan).Should(Equal(tiangan.Ren))
					Expect(board.Blocks[9].Location.TianGan).Should(Equal(tiangan.Gui))
					Expect(board.Blocks[10].Location.TianGan).Should(Equal(tiangan.Jia))
					Expect(board.Blocks[11].Location.TianGan).Should(Equal(tiangan.Yi))
				})
				It("should display correct location", func() {
					Expect(board.Blocks[0].Location.DiZhi).Should(Equal(dizhi.Zi))
					Expect(board.Blocks[1].Location.DiZhi).Should(Equal(dizhi.Chou))
					Expect(board.Blocks[2].Location.DiZhi).Should(Equal(dizhi.Yin))
					Expect(board.Blocks[3].Location.DiZhi).Should(Equal(dizhi.Mao))
					Expect(board.Blocks[4].Location.DiZhi).Should(Equal(dizhi.Chen))
					Expect(board.Blocks[5].Location.DiZhi).Should(Equal(dizhi.Si))
					Expect(board.Blocks[6].Location.DiZhi).Should(Equal(dizhi.Wu))
					Expect(board.Blocks[7].Location.DiZhi).Should(Equal(dizhi.Wei))
					Expect(board.Blocks[8].Location.DiZhi).Should(Equal(dizhi.Shen))
					Expect(board.Blocks[9].Location.DiZhi).Should(Equal(dizhi.You))
					Expect(board.Blocks[10].Location.DiZhi).Should(Equal(dizhi.Xu))
					Expect(board.Blocks[11].Location.DiZhi).Should(Equal(dizhi.Hai))
				})
				It("should returns correct Ming Ju", func() {
					Expect(board.MingJu.JuType).Should(Equal(mingju.Shui))
					Expect(board.MingJu.Number).Should(Equal(uint(2)))
				})
			})
			When("given a birthday with Wu year", func() {
				BeforeEach(func() {
					// luna date: 丙寅年 7 / 3 子時
					birthday = time.Date(1988, 8, 8, 0, 4, 0, 0, time.Local)
				})
				It("shouldn't returns any errors", func() {
					Expect(err).Should(BeNil())
				})
				It("should display Yin Shou by birthday", func() {
					Expect(board.Blocks[2].Location.TianGan).Should(Equal(tiangan.Jia))
					Expect(board.Blocks[3].Location.TianGan).Should(Equal(tiangan.Yi))
					Expect(board.Blocks[4].Location.TianGan).Should(Equal(tiangan.Bing))
					Expect(board.Blocks[5].Location.TianGan).Should(Equal(tiangan.Ding))
					Expect(board.Blocks[6].Location.TianGan).Should(Equal(tiangan.Wu))
					Expect(board.Blocks[7].Location.TianGan).Should(Equal(tiangan.Ji))
					Expect(board.Blocks[8].Location.TianGan).Should(Equal(tiangan.Geng))
					Expect(board.Blocks[9].Location.TianGan).Should(Equal(tiangan.Xin))
					Expect(board.Blocks[10].Location.TianGan).Should(Equal(tiangan.Ren))
					Expect(board.Blocks[11].Location.TianGan).Should(Equal(tiangan.Gui))
					Expect(board.Blocks[0].Location.TianGan).Should(Equal(tiangan.Jia))
					Expect(board.Blocks[1].Location.TianGan).Should(Equal(tiangan.Yi))
				})
				It("should returns correct Ming Ju", func() {
					Expect(board.MingJu.JuType).Should(Equal(mingju.Huo))
					Expect(board.MingJu.Number).Should(Equal(uint(6)))
				})
			})
			When("given a birthday with Bing year", func() {
				BeforeEach(func() {
					// luna date: 丙寅年 7 / 3 子時
					birthday = time.Date(1986, 8, 8, 0, 4, 0, 0, time.Local)
				})
				It("shouldn't returns any errors", func() {
					Expect(err).Should(BeNil())
				})
				It("should display Yin Shou by birthday", func() {
					Expect(board.Blocks[8].Location.TianGan).Should(Equal(tiangan.Bing))
					Expect(board.Blocks[9].Location.TianGan).Should(Equal(tiangan.Ding))
					Expect(board.Blocks[10].Location.TianGan).Should(Equal(tiangan.Wu))
					Expect(board.Blocks[11].Location.TianGan).Should(Equal(tiangan.Ji))
					Expect(board.Blocks[0].Location.TianGan).Should(Equal(tiangan.Wu))
					Expect(board.Blocks[1].Location.TianGan).Should(Equal(tiangan.Ji))
					Expect(board.Blocks[2].Location.TianGan).Should(Equal(tiangan.Geng))
					Expect(board.Blocks[3].Location.TianGan).Should(Equal(tiangan.Xin))
					Expect(board.Blocks[4].Location.TianGan).Should(Equal(tiangan.Ren))
					Expect(board.Blocks[5].Location.TianGan).Should(Equal(tiangan.Gui))
					Expect(board.Blocks[6].Location.TianGan).Should(Equal(tiangan.Jia))
					Expect(board.Blocks[7].Location.TianGan).Should(Equal(tiangan.Yi))
				})
				It("should returns correct Ming Ju", func() {
					Expect(board.MingJu.JuType).Should(Equal(mingju.Huo))
					Expect(board.MingJu.Number).Should(Equal(uint(6)))
				})
			})
		})
		Context("testing gong Wei order", func() {
			When("given a birthday start with Zhi hour", func() {
				BeforeEach(func() {
					// luna date: 丙寅年 7 / 3 子時
					birthday = time.Date(1984, 8, 8, 0, 4, 0, 0, time.Local)
				})
				It("shouldn't returns any errors", func() {
					Expect(err).Should(BeNil())
				})
				It("should display Gong Wei by birthday", func() {
					Expect(board.Blocks[7].GongWeiName).Should(Equal(gong.XiongDi))
					Expect(board.Blocks[8].GongWeiName).Should(Equal(gong.Ming))
					Expect(board.Blocks[9].GongWeiName).Should(Equal(gong.FuMu))
					Expect(board.Blocks[10].GongWeiName).Should(Equal(gong.FuDe))
					Expect(board.Blocks[11].GongWeiName).Should(Equal(gong.TianZhai))
					Expect(board.Blocks[0].GongWeiName).Should(Equal(gong.ShiYe))
					Expect(board.Blocks[1].GongWeiName).Should(Equal(gong.JiaoYou))
					Expect(board.Blocks[2].GongWeiName).Should(Equal(gong.QianYi))
					Expect(board.Blocks[3].GongWeiName).Should(Equal(gong.JiBing))
					Expect(board.Blocks[4].GongWeiName).Should(Equal(gong.CaiBo))
					Expect(board.Blocks[5].GongWeiName).Should(Equal(gong.ZiNu))
					Expect(board.Blocks[6].GongWeiName).Should(Equal(gong.FuQi))
				})
				It("should returns correct Ming Ju", func() {
					Expect(board.MingJu.JuType).Should(Equal(mingju.Shui))
					Expect(board.MingJu.Number).Should(Equal(uint(2)))
				})
			})
			When("given a birthday start with Chou hour", func() {
				BeforeEach(func() {
					// luna date: 丙寅年 7 / 3 丑時
					birthday = time.Date(1984, 8, 8, 2, 4, 0, 0, time.Local)
				})
				It("shouldn't returns any errors", func() {
					Expect(err).Should(BeNil())
				})
				It("should returns location", func() {
					Expect(len(board.Blocks)).Should(Equal(12))
					for i, block := range board.Blocks {
						Expect(block.Location.DiZhi.String()).Should(Equal(dizhi.DiZhi(i).String()))
					}
				})
				It("should display Gong Wei by birthday", func() {
					Expect(board.Blocks[6].GongWeiName).Should(Equal(gong.XiongDi))
					Expect(board.Blocks[7].GongWeiName).Should(Equal(gong.Ming))
					Expect(board.Blocks[8].GongWeiName).Should(Equal(gong.FuMu))
					Expect(board.Blocks[9].GongWeiName).Should(Equal(gong.FuDe))
					Expect(board.Blocks[10].GongWeiName).Should(Equal(gong.TianZhai))
					Expect(board.Blocks[11].GongWeiName).Should(Equal(gong.ShiYe))
					Expect(board.Blocks[0].GongWeiName).Should(Equal(gong.JiaoYou))
					Expect(board.Blocks[1].GongWeiName).Should(Equal(gong.QianYi))
					Expect(board.Blocks[2].GongWeiName).Should(Equal(gong.JiBing))
					Expect(board.Blocks[3].GongWeiName).Should(Equal(gong.CaiBo))
					Expect(board.Blocks[4].GongWeiName).Should(Equal(gong.ZiNu))
					Expect(board.Blocks[5].GongWeiName).Should(Equal(gong.FuQi))
				})
				It("should returns correct Ming Ju", func() {
					Expect(board.MingJu.JuType).Should(Equal(mingju.Jin))
					Expect(board.MingJu.Number).Should(Equal(uint(4)))
				})
			})
			When("given a birthday start with Chou hour", func() {
				BeforeEach(func() {
					// luna date: 丙寅年 7 / 3 子時
					birthday = time.Date(1984, 12, 8, 0, 4, 0, 0, time.Local)
				})
				It("shouldn't returns any errors", func() {
					Expect(err).Should(BeNil())
				})
				It("should display Gong Wei by birthday", func() {
					Expect(board.Blocks[10].GongWeiName).Should(Equal(gong.XiongDi))
					Expect(board.Blocks[11].GongWeiName).Should(Equal(gong.Ming))
					Expect(board.Blocks[0].GongWeiName).Should(Equal(gong.FuMu))
					Expect(board.Blocks[1].GongWeiName).Should(Equal(gong.FuDe))
					Expect(board.Blocks[2].GongWeiName).Should(Equal(gong.TianZhai))
					Expect(board.Blocks[3].GongWeiName).Should(Equal(gong.ShiYe))
					Expect(board.Blocks[4].GongWeiName).Should(Equal(gong.JiaoYou))
					Expect(board.Blocks[5].GongWeiName).Should(Equal(gong.QianYi))
					Expect(board.Blocks[6].GongWeiName).Should(Equal(gong.JiBing))
					Expect(board.Blocks[7].GongWeiName).Should(Equal(gong.CaiBo))
					Expect(board.Blocks[8].GongWeiName).Should(Equal(gong.ZiNu))
					Expect(board.Blocks[9].GongWeiName).Should(Equal(gong.FuQi))
				})
				It("should returns correct Ming Ju", func() {
					Expect(board.MingJu.JuType).Should(Equal(mingju.Huo))
					Expect(board.MingJu.Number).Should(Equal(uint(6)))
				})
			})
		})
		Context("should display correct fourteen main stars", func() {
			Context("zi wei star", func() {
				When("given a birthday shui ju and 12th", func() {
					BeforeEach(func() {
						// luna date: 甲子年 7 / 12 子時
						birthday = time.Date(1984, 8, 8, 0, 4, 0, 0, time.Local)
					})
					It("shouldn't returns any errors", func() {
						Expect(err).Should(BeNil())
					})
					It("should returns correct Ming Ju", func() {
						Expect(board.MingJu.JuType).Should(Equal(mingju.Shui))
						Expect(board.MingJu.Number).Should(Equal(uint(2)))
					})
					It("should dispaly zi wei in wei location", func() {
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZiWei.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
					})
					It("should dispaly tian fu in you location", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianFu.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
					})
					It("should display tian ji, tai yang, wu qu, tian tong, lian zhen", func() {
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianJi.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TaiYang.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WuQu.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianTong.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.LianZhen.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
					})
					It("should display tai yin, tan lang, ju men, tian xiang, tian liang, qi sha, po jun", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TaiYin.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TanLang.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JuMen.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianXiang.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianLiang.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.QiSha.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.PoJun.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
					})
				})
				When("given a birthday Huo ju and 26th", func() {
					BeforeEach(func() {
						// luna date: 戊辰年 6 / 26 子時
						birthday = time.Date(1988, 8, 8, 0, 4, 0, 0, time.Local)
					})
					It("shouldn't returns any errors", func() {
						Expect(err).Should(BeNil())
					})
					It("should returns correct Ming Ju", func() {
						Expect(board.MingJu.JuType).Should(Equal(mingju.Huo))
						Expect(board.MingJu.Number).Should(Equal(uint(6)))
					})
					It("should dispaly in xu location", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZiWei.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
					})
					It("should display tian ji, tai yang, wu qu, tian tong, lian zhen", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianJi.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TaiYang.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WuQu.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianTong.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.LianZhen.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
					})
					It("should display tai yin, tan lang, ju men, tian xiang, tian liang, qi sha, po jun", func() {
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TaiYin.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TanLang.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JuMen.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianXiang.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianLiang.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.QiSha.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.PoJun.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
					})
				})
				When("given a birthday Huo ju and 12th", func() {
					BeforeEach(func() {
						// luna date: 戊辰年 6 / 26 子時
						birthday = time.Date(1988, 8, 8, 0, 4, 0, 0, time.Local)
					})
					It("shouldn't returns any errors", func() {
						Expect(err).Should(BeNil())
					})
					It("should returns correct Ming Ju", func() {
						Expect(board.MingJu.JuType).Should(Equal(mingju.Huo))
						Expect(board.MingJu.Number).Should(Equal(uint(6)))
					})
					It("should dispaly in xu location", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZiWei.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
					})
					It("should display tian ji, tai yang, wu qu, tian tong, lian zhen", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianJi.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TaiYang.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WuQu.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianTong.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.LianZhen.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
					})
					It("should display tai yin, tan lang, ju men, tian xiang, tian liang, qi sha, po jun", func() {
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TaiYin.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TanLang.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JuMen.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianXiang.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianLiang.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.QiSha.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.PoJun.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
					})
				})
				When("given a birthday Tu ju and 12th", func() {
					BeforeEach(func() {
						// luna date: 庚子年 6 / 18 子時
						birthday = time.Date(1990, 8, 8, 0, 4, 0, 0, time.Local)
					})
					It("shouldn't returns any errors", func() {
						Expect(err).Should(BeNil())
					})
					It("should returns correct Ming Ju", func() {
						Expect(board.MingJu.JuType).Should(Equal(mingju.Tu))
						Expect(board.MingJu.Number).Should(Equal(uint(5)))
					})
					It("should dispaly in wei location", func() {
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZiWei.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
					})
					It("should display tian ji, tai yang, wu qu, tian tong, lian zhen", func() {
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianJi.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TaiYang.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WuQu.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianTong.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.LianZhen.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
					})
				})
				When("given a birthday Mu ju and 18th", func() {
					BeforeEach(func() {
						// luna date: 壬申年 6 / 30 巳時
						birthday = time.Date(1992, 7, 29, 9, 4, 0, 0, time.Local)
					})
					It("shouldn't returns any errors", func() {
						Expect(err).Should(BeNil())
					})
					It("should returns correct Ming Ju", func() {
						Expect(board.MingJu.JuType).Should(Equal(mingju.Jin))
						Expect(board.MingJu.Number).Should(Equal(uint(4)))
					})
					It("should dispaly in hei location", func() {
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZiWei.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
					})
					It("should display tian ji, tai yang, wu qu, tian tong, lian zhen", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianJi.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TaiYang.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WuQu.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianTong.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.LianZhen.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
					})
					It("should display tai yin, tan lang, ju men, tian xiang, tian liang, qi sha, po jun", func() {
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TaiYin.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TanLang.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JuMen.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianXiang.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianLiang.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.QiSha.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.PoJun.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
					})
				})
				When("given a birthday Mu ju and 18th", func() {
					BeforeEach(func() {
						// luna date: 癸酉年 5 / 1 申時
						birthday = time.Date(1993, 6, 20, 15, 4, 0, 0, time.Local)
					})
					It("should returns correct Ming Ju", func() {
						Expect(board.MingJu.JuType).Should(Equal(mingju.Shui))
						Expect(board.MingJu.Number).Should(Equal(uint(2)))
					})
					It("should dispaly in zi location", func() {
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZiWei.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
					})
					It("should display tian ji, tai yang, wu qu, tian tong, lian zhen", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianJi.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TaiYang.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WuQu.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianTong.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.LianZhen.String(),
							StarType: startype.FourteenMainStars,
							Location: 0,
							MiaoXian: nil,
						}))
					})
				})
				It("should display tai yin, tan lang, ju men, tian xiang, tian liang, qi sha, po jun", func() {
					Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TaiYin.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TanLang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JuMen.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianXiang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianLiang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.QiSha.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.PoJun.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
				})
			})
		})
		Context("tain fu star", func() {
			When("zi wei is in si location", func() {
				BeforeEach(func() {
					// luna date: 癸酉年 5 / 19 戌時
					birthday = time.Date(1993, 7, 8, 19, 4, 0, 0, time.Local)
				})
				It("should display zi wei star in si location", func() {
					Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.ZiWei.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
				})
				It("should display tain fu star in hai location", func() {
					Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianFu.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
				})
				It("should display tian ji, tai yang, wu qu, tian tong, lian zhen", func() {
					Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianJi.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TaiYang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.WuQu.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianTong.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LianZhen.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
				})
				It("should display tai yin, tan lang, ju men, tian xiang, tian liang, qi sha, po jun", func() {
					Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TaiYin.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TanLang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JuMen.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianXiang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianLiang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.QiSha.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.PoJun.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
				})
			})
			When("zi wei is in hai location", func() {
				BeforeEach(func() {
					// luna date: 癸酉年 5 / 1 戌時
					birthday = time.Date(1993, 6, 20, 19, 4, 0, 0, time.Local)
				})
				It("should display zi wei star in hai location", func() {
					Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.ZiWei.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
				})
				It("should display tain fu star in si location", func() {
					Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianFu.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
				})
				It("should display tian ji, tai yang, wu qu, tian tong, lian zhen", func() {
					Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianJi.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TaiYang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.WuQu.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianTong.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LianZhen.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
				})
				It("should display tai yin, tan lang, ju men, tian xiang, tian liang, qi sha, po jun", func() {
					Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TaiYin.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TanLang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JuMen.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianXiang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianLiang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.QiSha.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.PoJun.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
				})
			})
			When("zi wei is in yin location", func() {
				BeforeEach(func() {
					// luna date: 癸酉年 5 / 2 辰時
					birthday = time.Date(1993, 6, 21, 3, 4, 0, 0, time.Local)
				})
				It("should display zi wei star in yin location", func() {
					Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.ZiWei.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
				})
				It("should display tain fu star in si location", func() {
					Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianFu.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
				})
				It("should display tian ji, tai yang, wu qu, tian tong, lian zhen", func() {
					Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianJi.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TaiYang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.WuQu.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianTong.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LianZhen.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
				})
				It("should display tai yin, tan lang, ju men, tian xiang, tian liang, qi sha, po jun", func() {
					Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TaiYin.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TanLang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JuMen.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianXiang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianLiang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.QiSha.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.PoJun.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
				})
			})
			When("zi wei is in shen location", func() {
				BeforeEach(func() {
					// luna date: 癸酉年 5 / 13 午時
					birthday = time.Date(1993, 7, 2, 0, 4, 0, 0, time.Local)
				})
				It("should display zi wei star in shen location", func() {
					Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.ZiWei.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
				})
				It("should display tain fu star in shen location", func() {
					Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianFu.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
				})
				It("should display tian ji, tai yang, wu qu, tian tong, lian zhen", func() {
					Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianJi.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TaiYang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.WuQu.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianTong.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LianZhen.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
				})
				It("should display tai yin, tan lang, ju men, tian xiang, tian liang, qi sha, po jun", func() {
					Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TaiYin.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TanLang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JuMen.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianXiang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianLiang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.QiSha.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.PoJun.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
				})
			})
			When("zi wei is in chou location", func() {
				BeforeEach(func() {
					// luna date: 癸酉年 5 / 3 子時
					birthday = time.Date(1993, 6, 22, 6, 4, 0, 0, time.Local)
				})
				It("should display zi wei star in shen location", func() {
					Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.ZiWei.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
				})
				It("should display tain fu star in shen location", func() {
					Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianFu.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
				})
				It("should display tian ji, tai yang, wu qu, tian tong, lian zhen", func() {
					Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianJi.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TaiYang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.WuQu.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianTong.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LianZhen.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
				})
				It("should display tai yin, tan lang, ju men, tian xiang, tian liang, qi sha, po jun", func() {
					Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TaiYin.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TanLang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JuMen.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianXiang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianLiang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.QiSha.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.PoJun.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
				})
			})
			When("zi wei is in zi location", func() {
				BeforeEach(func() {
					// luna date: 癸酉年 5 / 25 午時
					birthday = time.Date(1993, 7, 14, 12, 4, 0, 0, time.Local)
				})
				It("should display zi wei star in shen location", func() {
					Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.ZiWei.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
				})
				It("should display tain fu star in shen location", func() {
					Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianFu.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
				})

				It("should display tian ji, tai yang, wu qu, tian tong, lian zhen", func() {
					Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianJi.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TaiYang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.WuQu.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianTong.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LianZhen.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
				})

				It("should display tai yin, tan lang, ju men, tian xiang, tian liang, qi sha, po jun", func() {
					Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TaiYin.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TanLang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JuMen.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianXiang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianLiang.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.QiSha.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
					Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.PoJun.String(),
						StarType: startype.FourteenMainStars,
						Location: 0,
						MiaoXian: nil,
					}))
				})
			})
		})
		Context("should display correct nain gan stars", func() {
			BeforeEach(func() {
				// luna date: 癸酉年 5 / 25 午時
				birthday = time.Date(1993, 7, 14, 12, 4, 0, 0, time.Local)
			})
			It("should display correct stars location", func() {
				Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
					Name:     stars.LuCun.String(),
					StarType: startype.NianGanXiZhuXing,
				}))
				Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
					Name:     stars.QingYang.String(),
					StarType: startype.NianGanXiZhuXing,
				}))
				Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
					Name:     stars.TuoLuo.String(),
					StarType: startype.NianGanXiZhuXing,
				}))
				Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
					Name:     stars.TianKui.String(),
					StarType: startype.NianGanXiZhuXing,
				}))
				Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
					Name:     stars.TianYue.String(),
					StarType: startype.NianGanXiZhuXing,
				}))
				Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
					Name:     stars.TianGuan.String(),
					StarType: startype.NianGanXiZhuXing,
				}))
				Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
					Name:     stars.TianFu.String(),
					StarType: startype.NianGanXiZhuXing,
				}))
			})
		})
		Context("setXunKong()", func() {
			When("birth year is in 癸酉", func() {
				BeforeEach(func() {
					// luna date: 癸酉年 5 / 25 午時
					birthday = time.Date(1993, 7, 14, 12, 4, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.XunKong.String(),
						StarType: startype.NianGanXiZhuXing,
					}))
				})
			})
			When("birth year is in 乙亥", func() {
				BeforeEach(func() {
					// luna date: 乙亥年 5 / 25 午時
					birthday = time.Date(1995, 6, 22, 12, 4, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.XunKong.String(),
						StarType: startype.NianGanXiZhuXing,
					}))
				})
			})
		})
		Context("setJieKong()", func() {
			When("birth year is in 甲戌", func() {
				BeforeEach(func() {
					// luna date: 甲戌年 5 / 14 午時
					birthday = time.Date(1994, 6, 22, 12, 4, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JieKong.String(),
						StarType: startype.NianGanXiZhuXing,
					}))
					Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JieKong.String(),
						StarType: startype.NianGanXiZhuXing,
					}))
				})
			})
			When("birth year is in 乙亥", func() {
				BeforeEach(func() {
					// luna date: 乙亥年 5 / 25 午時
					birthday = time.Date(1995, 6, 22, 12, 4, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JieKong.String(),
						StarType: startype.NianGanXiZhuXing,
					}))
					Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JieKong.String(),
						StarType: startype.NianGanXiZhuXing,
					}))
				})
			})
			When("birth year is in 丙子", func() {
				BeforeEach(func() {
					// luna date: 丙子年 5 / 7 午時
					birthday = time.Date(1996, 6, 22, 12, 4, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JieKong.String(),
						StarType: startype.NianGanXiZhuXing,
					}))
					Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JieKong.String(),
						StarType: startype.NianGanXiZhuXing,
					}))
				})
			})
			When("birth year is in 丁丑", func() {
				BeforeEach(func() {
					// luna date: 丁丑年 5 / 18 午時
					birthday = time.Date(1997, 6, 22, 12, 4, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JieKong.String(),
						StarType: startype.NianGanXiZhuXing,
					}))
					Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JieKong.String(),
						StarType: startype.NianGanXiZhuXing,
					}))
				})
			})
			When("birth year is in 戊寅", func() {
				BeforeEach(func() {
					// luna date: 戊寅年 5 / 28 午時
					birthday = time.Date(1998, 6, 22, 12, 4, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JieKong.String(),
						StarType: startype.NianGanXiZhuXing,
					}))
					Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JieKong.String(),
						StarType: startype.NianGanXiZhuXing,
					}))
				})
			})
			When("birth year is in 己卯", func() {
				BeforeEach(func() {
					// luna date: 己卯年 5 / 9 午時
					birthday = time.Date(1999, 6, 22, 12, 4, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JieKong.String(),
						StarType: startype.NianGanXiZhuXing,
					}))
					Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JieKong.String(),
						StarType: startype.NianGanXiZhuXing,
					}))
				})
			})
			When("birth year is in 庚辰", func() {
				BeforeEach(func() {
					// luna date: 庚辰年 5 / 21 午時
					birthday = time.Date(2000, 6, 22, 12, 4, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JieKong.String(),
						StarType: startype.NianGanXiZhuXing,
					}))
					Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JieKong.String(),
						StarType: startype.NianGanXiZhuXing,
					}))
				})
			})
			When("birth year is in 辛巳", func() {
				BeforeEach(func() {
					// luna date: 辛巳年 5 / 2 午時
					birthday = time.Date(2001, 6, 22, 12, 4, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JieKong.String(),
						StarType: startype.NianGanXiZhuXing,
					}))
					Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JieKong.String(),
						StarType: startype.NianGanXiZhuXing,
					}))
				})
			})
			When("birth year is in 壬午", func() {
				BeforeEach(func() {
					// luna date: 壬午年 5 / 12 午時
					birthday = time.Date(2002, 6, 22, 12, 4, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JieKong.String(),
						StarType: startype.NianGanXiZhuXing,
					}))
					Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JieKong.String(),
						StarType: startype.NianGanXiZhuXing,
					}))
				})
			})
			When("birth year is in 癸酉", func() {
				BeforeEach(func() {
					// luna date: 癸酉年 5 / 25 午時
					birthday = time.Date(1993, 7, 14, 12, 4, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JieKong.String(),
						StarType: startype.NianGanXiZhuXing,
					}))
					Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.JieKong.String(),
						StarType: startype.NianGanXiZhuXing,
					}))
				})
			})
		})
		Context("setNianZhiXiZhuXing()", func() {
			Context("setTainKu()", func() {
				When("birth year is in 癸酉", func() {
					BeforeEach(func() {
						// luna date: 癸酉年 5 / 25 午時
						birthday = time.Date(1993, 7, 14, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianKu.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 乙亥", func() {
					BeforeEach(func() {
						// luna date: 乙亥年 5 / 25 午時
						birthday = time.Date(1995, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianKu.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 辛巳", func() {
					BeforeEach(func() {
						// luna date: 辛巳年 5 / 2 午時
						birthday = time.Date(2001, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianKu.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
			})
			Context("setTainXu()", func() {
				When("birth year is in 癸酉", func() {
					BeforeEach(func() {
						// luna date: 癸酉年 5 / 25 午時
						birthday = time.Date(1993, 7, 14, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianXu.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 乙亥", func() {
					BeforeEach(func() {
						// luna date: 乙亥年 5 / 25 午時
						birthday = time.Date(1995, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianXu.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 辛巳", func() {
					BeforeEach(func() {
						// luna date: 辛巳年 5 / 2 午時
						birthday = time.Date(2001, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianXu.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
			})
			Context("setLongChi()", func() {
				When("birth year is in 癸酉", func() {
					BeforeEach(func() {
						// luna date: 癸酉年 5 / 25 午時
						birthday = time.Date(1993, 7, 14, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.LongChi.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 乙亥", func() {
					BeforeEach(func() {
						// luna date: 乙亥年 5 / 25 午時
						birthday = time.Date(1995, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.LongChi.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 辛巳", func() {
					BeforeEach(func() {
						// luna date: 辛巳年 5 / 2 午時
						birthday = time.Date(2001, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.LongChi.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
			})
			Context("setFengGe()", func() {
				When("birth year is in 癸酉", func() {
					BeforeEach(func() {
						// luna date: 癸酉年 5 / 25 午時
						birthday = time.Date(1993, 7, 14, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FengGe.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 乙亥", func() {
					BeforeEach(func() {
						// luna date: 乙亥年 5 / 25 午時
						birthday = time.Date(1995, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FengGe.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 辛巳", func() {
					BeforeEach(func() {
						// luna date: 辛巳年 5 / 2 午時
						birthday = time.Date(2001, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FengGe.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
			})
			Context("setHongLuan()", func() {
				When("birth year is in 癸酉", func() {
					BeforeEach(func() {
						// luna date: 癸酉年 5 / 25 午時
						birthday = time.Date(1993, 7, 14, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.HongLuan.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 乙亥", func() {
					BeforeEach(func() {
						// luna date: 乙亥年 5 / 25 午時
						birthday = time.Date(1995, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.HongLuan.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 辛巳", func() {
					BeforeEach(func() {
						// luna date: 辛巳年 5 / 2 午時
						birthday = time.Date(2001, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.HongLuan.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
			})
			Context("setTianXi()", func() {
				When("birth year is in 癸酉", func() {
					BeforeEach(func() {
						// luna date: 癸酉年 5 / 25 午時
						birthday = time.Date(1993, 7, 14, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianXi.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 乙亥", func() {
					BeforeEach(func() {
						// luna date: 乙亥年 5 / 25 午時
						birthday = time.Date(1995, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianXi.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 辛巳", func() {
					BeforeEach(func() {
						// luna date: 辛巳年 5 / 2 午時
						birthday = time.Date(2001, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianXi.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
			})
			Context("setGuChen()", func() {
				When("birth year is in 癸酉", func() {
					BeforeEach(func() {
						// luna date: 癸酉年 5 / 25 午時
						birthday = time.Date(1993, 7, 14, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.GuChen.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 乙亥", func() {
					BeforeEach(func() {
						// luna date: 乙亥年 5 / 25 午時
						birthday = time.Date(1995, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.GuChen.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 辛巳", func() {
					BeforeEach(func() {
						// luna date: 辛巳年 5 / 2 午時
						birthday = time.Date(2001, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.GuChen.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 戊寅", func() {
					BeforeEach(func() {
						// luna date: 戊寅年 5 / 8 午時
						birthday = time.Date(1998, 6, 2, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.GuChen.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 丙戌", func() {
					BeforeEach(func() {
						// luna date: 丙戌年 5 / 12 午時
						birthday = time.Date(2006, 6, 7, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.GuChen.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
			})
			Context("setGuaXiu()", func() {
				When("birth year is in 癸酉", func() {
					BeforeEach(func() {
						// luna date: 癸酉年 5 / 25 午時
						birthday = time.Date(1993, 7, 14, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.GuaXiu.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 乙亥", func() {
					BeforeEach(func() {
						// luna date: 乙亥年 5 / 25 午時
						birthday = time.Date(1995, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.GuaXiu.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 辛巳", func() {
					BeforeEach(func() {
						// luna date: 辛巳年 5 / 2 午時
						birthday = time.Date(2001, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.GuaXiu.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 戊寅", func() {
					BeforeEach(func() {
						// luna date: 戊寅年 5 / 8 午時
						birthday = time.Date(1998, 6, 2, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.GuaXiu.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 丙戌", func() {
					BeforeEach(func() {
						// luna date: 丙戌年 5 / 12 午時
						birthday = time.Date(2006, 6, 7, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.GuaXiu.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
			})
			Context("setJieShen()", func() {
				When("birth year is in 癸酉", func() {
					BeforeEach(func() {
						// luna date: 癸酉年 5 / 25 午時
						birthday = time.Date(1993, 7, 14, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JieShen.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 乙亥", func() {
					BeforeEach(func() {
						// luna date: 乙亥年 5 / 25 午時
						birthday = time.Date(1995, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JieShen.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 辛巳", func() {
					BeforeEach(func() {
						// luna date: 辛巳年 5 / 2 午時
						birthday = time.Date(2001, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JieShen.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
			})
			Context("setPoSui()", func() {
				When("birth year is in 癸酉", func() {
					BeforeEach(func() {
						// luna date: 癸酉年 5 / 25 午時
						birthday = time.Date(1993, 7, 14, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.PoSui.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 乙亥", func() {
					BeforeEach(func() {
						// luna date: 乙亥年 5 / 25 午時
						birthday = time.Date(1995, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.PoSui.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 辛巳", func() {
					BeforeEach(func() {
						// luna date: 辛巳年 5 / 2 午時
						birthday = time.Date(2001, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.PoSui.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 庚辰", func() {
					BeforeEach(func() {
						// luna date: 庚辰年 9 / 8 午時
						birthday = time.Date(2000, 10, 5, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.PoSui.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
			})
			Context("setTianMa()", func() {
				When("birth year is in 癸酉", func() {
					BeforeEach(func() {
						// luna date: 癸酉年 5 / 25 午時
						birthday = time.Date(1993, 7, 14, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianMa.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 乙亥", func() {
					BeforeEach(func() {
						// luna date: 乙亥年 5 / 25 午時
						birthday = time.Date(1995, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianMa.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 辛巳", func() {
					BeforeEach(func() {
						// luna date: 辛巳年 5 / 2 午時
						birthday = time.Date(2001, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianMa.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 庚辰", func() {
					BeforeEach(func() {
						// luna date: 庚辰年 9 / 8 午時
						birthday = time.Date(2000, 10, 5, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianMa.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 庚寅", func() {
					BeforeEach(func() {
						// luna date: 庚寅年 8 / 28 午時
						birthday = time.Date(2010, 10, 5, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianMa.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
			})
			Context("setDaHao()", func() {
				When("birth year is in 癸酉", func() {
					BeforeEach(func() {
						// luna date: 癸酉年 5 / 25 午時
						birthday = time.Date(1993, 7, 14, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DaHao.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 乙亥", func() {
					BeforeEach(func() {
						// luna date: 乙亥年 5 / 25 午時
						birthday = time.Date(1995, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DaHao.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 辛巳", func() {
					BeforeEach(func() {
						// luna date: 辛巳年 5 / 2 午時
						birthday = time.Date(2001, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DaHao.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 庚辰", func() {
					BeforeEach(func() {
						// luna date: 庚辰年 9 / 8 午時
						birthday = time.Date(2000, 10, 5, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DaHao.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 庚寅", func() {
					BeforeEach(func() {
						// luna date: 庚寅年 8 / 28 午時
						birthday = time.Date(2010, 10, 5, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DaHao.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
			})
			Context("setTainDe()", func() {
				When("birth year is in 癸酉", func() {
					BeforeEach(func() {
						// luna date: 癸酉年 5 / 25 午時
						birthday = time.Date(1993, 7, 14, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianDe.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 乙亥", func() {
					BeforeEach(func() {
						// luna date: 乙亥年 5 / 25 午時
						birthday = time.Date(1995, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianDe.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 辛巳", func() {
					BeforeEach(func() {
						// luna date: 辛巳年 5 / 2 午時
						birthday = time.Date(2001, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianDe.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
			})
			Context("setJieSha()", func() {
				When("birth year is in 癸酉", func() {
					BeforeEach(func() {
						// luna date: 癸酉年 5 / 25 午時
						birthday = time.Date(1993, 7, 14, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JieSha.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 乙亥", func() {
					BeforeEach(func() {
						// luna date: 乙亥年 5 / 25 午時
						birthday = time.Date(1995, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JieSha.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 辛巳", func() {
					BeforeEach(func() {
						// luna date: 辛巳年 5 / 2 午時
						birthday = time.Date(2001, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JieSha.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 庚辰", func() {
					BeforeEach(func() {
						// luna date: 庚辰年 9 / 8 午時
						birthday = time.Date(2000, 10, 5, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JieSha.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 庚寅", func() {
					BeforeEach(func() {
						// luna date: 庚寅年 8 / 28 午時
						birthday = time.Date(2010, 10, 5, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JieSha.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
			})
			Context("setHuaGai()", func() {
				When("birth year is in 癸酉", func() {
					BeforeEach(func() {
						// luna date: 癸酉年 5 / 25 午時
						birthday = time.Date(1993, 7, 14, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.HuaGai.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 乙亥", func() {
					BeforeEach(func() {
						// luna date: 乙亥年 5 / 25 午時
						birthday = time.Date(1995, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.HuaGai.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 辛巳", func() {
					BeforeEach(func() {
						// luna date: 辛巳年 5 / 2 午時
						birthday = time.Date(2001, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.HuaGai.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 庚辰", func() {
					BeforeEach(func() {
						// luna date: 庚辰年 9 / 8 午時
						birthday = time.Date(2000, 10, 5, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.HuaGai.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 庚寅", func() {
					BeforeEach(func() {
						// luna date: 庚寅年 8 / 28 午時
						birthday = time.Date(2010, 10, 5, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.HuaGai.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
			})
			Context("setXianChi()", func() {
				When("birth year is in 癸酉", func() {
					BeforeEach(func() {
						// luna date: 癸酉年 5 / 25 午時
						birthday = time.Date(1993, 7, 14, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XianChi.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 乙亥", func() {
					BeforeEach(func() {
						// luna date: 乙亥年 5 / 25 午時
						birthday = time.Date(1995, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XianChi.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 辛巳", func() {
					BeforeEach(func() {
						// luna date: 辛巳年 5 / 2 午時
						birthday = time.Date(2001, 6, 22, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XianChi.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 庚辰", func() {
					BeforeEach(func() {
						// luna date: 庚辰年 9 / 8 午時
						birthday = time.Date(2000, 10, 5, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XianChi.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 庚寅", func() {
					BeforeEach(func() {
						// luna date: 庚寅年 8 / 28 午時
						birthday = time.Date(2010, 10, 5, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XianChi.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
			})
			Context("setTianCai()", func() {
				When("birth year is in 癸酉", func() {
					BeforeEach(func() {
						// luna date: 癸酉年 5 / 25 午時
						// 命宮: 子
						// 身宮: 子
						birthday = time.Date(1993, 7, 14, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianCai.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 辛巳", func() {
					BeforeEach(func() {
						// luna date: 辛巳年 5 / 2 寅時
						// 命宮: 辰
						// 身宮: 申
						birthday = time.Date(2001, 6, 22, 4, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianCai.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 庚寅", func() {
					BeforeEach(func() {
						// luna date: 庚寅年 8 / 28 未時
						// 命宮: 寅
						// 身宮: 辰
						birthday = time.Date(2010, 10, 5, 13, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianCai.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 辛巳", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 9 / 2 巳時
						// 命宮: 巳
						// 身宮: 卯
						birthday = time.Date(1970, 10, 1, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianCai.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
			})
			Context("setTianShou()", func() {
				When("birth year is in 癸酉", func() {
					BeforeEach(func() {
						// luna date: 癸酉年 5 / 25 午時
						// 命宮: 子
						// 身宮: 子
						birthday = time.Date(1993, 7, 14, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianShou.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 辛巳", func() {
					BeforeEach(func() {
						// luna date: 辛巳年 5 / 2 寅時
						// 命宮: 辰
						// 身宮: 申
						birthday = time.Date(2001, 6, 22, 4, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianShou.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 庚寅", func() {
					BeforeEach(func() {
						// luna date: 庚寅年 8 / 28 未時
						// 命宮: 寅
						// 身宮: 辰
						birthday = time.Date(2010, 10, 5, 13, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianShou.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 辛巳", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 9 / 2 巳時
						// 命宮: 巳
						// 身宮: 卯
						birthday = time.Date(1970, 10, 1, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianShou.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
			})
			Context("setTianKong()", func() {
				When("birth year is in 癸酉", func() {
					BeforeEach(func() {
						// luna date: 癸酉年 5 / 25 午時
						birthday = time.Date(1993, 7, 14, 12, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianKong.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 辛巳", func() {
					BeforeEach(func() {
						// luna date: 辛巳年 5 / 2 寅時
						birthday = time.Date(2001, 6, 22, 4, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianKong.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 庚寅", func() {
					BeforeEach(func() {
						// luna date: 庚寅年 8 / 28 未時
						birthday = time.Date(2010, 10, 5, 13, 4, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianKong.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
				When("birth year is in 辛巳", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 9 / 2 巳時
						birthday = time.Date(1970, 10, 1, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianKong.String(),
							StarType: startype.NianZhiXiZhuXing,
						}))
					})
				})
			})
		})
		Context("setYueXiXing()", func() {
			Context("setZuoFu()", func() {
				When("birth month is in January", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 1 / 28 巳時
						birthday = time.Date(1970, 3, 5, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZuoFu.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in Febuary", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 2 / 29 巳時
						birthday = time.Date(1970, 4, 5, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZuoFu.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in March", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 3 / 28 巳時
						birthday = time.Date(1970, 5, 3, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZuoFu.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in April", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 4 / 1 巳時
						birthday = time.Date(1970, 5, 5, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZuoFu.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in May", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 5 / 7 巳時
						birthday = time.Date(1970, 6, 10, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZuoFu.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in June", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 6 / 6 巳時
						birthday = time.Date(1970, 7, 8, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZuoFu.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in July", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 7 / 11 巳時
						birthday = time.Date(1970, 8, 12, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZuoFu.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in Auguest", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 8 / 9 巳時
						birthday = time.Date(1970, 9, 9, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZuoFu.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in September", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 9 / 9 巳時
						birthday = time.Date(1970, 10, 8, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZuoFu.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in October", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 10 / 13 巳時
						birthday = time.Date(1970, 11, 11, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZuoFu.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in November", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 11 / 11 巳時
						birthday = time.Date(1970, 12, 9, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZuoFu.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in December", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 巳時
						birthday = time.Date(1971, 1, 8, 9, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZuoFu.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
			})
			Context("setYouBi()", func() {
				When("birth month is in January", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 1 / 28 巳時
						birthday = time.Date(1970, 3, 5, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YouBi.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in Febuary", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 2 / 29 巳時
						birthday = time.Date(1970, 4, 5, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YouBi.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in March", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 3 / 28 巳時
						birthday = time.Date(1970, 5, 3, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YouBi.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in April", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 4 / 1 巳時
						birthday = time.Date(1970, 5, 5, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YouBi.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in May", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 5 / 7 巳時
						birthday = time.Date(1970, 6, 10, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YouBi.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in June", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 6 / 6 巳時
						birthday = time.Date(1970, 7, 8, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YouBi.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in July", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 7 / 11 巳時
						birthday = time.Date(1970, 8, 12, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YouBi.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in Auguest", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 8 / 9 巳時
						birthday = time.Date(1970, 9, 9, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YouBi.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in September", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 9 / 9 巳時
						birthday = time.Date(1970, 10, 8, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YouBi.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in October", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 10 / 13 巳時
						birthday = time.Date(1970, 11, 11, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YouBi.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in November", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 11 / 11 巳時
						birthday = time.Date(1970, 12, 9, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YouBi.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in December", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 巳時
						birthday = time.Date(1971, 1, 8, 9, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YouBi.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
			})
			Context("setTianXing()", func() {
				When("birth month is in January", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 1 / 28 巳時
						birthday = time.Date(1970, 3, 5, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianXing.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in Febuary", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 2 / 29 巳時
						birthday = time.Date(1970, 4, 5, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianXing.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in March", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 3 / 28 巳時
						birthday = time.Date(1970, 5, 3, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianXing.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in April", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 4 / 1 巳時
						birthday = time.Date(1970, 5, 5, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianXing.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in May", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 5 / 7 巳時
						birthday = time.Date(1970, 6, 10, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianXing.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in June", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 6 / 6 巳時
						birthday = time.Date(1970, 7, 8, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianXing.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in July", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 7 / 11 巳時
						birthday = time.Date(1970, 8, 12, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianXing.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in Auguest", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 8 / 9 巳時
						birthday = time.Date(1970, 9, 9, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianXing.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in September", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 9 / 9 巳時
						birthday = time.Date(1970, 10, 8, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianXing.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in October", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 10 / 13 巳時
						birthday = time.Date(1970, 11, 11, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianXing.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in November", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 11 / 11 巳時
						birthday = time.Date(1970, 12, 9, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianXing.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in December", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 巳時
						birthday = time.Date(1971, 1, 8, 9, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianXing.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
			})
			Context("setTianYao()", func() {
				When("birth month is in January", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 1 / 28 巳時
						birthday = time.Date(1970, 3, 5, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianYao.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in Febuary", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 2 / 29 巳時
						birthday = time.Date(1970, 4, 5, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianYao.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in March", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 3 / 28 巳時
						birthday = time.Date(1970, 5, 3, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianYao.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in April", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 4 / 1 巳時
						birthday = time.Date(1970, 5, 5, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianYao.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in May", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 5 / 7 巳時
						birthday = time.Date(1970, 6, 10, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianYao.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in June", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 6 / 6 巳時
						birthday = time.Date(1970, 7, 8, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianYao.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in July", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 7 / 11 巳時
						birthday = time.Date(1970, 8, 12, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianYao.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in Auguest", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 8 / 9 巳時
						birthday = time.Date(1970, 9, 9, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianYao.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in September", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 9 / 9 巳時
						birthday = time.Date(1970, 10, 8, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianYao.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in October", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 10 / 13 巳時
						birthday = time.Date(1970, 11, 11, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianYao.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in November", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 11 / 11 巳時
						birthday = time.Date(1970, 12, 9, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianYao.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in December", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 巳時
						birthday = time.Date(1971, 1, 8, 9, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianYao.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
			})
			Context("setTianWu()", func() {
				When("birth month is in January", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 1 / 28 巳時
						birthday = time.Date(1970, 3, 5, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianWu.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in Febuary", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 2 / 29 巳時
						birthday = time.Date(1970, 4, 5, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianWu.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in March", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 3 / 28 巳時
						birthday = time.Date(1970, 5, 3, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianWu.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in April", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 4 / 1 巳時
						birthday = time.Date(1970, 5, 5, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianWu.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in May", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 5 / 7 巳時
						birthday = time.Date(1970, 6, 10, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianWu.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in June", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 6 / 6 巳時
						birthday = time.Date(1970, 7, 8, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianWu.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in July", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 7 / 11 巳時
						birthday = time.Date(1970, 8, 12, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianWu.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in Auguest", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 8 / 9 巳時
						birthday = time.Date(1970, 9, 9, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianWu.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in September", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 9 / 9 巳時
						birthday = time.Date(1970, 10, 8, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianWu.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in October", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 10 / 13 巳時
						birthday = time.Date(1970, 11, 11, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianWu.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in November", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 11 / 11 巳時
						birthday = time.Date(1970, 12, 9, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianWu.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in December", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 巳時
						birthday = time.Date(1971, 1, 8, 9, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianWu.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
			})
			Context("setYueXiXingTianYue()", func() {
				When("birth month is in January", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 1 / 28 巳時
						birthday = time.Date(1970, 3, 5, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YueXiXingTianYue.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in Febuary", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 2 / 29 巳時
						birthday = time.Date(1970, 4, 5, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YueXiXingTianYue.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in March", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 3 / 28 巳時
						birthday = time.Date(1970, 5, 3, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YueXiXingTianYue.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in April", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 4 / 1 巳時
						birthday = time.Date(1970, 5, 5, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YueXiXingTianYue.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in May", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 5 / 7 巳時
						birthday = time.Date(1970, 6, 10, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YueXiXingTianYue.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in June", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 6 / 6 巳時
						birthday = time.Date(1970, 7, 8, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YueXiXingTianYue.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in July", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 7 / 11 巳時
						birthday = time.Date(1970, 8, 12, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YueXiXingTianYue.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in Auguest", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 8 / 9 巳時
						birthday = time.Date(1970, 9, 9, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YueXiXingTianYue.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in September", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 9 / 9 巳時
						birthday = time.Date(1970, 10, 8, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YueXiXingTianYue.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in October", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 10 / 13 巳時
						birthday = time.Date(1970, 11, 11, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YueXiXingTianYue.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in November", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 11 / 11 巳時
						birthday = time.Date(1970, 12, 9, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YueXiXingTianYue.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in December", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 巳時
						birthday = time.Date(1971, 1, 8, 9, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YueXiXingTianYue.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
			})
			Context("setYinSha()", func() {
				When("birth month is in January", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 1 / 28 巳時
						birthday = time.Date(1970, 3, 5, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YinSha.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in Febuary", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 2 / 29 巳時
						birthday = time.Date(1970, 4, 5, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YinSha.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in March", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 3 / 28 巳時
						birthday = time.Date(1970, 5, 3, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YinSha.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in April", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 4 / 1 巳時
						birthday = time.Date(1970, 5, 5, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YinSha.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in May", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 5 / 7 巳時
						birthday = time.Date(1970, 6, 10, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YinSha.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in June", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 6 / 6 巳時
						birthday = time.Date(1970, 7, 8, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YinSha.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in July", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 7 / 11 巳時
						birthday = time.Date(1970, 8, 12, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YinSha.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in Auguest", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 8 / 9 巳時
						birthday = time.Date(1970, 9, 9, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YinSha.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in September", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 9 / 9 巳時
						birthday = time.Date(1970, 10, 8, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YinSha.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in October", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 10 / 13 巳時
						birthday = time.Date(1970, 11, 11, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YinSha.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in November", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 11 / 11 巳時
						birthday = time.Date(1970, 12, 9, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YinSha.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
				When("birth month is in December", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 巳時
						birthday = time.Date(1971, 1, 8, 9, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.YinSha.String(),
							StarType: startype.YueXiXing,
						}))
					})
				})
			})
		})
		Context("setShiXiZhuXing", func() {
			Context("setWenChang", func() {
				When("birth hour is in 子時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 子時
						birthday = time.Date(1971, 1, 8, 24, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WenChang.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 丑時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 丑時
						birthday = time.Date(1971, 1, 8, 2, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WenChang.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 寅時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 寅時
						birthday = time.Date(1971, 1, 8, 4, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WenChang.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 卯時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 卯時
						birthday = time.Date(1971, 1, 8, 6, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WenChang.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 辰時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 辰時
						birthday = time.Date(1971, 1, 8, 8, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WenChang.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 巳時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 巳時
						birthday = time.Date(1971, 1, 8, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WenChang.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 午時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 午時
						birthday = time.Date(1971, 1, 8, 12, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WenChang.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 未時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 未時
						birthday = time.Date(1971, 1, 8, 14, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WenChang.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 申時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 申時
						birthday = time.Date(1971, 1, 8, 16, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WenChang.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 酉時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 酉時
						birthday = time.Date(1971, 1, 8, 18, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WenChang.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 戌時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 戌時
						birthday = time.Date(1971, 1, 8, 20, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WenChang.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 亥時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 亥時
						birthday = time.Date(1971, 1, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WenChang.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
			})
			Context("setWenQu()", func() {
				When("birth hour is in 子時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 子時
						birthday = time.Date(1971, 1, 8, 24, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WenQu.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 丑時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 丑時
						birthday = time.Date(1971, 1, 8, 2, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WenQu.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 寅時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 寅時
						birthday = time.Date(1971, 1, 8, 4, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WenQu.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 卯時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 卯時
						birthday = time.Date(1971, 1, 8, 6, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WenQu.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 辰時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 辰時
						birthday = time.Date(1971, 1, 8, 8, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WenQu.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 巳時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 巳時
						birthday = time.Date(1971, 1, 8, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WenQu.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 午時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 午時
						birthday = time.Date(1971, 1, 8, 12, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WenQu.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 未時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 未時
						birthday = time.Date(1971, 1, 8, 14, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WenQu.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 申時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 申時
						birthday = time.Date(1971, 1, 8, 16, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WenQu.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 酉時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 酉時
						birthday = time.Date(1971, 1, 8, 18, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WenQu.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 戌時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 戌時
						birthday = time.Date(1971, 1, 8, 20, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WenQu.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 亥時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 亥時
						birthday = time.Date(1971, 1, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.WenQu.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
			})
			Context("setDiJie()", func() {
				When("birth hour is in 子時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 子時
						birthday = time.Date(1971, 1, 8, 24, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DiJie.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 丑時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 丑時
						birthday = time.Date(1971, 1, 8, 2, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DiJie.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 寅時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 寅時
						birthday = time.Date(1971, 1, 8, 4, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DiJie.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 卯時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 卯時
						birthday = time.Date(1971, 1, 8, 6, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DiJie.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 辰時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 辰時
						birthday = time.Date(1971, 1, 8, 8, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DiJie.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 巳時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 巳時
						birthday = time.Date(1971, 1, 8, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DiJie.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 午時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 午時
						birthday = time.Date(1971, 1, 8, 12, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DiJie.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 未時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 未時
						birthday = time.Date(1971, 1, 8, 14, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DiJie.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 申時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 申時
						birthday = time.Date(1971, 1, 8, 16, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DiJie.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 酉時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 酉時
						birthday = time.Date(1971, 1, 8, 18, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DiJie.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 戌時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 戌時
						birthday = time.Date(1971, 1, 8, 20, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DiJie.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 亥時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 亥時
						birthday = time.Date(1971, 1, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DiJie.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
			})
			Context("setDiKong()", func() {
				When("birth hour is in 子時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 子時
						birthday = time.Date(1971, 1, 8, 24, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DiKong.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 丑時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 丑時
						birthday = time.Date(1971, 1, 8, 2, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DiKong.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 寅時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 寅時
						birthday = time.Date(1971, 1, 8, 4, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DiKong.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 卯時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 卯時
						birthday = time.Date(1971, 1, 8, 6, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DiKong.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 辰時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 辰時
						birthday = time.Date(1971, 1, 8, 8, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DiKong.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 巳時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 巳時
						birthday = time.Date(1971, 1, 8, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DiKong.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 午時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 午時
						birthday = time.Date(1971, 1, 8, 12, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DiKong.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 未時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 未時
						birthday = time.Date(1971, 1, 8, 14, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DiKong.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 申時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 申時
						birthday = time.Date(1971, 1, 8, 16, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DiKong.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 酉時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 酉時
						birthday = time.Date(1971, 1, 8, 18, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DiKong.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 戌時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 戌時
						birthday = time.Date(1971, 1, 8, 20, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DiKong.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 亥時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 亥時
						birthday = time.Date(1971, 1, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.DiKong.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
			})
			Context("setTaiFu()", func() {
				When("birth hour is in 子時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 子時
						birthday = time.Date(1971, 1, 8, 24, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TaiFu.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 丑時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 丑時
						birthday = time.Date(1971, 1, 8, 2, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TaiFu.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 寅時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 寅時
						birthday = time.Date(1971, 1, 8, 4, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TaiFu.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 卯時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 卯時
						birthday = time.Date(1971, 1, 8, 6, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TaiFu.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 辰時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 辰時
						birthday = time.Date(1971, 1, 8, 8, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TaiFu.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 巳時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 巳時
						birthday = time.Date(1971, 1, 8, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TaiFu.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 午時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 午時
						birthday = time.Date(1971, 1, 8, 12, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TaiFu.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 未時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 未時
						birthday = time.Date(1971, 1, 8, 14, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TaiFu.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 申時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 申時
						birthday = time.Date(1971, 1, 8, 16, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TaiFu.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 酉時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 酉時
						birthday = time.Date(1971, 1, 8, 18, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TaiFu.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 戌時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 戌時
						birthday = time.Date(1971, 1, 8, 20, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TaiFu.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 亥時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 亥時
						birthday = time.Date(1971, 1, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TaiFu.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
			})
			Context("setFengGao()", func() {
				When("birth hour is in 子時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 子時
						birthday = time.Date(1971, 1, 8, 24, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FengGao.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 丑時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 丑時
						birthday = time.Date(1971, 1, 8, 2, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FengGao.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 寅時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 寅時
						birthday = time.Date(1971, 1, 8, 4, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FengGao.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 卯時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 卯時
						birthday = time.Date(1971, 1, 8, 6, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FengGao.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 辰時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 辰時
						birthday = time.Date(1971, 1, 8, 8, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FengGao.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 巳時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 巳時
						birthday = time.Date(1971, 1, 8, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FengGao.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 午時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 午時
						birthday = time.Date(1971, 1, 8, 12, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FengGao.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 未時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 未時
						birthday = time.Date(1971, 1, 8, 14, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FengGao.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 申時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 申時
						birthday = time.Date(1971, 1, 8, 16, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FengGao.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 酉時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 酉時
						birthday = time.Date(1971, 1, 8, 18, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FengGao.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 戌時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 戌時
						birthday = time.Date(1971, 1, 8, 20, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FengGao.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth hour is in 亥時", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 12 亥時
						birthday = time.Date(1971, 1, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FengGao.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
			})
			Context("setHuo()", func() {
				// group 1
				When("birth year is in 寅 and birth hour is 子", func() {
					BeforeEach(func() {
						// luna date: 庚寅年 1 / 20 亥時
						birthday = time.Date(1950, 3, 8, 0, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.Huo.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth year is in 午 and birth hour is 亥", func() {
					BeforeEach(func() {
						// luna date: 戊午年 1 / 30 亥時
						birthday = time.Date(1978, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.Huo.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth year is in 戌 and birth hour is 巳", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 3 巳時
						birthday = time.Date(1971, 1, 11, 10, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.Huo.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				//group 2
				When("birth year is in 申 and birth hour is 子", func() {
					BeforeEach(func() {
						// luna date: 丙申年 1 / 26 亥時
						birthday = time.Date(1956, 3, 8, 0, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.Huo.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth year is in 子 and birth hour is 亥", func() {
					BeforeEach(func() {
						// luna date: 庚子年 2 / 11 亥時
						birthday = time.Date(1960, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.Huo.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth year is in 辰 and birth hour is 巳", func() {
					BeforeEach(func() {
						// luna date: 戊辰年 2 / 17 巳時
						birthday = time.Date(1928, 3, 8, 10, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.Huo.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				// group 3
				When("birth year is in 巳 and birth hour is 子", func() {
					BeforeEach(func() {
						// luna date: 辛巳年 2 / 11 亥時
						birthday = time.Date(1941, 3, 8, 0, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.Huo.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth year is in 酉 and birth hour is 亥", func() {
					BeforeEach(func() {
						// luna date: 乙酉年 1 / 24 亥時
						birthday = time.Date(1945, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.Huo.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth year is in 丑 and birth hour is 巳", func() {
					BeforeEach(func() {
						// luna date: 己丑年 2 / 9 巳時
						birthday = time.Date(1949, 3, 8, 10, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.Huo.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				// group 4
				When("birth year is in 亥 and birth hour is 子", func() {
					BeforeEach(func() {
						// luna date: 己亥年 1 / 29 亥時
						birthday = time.Date(1959, 3, 8, 0, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.Huo.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth year is in 卯 and birth hour is 亥", func() {
					BeforeEach(func() {
						// luna date: 癸卯年 1 / 24 亥時
						birthday = time.Date(1963, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.Huo.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth year is in 未 and birth hour is 巳", func() {
					BeforeEach(func() {
						// luna date: 丁未年 1 / 28 巳時
						birthday = time.Date(1967, 3, 8, 10, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.Huo.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
			})
			Context("setLing()", func() {
				// group 1
				When("birth year is in 寅 and birth hour is 子", func() {
					BeforeEach(func() {
						// luna date: 庚寅年 1 / 20 亥時
						birthday = time.Date(1950, 3, 8, 0, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.Ling.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth year is in 午 and birth hour is 亥", func() {
					BeforeEach(func() {
						// luna date: 戊午年 1 / 30 亥時
						birthday = time.Date(1978, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.Ling.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth year is in 戌 and birth hour is 巳", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 12 / 3 巳時
						birthday = time.Date(1971, 1, 11, 10, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.Ling.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				//group 2
				When("birth year is in 申 and birth hour is 子", func() {
					BeforeEach(func() {
						// luna date: 丙申年 1 / 26 亥時
						birthday = time.Date(1956, 3, 8, 0, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.Ling.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth year is in 子 and birth hour is 亥", func() {
					BeforeEach(func() {
						// luna date: 庚子年 2 / 11 亥時
						birthday = time.Date(1960, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.Ling.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth year is in 辰 and birth hour is 巳", func() {
					BeforeEach(func() {
						// luna date: 戊辰年 2 / 17 巳時
						birthday = time.Date(1928, 3, 8, 10, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.Ling.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				// group 3
				When("birth year is in 巳 and birth hour is 子", func() {
					BeforeEach(func() {
						// luna date: 辛巳年 2 / 11 亥時
						birthday = time.Date(1941, 3, 8, 0, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.Ling.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth year is in 酉 and birth hour is 亥", func() {
					BeforeEach(func() {
						// luna date: 乙酉年 1 / 24 亥時
						birthday = time.Date(1945, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.Ling.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth year is in 丑 and birth hour is 巳", func() {
					BeforeEach(func() {
						// luna date: 己丑年 2 / 9 巳時
						birthday = time.Date(1949, 3, 8, 10, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.Ling.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				// group 4
				When("birth year is in 亥 and birth hour is 子", func() {
					BeforeEach(func() {
						// luna date: 己亥年 1 / 29 亥時
						birthday = time.Date(1959, 3, 8, 0, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.Ling.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth year is in 卯 and birth hour is 亥", func() {
					BeforeEach(func() {
						// luna date: 癸卯年 1 / 24 亥時
						birthday = time.Date(1963, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.Ling.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth year is in 未 and birth hour is 巳", func() {
					BeforeEach(func() {
						// luna date: 丁未年 1 / 28 巳時
						birthday = time.Date(1967, 3, 8, 10, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.Ling.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
			})
			Context("setSanTai()", func() {
				When("birth year is in 戌 and birth hour is 巳", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 9 / 2 巳時
						birthday = time.Date(1970, 10, 1, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.SanTai.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth year is in 卯 and birth hour is 亥", func() {
					BeforeEach(func() {
						// luna date: 癸卯年 1 / 24 亥時
						birthday = time.Date(1963, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.SanTai.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
			})
			Context("setBaZuo()", func() {
				When("birth year is in 戌 and birth hour is 巳", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 9 / 2 巳時
						birthday = time.Date(1970, 10, 1, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BaZuo.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth year is in 卯 and birth hour is 亥", func() {
					BeforeEach(func() {
						// luna date: 癸卯年 1 / 13 亥時
						birthday = time.Date(1963, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BaZuo.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
			})
			Context("setEnGuang()", func() {
				When("birth year is in 戌 and birth hour is 巳", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 9 / 2 巳時
						birthday = time.Date(1970, 10, 1, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.EnGuang.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth year is in 卯 and birth hour is 亥", func() {
					BeforeEach(func() {
						// luna date: 癸卯年 1 / 13 亥時
						birthday = time.Date(1963, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.EnGuang.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
			})
			Context("setTianGui()", func() {
				When("birth year is in 戌 and birth hour is 巳", func() {
					BeforeEach(func() {
						// luna date: 庚戌年 9 / 2 巳時
						birthday = time.Date(1970, 10, 1, 10, 10, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianGui.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
				When("birth year is in 卯 and birth hour is 亥", func() {
					BeforeEach(func() {
						// luna date: 癸卯年 1 / 13 亥時
						birthday = time.Date(1963, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct star location", func() {
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.TianGui.String(),
							StarType: startype.ShiXiZhuXing,
						}))
					})
				})
			})
		})
		Context("setMingZhu()", func() {
			When("birth year is in 子 and birth hour is 亥", func() {
				BeforeEach(func() {
					// luna date: 庚子年 2 / 11 亥時
					birthday = time.Date(1960, 3, 8, 22, 12, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[board.StarsMap[stars.TanLang]].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.MingZhu.String(),
						StarType: startype.ShenMing,
					}))
				})
			})
			When("birth year is in 丑 and birth hour is 巳", func() {
				BeforeEach(func() {
					// luna date: 己丑年 2 / 9 巳時
					birthday = time.Date(1949, 3, 8, 10, 12, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[board.StarsMap[stars.JuMen]].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.MingZhu.String(),
						StarType: startype.ShenMing,
					}))
				})
			})
			When("birth year is in 寅 and birth hour is 子", func() {
				BeforeEach(func() {
					// luna date: 庚寅年 1 / 20 亥時
					birthday = time.Date(1950, 3, 8, 0, 12, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[board.StarsMap[stars.LuCun]].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.MingZhu.String(),
						StarType: startype.ShenMing,
					}))
				})
			})
			When("birth year is in 卯 and birth hour is 亥", func() {
				BeforeEach(func() {
					// luna date: 癸卯年 1 / 24 亥時
					birthday = time.Date(1963, 3, 8, 22, 12, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[board.StarsMap[stars.WenQu]].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.MingZhu.String(),
						StarType: startype.ShenMing,
					}))
				})
			})
			When("birth year is in 辰 and birth hour is 巳", func() {
				BeforeEach(func() {
					// luna date: 戊辰年 2 / 17 巳時
					birthday = time.Date(1928, 3, 8, 10, 12, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[board.StarsMap[stars.LianZhen]].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.MingZhu.String(),
						StarType: startype.ShenMing,
					}))
				})
			})
			When("birth year is in 巳 and birth hour is 子", func() {
				BeforeEach(func() {
					// luna date: 辛巳年 2 / 11 亥時
					birthday = time.Date(1941, 3, 8, 0, 12, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[board.StarsMap[stars.WuQu]].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.MingZhu.String(),
						StarType: startype.ShenMing,
					}))
				})
			})
			When("birth year is in 午 and birth hour is 亥", func() {
				BeforeEach(func() {
					// luna date: 戊午年 1 / 30 亥時
					birthday = time.Date(1978, 3, 8, 22, 12, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[board.StarsMap[stars.PoJun]].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.MingZhu.String(),
						StarType: startype.ShenMing,
					}))
				})
			})
			When("birth year is in 未 and birth hour is 巳", func() {
				BeforeEach(func() {
					// luna date: 丁未年 1 / 28 巳時
					birthday = time.Date(1967, 3, 8, 10, 12, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[board.StarsMap[stars.WuQu]].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.MingZhu.String(),
						StarType: startype.ShenMing,
					}))
				})
			})
			When("birth year is in 申 and birth hour is 子", func() {
				BeforeEach(func() {
					// luna date: 丙申年 1 / 26 亥時
					birthday = time.Date(1956, 3, 8, 0, 12, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[board.StarsMap[stars.LianZhen]].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.MingZhu.String(),
						StarType: startype.ShenMing,
					}))
				})
			})
			When("birth year is in 酉 and birth hour is 亥", func() {
				BeforeEach(func() {
					// luna date: 乙酉年 1 / 24 亥時
					birthday = time.Date(1945, 3, 8, 22, 12, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[board.StarsMap[stars.WenQu]].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.MingZhu.String(),
						StarType: startype.ShenMing,
					}))
				})
			})
			When("birth year is in 戌 and birth hour is 巳", func() {
				BeforeEach(func() {
					// luna date: 庚戌年 12 / 3 巳時
					birthday = time.Date(1971, 1, 11, 10, 12, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[board.StarsMap[stars.LuCun]].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.MingZhu.String(),
						StarType: startype.ShenMing,
					}))
				})
			})
			When("birth year is in 亥 and birth hour is 子", func() {
				BeforeEach(func() {
					// luna date: 己亥年 1 / 29 亥時
					birthday = time.Date(1959, 3, 8, 0, 12, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[board.StarsMap[stars.JuMen]].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.MingZhu.String(),
						StarType: startype.ShenMing,
					}))
				})
			})
		})
		Context("setShenZhu()", func() {
			When("birth year is in 子 and birth hour is 亥", func() {
				BeforeEach(func() {
					// luna date: 庚子年 2 / 11 亥時
					birthday = time.Date(1960, 3, 8, 22, 12, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[board.StarsMap[stars.Ling]].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.ShenZhu.String(),
						StarType: startype.ShenMing,
					}))
				})
			})
			When("birth year is in 丑 and birth hour is 巳", func() {
				BeforeEach(func() {
					// luna date: 己丑年 2 / 9 巳時
					birthday = time.Date(1949, 3, 8, 10, 12, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[board.StarsMap[stars.TianXiang]].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.ShenZhu.String(),
						StarType: startype.ShenMing,
					}))
				})
			})
			When("birth year is in 寅 and birth hour is 子", func() {
				BeforeEach(func() {
					// luna date: 庚寅年 1 / 20 亥時
					birthday = time.Date(1950, 3, 8, 0, 12, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[board.StarsMap[stars.TianLiang]].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.ShenZhu.String(),
						StarType: startype.ShenMing,
					}))
				})
			})
			When("birth year is in 卯 and birth hour is 亥", func() {
				BeforeEach(func() {
					// luna date: 癸卯年 1 / 24 亥時
					birthday = time.Date(1963, 3, 8, 22, 12, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[board.StarsMap[stars.TianTong]].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.ShenZhu.String(),
						StarType: startype.ShenMing,
					}))
				})
			})
			When("birth year is in 辰 and birth hour is 巳", func() {
				BeforeEach(func() {
					// luna date: 戊辰年 2 / 17 巳時
					birthday = time.Date(1928, 3, 8, 10, 12, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[board.StarsMap[stars.WenChang]].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.ShenZhu.String(),
						StarType: startype.ShenMing,
					}))
				})
			})
			When("birth year is in 巳 and birth hour is 子", func() {
				BeforeEach(func() {
					// luna date: 辛巳年 2 / 11 亥時
					birthday = time.Date(1941, 3, 8, 0, 12, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[board.StarsMap[stars.TianJi]].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.ShenZhu.String(),
						StarType: startype.ShenMing,
					}))
				})
			})
			When("birth year is in 午 and birth hour is 亥", func() {
				BeforeEach(func() {
					// luna date: 戊午年 1 / 30 亥時
					birthday = time.Date(1978, 3, 8, 22, 12, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[board.StarsMap[stars.Huo]].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.ShenZhu.String(),
						StarType: startype.ShenMing,
					}))
				})
			})
			When("birth year is in 未 and birth hour is 巳", func() {
				BeforeEach(func() {
					// luna date: 丁未年 1 / 28 巳時
					birthday = time.Date(1967, 3, 8, 10, 12, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[board.StarsMap[stars.TianXiang]].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.ShenZhu.String(),
						StarType: startype.ShenMing,
					}))
				})
			})
			When("birth year is in 申 and birth hour is 子", func() {
				BeforeEach(func() {
					// luna date: 丙申年 1 / 26 亥時
					birthday = time.Date(1956, 3, 8, 0, 12, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[board.StarsMap[stars.TianLiang]].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.ShenZhu.String(),
						StarType: startype.ShenMing,
					}))
				})
			})
			When("birth year is in 酉 and birth hour is 亥", func() {
				BeforeEach(func() {
					// luna date: 乙酉年 1 / 24 亥時
					birthday = time.Date(1945, 3, 8, 22, 12, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[board.StarsMap[stars.TianTong]].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.ShenZhu.String(),
						StarType: startype.ShenMing,
					}))
				})
			})
			When("birth year is in 戌 and birth hour is 巳", func() {
				BeforeEach(func() {
					// luna date: 庚戌年 12 / 3 巳時
					birthday = time.Date(1971, 1, 11, 10, 12, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[board.StarsMap[stars.WenChang]].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.ShenZhu.String(),
						StarType: startype.ShenMing,
					}))
				})
			})
			When("birth year is in 亥 and birth hour is 子", func() {
				BeforeEach(func() {
					// luna date: 己亥年 1 / 29 亥時
					birthday = time.Date(1959, 3, 8, 0, 12, 0, 0, time.Local)
				})
				It("should display correct star location", func() {
					Expect(board.Blocks[board.StarsMap[stars.TianJi]].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.ShenZhu.String(),
						StarType: startype.ShenMing,
					}))
				})
			})
		})
		Context("setGender()", func() {
			When("gender is male", func() {
				BeforeEach(func() {
					gender = genders.Male
				})
				When("birth year is in 甲", func() {
					BeforeEach(func() {
						// luna date: 甲午年 2 / 4 亥時
						birthday = time.Date(1954, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct gender", func() {
						Expect(board.Gender).Should(Equal(genders.YangMale))
					})
				})
				When("birth year is in 乙", func() {
					BeforeEach(func() {
						// luna date: 乙未年 2 / 15 亥時
						birthday = time.Date(1955, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct gender", func() {
						Expect(board.Gender).Should(Equal(genders.YinMale))
					})
				})
				When("birth year is in 丙", func() {
					BeforeEach(func() {
						// luna date: 丙申年 2 / 7 亥時
						birthday = time.Date(1956, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct gender", func() {
						Expect(board.Gender).Should(Equal(genders.YangMale))
					})
				})
				When("birth year is in 丁", func() {
					BeforeEach(func() {
						// luna date: 丁酉年 2 / 7 亥時
						birthday = time.Date(1957, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct gender", func() {
						Expect(board.Gender).Should(Equal(genders.YinMale))
					})
				})
				When("birth year is in 戊", func() {
					BeforeEach(func() {
						// luna date: 戊戌年 1 / 19 亥時
						birthday = time.Date(1958, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct gender", func() {
						Expect(board.Gender).Should(Equal(genders.YangMale))
					})
				})
				When("birth year is in 己", func() {
					BeforeEach(func() {
						// luna date: 己亥年 1 / 29 亥時
						birthday = time.Date(1959, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct gender", func() {
						Expect(board.Gender).Should(Equal(genders.YinMale))
					})
				})
				When("birth year is in 庚", func() {
					BeforeEach(func() {
						// luna date: 庚子年 2 / 11 亥時
						birthday = time.Date(1960, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct gender", func() {
						Expect(board.Gender).Should(Equal(genders.YangMale))
					})
				})
				When("birth year is in 辛", func() {
					BeforeEach(func() {
						// luna date: 辛丑年 1 / 22 亥時
						birthday = time.Date(1961, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct gender", func() {
						Expect(board.Gender).Should(Equal(genders.YinMale))
					})
				})
				When("birth year is in 壬", func() {
					BeforeEach(func() {
						// luna date: 壬寅年 2 / 3 亥時
						birthday = time.Date(1962, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct gender", func() {
						Expect(board.Gender).Should(Equal(genders.YangMale))
					})
				})
				When("birth year is in 癸", func() {
					BeforeEach(func() {
						// luna date: 癸卯年 2 / 13 亥時
						birthday = time.Date(1963, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct gender", func() {
						Expect(board.Gender).Should(Equal(genders.YinMale))
					})
				})
			})
			When("gender is female", func() {
				BeforeEach(func() {
					gender = genders.Female
				})
				When("birth year is in 甲", func() {
					BeforeEach(func() {
						// luna date: 甲午年 2 / 4 亥時
						birthday = time.Date(1954, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct gender", func() {
						Expect(board.Gender).Should(Equal(genders.YangFemale))
					})
				})
				When("birth year is in 乙", func() {
					BeforeEach(func() {
						// luna date: 乙未年 2 / 15 亥時
						birthday = time.Date(1955, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct gender", func() {
						Expect(board.Gender).Should(Equal(genders.YinFemale))
					})
				})
				When("birth year is in 丙", func() {
					BeforeEach(func() {
						// luna date: 丙申年 2 / 7 亥時
						birthday = time.Date(1956, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct gender", func() {
						Expect(board.Gender).Should(Equal(genders.YangFemale))
					})
				})
				When("birth year is in 丁", func() {
					BeforeEach(func() {
						// luna date: 丁酉年 2 / 7 亥時
						birthday = time.Date(1957, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct gender", func() {
						Expect(board.Gender).Should(Equal(genders.YinFemale))
					})
				})
				When("birth year is in 戊", func() {
					BeforeEach(func() {
						// luna date: 戊戌年 1 / 19 亥時
						birthday = time.Date(1958, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct gender", func() {
						Expect(board.Gender).Should(Equal(genders.YangFemale))
					})
				})
				When("birth year is in 己", func() {
					BeforeEach(func() {
						// luna date: 己亥年 1 / 29 亥時
						birthday = time.Date(1959, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct gender", func() {
						Expect(board.Gender).Should(Equal(genders.YinFemale))
					})
				})
				When("birth year is in 庚", func() {
					BeforeEach(func() {
						// luna date: 庚子年 2 / 11 亥時
						birthday = time.Date(1960, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct gender", func() {
						Expect(board.Gender).Should(Equal(genders.YangFemale))
					})
				})
				When("birth year is in 辛", func() {
					BeforeEach(func() {
						// luna date: 辛丑年 1 / 22 亥時
						birthday = time.Date(1961, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct gender", func() {
						Expect(board.Gender).Should(Equal(genders.YinFemale))
					})
				})
				When("birth year is in 壬", func() {
					BeforeEach(func() {
						// luna date: 壬寅年 2 / 3 亥時
						birthday = time.Date(1962, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct gender", func() {
						Expect(board.Gender).Should(Equal(genders.YangFemale))
					})
				})
				When("birth year is in 癸", func() {
					BeforeEach(func() {
						// luna date: 癸卯年 2 / 13 亥時
						birthday = time.Date(1963, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct gender", func() {
						Expect(board.Gender).Should(Equal(genders.YinFemale))
					})
				})
			})
		})
		Context("setBoShiTwelveStars()", func() {
			When("LuCen is in 寅 and gender", func() {
				When("gender is YangMale", func() {
					BeforeEach(func() {
						// luna date: 甲午年 2 / 4 亥時
						birthday = time.Date(1954, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct stars", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FuBing.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.GuanFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.LiShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.QingLong.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiaoHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JiangJun.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZouShu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FeiLian.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiShen.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BingFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShiDaHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
					})
				})
				When("gender is YangFemale", func() {
					BeforeEach(func() {
						gender = genders.Female
					})
					It("should display correct stars", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.QingLong.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.LiShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.GuanFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FuBing.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShiDaHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BingFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiShen.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FeiLian.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZouShu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JiangJun.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiaoHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
					})
				})
			})
			When("LuCen is in 卯", func() {
				BeforeEach(func() {
					// luna date: 乙未年 2 / 15 亥時
					birthday = time.Date(1955, 3, 8, 22, 12, 0, 0, time.Local)
				})
				When("gender is YingMale", func() {
					It("should display correct stars", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiaoHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.QingLong.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.LiShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.GuanFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FuBing.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShiDaHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BingFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiShen.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FeiLian.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZouShu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JiangJun.String(),
							StarType: startype.BoShiTwelveStars,
						}))
					})
				})
				When("gender is YingFemale", func() {
					BeforeEach(func() {
						gender = genders.Female
					})
					It("should display correct stars", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShiDaHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FuBing.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.GuanFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.LiShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.QingLong.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiaoHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JiangJun.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZouShu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FeiLian.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiShen.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BingFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
					})
				})
			})
			When("LuCen is in 巳", func() {
				BeforeEach(func() {
					// luna date: 丙申年 2 / 7 亥時
					birthday = time.Date(1956, 3, 8, 22, 12, 0, 0, time.Local)
				})
				When("gender is YingMale", func() {
					It("should display correct stars", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiShen.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BingFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShiDaHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FuBing.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.GuanFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.LiShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.QingLong.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiaoHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JiangJun.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZouShu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FeiLian.String(),
							StarType: startype.BoShiTwelveStars,
						}))
					})
				})
				When("gender is YangFemale", func() {
					BeforeEach(func() {
						gender = genders.Female
					})
					It("should display correct stars", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZouShu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JiangJun.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiaoHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.QingLong.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.LiShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.GuanFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FuBing.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShiDaHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BingFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiShen.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FeiLian.String(),
							StarType: startype.BoShiTwelveStars,
						}))
					})
				})
			})
			When("LuCen is in 午", func() {
				BeforeEach(func() {
					// luna date: 丁酉年 2 / 7 亥時
					birthday = time.Date(1957, 3, 8, 22, 12, 0, 0, time.Local)
				})
				When("gender is YingMale", func() {
					It("should display correct stars", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FeiLian.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZouShu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JiangJun.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiaoHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.QingLong.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.LiShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.GuanFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FuBing.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShiDaHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BingFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiShen.String(),
							StarType: startype.BoShiTwelveStars,
						}))
					})
				})
				When("gender is YingFemale", func() {
					BeforeEach(func() {
						gender = genders.Female
					})
					It("should display correct stars", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FeiLian.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiShen.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BingFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShiDaHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FuBing.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.GuanFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.LiShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.QingLong.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiaoHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JiangJun.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZouShu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
					})
				})
			})
			When("LuCen is in 申", func() {
				BeforeEach(func() {
					// luna date: 庚子年 2 / 11 亥時
					birthday = time.Date(1960, 3, 8, 22, 12, 0, 0, time.Local)
				})
				When("gender is YangMale", func() {
					It("should display correct stars", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JiangJun.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZouShu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FeiLian.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiShen.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BingFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShiDaHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FuBing.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.GuanFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.LiShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.QingLong.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiaoHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
					})
				})
				When("gender is YangFemale", func() {
					BeforeEach(func() {
						gender = genders.Female
					})
					It("should display correct stars", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BingFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiShen.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FeiLian.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZouShu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JiangJun.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiaoHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.QingLong.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.LiShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.GuanFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FuBing.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShiDaHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
					})
				})
			})
			When("LuCen is in 酉", func() {
				BeforeEach(func() {
					// luna date: 辛丑年 1 / 22 亥時
					birthday = time.Date(1961, 3, 8, 22, 12, 0, 0, time.Local)
				})
				When("gender is YingMale", func() {
					It("should display correct stars", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShiDaHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BingFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiShen.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FeiLian.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZouShu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JiangJun.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiaoHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.QingLong.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.LiShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.GuanFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FuBing.String(),
							StarType: startype.BoShiTwelveStars,
						}))
					})
				})
				When("gender is YingFemale", func() {
					BeforeEach(func() {
						gender = genders.Female
					})
					It("should display correct stars", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiaoHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JiangJun.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZouShu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FeiLian.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiShen.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BingFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShiDaHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FuBing.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.GuanFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.LiShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.QingLong.String(),
							StarType: startype.BoShiTwelveStars,
						}))
					})
				})
			})
			When("LuCen is in 亥", func() {
				BeforeEach(func() {
					// luna date: 壬寅年 2 / 3 亥時
					birthday = time.Date(1962, 3, 8, 22, 12, 0, 0, time.Local)
				})
				When("gender is YangMale", func() {
					It("should display correct stars", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.LiShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.QingLong.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiaoHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JiangJun.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZouShu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FeiLian.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiShen.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BingFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShiDaHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FuBing.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.GuanFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
					})
				})
				When("gender is YangFemale", func() {
					BeforeEach(func() {
						gender = genders.Female
					})
					It("should display correct stars", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.GuanFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FuBing.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShiDaHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BingFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiShen.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FeiLian.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZouShu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JiangJun.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiaoHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.QingLong.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.LiShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
					})
				})
			})
			When("LuCen is in　子", func() {
				BeforeEach(func() {
					// luna date: 癸卯年 2 / 13 亥時
					birthday = time.Date(1963, 3, 8, 22, 12, 0, 0, time.Local)
				})
				When("gender is YingMale", func() {
					It("should display correct stars", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.GuanFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FuBing.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShiDaHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BingFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiShen.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FeiLian.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZouShu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JiangJun.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiaoHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.QingLong.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.LiShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
					})
				})
				When("gender is YingFemale", func() {
					BeforeEach(func() {
						gender = genders.Female
					})
					It("should display correct stars", func() {
						Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.LiShi.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.QingLong.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiaoHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.JiangJun.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.ZouShu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FeiLian.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.XiShen.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BingFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.BoShiDaHao.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.FuBing.String(),
							StarType: startype.BoShiTwelveStars,
						}))
						Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.GuanFu.String(),
							StarType: startype.BoShiTwelveStars,
						}))
					})
				})
			})
		})
		Context("setLiuNianSuiQianZhuXing()", func() {
			var patch *monkey.PatchGuard
			When("current year is in 子", func() {
				BeforeEach(func() {
					// luna date: 庚子年 2 / 11 亥時
					birthday = time.Date(1960, 3, 8, 22, 12, 0, 0, time.Local)
					patch = monkey.Patch(time.Now, func() time.Time { return birthday })
				})
				AfterEach(func() {
					patch.Unpatch()
				})
				It("should display correct stars", func() {
					Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.SuiJian.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.HuiQi.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.SangMen.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.GuanSuo.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.GuanFu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianXiaoHao.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianDaHao.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LongDe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.BaiHu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianDe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.DiKe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianBingFu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
				})
			})
			When("current year is in 丑", func() {
				BeforeEach(func() {
					// luna date: 己丑年 2 / 9 巳時
					birthday = time.Date(1949, 3, 8, 10, 12, 0, 0, time.Local)
					patch = monkey.Patch(time.Now, func() time.Time { return birthday })
				})
				AfterEach(func() {
					patch.Unpatch()
				})
				It("should display correct stars", func() {
					Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianBingFu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.SuiJian.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.HuiQi.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.SangMen.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.GuanSuo.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.GuanFu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianXiaoHao.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianDaHao.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LongDe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.BaiHu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianDe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.DiKe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
				})
			})
			When("current year is in 寅", func() {
				BeforeEach(func() {
					// luna date: 庚寅年 1 / 20 亥時
					birthday = time.Date(1950, 3, 8, 0, 12, 0, 0, time.Local)
					patch = monkey.Patch(time.Now, func() time.Time { return birthday })
				})
				AfterEach(func() {
					patch.Unpatch()
				})
				It("should display correct stars", func() {
					Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.DiKe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianBingFu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.SuiJian.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.HuiQi.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.SangMen.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.GuanSuo.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.GuanFu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianXiaoHao.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianDaHao.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LongDe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.BaiHu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianDe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
				})
			})
			When("current year is in 卯", func() {
				BeforeEach(func() {
					// luna date: 癸卯年 1 / 24 亥時
					birthday = time.Date(1963, 3, 8, 22, 12, 0, 0, time.Local)
					patch = monkey.Patch(time.Now, func() time.Time { return birthday })
				})
				AfterEach(func() {
					patch.Unpatch()
				})
				It("should display correct stars", func() {
					Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianDe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.DiKe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianBingFu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.SuiJian.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.HuiQi.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.SangMen.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.GuanSuo.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.GuanFu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianXiaoHao.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianDaHao.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LongDe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.BaiHu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
				})
			})
			When("current year is in 辰", func() {
				BeforeEach(func() {
					// luna date: 戊辰年 2 / 17 巳時
					birthday = time.Date(1928, 3, 8, 10, 12, 0, 0, time.Local)
					patch = monkey.Patch(time.Now, func() time.Time { return birthday })
				})
				AfterEach(func() {
					patch.Unpatch()
				})
				It("should display correct stars", func() {
					Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.BaiHu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianDe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.DiKe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianBingFu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.SuiJian.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.HuiQi.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.SangMen.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.GuanSuo.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.GuanFu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianXiaoHao.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianDaHao.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LongDe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
				})
			})
			When("current year is in 巳", func() {
				BeforeEach(func() {
					// luna date: 辛巳年 2 / 11 亥時
					birthday = time.Date(1941, 3, 8, 0, 12, 0, 0, time.Local)
					patch = monkey.Patch(time.Now, func() time.Time { return birthday })
				})
				AfterEach(func() {
					patch.Unpatch()
				})
				It("should display correct stars", func() {
					Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LongDe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.BaiHu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianDe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.DiKe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianBingFu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.SuiJian.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.HuiQi.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.SangMen.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.GuanSuo.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.GuanFu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianXiaoHao.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianDaHao.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
				})
			})
			When("current year is in 午", func() {
				BeforeEach(func() {
					// luna date: 戊午年 1 / 30 亥時
					birthday = time.Date(1978, 3, 8, 22, 12, 0, 0, time.Local)
					patch = monkey.Patch(time.Now, func() time.Time { return birthday })
				})
				AfterEach(func() {
					patch.Unpatch()
				})
				It("should display correct stars", func() {
					Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianDaHao.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LongDe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.BaiHu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianDe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.DiKe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianBingFu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.SuiJian.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.HuiQi.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.SangMen.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.GuanSuo.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.GuanFu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianXiaoHao.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
				})
			})
			When("current year is in 未", func() {
				BeforeEach(func() {
					// luna date: 丁未年 1 / 28 巳時
					birthday = time.Date(1967, 3, 8, 10, 12, 0, 0, time.Local)
					patch = monkey.Patch(time.Now, func() time.Time { return birthday })
				})
				AfterEach(func() {
					patch.Unpatch()
				})
				It("should display correct stars", func() {
					Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianXiaoHao.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianDaHao.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LongDe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.BaiHu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianDe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.DiKe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianBingFu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.SuiJian.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.HuiQi.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.SangMen.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.GuanSuo.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.GuanFu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
				})
			})
			When("current year is in 申", func() {
				BeforeEach(func() {
					// luna date: 丙申年 1 / 26 亥時
					birthday = time.Date(1956, 3, 8, 0, 12, 0, 0, time.Local)
					patch = monkey.Patch(time.Now, func() time.Time { return birthday })
				})
				AfterEach(func() {
					patch.Unpatch()
				})
				It("should display correct stars", func() {
					Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.GuanFu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianXiaoHao.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianDaHao.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LongDe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.BaiHu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianDe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.DiKe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianBingFu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.SuiJian.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.HuiQi.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.SangMen.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.GuanSuo.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
				})
			})
			When("current year is in 酉", func() {
				BeforeEach(func() {
					// luna date: 乙酉年 1 / 24 亥時
					birthday = time.Date(1945, 3, 8, 22, 12, 0, 0, time.Local)
					patch = monkey.Patch(time.Now, func() time.Time { return birthday })
				})
				AfterEach(func() {
					patch.Unpatch()
				})
				It("should display correct stars", func() {
					Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.GuanSuo.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.GuanFu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianXiaoHao.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianDaHao.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LongDe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.BaiHu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianDe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.DiKe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianBingFu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.SuiJian.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.HuiQi.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.SangMen.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
				})
			})
			When("current year is in 戌", func() {
				BeforeEach(func() {
					// luna date: 庚戌年 12 / 3 巳時
					birthday = time.Date(1971, 1, 11, 10, 12, 0, 0, time.Local)
					patch = monkey.Patch(time.Now, func() time.Time { return birthday })
				})
				AfterEach(func() {
					patch.Unpatch()
				})
				It("should display correct stars", func() {
					Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.SangMen.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.GuanSuo.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.GuanFu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianXiaoHao.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianDaHao.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LongDe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.BaiHu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianDe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.DiKe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianBingFu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.SuiJian.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.HuiQi.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
				})
			})
			When("current year is in 亥", func() {
				BeforeEach(func() {
					// luna date: 己亥年 1 / 29 亥時
					birthday = time.Date(1959, 3, 8, 0, 12, 0, 0, time.Local)
					patch = monkey.Patch(time.Now, func() time.Time { return birthday })
				})
				AfterEach(func() {
					patch.Unpatch()
				})
				It("should display correct stars", func() {
					Expect(board.Blocks[0].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.HuiQi.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[1].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.SangMen.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[2].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.GuanSuo.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[3].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.GuanFu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[4].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianXiaoHao.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[5].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianDaHao.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[6].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LongDe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[7].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.BaiHu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[8].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.TianDe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[9].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.DiKe.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[10].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.LiuNianBingFu.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
					Expect(board.Blocks[11].Stars).Should(ContainElement(&ziwei.Star{
						Name:     stars.SuiJian.String(),
						StarType: startype.LiuNianSuiQianZhuXing,
					}))
				})
			})
		})
		Context("setSiHua()", func() {
			Context("setHuaLu()", func() {
				When("birth year is in 甲", func() {
					BeforeEach(func() {
						// luna date: 甲午年 2 / 4 亥時
						birthday = time.Date(1954, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct hua lu", func() {
						Expect(board.Blocks[board.StarsMap[stars.LianZhen]].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.HuaLu.String(),
							StarType: startype.SiHua,
						}))
					})
				})
				When("birth year is in 乙", func() {
					BeforeEach(func() {
						// luna date: 乙未年 2 / 15 亥時
						birthday = time.Date(1955, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct hua lu", func() {
						Expect(board.Blocks[board.StarsMap[stars.TianJi]].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.HuaLu.String(),
							StarType: startype.SiHua,
						}))
					})
				})
				When("birth year is in 丙", func() {
					BeforeEach(func() {
						// luna date: 丙申年 2 / 7 亥時
						birthday = time.Date(1956, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct hua lu", func() {
						Expect(board.Blocks[board.StarsMap[stars.TianTong]].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.HuaLu.String(),
							StarType: startype.SiHua,
						}))
					})
				})
				When("birth year is in 丁", func() {
					BeforeEach(func() {
						// luna date: 丁酉年 2 / 7 亥時
						birthday = time.Date(1957, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct hua lu", func() {
						Expect(board.Blocks[board.StarsMap[stars.TaiYin]].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.HuaLu.String(),
							StarType: startype.SiHua,
						}))
					})
				})
				When("birth year is in 戊", func() {
					BeforeEach(func() {
						// luna date: 戊戌年 1 / 19 亥時
						birthday = time.Date(1958, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct hua lu", func() {
						Expect(board.Blocks[board.StarsMap[stars.TanLang]].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.HuaLu.String(),
							StarType: startype.SiHua,
						}))
					})
				})
				When("birth year is in 己", func() {
					BeforeEach(func() {
						// luna date: 己亥年 1 / 29 亥時
						birthday = time.Date(1959, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct hua lu", func() {
						Expect(board.Blocks[board.StarsMap[stars.WuQu]].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.HuaLu.String(),
							StarType: startype.SiHua,
						}))
					})
				})
				When("birth year is in 庚", func() {
					BeforeEach(func() {
						// luna date: 庚子年 2 / 11 亥時
						birthday = time.Date(1960, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct hua lu", func() {
						Expect(board.Blocks[board.StarsMap[stars.TaiYang]].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.HuaLu.String(),
							StarType: startype.SiHua,
						}))
					})
				})
				When("birth year is in 辛", func() {
					BeforeEach(func() {
						// luna date: 辛丑年 1 / 22 亥時
						birthday = time.Date(1961, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct hua lu", func() {
						Expect(board.Blocks[board.StarsMap[stars.JuMen]].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.HuaLu.String(),
							StarType: startype.SiHua,
						}))
					})
				})
				When("birth year is in 壬", func() {
					BeforeEach(func() {
						// luna date: 壬寅年 2 / 3 亥時
						birthday = time.Date(1962, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct hua lu", func() {
						Expect(board.Blocks[board.StarsMap[stars.TianLiang]].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.HuaLu.String(),
							StarType: startype.SiHua,
						}))
					})
				})
				When("birth year is in 癸", func() {
					BeforeEach(func() {
						// luna date: 癸卯年 2 / 13 亥時
						birthday = time.Date(1963, 3, 8, 22, 12, 0, 0, time.Local)
					})
					It("should display correct hua lu", func() {
						Expect(board.Blocks[board.StarsMap[stars.PoJun]].Stars).Should(ContainElement(&ziwei.Star{
							Name:     stars.HuaLu.String(),
							StarType: startype.SiHua,
						}))
					})
				})
			})
		})
	})
})
