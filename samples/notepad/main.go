package main

import (
	"github.com/energye/golcl/lcl"
	_ "github.com/energye/golcl/pkgs/winappres"
)

func main() {

	lcl.Application.SetOnException(func(sender lcl.IObject, e *lcl.Exception) {
		// 在这里自行处理VCL中的异常
	})

	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.CreateForm(&MainForm)
	lcl.Application.Run()

}
