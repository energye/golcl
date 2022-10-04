package main

import (
	"fmt"

	"github.com/energye/golcl/lcl/win"

	"github.com/energye/golcl/lcl/types/keys"

	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

//::private::
type TForm1Fields struct {
}

func (f *TForm1) OnFormCreate(sender lcl.IObject) {
	win.OutputDebugString(Form1.Caption(), f.PixelsPerInch())
}

func (f *TForm1) OnFormDestroy(sender lcl.IObject) {
	win.OutputDebugString("----------------------OnFormDestroy")
}

func (f *TForm1) memoOnKeyup(sender lcl.IObject, key *types.Char, shift types.TShiftState) {
	if shift.In(types.SsCtrl) && *key == /*keys.VkA*/ keys.VkB {
		lcl.AsMemo(sender).SelectAll()
	}
}

func (f *TForm1) OnActExitExecute(lcl.IObject) {

}

func (f *TForm1) OnButton2Click(lcl.IObject) {
	win.OutputDebugString("currentid-1: ", win.GetCurrentThreadId())
	go func() {
		win.OutputDebugString("OnButton2Click,  currentid-2", win.GetCurrentThreadId())

		lcl.ThreadSync(func() {
			win.OutputDebugString("ThreadSync OnButton2Click-3", win.GetCurrentThreadId())
			lcl.ShowMessage("测试")
		})
	}()
}

func (f *TForm1) OnActFileNewExecute(lcl.IObject) {
	lcl.ShowMessage("ActFileNew Execute.")
}

func (f *TForm1) OnFormCloseQuery(sender lcl.IObject, canClose *bool) {
	fmt.Println("关闭。")
}

func (f *TForm1) OnCheckBox1Click(sender lcl.IObject) {
	f.Button1.SetEnabled(f.CheckBox1.Checked())
}

func (f *TForm1) OnButton1Click(sender lcl.IObject) {
	lcl.ShowMessage("Hello!")

}
