package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/plh97/wire-learn/api"
	"github.com/plh97/wire-learn/core"
	"github.com/plh97/wire-learn/router"
	"github.com/plh97/wire-learn/service"
)

func InitWire() *gin.Engine {
	DB := core.NewDB()
	LOG := core.NewLogger()
	userService := service.NewUserService(DB, LOG)
	userApi := api.NewUserApi(userService)
	r := router.NewUserRouter(userApi)
	return r
}
