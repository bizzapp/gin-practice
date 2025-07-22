package router

import (
	"log"
	"myapp/internal/handler"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default() // Set trusted proxies secara aman
	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	v1 := r.Group("/api/v1")
	{
		v1.GET("/", handler.HelloWorld)
		v1.GET("/ping", handler.Ping)
		v1.GET("/users", handler.ListUsers)

	}

	// Example route

	return r
}
