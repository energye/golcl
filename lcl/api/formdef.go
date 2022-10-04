//----------------------------------------
//
// Copyright © sxm. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

func Form_Create2(owner uintptr, initScale bool) uintptr {
	if IsNilProcApi(form_Create2) {
		form_Create2 = liblcl.NewProc("Form_Create2")
	}
	ret, _, _ := form_Create2.Call(owner, GoBoolToDBool(initScale))
	return ret
}

func Form_SetOnWndProc(obj uintptr, fn interface{}) {
	if IsNilProcApi(form_SetOnWndProc) {
		form_SetOnWndProc = liblcl.NewProc("Form_SetOnWndProc")
	}
	form_SetOnWndProc.Call(obj, addMessageEventToMap(obj, fn))
}

func Form_SetGoPtr(obj uintptr, ptr uintptr) {
	if IsNilProcApi(form_SetGoPtr) {
		form_SetGoPtr = liblcl.NewProc("Form_SetGoPtr")
	}
	form_SetGoPtr.Call(obj, ptr)
}
