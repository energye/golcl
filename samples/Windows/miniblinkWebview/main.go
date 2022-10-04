package main

import (
	"github.com/energye/golcl/inits"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/pkgs/miniblink"
	_ "github.com/energye/golcl/pkgs/winappres"
)

func main() {
	inits.Init(nil, nil)
	miniblink.Init()
	defer miniblink.Finalize()

	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.CreateForm(&MainForm)
	lcl.Application.CreateForm(&HTMLForm)
	lcl.Application.CreateForm(&JsForm)
	lcl.Application.Run()

}
