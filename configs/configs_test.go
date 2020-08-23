package configs_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
	"github.com/zi-wei-dou-shu-gin/configs"
)

var _ = Describe("configs", func() {
	Describe("NewConfigs", func() {
		var (
			err      error
			path     string
			fileType string
			config   *configs.Config
		)
		BeforeEach(func() {
			os.Setenv("GIN_MODE", "test")
			path = "./tests"
			fileType = "yaml"
		})
		JustBeforeEach(func() {
			config, err = configs.NewConfigs(path, fileType)
		})
		AfterEach(func() {
			viper.Reset()
		})
		When("successful read the config", func() {
			It("should got no errors", func() {
				Expect(err).Should(BeNil())
				Expect(config.DB.Host).Should(Equal("db"))
				Expect(config.DB.Port).Should(Equal(int32(5432)))
				Expect(config.DB.User).Should(Equal("postgres"))
				Expect(config.DB.Password).Should(Equal(""))
				Expect(config.DB.Database).Should(Equal("postgres"))
				Expect(config.DB.SSLMode).Should(Equal("disable"))

			})
		})
		When("config not found", func() {
			BeforeEach(func() {
				path = "./not_found_path"
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(config).Should(BeNil())
			})
		})
		When("given invalid env", func() {
			BeforeEach(func() {
				os.Setenv("GIN_MODE", "unknown")
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(config).Should(BeNil())
			})
		})
	})
	Describe("GetGINMode", func() {
		var (
			env      string
			err      error
			envInput string
		)
		BeforeEach(func() {
			viper.SetEnvPrefix("GIN")
			viper.BindEnv("MODE")
		})
		JustBeforeEach(func() {
			os.Setenv("GIN_MODE", envInput)
			env, err = configs.GetGINMode()
		})
		AfterEach(func() {
			viper.Reset()
		})
		When("give test env", func() {
			BeforeEach(func() {
				envInput = "test"
			})
			It("should returns test", func() {
				Expect(env).Should(Equal("test"))
				Expect(err).Should(BeNil())
			})
		})
		When("give debug env", func() {
			BeforeEach(func() {
				envInput = "debug"
			})
			It("should returns test", func() {
				Expect(env).Should(Equal("debug"))
				Expect(err).Should(BeNil())
			})
		})
		When("give release env", func() {
			BeforeEach(func() {
				envInput = "release"
			})
			It("should returns test", func() {
				Expect(env).Should(Equal("release"))
				Expect(err).Should(BeNil())
			})
		})
		When("give unknown env", func() {
			BeforeEach(func() {
				envInput = "unknown"
			})
			It("should returns test", func() {
				Expect(env).Should(Equal(""))
				Expect(err).ShouldNot(BeNil())
			})
		})
	})
	Describe("ResetConfig", func() {
		BeforeEach(func() {
			viper.SetEnvPrefix("BEARNERS")
			viper.BindEnv("ENV")
		})
		JustBeforeEach(func() {
			configs.NewConfigs(".", "yaml")
			configs.ResetConfig()
		})
		When("give unknown env", func() {
			It("should returns test", func() {
				Expect(viper.Get("db")).Should(BeNil())
			})
		})
	})
})
