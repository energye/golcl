//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package main

import (
	"fmt"
	"github.com/energye/golcl/energy/inits"
	"github.com/energye/golcl/lcl"
)
import _ "github.com/energye/golcl/pkgs/winappres"

type TMainForm struct {
	*lcl.TForm
	btn  *lcl.TButton
	btn2 *lcl.TButton
	btn3 *lcl.TButton
}

var mainForm *TMainForm

func main() {
	inits.Init(nil, nil)
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.CreateForm(&mainForm)
	lcl.Application.Run()
}

func (f *TMainForm) OnFormCreate(sender lcl.IObject) {
	f.ScreenCenter()
	f.btn = lcl.NewButton(f)
	f.btn.SetParent(f)
	f.btn.SetCaption("GetText")
	f.btn.SetOnClick(f.onBtnClick)

	f.btn2 = lcl.NewButton(f)
	f.btn2.SetParent(f)
	f.btn2.SetTop(40)
	f.btn2.SetWidth(120)
	f.btn2.SetCaption("GetTextBuffer")
	f.btn2.SetOnClick(f.onBtn2Click)

	f.btn3 = lcl.NewButton(f)
	f.btn3.SetParent(f)
	f.btn3.SetTop(80)
	f.btn3.SetWidth(120)
	f.btn3.SetCaption("SetText")
	f.btn3.SetOnClick(f.onBtn3Click)
}

func (f *TMainForm) onBtnClick(sender lcl.IObject) {
	str := lcl.Clipboard.AsText()
	fmt.Println("len1:", len(str))
}

func (f *TMainForm) onBtn2Click(sender lcl.IObject) {
	str := ""
	lcl.Clipboard.GetTextBuf(&str, 1000) // buff不够长
	fmt.Println("len2:", len(str))
}

func (f *TMainForm) onBtn3Click(sender lcl.IObject) {

	lcl.Clipboard.SetAsText("afdsdf")
}
