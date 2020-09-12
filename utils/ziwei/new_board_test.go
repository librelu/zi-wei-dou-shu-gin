package ziwei_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/dizhi"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/gong"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/mingju"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/tiangan"
)

var _ = Describe("ziwei", func() {
	Describe("NewBoard", func() {
		var (
			board    *ziwei.Board
			birthday time.Time
		)
		JustBeforeEach(func() {
			board = ziwei.NewBoard(birthday)
		})
		FWhen("init board", func() {
			BeforeEach(func() {
				birthday = time.Date(1984, 8, 8, 0, 4, 0, 0, time.Local)
			})
			It("should returns location", func() {
				Expect(len(board.Blocks)).Should(Equal(12))
				for i, block := range board.Blocks {
					Expect(block.Location.DiZhi.String()).Should(Equal(dizhi.DiZhi(i).String()))
				}
			})
		})
		Context("testing Yin Shou", func() {
			When("given a birthday with Jia year", func() {
				BeforeEach(func() {
					birthday = time.Date(1984, 8, 8, 0, 4, 0, 0, time.Local)
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
					birthday = time.Date(1988, 8, 8, 0, 4, 0, 0, time.Local)
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
					birthday = time.Date(1986, 8, 8, 0, 4, 0, 0, time.Local)
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
					birthday = time.Date(1984, 8, 8, 0, 4, 0, 0, time.Local)
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
					birthday = time.Date(1984, 8, 8, 2, 4, 0, 0, time.Local)
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
					birthday = time.Date(1984, 12, 8, 0, 4, 0, 0, time.Local)
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
	})
})
