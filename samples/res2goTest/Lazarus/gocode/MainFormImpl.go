// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
	"github.com/energye/golcl/lcl"
)

//::private::
type TMainFormFields struct {
	test string
}

func (f *TMainForm) OnFormCreate(sender lcl.IObject) {
	f.test = "aaaaaa"
	f.ScreenCenter()
	// Lazarus 无此属性，所以手动设置下
	f.Panel1.SetParentBackground(false)
	f.Panel2.SetParentBackground(false)
	f.Panel3.SetParentBackground(false)
}

func (f *TMainForm) OnButton1Click(sender lcl.IObject) {
	About.SetCaption(f.test)
	About.Show()
}

func (f *TMainForm) OnButton2Click(sender lcl.IObject) {
	btn := lcl.AsButton(sender)
	lcl.ShowMessage("Name:" + btn.Name())
}

func (f *TMainForm) OnPanel1MouseEnter(sender lcl.IObject) {
	pnl := lcl.AsPanel(sender)
	pnl.SetLeft(pnl.Left() + 5)
}

func (f *TMainForm) OnPanel1MouseLeave(sender lcl.IObject) {
	pnl := lcl.AsPanel(sender)
	pnl.SetLeft(pnl.Left() - 5)
}
