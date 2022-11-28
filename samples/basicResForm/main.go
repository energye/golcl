package main

import (
	"github.com/energye/golcl/inits"
	"github.com/energye/golcl/lcl"
	_ "github.com/energye/golcl/pkgs/winappres"
)

func main() {
	inits.Init(nil, nil)
	lcl.RunApp(&Form1, &Form2)
}
