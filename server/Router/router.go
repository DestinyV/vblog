package Router

import (
	"server/Controllers"
	"server/Middlewares"
	"server/Sessions"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(Middlewares.Cors())
	// 使用 session(cookie-based)
	router.Use(sessions.Sessions("mysession", Sessions.Store))
	v1 := router.Group("v1")
	{
		v1.GET("/testGet", Controllers.TestGet)
	}

	router.Run(":8080")
}
