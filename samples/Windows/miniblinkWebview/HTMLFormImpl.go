// 在这里写你的事件

package main

import "github.com/energye/golcl/lcl"

//::private::
type THTMLFormFields struct {
}

func (f *THTMLForm) OnFormCreate(sender lcl.IObject) {
	f.ScreenCenter()
}
