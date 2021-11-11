package bootstrap

import (
	"blog/routes"

	"github.com/gin-gonic/gin"
)

// SetupRoute 路由初始化
func SetupRoute(engine *gin.Engine) {
	//加载前台路由
	routes.RegisterApiRoute(engine)
}
