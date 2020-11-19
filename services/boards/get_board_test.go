package boards_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zi-wei-dou-shu-gin/db/dao"
	"github.com/zi-wei-dou-shu-gin/services/boards"
)

var _ = Describe("BoardHandler()", func() {
	var (
		engine *gin.Engine
	)
	BeforeEach(func() {
		engine = buildTestingServer()
	})
	Context("GetBoard()", func() {
		var (
			path     string
			recorder *httptest.ResponseRecorder
			method   string
			dbClient dao.DaoHandler
			handler  boards.Handler
		)
		BeforeEach(func() {
			method = http.MethodGet
			handler = boards.NewBoardHandler(dbClient)
			boards.BoardRegister(engine.Group(""), handler)
		})
		JustBeforeEach(func() {
			var err error
			recorder, err = recordBodyResponse(engine, path, method)
			if err != nil {
				panic(err)
			}
		})
		When("in success", func() {
			BeforeEach(func() {
				path = "/board?birthday=1602658277&gender=0"
			})
			It("should returns correct response", func() {
				Expect(recorder.Code).Should(Equal(http.StatusOK))
			})
		})
		When("in invalidate request", func() {
			When("given wrong gender number", func() {
				BeforeEach(func() {
					path = "/board?birthday=1602658277&gender=3"
				})
				It("should returns correct response", func() {
					Expect(recorder.Code).Should(Equal(http.StatusBadRequest))
				})
			})
			When("given none query string", func() {
				BeforeEach(func() {
					path = "/board"
				})
				It("should returns correct response", func() {
					Expect(recorder.Code).Should(Equal(http.StatusBadRequest))
				})
			})
		})
	})
})
