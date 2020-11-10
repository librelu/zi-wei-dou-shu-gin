package boards_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBoards(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Boards Suite")
}

func buildTestingServer() *gin.Engine {
	gin.SetMode(gin.TestMode)
	engine := gin.New()
	return engine
}

func recordBodyResponse(engine *gin.Engine, path, method string) (*httptest.ResponseRecorder, error) {
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		return nil, err
	}
	engine.ServeHTTP(recorder, req)
	return recorder, nil
}
