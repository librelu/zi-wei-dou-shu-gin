package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zi-wei-dou-shu-gin/configs"
	"github.com/zi-wei-dou-shu-gin/db"
	"github.com/zi-wei-dou-shu-gin/db/dao"
	"github.com/zi-wei-dou-shu-gin/services/boards"
)

type clients struct {
}

func main() {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	config, err := configs.NewConfigs(
		configs.DefaultConfigYamlPath,
		configs.DefaultConfigType,
	)
	if err != nil {
		panic(err)
	}
	dbClient, err := db.NewDBClient(
		config.DB.Host,
		config.DB.User,
		config.DB.Password,
		config.DB.Database,
		config.DB.SSLMode,
		config.DB.Port,
	)
	if err != nil {
		panic(err)
	}
	dao := dao.NewDao(dbClient)
	// init clients
	// init endpoints
	initEndpoints(engine, dao)
	// startup gin server
	// close connection if server down
	engine.Run()
}

func initEndpoints(engine *gin.Engine, dao dao.DaoHandler) {
	engine.GET("/healthcheck", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	groupAPI := engine.Group("/api")
	v1 := groupAPI.Group("/v1")
	v1.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	boards.BoardRegister(v1, boards.NewBoardHandler(dao))
}
