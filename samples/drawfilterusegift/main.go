package main

import (
	"github.com/energye/golcl/inits"
	"github.com/energye/golcl/lcl"
	_ "github.com/energye/golcl/pkgs/winappres"
	"github.com/energye/golcl/samples/drawfilterusegift/src"
)

func main() {
	inits.Init(nil, nil)
	lcl.RunApp(&src.Form1)
}
