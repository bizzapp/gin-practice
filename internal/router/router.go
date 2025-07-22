package router

import (
	"github.com/gin-gonic/gin"
	"myapp/internal/handler"
)

func New() *gin.Engine {
	r := gin.Default()

	// Example route
	r.GET("/ping", handler.Ping)
	r.GET("/get_user", handler.GetUser)
	r.GET("/", handler.HelloWorld)
	r.GET("/hello", handler.HelloWorld)

	return r
}
