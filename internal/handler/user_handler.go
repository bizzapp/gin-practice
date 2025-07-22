package handler

import (
	"net/http"
	"strconv"

	"myapp/internal/config"
	"myapp/internal/model"
	"myapp/internal/service"
	"myapp/pkg/response"

	"github.com/gin-gonic/gin"
)

func ListUsers(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, _ := strconv.Atoi(limitStr)
	users, err := service.GetAllUsers(config.DB, limit)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "USER_QUERY_FAILED", "Failed to fetch users", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, "Users fetched successfully", users)
}

func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_PAYLOAD", "Invalid input", err.Error())
		return
	}
	if err := service.RegisterUser(config.DB, &user); err != nil {
		response.Error(c, http.StatusInternalServerError, "USER_CREATION_FAILED", "Failed to create user", err.Error())
		return
	}
	response.JSON(c, http.StatusCreated, "User created successfully", user)
}
