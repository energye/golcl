package main

import (
	"github.com/energye/golcl/energy/inits"
	"github.com/energye/golcl/lcl"
)

func main() {
	inits.Init(nil, nil)
	lcl.RunApp(&Form1)
}
