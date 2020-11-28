package ziwei_test

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/genders"
)

var _ = Describe("ziwei", func() {
	var (
		birthday  time.Time
		yearBoard *ziwei.YearBoard
		board     *ziwei.Board
		err       error
		index     int
		gender    genders.Gender
	)
	JustBeforeEach(func() {
		b, err := ziwei.NewBoard(birthday, gender)
		if err != nil {
			panic(err)
		}
		board, err = b.CreateTianBoard()
		yearBoard, err = ziwei.NewTenYearsBoard(birthday, gender, index)
	})
	Describe("NewTenYearsBoard()", func() {
		BeforeEach(func() {
			birthday = time.Date(1984, 8, 8, 0, 4, 0, 0, time.Local)
			gender = genders.Male
			index = 0
		})
		When("given correct birthday date", func() {
			It("should retruns correct data", func() {
				Expect(err).To(BeNil())
				Expect(yearBoard).NotTo(BeNil())
			})
		})
	})
	Describe("rotateGongWeiNameByIndex()", func() {
		actualBlocksShouldMatchOriginWithIndex := func(index int) bool {
			for i, actualBlock := range yearBoard.Blocks {
				idx := i + index
				if idx >= len(board.Blocks) {
					idx = idx - len(board.Blocks)
				}
				if !Expect(actualBlock.GongWeiName).To(Equal(board.Blocks[idx].GongWeiName),
					fmt.Sprintf("testing input: %d, actual block index:%d, current block index:%d", index, i, idx),
				) {
					return false
				}
			}
			return true
		}
		BeforeEach(func() {
			birthday = time.Date(2019, 9, 13, 19, 4, 0, 0, time.Local)
			gender = genders.Male
		})

		When("given an index 1", func() {
			BeforeEach(func() {
				index = 1
			})
			It("should display correct blocks location", func() {
				Expect(actualBlocksShouldMatchOriginWithIndex(index)).To(BeTrue())
			})
		})

		When("given an index 2", func() {
			BeforeEach(func() {
				index = 2
			})
			It("should display correct blocks location", func() {
				Expect(actualBlocksShouldMatchOriginWithIndex(index)).To(BeTrue())
			})
		})

		When("given an index 3", func() {
			BeforeEach(func() {
				index = 3
			})
			It("should display correct blocks location", func() {
				Expect(actualBlocksShouldMatchOriginWithIndex(index)).To(BeTrue())
			})
		})

		When("given an index 4", func() {
			BeforeEach(func() {
				index = 4
			})
			It("should display correct blocks location", func() {
				Expect(actualBlocksShouldMatchOriginWithIndex(index)).To(BeTrue())
			})
		})
		When("given an index 5", func() {
			BeforeEach(func() {
				index = 5
			})
			It("should display correct blocks location", func() {
				Expect(actualBlocksShouldMatchOriginWithIndex(index)).To(BeTrue())
			})
		})

		When("given an index 6", func() {
			BeforeEach(func() {
				index = 6
			})
			It("should display correct blocks location", func() {
				Expect(actualBlocksShouldMatchOriginWithIndex(index)).To(BeTrue())
			})
		})

		When("given an index 7", func() {
			BeforeEach(func() {
				index = 7
			})
			It("should display correct blocks location", func() {
				Expect(actualBlocksShouldMatchOriginWithIndex(index)).To(BeTrue())
			})
		})

		When("given an index 8", func() {
			BeforeEach(func() {
				index = 8
			})
			It("should display correct blocks location", func() {
				Expect(actualBlocksShouldMatchOriginWithIndex(index)).To(BeTrue())
			})
		})

		When("given an index 9", func() {
			BeforeEach(func() {
				index = 9
			})
			It("should display correct blocks location", func() {
				Expect(actualBlocksShouldMatchOriginWithIndex(index)).To(BeTrue())
			})
		})

		When("given an index 10", func() {
			BeforeEach(func() {
				index = 10
			})
			It("should display correct blocks location", func() {
				Expect(actualBlocksShouldMatchOriginWithIndex(index)).To(BeTrue())
			})
		})

		When("given an index 11", func() {
			BeforeEach(func() {
				index = 11
			})
			It("should display correct blocks location", func() {
				Expect(actualBlocksShouldMatchOriginWithIndex(index)).To(BeTrue())
			})
		})

		When("given an index 12", func() {
			BeforeEach(func() {
				index = 12
			})
			It("should display correct blocks location", func() {
				Expect(actualBlocksShouldMatchOriginWithIndex(index)).To(BeTrue())
			})
		})
	})
})
