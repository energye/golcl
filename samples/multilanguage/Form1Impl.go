// 在这里写你的事件

package main

import (
	"fmt"

	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/i18n"
	"github.com/energye/golcl/lcl/types"
)

var (
	testMessage  = "这是一个测试消息！"
	testMessage2 = "你确定么？"
)

//::private::
type TForm1Fields struct {
}

func (f *TForm1) OnFormCreate(sender lcl.IObject) {
	i18n.InitComponentLang(f)
	f.ScreenCenter()

	item := lcl.NewMenuItem(f)
	item.SetCaption("Languages")
	for key, val := range i18n.LocalLangs {
		subitem := lcl.NewMenuItem(f)
		subitem.SetGroupIndex(1)
		subitem.SetRadioItem(true)
		subitem.SetCaption(fmt.Sprintf("%d - %s", val.Language.Id, val.Language.Description))
		subitem.SetOnClick(f.OnLanguageMenuItemClick)
		subitem.SetTag(key)
		if i18n.CurrentLang == val.Language.Name {
			subitem.SetChecked(true)
		}
		item.Add(subitem)
	}
	f.MainMenu1.Items().Add(item)
}

func (f *TForm1) OnLanguageMenuItemClick(sender lcl.IObject) {
	id := lcl.AsMenuItem(sender).Tag()
	if lang, ok := i18n.LocalLangs[id]; ok {
		fmt.Println(lang)
		i18n.ChangeLang(lang.Language.Name)
		i18n.WriteSetLang(lang.Language.Name)
		lcl.AsMenuItem(sender).SetChecked(true)
	}
}

func (f *TForm1) OnButton1Click(sender lcl.IObject) {
	lcl.ShowMessage(testMessage)
}

func (f *TForm1) OnButton2Click(sender lcl.IObject) {
	if lcl.MessageDlg(testMessage2, types.MtConfirmation, types.MbYes, types.MbNo) == types.IdYes {

	}
}

func (f *TForm1) OnButton3Click(sender lcl.IObject) {
	lcl.ShowMessage(i18n.IdRes("testMessage3"))
}

// 初始就注册
func init() {
	i18n.RegsiterVarString("testMessage", &testMessage)
	i18n.RegsiterVarString("testMessage2", &testMessage2)
	//multilang.RegsiterVar(&testMessage)
}
