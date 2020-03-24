package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/panwenbin/ghttplog/actions"
)

var r *gin.Engine

func Load() *gin.Engine {
	r.POST("/log", actions.LogPost)

	return r
}

func init() {
	r = gin.Default()
}
