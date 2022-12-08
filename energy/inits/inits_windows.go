//go:build windows
// +build windows

//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// energy extension
//
//govcl -> golcl/energy/inits/inits.go

package inits

import "github.com/energye/golcl/lcl/win"

func winInit() {
	//win
	win.Init()
}
