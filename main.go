package main

import (
	"github.com/plh97/wire-learn/wire"
)

func main() {
	// global.DB = core.NewDB()
	// global.Logger = core.NewLogger()
	// userService := service.NewUserService(
	// 	global.DB,
	// 	global.Logger,
	// )
	// api := api.NewUserApi(userService)
	// r := router.NewUserRouter(*api)
	r := wire.InitWire()
	r.Run(":80")
}
