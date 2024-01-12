//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------
// energy extension
//
//govcl/vcl/api/dylib.go -> golcl/lcl/api/dylib.go

package api

import (
	"github.com/energye/golcl/lcl/api/dllimports"
	"github.com/energye/golcl/pkgs/libname"
	"unsafe"
)

var (
	uiLib           dllimports.DLL                 // 全局导入库
	loadLibCallback func() (dllimports.DLL, error) // 自定义加载liblcl动态库回调函数
)

// SetLoadLibCallback
//  设置加载liblcl动态库回调函数
//  如果设置该回调函数我们可以自定义加载动态链接库
//  在调用 inits.Init 之前设置
func SetLoadLibCallback(fn func() (dllimports.DLL, error)) {
	if loadLibCallback == nil {
		loadLibCallback = fn
	}
}

func loadUILib() dllimports.DLL {
	if loadLibCallback != nil {
		dll, _ := loadLibCallback()
		return dll
	} else {
		lib, err := dllimports.NewDLL(libname.LibName)
		if lib == 0 && err != nil {
			panic(err)
		}
		return lib
	}
}

func closeLib() {
	if uiLib != 0 {
		uiLib.Release()
		uiLib = 0
	}
}

// 调用自动生成的API列表中的函数
func syscallN(trap int, args ...uintptr) uintptr {
	r1, _, _ := dllimports.GetImportFunc(uiLib, trap).Call(args...)
	return r1
}

func syscallGetTextBuf(trap int, obj uintptr, buffer *string, bufSize uintptr) uintptr {
	if buffer == nil || bufSize == 0 {
		return 0
	}
	strPtr := make([]uint8, bufSize+1)
	sLen := syscallN(trap, obj, uintptr(unsafe.Pointer(&strPtr[0])), bufSize)
	if sLen > 0 {
		*buffer = string(strPtr[:sLen])
	}
	return sLen
}

// 调用手动导入的API列表中的函数
func defSyscallN(trap int, args ...uintptr) uintptr {
	r1, _, _ := dllimports.GetImportDefFunc(uiLib, trap).Call(args...)
	return r1
}
