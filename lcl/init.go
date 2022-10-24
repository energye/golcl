//----------------------------------------
//
// Copyright © sxm. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"os"
	"runtime"

	. "github.com/energye/golcl/lcl/api"
)

var (
	// 几个实例类，不需要Create即可访问，同时也不需要手动Free

	Application *TApplication // 应用程序管理
	Screen      *TScreen      // 屏幕
	Mouse       *TMouse       // 鼠标
	Clipboard   *TClipboard   // 剪切板
	Printer     *TPrinter     // 打印机
)

func LCLInit() {
	if !DEBUG {
		defer func() {
			if err := recover(); err != nil {
				showError(err)
				os.Exit(1)
			}
		}()
	}
	// 这个似乎得默认加上，锁定主线程，防止中间被改变
	runtime.LockOSThread()
	// 设置事件的回调函数，因go中callback数量有限，只好折中处理
	SetEventCallback(eventCallback)
	// 消息回调
	SetMessageCallback(messageCallback)
	// 线程同步回调
	SetThreadSyncCallback(threadSyncCallback)
	// 调求回调CreateParams方法
	SetRequestCallCreateParamsCallback(requestCallCreateParamsCallback)

	// 导入几个实例类
	Application = AsApplication(Application_Instance())
	Screen = AsScreen(Screen_Instance())
	Mouse = AsMouse(Mouse_Instance())
	Clipboard = AsClipboard(Clipboard_Instance())
	Printer = AsPrinter(Printer_Instance())

	// 尝试加载ICON，仅Windows下有效，尝试加载名为MAINICON的图标
	tryLoadAppIcon()
}
