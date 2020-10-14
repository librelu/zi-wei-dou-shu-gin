package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zi-wei-dou-shu-gin/services/boards"
)

type clients struct {
}

func main() {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	// init clients
	// init endpoints
	initEndpoints(engine)
	// startup gin server
	// close connection if server down
	engine.Run()
}

func initClients() *clients {
	return nil
}

func initEndpoints(engine *gin.Engine) {
	groupAPI := engine.Group("/api")
	v1 := groupAPI.Group("/v1")
	boards.BoardRegister(v1, boards.NewBoardHandler())
}
