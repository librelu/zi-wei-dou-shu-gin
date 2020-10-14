package boards_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBoards(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Boards Suite")
}
