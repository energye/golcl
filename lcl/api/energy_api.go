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

func APIInit() {
	uiLib = loadUILib()
}

// Get Import
func ImportDefFunc(importTable []*dllimports.ImportTable, index int) dllimports.ProcAddr {
	return dllimports.ImportDefFunc(uiLib, importTable, index)
}

// EnergyLibRelease 在energy中释放
func EnergyLibRelease() {
	callGC()
	closeLib()
}
