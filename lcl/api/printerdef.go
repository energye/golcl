//----------------------------------------
//
// Copyright © sxm. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

// Printer
func Printer_Instance() uintptr {
	if IsNilProcApi(printer_Instance) {
		printer_Instance = liblcl.NewProc("Printer_Instance")
	}
	r, _, _ := printer_Instance.Call()
	return r
}

func Printer_SetPrinter(obj uintptr, aName string) {
	if IsNilProcApi(printer_SetPrinter) {
		printer_SetPrinter = liblcl.NewProc("Printer_SetPrinter")
	}
	printer_SetPrinter.Call(obj, GoStrToDStr(aName))
}
