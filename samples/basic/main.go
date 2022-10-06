package main

import (
	"fmt"
	"github.com/energye/golcl/inits"
	"github.com/energye/golcl/lcl"

	"github.com/energye/golcl/lcl/types"
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

	x := lcl.NewFloatSpinEdit(f)
	x.SetParent(f)
	x.SetLeft(100)
	x.SetTop(100)
	x.SetMaxValue(10.0)
	x.SetIncrement(0.2)
	x.SetMinValue(0.1)
	x.SetValue(3.0)
	x.SetWidth(100)

	d := lcl.NewDirectoryEdit(f)
	d.SetParent(f)
	d.SetLeft(100)
	d.SetTop(150)
	d.SetWidth(200)
	d.SetDirectory("/home/sxm/app/swt")

	c := lcl.NewColorButton(f)
	c.SetParent(f)
	c.SetOnColorChanged(func(sender lcl.IObject) {
		lcl.QueueAsyncCall(func(id int) {
			fmt.Println(c.ButtonColor())
		})
	})

}

func (f *TMainForm) OnFormCloseQuery(Sender lcl.IObject, CanClose *bool) {
	*CanClose = lcl.MessageDlg("是否退出？", types.MtConfirmation, types.MbYes, types.MbNo) == types.IdYes
}

func (f *TMainForm) OnButton1Click(object lcl.IObject) {
	lcl.QueueAsyncCall(func(id int) {
		fmt.Println("OnButton1Click")
		form1.Show()
	})
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
