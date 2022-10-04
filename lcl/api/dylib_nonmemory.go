//----------------------------------------
//
// Copyright © sxm. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build !memorydll
// +build !memorydll

package api

import (
	"github.com/energye/dylib"
	"github.com/energye/golcl/pkgs/libname"
)

// 加载库
func loadUILib() *dylib.LazyDLL {
	libName := libname.GetDLLName()
	// 如果支持运行时释放，则使用此种方法
	if support, newDLLPath := checkAndReleaseDLL(); support {
		libName = newDLLPath
	} else {
		if libname.LibName != "" {
			libName = libname.LibName
		}
	}
	lib := dylib.NewLazyDLL(libName)
	err := lib.Load()
	if err != nil {
		panic(err)
	}
	if getLibType(lib) != 1 {
		// 当前已经不再支持VCL库了。如果有需要，请使用最后一个支持VCL版本的代码：https://github.com/sxm/govcl/tree/last-vcl-support。
		panic("The VCL library is no longer supported. If necessary, please use the last code that supports VCL version: https://github.com/sxm/govcl/tree/last-vcl-support.")
	}
	return lib
}

func getLibType(lib *dylib.LazyDLL) int32 {
	proc := lib.NewProc("DGetLibType")
	r, _, _ := proc.Call()
	return int32(r)
}

// 获取dll库实例，用于在外扩展第三方组件的。移动来自dfuncs.go
func GetLibLCL() *dylib.LazyDLL {
	return liblcl
}

func closeLib() {
	if liblcl != nil {
		liblcl.Close()
		liblcl = nil
	}
}
