package main

import (
	"github.com/liom-source/wire-learn/wire"
)

func main() {
	r := wire.InitWire()
	r.Run(":80")
}
