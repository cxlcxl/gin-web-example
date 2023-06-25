package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"silent-cxl.top/app/validators/valid"
	"silent-cxl.top/app/vars"
)

func Router() error {
	_router := gin.Default()

	if vars.YmlConfig.GetBool("HttpServer.AllowCrossDomain") {
		_router.Use(corsNext())
	}
	if !vars.YmlConfig.GetBool("Debug") {
		gin.SetMode(gin.ReleaseMode)
	}

	_router.GET("", valid.VUserQuery)

	return _router.Run(vars.YmlConfig.GetString("HttpServer.Port"))
}

// 允许跨域
func corsNext() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers,Authorization,User-Agent, Keep-Alive, Content-Type, X-Requested-With,X-CSRF-Token")
		c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		method := c.Request.Method
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusAccepted)
		}
		c.Next()
	}
}
