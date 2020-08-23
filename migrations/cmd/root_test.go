package cmd

import (
	"bytes"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
)

var _ = Describe("db", func() {
	Describe("root", func() {
		var (
			output             *bytes.Buffer
			envArg             string
			configArg          string
			migrationPathArg   string
			migrationDriverArg string
			cobraExecutor      func() error
		)
		BeforeEach(func() {
			output = bytes.NewBufferString("")
			envArg = "test"
			configArg = "./test/"
			migrationPathArg = "test"
			migrationDriverArg = "postgres"
		})
		AfterEach(func() {
			viper.Reset()
			os.Unsetenv("GIN_MODE")
		})
		JustBeforeEach(func() {
			rootCmd.SetOutput(output)
			rootCmd.SetArgs([]string{"-e", envArg, "-c", configArg, "-m", migrationPathArg, "-d", migrationDriverArg})
			cobraExecutor = rootCmd.Execute
		})
		When("given valid args", func() {
			It("should pass without error", func() {
				Expect(cobraExecutor()).Should(BeNil())
			})
		})
		When("given wrong config path", func() {
			BeforeEach(func() {
				configArg = "./invalid path"
			})
			It("should panic", func() {
				Expect(func() { cobraExecutor() }).Should(Panic())
			})
		})
		When("given wrong env", func() {
			BeforeEach(func() {
				envArg = "invalid env"
			})
			It("should panic", func() {
				Expect(func() { cobraExecutor() }).Should(Panic())
			})
		})
		When("given wrong driver", func() {
			BeforeEach(func() {
				migrationDriverArg = "invalid driver"
			})
			It("should panic", func() {
				Expect(func() { cobraExecutor() }).Should(Panic())
			})
		})
	})
})
