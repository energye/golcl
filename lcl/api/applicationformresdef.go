//----------------------------------------
//
// Copyright © sxm. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//----------------------------------------
// 加载文件或者内存中的窗口资源文件功能
// 需要配合窗口设计器使用
//----------------------------------------
package api

import "github.com/energye/dylib"

var (
	resFormLoadFromStream       *dylib.LazyProc //= liblcl.NewProc("ResFormLoadFromStream")
	resFormLoadFromFile         *dylib.LazyProc //= liblcl.NewProc("ResFormLoadFromFile")
	resFormLoadFromResourceName *dylib.LazyProc //= liblcl.NewProc("ResFormLoadFromResourceName")
)

// ResFormLoadFromStream
func ResFormLoadFromStream(obj, root uintptr) {
	if resFormLoadFromStream == nil {
		resFormLoadFromStream = liblcl.NewProc("ResFormLoadFromStream")
	}
	resFormLoadFromStream.Call(obj, root)
}

func ResFormLoadFromFile(filename string, root uintptr) {
	if resFormLoadFromFile == nil {
		resFormLoadFromFile = liblcl.NewProc("ResFormLoadFromFile")
	}
	resFormLoadFromFile.Call(GoStrToDStr(filename), root)
}

func ResFormLoadFromResourceName(instance uintptr, resName string, root uintptr) {
	if resFormLoadFromResourceName == nil {
		resFormLoadFromResourceName = liblcl.NewProc("ResFormLoadFromResourceName")
	}
	resFormLoadFromResourceName.Call(instance, GoStrToDStr(resName), root)
}
