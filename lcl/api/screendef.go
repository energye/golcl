//----------------------------------------
//
// Copyright © sxm. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

func Screen_Instance() uintptr {
	if IsNilProcApi(screen_Instance) {
		screen_Instance = liblcl.NewProc("Screen_Instance")
	}
	ret, _, _ := screen_Instance.Call()
	return ret
}
