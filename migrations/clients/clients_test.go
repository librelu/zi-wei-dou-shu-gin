package migrations_test

import (
	migrations "github.com/bearners-gin/migrations/clients"
	"github.com/golang-migrate/migrate/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("clients", func() {
	Describe("NewDBMigration", func() {
		var (
			err               error
			migrateResult     *migrate.Migrate
			host              string
			user              string
			password          string
			database          string
			sslmode           string
			migrationFilePath string
			driverName        string
			port              int32
		)
		BeforeEach(func() {
			host = "0.0.0.0"
			user = "postgres"
			password = ""
			database = "postgres"
			sslmode = "disable"
			migrationFilePath = "."
			driverName = migrations.PostgresDriverName
			port = 5432
		})
		JustBeforeEach(func() {
			migrateResult, err = migrations.NewDBMigration(
				host, user, password, database, sslmode, migrationFilePath, driverName, port,
			)
		})
		When("given valid input", func() {
			It("should returns migrate object without error", func() {
				Expect(err).Should(BeNil())
				Expect(migrateResult).ShouldNot(BeNil())
			})
		})
		When("given an inncorrect driver name", func() {
			BeforeEach(func() {
				driverName = "incorrect driver"
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(migrateResult).Should(BeNil())
			})
		})
		When("given an inncorrect file path", func() {
			BeforeEach(func() {
				migrationFilePath = "incorrect path"
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(migrateResult).Should(BeNil())
			})
		})
		When("given an inncorrect config", func() {
			BeforeEach(func() {
				host = "incorrect host"
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(migrateResult).Should(BeNil())
			})
		})
		When("given a blank host", func() {
			BeforeEach(func() {
				host = ""
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(migrateResult).Should(BeNil())
			})
		})
		When("given a blank user", func() {
			BeforeEach(func() {
				user = ""
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(migrateResult).Should(BeNil())
			})
		})
		When("given a blank database", func() {
			BeforeEach(func() {
				database = ""
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(migrateResult).Should(BeNil())
			})
		})
		When("given a blank sslmode", func() {
			BeforeEach(func() {
				sslmode = ""
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(migrateResult).Should(BeNil())
			})
		})
		When("given a blank migrationFilePath", func() {
			BeforeEach(func() {
				migrationFilePath = ""
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(migrateResult).Should(BeNil())
			})
		})
		When("given a blank driverName", func() {
			BeforeEach(func() {
				driverName = ""
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(migrateResult).Should(BeNil())
			})
		})
		When("given a zero port", func() {
			BeforeEach(func() {
				port = 0
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(migrateResult).Should(BeNil())
			})
		})
		When("given a negative port", func() {
			BeforeEach(func() {
				port = -100
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(migrateResult).Should(BeNil())
			})
		})
	})
})
