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
		var judgeDto JudgeDto
		bindErr := c.ShouldBind(&judgeDto)
		if bindErr == nil {
			resDto, judgeErr := Judge(judgeDto)
			if judgeErr == nil {
				c.JSON(http.StatusOK, utils.H[JudgeResponseDto]{
					Data: resDto,
				})
			} else {
				c.JSON(http.StatusOK, utils.H[string]{
					Err: judgeErr.Error(),
				})
			}
		} else {
			c.JSON(http.StatusOK, utils.H[string]{
				Err: bindErr.Error(),
			})
		}
	})

	r.POST("/compile_spj", func(c *gin.Context) {
		var json SpjCompileDto
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
