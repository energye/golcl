package main

import (
	"fmt"

	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/messages"
	_ "github.com/energye/golcl/pkgs/winappres"
)

type TForm1 struct {
	*lcl.TForm
}

var form1 *TForm1

func main() {
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.CreateForm(&form1)
	lcl.Application.Run()
}

func (f *TForm1) OnFormCreate(sender lcl.IObject) {
	form1.SetCaption("Message Test")
	form1.SetPosition(types.PoScreenCenter)
	form1.EnabledMaximize(false)
	form1.SetWidth(500)
	form1.SetHeight(400)
	form1.SetOnWndProc(f.OnFormWndProc)
}

func (f *TForm1) OnFormWndProc(msg *types.TMessage) {
	// 这句一定要
	f.InheritedWndProc(msg)
	switch msg.Msg {
	case messages.WM_MOUSEMOVE:

	case messages.WM_LBUTTONDOWN:
		fmt.Println("左键接下")

	case messages.WM_LBUTTONUP:
		fmt.Println("左键抬起")

	case messages.WM_LBUTTONDBLCLK:
		fmt.Println("左键双击")

	case messages.WM_RBUTTONDOWN:
		fmt.Println("右键接下")

	case messages.WM_RBUTTONUP:
		fmt.Println("右键抬起")

	case messages.WM_RBUTTONDBLCLK:
		fmt.Println("右键双击")

	case messages.WM_MOUSEWHEEL:
		fmt.Println("鼠标滚轮")

	}
}
