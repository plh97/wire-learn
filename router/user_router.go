package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/plh97/wire-learn/api"
)

var RouterProvider = wire.NewSet(NewUserRouter, api.UserApiProvider, api.VideoApiProvider)

func NewUserRouter(userApi *api.UserApi, videoApi *api.VideoApi) *gin.Engine {
	r := gin.Default()
	g := r.Group("/api")
	{
		g.GET("/user", userApi.GetUser)
		g.GET("/video", videoApi.GetVideo)
	}
	return r
}
