//----------------------------------------
//
// Copyright © sxm. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

func Application_Instance() uintptr {
	if IsNilProcApi(application_Instance) {
		application_Instance = liblcl.NewProc("Application_Instance")
	}
	ret, _, _ := application_Instance.Call()
	return ret
}

func Application_CreateForm(app uintptr, initScale bool) uintptr {
	if IsNilProcApi(application_CreateForm) {
		application_CreateForm = liblcl.NewProc("Application_CreateForm")
	}
	ret, _, _ := application_CreateForm.Call(app, GoBoolToDBool(initScale))
	return ret
}

func Application_Run(app uintptr) {
	defer func() {
		// 开启了finalizerOn选项后，以防止关闭库后GC还没开始调用。
		callGC()
		// 运行结束后就结束close掉lib，不然他不会关掉的
		closeLib()
	}()
	if IsNilProcApi(application_Run) {
		application_Run = liblcl.NewProc("Application_Run")
	}
	application_Run.Call(app)
}

func Application_Initialize(obj uintptr) {
	if IsNilProcApi(application_Initialize) {
		application_Initialize = liblcl.NewProc("Application_Initialize")
	}
	application_Initialize.Call(obj)
}
