package main

import (
	"github.com/energye/golcl/inits"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/win"

	"fmt"

	_ "github.com/energye/golcl/pkgs/winappres"
)

type TMainForm struct {
	*lcl.TForm
	Button1 *lcl.TButton
}

type TForm1 struct {
	*lcl.TForm
	Button1 *lcl.TButton
}

var (
	mainForm *TMainForm
	form1    *TForm1
)

func main() {
	inits.Init(nil, nil)
	lcl.RunApp(&mainForm, &form1)
}

// --------------MainForm -----------------
func (f *TMainForm) OnFormCreate(sender lcl.IObject) {

	f.SetCaption("Hello")
	f.EnabledMaximize(false)
	f.SetWidth(300)
	f.SetHeight(200)
	f.ScreenCenter()

	f.Button1 = lcl.NewButton(f)
	f.Button1.SetParent(f)
	f.Button1.SetCaption("窗口1")
	f.Button1.SetLeft(50)
	f.Button1.SetTop(50)
	f.Button1.SetOnClick(f.OnButton1Click)
}

func (f *TMainForm) OnButton1Click(object lcl.IObject) {
	form1.Show()
}

// 可以窗口创建前修改一些样式等操作，被动调用
func (f *TMainForm) CreateParams(params *types.TCreateParams) {
	fmt.Println("调用此过程  TMainForm.CreateParams")

}

// ---------- Form1 ----------------

func (f *TForm1) OnFormCreate(sender lcl.IObject) {

	fmt.Println("onCreate")
	f.Button1 = lcl.NewButton(f)
	fmt.Println("f.Button1:", f.Button1.Instance())
	f.Button1.SetParent(f)
	//f.Button1.SetName("Button1")
	f.Button1.SetCaption("我是按钮")
	f.Button1.SetOnClick(f.OnButton1Click)
}

func (f *TForm1) OnButton1Click(object lcl.IObject) {
	lcl.ShowMessage("Click")
}

// 可以窗口创建前修改一些样式等操作，被动调用
func (f *TForm1) CreateParams(params *types.TCreateParams) {
	fmt.Println("调用此过程 TForm1.CreateParams")
	params.WndParent = win.GetDesktopWindow()

}
