package server

import (
	"ginrtsp/api"

	"github.com/gin-gonic/gin"
)

// NewRouter Gin 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 路由
	r.GET("/ping", api.Ping)
	route := r.Group("/stream")
	{
		route.POST("/play", api.PlayRTSP)
		route.POST("/upload/:channel", api.Mpeg1Video)
		route.GET("/live/:channel", api.Wsplay)
	}

	return r
}
