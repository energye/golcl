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
	// 全局导入库
	uiLib dllimports.DLL
)

func loadUILib() dllimports.DLL {
	libName := libname.LibName
	lib, err := dllimports.NewDLL(libName)
	if err != nil {
		panic(err)
	}
	return lib
}

func closeLib() {
	if uiLib != 0 {
		uiLib.Release()
		uiLib = 0
	}
}

func getDLLName() string {
	return libname.GetDLLName()
}

// 调用自动生成的API列表中的函数
func syscallN(trap int, args ...uintptr) uintptr {
	r1, r2, err := dllimports.GetImportFunc(uiLib, trap).Call(args...)
	println("syscall-n:", r1, r2, err)
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
