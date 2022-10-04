//go:build windows
// +build windows

// miniblink及wke头文件导入
// 由sxm翻译，应用于golcl

package miniblink

import (
	"runtime"
	"syscall"
)

var (
	wkedll = LoadMiniBlinkDLL()
	is386  = runtime.GOARCH == "386"
)

func LoadMiniBlinkDLL() *syscall.LazyDLL {
	if !is386 {
		return syscall.NewLazyDLL("miniblink_x64.dll")
	}
	return syscall.NewLazyDLL("node.dll")
}
