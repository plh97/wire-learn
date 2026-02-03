package fields_of

type Info struct {
	Name string
}

func NewInfo() *Info {
	return &Info{
		Name: "fields_of",
	}
}