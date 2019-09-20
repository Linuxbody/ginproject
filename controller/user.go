package controller

import(
	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	c.String(200, "Hello Controller")
	
}