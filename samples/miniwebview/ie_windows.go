//go:build windows
// +build windows

package main

import (
	"fmt"

	"github.com/energye/golcl/lcl/win"

	"github.com/energye/golcl/lcl"
)

func SetIEVersion(webview *lcl.TMiniWebview) {
	if win.IsAdministrator() {
		v := webview.GetIEVersion()
		fmt.Println("IE Version:", v)
		if v >= 7 {
			webview.SetIEVersion(v)
		}
	} else {
		fmt.Println("需要Administrator权限才可使用。")
	}
}
