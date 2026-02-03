//go:build wireinject

package build

import (
	"github.com/google/wire"
	"github.com/plh97/wire-learn/core"
	"gorm.io/gorm"
)

func InitUserApi() (*UserApi, error) {
	wire.Build(
		NewUserApi,
		NewUserService,
		core.NewDB,
		NewRedis,
		core.NewLogger,
	)
	return nil, nil
}

func InitUserService() (*UserService, error) {
	wire.Build(
		NewUserService,
		core.NewDB,
		NewRedis,
		core.NewLogger,
	)
	return nil, nil
}

func InitInfo() (*Info, error) {
	wire.Build(
		NewInfo,
		core.NewDB,
	)
	return nil, nil
}


func LoadDB(db *gorm.DB) (*Info, error) {
	wire.Build(
		NewInfo,
	)
	return nil, nil
}