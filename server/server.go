package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/helsonxiao/JudgeServer/utils"
)

// TODO: check X-Judge-Server-Token in the middleware
func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.POST("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, utils.H{
			Err:  nil,
			Data: utils.GetServerInfo(),
		})
	})

	r.POST("/judge", func(c *gin.Context) {
		c.JSON(http.StatusOK, utils.H{
			Err:  nil,
			Data: "pong",
		})
	})

	r.POST("/compile_spj", func(c *gin.Context) {
		c.JSON(http.StatusOK, utils.H{
			Err:  nil,
			Data: "pong",
		})
	})

	return r
}
