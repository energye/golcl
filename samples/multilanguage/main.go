package main

import (
	"github.com/energye/golcl/lcl/i18n"
	_ "github.com/energye/golcl/pkgs/winappres"
)

func main() {

	i18n.InitDefaultLang()
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.CreateForm(&Form1)
	lcl.Application.Run()

}
