// 由res2go自动生成。
package main

import (
	"github.com/energye/golcl/lcl"
	_ "github.com/energye/golcl/pkgs/winappres"
)

const Lazarus = true

func main() {

	lcl.Application.Initialize()
	lcl.Application.CreateForm(&MainForm)
	lcl.Application.CreateForm(&About)
	lcl.Application.Run()
}
