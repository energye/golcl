package main

import (
	"github.com/energye/golcl/energy/inits"
	"github.com/energye/golcl/lcl"
	_ "github.com/energye/golcl/pkgs/winappres"
	"github.com/energye/golcl/samples/miniwebview/src"
)

func main() {
	inits.Init(nil, nil)
	lcl.Application.Initialize()
	lcl.Application.CreateForm(&src.MainForm)
	lcl.Application.Run()
}
