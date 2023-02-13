package main

import (
	"embed"
	"github.com/energye/golcl/energy/inits"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/samples/action/src"
)
import _ "github.com/energye/golcl/pkgs/winappres"

//go:embed resources
var resources embed.FS

func main() {
	inits.Init(nil, nil)
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.CreateForm(&src.MainForm, true)
	lcl.Application.Run()
}
