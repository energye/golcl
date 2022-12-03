package main

import "C"
import (
	"runtime"
	"syscall"
	"unsafe"

	"github.com/energye/golcl/lcl/types"

	"github.com/energye/golcl/lcl/types/messages"

	"github.com/energye/golcl/lcl/rtl"

	"github.com/energye/golcl/lcl"

	"github.com/energye/golcl/lcl/win"
)

var (
	nppData           NppData
	nppPluginFuncItem []FuncItem
	nppPluginName     = "Go语言NPP插件测试"
)

func Test12() uintptr {
	defer func() {
		if err := recover(); err != nil {
			win.OutputDebugString("err: ", err)
		}
	}()
	//Form1.ScreenCenter()
	Form1.Show()
	return 0
}

func addFuncItem(name string, fn func() uintptr) {
	defer func() {
		if err := recover(); err != nil {
			win.OutputDebugString("addFuncItem err: ", err)
		}
	}()
	var item FuncItem
	copy(item.ItemName[:], syscall.StringToUTF16(name))
	item.Init2Check = false
	item.PShKey = nil
	item.PFunc = syscall.NewCallbackCDecl(fn)
	nppPluginFuncItem = append(nppPluginFuncItem, item)
}

func commandMenuInit() {
	addFuncItem("测试函数", Test12)
}

//export isUnicode
func isUnicode() int32 {
	return 1
}

//export getName
func getName() PTCHAR {
	return PTCHAR(unsafe.Pointer(syscall.StringToUTF16Ptr(nppPluginName)))
}

//export setInfo
func setInfo(nppHandle, scintillaMainHandle, scintillaSecondHandle HWND) {
	defer func() {
		if err := recover(); err != nil {
			win.OutputDebugString("addFuncItem err: ", err)
		}
	}()
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	nppData.NppHandle = nppHandle
	nppData.ScintillaMainHandle = scintillaMainHandle
	nppData.ScintillaSecondHandle = scintillaSecondHandle
	win.OutputDebugString("setInfo: ", nppData)
	commandMenuInit()

	// 主要是setInfo测试出来是在主线程调用的，而liblcl中的值则不为主线程
	// liblcl中的MainThreadId不为主线程id的原因是go的dll会另起一条线程初始化

	// 不要在go的dll中使用以下方法
	//lcl.Application.Initialize()
	//lcl.Application.SetMainFormOnTaskBar()
	//lcl.Application.CreateForm()
	//lcl.Application.Run()

	// 在go的dll中应该使用专用的初始函数
	// 此方法也只适合第三方非govcl程序，不适合go+govcl+godll+govcl方式
	rtl.InitGoDll(0) //0则自动获取当前线程Id

	// 设置application的icon为notepad++的icon
	lcl.Application.Icon().SetHandle(win.LoadIcon(win.GetSelfModuleHandle(), 100))

	// 新建一个窗口
	if Form1 == nil {
		// Form1的OnDestroy失效
		// Form1估计也有内存泄露
		Form1 = NewForm1(lcl.Application)
		Form1.SetShowInTaskBar(types.StAlways)
		Form1.SetFormStyle(types.FsStayOnTop)
	}
}

//export messageProc
func messageProc(iMessage UINT, wParam WPARAM, lParam LPARAM) LRESULT {
	if iMessage == messages.WM_DESTROY {
		win.OutputDebugString("------------------释放")
	}
	return 1
}

//export getFuncsArray
func getFuncsArray(nbF *int32) uintptr {
	*nbF = int32(len(nppPluginFuncItem))
	return uintptr(unsafe.Pointer(&nppPluginFuncItem[0]))
}

//export beNotified
func beNotified(notifyCode /*TSCNotification*/ uintptr) {

	// MyNppPlugin.beNotified(notifyCode);
}
