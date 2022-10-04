//----------------------------------------
//
// Copyright © sxm. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build !tempdll && !memorydll
// +build !tempdll,!memorydll

// 指令为：!tempdll && !memorydll

package api

func checkAndReleaseDLL() (bool, string) {
	return false, ""
}
