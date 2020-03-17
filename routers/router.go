package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/panwenbin/ghttplog/actions"
)

var r *gin.Engine

func Load() *gin.Engine {
	r.GET("/hello", actions.Hello)

	return r
}

func init() {
	r = gin.Default()
}
