//----------------------------------------
//
// Copyright © sxm. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
	"github.com/energye/golcl/dylib"
	"github.com/energye/golcl/pkgs/libname"
)

// 加载库
func loadUILib() *dylib.LazyDLL {
	lib := dylib.NewLazyDLL(libname.LibName)
	err := lib.Load()
	if err != nil {
		panic(err)
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
