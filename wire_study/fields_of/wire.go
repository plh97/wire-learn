package fields_of

import "github.com/google/wire"

func GetInfoName() string {
	wire.Build(
		NewInfo,
		wire.FieldsOf(new(Info), "Name"),
	)
	return "fields_of"
}