// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
	"github.com/energye/golcl/lcl"
)

//::private::
type TFrame1Fields struct {
}

func (f *TFrame1) OnButton1Click(sender lcl.IObject) {
	lcl.ShowMessage("Frame1")
}
