package routers

import (
	"github.com/gin-gonic/gin"
	"ginproject/controller"
	"ginproject/middleware"
	"ginproject/utils"
)

func InitRouter() {
	r := gin.New()
	r.Use(middleware.CustomLogger())
	r.Use(gin.Recovery())


	r.GET("/index", controller.GetIndex)
	
	r.Run(utils.ServerInfo.ServerAddr)

}