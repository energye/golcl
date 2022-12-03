// 在这里写你的事件

package main

import (
	"fmt"
	"strings"

	"github.com/energye/golcl/lcl/types/keys"
	"github.com/energye/golcl/lcl/win"

	"github.com/energye/golcl/lcl/types"

	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/pkgs/miniblink"
)

//::private::
type TMainFormFields struct {
	web *miniblink.TMiniBlinkWebview
}

func (f *TMainForm) OnFormCreate(sender lcl.IObject) {
	f.ScreenCenter()
	fmt.Println("OnFormCreate:", win.GetCurrentThreadId())
	f.web = miniblink.NewMiniBlinkWebview(f.PnlWeb.Handle())
	f.web.Show(true)
	f.web.OnTitleChanged = func(sender *miniblink.TMiniBlinkWebview, title string) {
		//fmt.Println("标题改变:", title)
		fmt.Println("onTitleChanged:", win.GetCurrentThreadId())
		f.SetCaption(title + " - ying32")
	}

	f.web.OnURLChanged = func(sender *miniblink.TMiniBlinkWebview, url string) {
		fmt.Println("URL改变:", url)
	}

	f.web.OnDocumentReady = func(sender *miniblink.TMiniBlinkWebview) {
		fmt.Println("文档已准备。")
	}

	f.web.OnLoadingFinish = func(sender *miniblink.TMiniBlinkWebview, url string, result miniblink.WkeLoadingResult, failedReason string) {
		fmt.Println("加载完成：", url, ", result:", result, ", failedReason:", failedReason)
	}
}

func (f *TMainForm) OnFormDestroy(sender lcl.IObject) {
	if f.web != nil && f.web.IsValid() {
		f.web.Free()
	}
}

func (f *TMainForm) OnPnlWebResize(sender lcl.IObject) {
	if f.web != nil && f.web.IsValid() {
		f.web.MoveWindow(0, 0, int(f.Width()), int(f.Height()))
	}
}

func (f *TMainForm) OnBtnReloadClick(sender lcl.IObject) {
	f.web.Webview.Reload()
}

func (f *TMainForm) OnBtnLoadFromFileClick(sender lcl.IObject) {
	if f.OpenDialog1.Execute() {
		fileName := f.OpenDialog1.FileName()
		f.web.Webview.LoadFile(fileName)
	}
}

func (f *TMainForm) OnBtnLoadFromStringClick(sender lcl.IObject) {
	result := HTMLForm.ShowModal()
	if result == types.MrOk {
		f.web.Webview.LoadHTML(HTMLForm.Memo1.Text())
	}
}

//func (f *TMainForm) OnFormKeyPress(sender lcl.IObject, key *types.Char) {
//
//}

func (f *TMainForm) OnActGoBackExecute(sender lcl.IObject) {
	f.web.Webview.GoBack()
}

func (f *TMainForm) OnActGoBackUpdate(sender lcl.IObject) {
	lcl.AsAction(sender).SetEnabled(f.web != nil && f.web.IsValid() && f.web.Webview.CanGoBack())
}

func (f *TMainForm) OnActGoForwardExecute(sender lcl.IObject) {
	f.web.Webview.GoForward()
}

func (f *TMainForm) OnActGoForwardUpdate(sender lcl.IObject) {
	lcl.AsAction(sender).SetEnabled(f.web != nil && f.web.IsValid() && f.web.Webview.CanGoForward())
}

func (f *TMainForm) OnActNavExecute(sender lcl.IObject) {
	text := strings.TrimSpace(f.EdtURL.Text())
	if text == "" {
		lcl.ShowMessage("请输入一个URL！")
		return
	}
	f.web.Webview.LoadURL(text)
}

func (f *TMainForm) OnActNavUpdate(sender lcl.IObject) {
	//lcl.AsAction(sender).SetEnabled(f.web != nil && f.web.IsValid() && f.web.Webview.CanGoForward())
}

func (f *TMainForm) OnFormKeyDown(sender lcl.IObject, key *types.Char, shift types.TShiftState) {
	if *key == keys.VkReturn {
		f.ActNav.Execute()
		*key = 0
	}
}

func (f *TMainForm) OnBtnExecJSClick(sender lcl.IObject) {
	result := JsForm.ShowModal()
	if result == types.MrOk && f.web.IsValid() {
		val := f.web.Webview.RunJS(JsForm.Memo1.Text())
		switch val.TypeOf() {
		case miniblink.JSTYPE_NUMBER:
			fmt.Println("Number")
		case miniblink.JSTYPE_STRING:
			fmt.Println("String")
			fmt.Println("返回值：", f.web.Webview.GlobalExec().ToTempString(val))
		case miniblink.JSTYPE_BOOLEAN:
			fmt.Println("Boolean")
		case miniblink.JSTYPE_OBJECT:
			fmt.Println("Object")
		case miniblink.JSTYPE_FUNCTION:
			fmt.Println("Function")
		case miniblink.JSTYPE_UNDEFINED:
			fmt.Println("Undefined")
		case miniblink.JSTYPE_ARRAY:
			fmt.Println("Array")
		case miniblink.JSTYPE_NULL:
			fmt.Println("Null")
		default:
			fmt.Println("未知")

		}

	}
}
