package main

import (
	"github.com/energye/golcl/lcl"
	_ "github.com/energye/golcl/pkgs/winappres"
)

// windows : stdcall;  其他平台： cdecl
//var runLoopReceivedProc = syscall.NewCallback(runLoop)

// data 为 Application 实例
/*func runLoop(data uintptr) uintptr {
	// 这里不能产生异常，否则程序会崩溃

	return 0
	//lcl.Application.HandleMessage()
	//return 1 // 返回1表示不由Lazarus处理后面的，必须要加上Application.HandleMessage()，一般情况下最好不自己处理
}*/

func main() {
	//lcl.Application.SetRunLoopReceived(runLoopReceivedProc)
	lcl.RunApp(&Form1, &Form2)

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

}
