package routers

import (
	"gvb_server/global"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.Engine
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	RouterGroupApp := RouterGroup{router}
	RouterGroupApp.SettingsRouter()
	return router

}
