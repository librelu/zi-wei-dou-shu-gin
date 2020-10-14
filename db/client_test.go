package db_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zi-wei-dou-shu-gin/db"
)

var _ = PDescribe("db", func() {
	Describe("NewDBClient", func() {
		var (
			host     string
			user     string
			password string
			database string
			sslmode  string
			port     int32
			dbClient *db.Client
			err      error
		)
		BeforeEach(func() {
			host = "0.0.0.0"
			user = "postgres"
			database = "postgres"
			port = 5432
			sslmode = "disable"
		})
		JustBeforeEach(func() {
			dbClient, err = db.NewDBClient(host, user, password, database, sslmode, port)
		})
		When("given valid params", func() {
			It("should init the db object", func() {
				Expect(err).Should(BeNil())
				Expect(dbClient).ShouldNot(BeNil())
			})
		})
		When("given a blank host", func() {
			BeforeEach(func() {
				host = ""
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(dbClient).Should(BeNil())
			})
		})
		When("given a blank user", func() {
			BeforeEach(func() {
				user = ""
			})
			It("should init the db object", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(dbClient).Should(BeNil())
			})
		})
		When("given a blank database", func() {
			BeforeEach(func() {
				database = ""
			})
			It("should init the db object", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(dbClient).Should(BeNil())
			})
		})
		When("given a zero port", func() {
			BeforeEach(func() {
				port = 0
			})
			It("should init the db object", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(dbClient).Should(BeNil())
			})
		})
		When("given a negative port", func() {
			BeforeEach(func() {
				port = -100
			})
			It("should init the db object", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(dbClient).Should(BeNil())
			})
		})
		When("given a blank sslmode", func() {
			BeforeEach(func() {
				sslmode = ""
			})
			It("should init the db object", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(dbClient).Should(BeNil())
			})
		})
	})
})
