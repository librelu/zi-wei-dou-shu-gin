package configs_test

import (
	"bytes"

	"github.com/bearners-gin/configs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
)

var _ = Describe("db", func() {
	Describe("NewDBConfig", func() {
		var (
			yamlExample []byte
			dbConfig    *configs.DBConfig
			err         error
		)
		BeforeEach(func() {
			yamlExample = []byte(`
db:
  host: "10.100.100.10"
  port: 5432
  user: "postgresUser"
  password: "testpassword"
  database: "postgres"
  sslmode: "disable"
`)
		})

		JustBeforeEach(func() {
			viper.SetConfigType("yaml")
			viper.ReadConfig(bytes.NewBuffer(yamlExample))
			dbConfig, err = configs.NewDBConfig()
		})
		When("given valid input", func() {
			It("should returns correct db configs", func() {
				Expect(err).Should(BeNil())
				Expect(dbConfig.Host).Should(Equal("10.100.100.10"))
				Expect(dbConfig.Port).Should(Equal(int32(5432)))
				Expect(dbConfig.User).Should(Equal("postgresUser"))
				Expect(dbConfig.Password).Should(Equal("testpassword"))
				Expect(dbConfig.Database).Should(Equal("postgres"))
				Expect(dbConfig.SSLMode).Should(Equal("disable"))
			})
		})
		When("host is blank", func() {
			BeforeEach(func() {
				yamlExample = []byte(`
db:
  host: ""
  port: 5432
  user: "postgresUser"
  password: "testpassword"
  database: "postgres"
  sslmode: "disable"
`)
			})
			It("should returns error", func() {
				Expect(dbConfig).Should(BeNil())
				Expect(err).ShouldNot(BeNil())
			})
		})
		When("port is blank", func() {
			BeforeEach(func() {
				yamlExample = []byte(`
db:
  host: "10.100.100.10"
  port: 0
  user: "postgresUser"
  password: "testpassword"
  database: "postgres"
  sslmode: "disable"
`)
			})
			It("should returns error", func() {
				Expect(dbConfig).Should(BeNil())
				Expect(err).ShouldNot(BeNil())
			})
		})
		When("port is negative", func() {
			BeforeEach(func() {
				yamlExample = []byte(`
db:
  host: "10.100.100.10"
  port: -100
  user: "postgresUser"
  password: "testpassword"
  database: "postgres"
  sslmode: "disable"
`)
			})
			It("should returns error", func() {
				Expect(dbConfig).Should(BeNil())
				Expect(err).ShouldNot(BeNil())
			})
		})
		When("user is blank", func() {
			BeforeEach(func() {
				yamlExample = []byte(`
db:
  host: "10.100.100.10"
  port: 5432
  user: ""
  password: "testpassword"
  database: "postgres"
  sslmode: "disable"
`)
			})
			It("should returns error", func() {
				Expect(dbConfig).Should(BeNil())
				Expect(err).ShouldNot(BeNil())
			})
		})
		When("database is blank", func() {
			BeforeEach(func() {
				yamlExample = []byte(`
db:
  host: "10.100.100.10"
  port: 5432
  user: "postgresUser"
  password: "testpassword"
  database: ""
  sslmode: "disable"
`)
			})
			It("should returns error", func() {
				Expect(dbConfig).Should(BeNil())
				Expect(err).ShouldNot(BeNil())
			})
		})
		When("sslmode is blank", func() {
			BeforeEach(func() {
				yamlExample = []byte(`
db:
  host: "10.100.100.10"
  port: 5432
  user: "postgresUser"
  password: "testpassword"
  database: "postgres"
  sslmode: ""
`)
			})
			It("should returns error", func() {
				Expect(dbConfig).Should(BeNil())
				Expect(err).ShouldNot(BeNil())
			})
		})
		When("config not found", func() {
			BeforeEach(func() {
				yamlExample = []byte(``)
			})
			It("should returns error", func() {
				Expect(dbConfig).Should(BeNil())
				Expect(err).ShouldNot(BeNil())
			})
		})
	})
})
