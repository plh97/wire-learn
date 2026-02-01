package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/liom-source/wire-learn/api/user_api"
	"github.com/liom-source/wire-learn/core"
	"github.com/liom-source/wire-learn/router"
	"github.com/liom-source/wire-learn/service/user_service"
)

func InitWire() *gin.Engine {
	DB := core.NewDB()
	LOG := core.NewLogger()
	userService := user_service.NewUserService(DB, LOG)
	userApi := user_api.NewUserApi(userService)
	r := router.NewRouter(userApi)
	return r
}
