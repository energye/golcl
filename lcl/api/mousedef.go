//----------------------------------------
//
// Copyright © sxm. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

func Mouse_Instance() uintptr {
	if IsNilProcApi(mouse_Instance) {
		mouse_Instance = liblcl.NewProc("Mouse_Instance")
	}
	ret, _, _ := mouse_Instance.Call()
	return ret
}
