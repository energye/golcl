package main

import (
	"github.com/energye/golcl/lcl"
)

// 本例程需要go1.16版本支持。

//::private::
type TForm1Fields struct {
}

func (f *TForm1) OnButton1Click(sender lcl.IObject) {
	lcl.ShowMessage("hello")
}
