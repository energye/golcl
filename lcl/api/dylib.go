//----------------------------------------
//
// Copyright © sxm. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
	"sync"

	"github.com/energye/golcl/dylib"
)

var (
	// 全局导入库
	liblcl *dylib.LazyDLL // = loadUILib()

	// 导出的DLL，考虑到导入的函数太多了，导致go无法编译通过
	// 只能动态作，这样可能牺牲一点性能吧，但文件大小会减小几M左右吧。
	functions sync.Map
)

func getLazyProc(name string) *dylib.LazyProc {
	if val, ok := functions.Load(name); !ok {
		proc := liblcl.NewProc(name)
		functions.Store(name, proc)
		return proc
	} else {
		return val.(*dylib.LazyProc)
	}
}

func GetLazyProc(name string) *dylib.LazyProc {
	return getLazyProc(name)
}

func APIInit() {
	liblcl = loadUILib()
}
