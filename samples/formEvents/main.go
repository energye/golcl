package main

import (
	"github.com/energye/golcl/lcl"
	_ "github.com/energye/golcl/pkgs/macapp"
	_ "github.com/energye/golcl/pkgs/winappres"
)

func main() {

	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.CreateForm(&Form1)
	lcl.Application.CreateForm(&Form2)

	lcl.Application.Run()

}
