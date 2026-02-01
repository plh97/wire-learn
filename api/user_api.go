package api

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/plh97/wire-learn/service"
	"gorm.io/gorm"
)

type GetUserRequest struct {
	ID uint `form:"id" binding:"required"`
}

type UserApi struct{
	*service.UserService
}

func NewUserApi(userService *service.UserService) *UserApi {
	return &UserApi{
		UserService: userService,
	}
}

func (api *UserApi) GetUser(c *gin.Context) {
	var req GetUserRequest

	err := c.ShouldBindQuery(&req)

	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error binding query: " + err.Error(),
		})
		return
	}

	user, err := api.UserService.GetUser(req.ID)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error getting user: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User Router",
		"data":    user,
	})
}
