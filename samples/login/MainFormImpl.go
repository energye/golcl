// 在这里写你的事件

package main

import (
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

//::private::
type TMainFormFields struct {
}

var (
	// 是否登录
	isLogin bool
)

func (f *TMainForm) OnFormCreate(sender lcl.IObject) {

}

func (f *TMainForm) OnFormDestroy(sender lcl.IObject) {

}

func (f *TMainForm) OnFormCloseQuery(sender lcl.IObject, canClose *bool) {
	if isLogin {
		*canClose = lcl.MessageDlg("是否退出?", types.MtConfirmation, types.MbYes, types.MbNo) == types.MrYes
	}
}
