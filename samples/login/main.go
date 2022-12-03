package main

import (
	"github.com/energye/golcl/lcl"
	_ "github.com/energye/golcl/pkgs/winappres"
)

func main() {

	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.CreateForm(&MainForm)
	lcl.Application.CreateForm(&LoginForm)

	// 这里可以决定是不是显示主窗口
	lcl.Application.SetShowMainForm(false)
	// 这里显示出第二个窗口
	LoginForm.Show()

	lcl.Application.Run()

}
