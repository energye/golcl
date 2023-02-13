//go:build windows
// +build windows

package main

import (
	"github.com/energye/golcl/energy/inits"
	"github.com/energye/golcl/lcl"
	_ "github.com/energye/golcl/pkgs/winappres"
)

func main() {
	inits.Init(nil, nil)

	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.CreateForm(&Form1)
	lcl.Application.Run()
}
