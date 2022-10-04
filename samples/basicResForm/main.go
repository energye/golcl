package main

import (
	"github.com/energye/golcl/inits"
	"github.com/energye/golcl/lcl"
	_ "github.com/energye/golcl/pkgs/winappres"
)

func main() {
	//libname.LibName = "/Users/zhangli/go/bin/liblcl.dylib"
	//fmt.Println("os.Args[0]", os.Args[0])
	inits.Init(nil, nil)
	//lcl.Application.SetScaled(true)
	//lcl.Application.Initialize()
	//lcl.Application.SetMainFormOnTaskBar(true)
	//lcl.Application.SetOnException(func(sender lcl.IObject, e *lcl.Exception) {
	//	lcl.ShowMessage(e.Message())
	//})
	////   Form1.gfm
	//lcl.Application.CreateForm(&Form1)
	//// 字节加载方式
	//lcl.Application.CreateForm(&Form2)
	//
	//lcl.Application.Run()
	lcl.RunApp(&Form1, &Form2)
}
