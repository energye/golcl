package main

import (
	"github.com/energye/golcl/lcl"
)

//::private::
type TForm1Fields struct {
}

func (f *TForm1) OnFormCreate(sender lcl.IObject) {
	f.ScreenCenter()
}

func (f *TForm1) OnButton2Click(sender lcl.IObject) {
	f.ScaleForCurrentDpi()
}

func (f *TForm1) OnButton1Click(sender lcl.IObject) {
	f.ScaleForPPI(125)
}
