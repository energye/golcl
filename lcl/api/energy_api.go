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

var (
	releaseCallback  func()
	energyImportDefs []*dllimports.ImportTable
)

const (
	// CustomWidgetInterface for Linux
	interface_CustomWidgetSetInitialization = iota
	interface_CustomWidgetSetFinalization
)

func init() {
	energyImportDefs = []*dllimports.ImportTable{
		// CustomWidgetInterface for Linux
		dllimports.NewEnergyImport("Interface_CustomWidgetSetInitialization", 0),
		dllimports.NewEnergyImport("Interface_CustomWidgetSetFinalization", 0),
	} //end
}

func APIInit() {
	uiLib = loadUILib()
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

// CustomWidgetSetInitialization
// 自定义组件初始化 Linux
func CustomWidgetSetInitialization() uintptr {
	r1, _, _ := dllimports.ImportDefFunc(uiLib, energyImportDefs, interface_CustomWidgetSetInitialization).Call()
	return r1
}

// CustomWidgetSetFinalization
//  自定义组件销毁 Linux
func CustomWidgetSetFinalization() uintptr {
	r1, _, _ := dllimports.ImportDefFunc(uiLib, energyImportDefs, interface_CustomWidgetSetFinalization).Call()
	return r1
}
