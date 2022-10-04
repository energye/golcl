//go:build windows
// +build windows

package main

import (
	"fmt"
	"github.com/energye/golcl/inits"
	"github.com/go-ole/go-ole"

	"github.com/energye/golcl/lcl"
	_ "github.com/energye/golcl/pkgs/winappres"
	"github.com/go-ole/go-ole/oleutil"
)

type TMainForm struct {
	*lcl.TForm
	Btn1 *lcl.TButton
}

var mainForm *TMainForm

func (f *TMainForm) OnFormCreate(object lcl.IObject) {
	f.SetCaption("WinOLE")
	f.ScreenCenter()

	f.Btn1 = lcl.NewButton(f)
	f.Btn1.SetParent(f)
	f.Btn1.SetBounds(10, 10, 100, 30)
	f.Btn1.SetCaption("打开记事本")
	f.Btn1.SetOnClick(f.OnBtn1Click)
}

func (f *TMainForm) OnBtn1Click(object lcl.IObject) {
	ole.CoInitializeEx(0, 0)
	defer ole.CoUninitialize()

	unknown, err := oleutil.CreateObject("WScript.Shell")
	if err != nil {
		fmt.Println(err)
		return
	}
	shell, _ := unknown.QueryInterface(ole.IID_IDispatch)
	shell.CallMethod("Run", "notepad")
}

func main() {
	inits.Init(nil, nil)
	lcl.RunApp(&mainForm)
}
