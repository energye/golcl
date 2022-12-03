package main

import (
	"github.com/energye/golcl/lcl"
	_ "github.com/energye/golcl/pkgs/winappres"
)

func main() {
	lcl.Application.Initialize()
	lcl.Application.CreateForm(&MainForm)
	lcl.Application.Run()
}
