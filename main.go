package main

import (
	"github.com/gin-gonic/gin"
	"ginproject/routers"
	"ginproject/middleware"
	"ginproject/utils"
)

func main() {
	
	gin.SetMode(utils.ServerInfo.RunMode)
	
	routers.InitRouter()
	
	defer middleware.CloseLogFile()

	
}