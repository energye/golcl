package main

import (
	"github.com/energye/golcl/inits"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/i18n"
	_ "github.com/energye/golcl/pkgs/winappres"
)

func main() {
	inits.Init(nil, nil)
	i18n.InitDefaultLang()
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.CreateForm(&Form1)
	lcl.Application.Run()

}
