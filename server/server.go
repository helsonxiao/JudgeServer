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
		c.JSON(http.StatusOK, utils.H[utils.ServerInfo]{
			Data: utils.GetServerInfo(),
		})
	})

	r.POST("/judge", func(c *gin.Context) {
		c.JSON(http.StatusOK, utils.H[string]{
			Data: "pong",
		})
	})

	r.POST("/compile_spj", func(c *gin.Context) {
		var json utils.SpjCompileJson
		err := c.ShouldBind(&json)
		if err == nil {
			c.JSON(http.StatusOK, utils.H[string]{
				Data: "success",
			})
		} else {
			c.JSON(http.StatusOK, utils.H[any]{
				Err: err.Error(),
			})
		}
	})

	return r
}
