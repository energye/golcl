//go:build windows
// +build windows

package main

import (
	"fmt"

	"github.com/energye/golcl/lcl"
	_ "github.com/energye/golcl/pkgs/winappres"
	"github.com/tryor/gdiplus"
	"github.com/tryor/winapi"
)

var (
	gpToken winapi.ULONG_PTR
)

// 用了一个第三方的go gdi+库， 看样子是个国人的作品

func main() {

	_, err := gdiplus.Startup(&gpToken, nil, nil)
	if err != nil {
		if err != nil {
			lcl.ShowMessage("初始GDI+失败，错误:" + err.Error())
		}
	} else {
		fmt.Println("gdi+实始成功")
		defer gdiplus.Shutdown(gpToken)
	}

	lcl.Application.SetOnException(func(sender lcl.IObject, e *lcl.Exception) {
		// 在这里自行处理VCL中的异常
	})

	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.CreateForm(&GdipForm)
	lcl.Application.Run()

}
