package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/plh97/wire-learn/service"
)

var VideoApiProvider = wire.NewSet(NewVideoApi, service.NewVideoService)

type GetVideoRequest struct {
	ID uint `form:"id" binding:"required"`
}

type VideoApi struct {
	*service.UserService
	*service.VideoService
}

func NewVideoApi(videoService *service.VideoService, userService *service.UserService) *VideoApi {
	return &VideoApi{
		UserService: userService,
		VideoService: videoService,
	}
}

func (api *VideoApi) GetVideo(c *gin.Context) {
	var req GetVideoRequest

	err := c.ShouldBindQuery(&req)

	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error binding query: " + err.Error(),
		})
		return
	}

	video, err := api.VideoService.GetVideo(req.ID)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error getting video: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Video Router",
		"data":    video,
	})
}
