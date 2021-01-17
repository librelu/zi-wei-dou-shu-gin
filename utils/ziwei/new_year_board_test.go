package ziwei_test

import (
	"time"

	"bou.ke/monkey"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/genders"
)

var _ = Describe("ziwei", func() {
	var (
		birthday  time.Time
		yearBoard *ziwei.YearBoard
		err       error
		index     int
		gender    genders.Gender
	)
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	})
	JustBeforeEach(func() {
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
		BeforeEach(func() {
			birthday = time.Date(2019, 9, 13, 19, 4, 0, 0, time.Local)
			gender = genders.Male
		})

		When("given an index 1", func() {
			BeforeEach(func() {
				index = 1
			})
			It("should display correct blocks location", func() {
				Expect(yearBoard.Blocks[1].GongWeiName).To(Equal("命宮"))
			})
		})

		When("given an index 2", func() {
			BeforeEach(func() {
				index = 2
			})
			It("should display correct blocks location", func() {
				Expect(yearBoard.Blocks[2].GongWeiName).To(Equal("命宮"))
			})
		})

		When("given an index 3", func() {
			BeforeEach(func() {
				index = 3
			})
			It("should display correct blocks location", func() {
				Expect(yearBoard.Blocks[3].GongWeiName).To(Equal("命宮"))
			})
		})

		When("given an index 4", func() {
			BeforeEach(func() {
				index = 4
			})
			It("should display correct blocks location", func() {
				Expect(yearBoard.Blocks[4].GongWeiName).To(Equal("命宮"))
			})
		})
		When("given an index 5", func() {
			BeforeEach(func() {
				index = 5
			})
			It("should display correct blocks location", func() {
				Expect(yearBoard.Blocks[5].GongWeiName).To(Equal("命宮"))
			})
		})

		When("given an index 6", func() {
			BeforeEach(func() {
				index = 6
			})
			It("should display correct blocks location", func() {
				Expect(yearBoard.Blocks[6].GongWeiName).To(Equal("命宮"))
			})
		})

		When("given an index 7", func() {
			BeforeEach(func() {
				index = 7
			})
			It("should display correct blocks location", func() {
				Expect(yearBoard.Blocks[7].GongWeiName).To(Equal("命宮"))
			})
		})

		When("given an index 8", func() {
			BeforeEach(func() {
				index = 8
			})
			It("should display correct blocks location", func() {
				Expect(yearBoard.Blocks[8].GongWeiName).To(Equal("命宮"))
			})
		})

		When("given an index 9", func() {
			BeforeEach(func() {
				index = 9
			})
			It("should display correct blocks location", func() {
				Expect(yearBoard.Blocks[9].GongWeiName).To(Equal("命宮"))
			})
		})

		When("given an index 10", func() {
			BeforeEach(func() {
				index = 10
			})
			It("should display correct blocks location", func() {
				Expect(yearBoard.Blocks[10].GongWeiName).To(Equal("命宮"))
			})
		})

		When("given an index 11", func() {
			BeforeEach(func() {
				index = 11
			})
			It("should display correct blocks location", func() {
				Expect(yearBoard.Blocks[11].GongWeiName).To(Equal("命宮"))
			})
		})

		When("given an index 12", func() {
			BeforeEach(func() {
				index = 12
			})
			It("should display correct blocks location", func() {
				Expect(yearBoard.Blocks[0].GongWeiName).To(Equal("命宮"))
			})
		})
	})
})
