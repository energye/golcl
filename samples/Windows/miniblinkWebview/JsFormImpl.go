// 在这里写你的事件

package main

import "github.com/energye/golcl/lcl"

//::private::
type TJsFormFields struct {
}

func (f *TJsForm) OnFormCreate(sender lcl.IObject) {
	f.ScreenCenter()
}
