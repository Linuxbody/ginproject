package middleware

import (
	"os"
	"io"
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	file *os.File
	e error
)

func CustomLogger() gin.HandlerFunc {
	gin.DisableConsoleColor()

	// 需要提前在项目下创建好 logs 目录
	file, e = os.OpenFile("logs/ginWeb.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0664)
	if e != nil {
		panic(e)
	}

	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	g := gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - %s \"%s %s %s %d %s \"%s\" %s\"\n",
			param.TimeStamp.Format("2006-01-02 03:04:05"),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
	return g
}

func CloseLogFile() {
	if err := file.Close(); err != nil {
		return
	}
}
