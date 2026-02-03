package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/plh97/wire-learn/service"
)

type VideoApi struct {
	*service.UserService
	*service.VideoService
}

type GetVideoRequest struct {
	ID int `form:"id" binding:"required"`
}

var NewVideoApiProvider = wire.NewSet(NewVideoApi)

func NewVideoApi(userService *service.UserService, videoService *service.VideoService) *VideoApi {
	return &VideoApi{
		UserService:  userService,
		VideoService: videoService,
	}
}

func (u *VideoApi) GetVideo(c *gin.Context) {
	var req GetVideoRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	video, err := u.VideoService.GetVideo(req.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "Video fetched successfully",
		"video":   video,
	})
}
