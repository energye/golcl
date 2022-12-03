// 由res2go自动生成。
package main

import (
	"github.com/energye/golcl/lcl"
	_ "github.com/energye/golcl/pkgs/winappres"
)

func main() {
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.CreateForm(&Form11)
	lcl.Application.Run()
}
