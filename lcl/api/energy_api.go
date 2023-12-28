//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// energy extension
//
//govcl/vcl/api -> golcl/lcl/api/energy_api.go

package api

import "github.com/energye/golcl/lcl/api/dllimports"

var releaseCallback func()

func APIInit() {
	uiLib = loadUILib()
}

// Get Import
func ImportDefFunc(importTable []*dllimports.ImportTable, index int) dllimports.ProcAddr {
	return dllimports.ImportDefFunc(uiLib, importTable, index)
}

// EnergyLibRelease 在energy中释放
func EnergyLibRelease() {
	if releaseCallback != nil {
		releaseCallback()
	}
	callGC()
	closeLib()
}

// SetReleaseCallback 应用运行结束后释放资源之前执行
func SetReleaseCallback(fn func()) {
	if releaseCallback == nil {
		releaseCallback = fn
	}
}
