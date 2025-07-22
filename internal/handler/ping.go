package handler
import (
	"github.com/gin-gonic/gin"
	"myapp/pkg/response"
)

func HelloWorld(c *gin.Context) {
	response.JSON(c, 200, "Hello World", nil)
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}


func GetUser(c *gin.Context) {
	id := c.Param("id")

	if id != "1" {
		response.Error(c, 404, "USER_NOT_FOUND", "User not found", "No user with ID "+id)
		return
	}

	user := map[string]interface{}{
		"id":   1,
		"name": "Jerome",
	}

	response.JSON(c, 200, "User fetched successfully", user)
}
