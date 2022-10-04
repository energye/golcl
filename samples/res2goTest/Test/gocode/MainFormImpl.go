// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
	"github.com/energye/golcl/lcl"
)

//::private::
type TMainFormFields struct {
	frame1 *TFrame1
	frame2 *TFrame2
}

func (f *TMainForm) OnFormCreate(sender lcl.IObject) {
	f.frame1 = NewFrame1(f)
	f.frame2 = NewFrame2(f)

	f.frame1.Hide()
	f.frame2.Hide()

	f.frame1.SetParent(f.Panel2)
	f.frame2.SetParent(f.Panel2)
}

func (f *TMainForm) OnBtnShowAboutClick(sender lcl.IObject) {
	// 动态创建，完后释放掉，不使用Application.CreateForm
	about := NewAbout(f)
	defer about.Free()
	about.ShowModal()
}

func (f *TMainForm) OnBtnShowFrame1Click(sender lcl.IObject) {
	f.frame2.Hide()
	f.frame1.Show()
}

func (f *TMainForm) OnBtnShowFrame2Click(sender lcl.IObject) {
	f.frame1.Hide()
	f.frame2.Show()
}
