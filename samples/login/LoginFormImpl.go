// 在这里写你的事件

package main

import (
	"fmt"

	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/keys"
)

//::private::
type TLoginFormFields struct {
}

func (f *TLoginForm) OnFormCreate(sender lcl.IObject) {
	f.SetShowInTaskBar(types.StAlways)
}

func (f *TLoginForm) OnFormClose(sender lcl.IObject, action *types.TCloseAction) {
	if !isLogin {
		lcl.Application.Terminate()
	}
}

func (f *TLoginForm) OnButtonLoginClick(sender lcl.IObject) {
	usr := f.EditUserName.Text()
	if usr == "" {
		fmt.Println("输入用户名吧")
		f.EditUserName.SetFocus()
		return
	}
	pwd := f.EditPassword.Text()
	if pwd == "" {
		fmt.Println("输入密码吧。")
		f.EditPassword.SetFocus()
		return
	}
	if usr == "admin" && pwd == "admin" {
		isLogin = true
		f.Close()
		//MainForm.Show()
		lcl.Application.MainForm().Show()
	}
}

func (f *TLoginForm) OnFormKeyPress(sender lcl.IObject, key *types.Char) {
	fmt.Println("key:", *key)
	if *key == keys.VkReturn {
		f.ButtonLogin.Click()
	}
}
