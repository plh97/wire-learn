//go:build !wireinject

package _value

import "github.com/google/wire"

func InitWire() *Info {
	wire.Build(
		wire.Value(&Info{
			Name: "wire",
		}),
	)
	return nil
}
