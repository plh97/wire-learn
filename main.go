package main

import (
	"github.com/plh97/wire-learn/wire"
)

func main() {
	r := wire.InitWire()
	r.Run(":80")
}
