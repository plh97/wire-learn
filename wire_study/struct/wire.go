//go:build wireinject
// +build wireinject

package _struct

import (
	"github.com/google/wire"
	"github.com/plh97/wire-learn/core"
)

func InitUserApi() *UserApi {
	wire.Build(
		NewUserService,
		NewVideoService,
		core.NewDB,
		core.NewLogger,
		wire.Struct(new(UserApi), "*"),
	)
	return nil
}

func InitUserApi1() *UserApi {
	wire.Build(
		core.NewDB,
		core.NewLogger,
		wire.Struct(new(UserApi), "*"),
		wire.Struct(new(VideoService), "*"),
		wire.Struct(new(UserService), "*"),
	)
	return nil
}
