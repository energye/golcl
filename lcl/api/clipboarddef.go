//----------------------------------------
//
// Copyright © sxm. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
	"unsafe"

	"github.com/energye/golcl/lcl/types"
)

func Clipboard_Instance() uintptr {
	if IsNilProcApi(clipboard_Instance) {
		clipboard_Instance = liblcl.NewProc("Clipboard_Instance")
	}
	r, _, _ := clipboard_Instance.Call()
	return r
}

func Clipboard_HasFormat(obj uintptr, aFormatID types.TClipboardFormat) bool {
	if IsNilProcApi(clipboard_HasFormat) {
		clipboard_HasFormat = liblcl.NewProc("Clipboard_HasFormat")
	}
	r, _, _ := clipboard_HasFormat.Call(obj, uintptr(aFormatID))
	return r != 0
}

func Clipboard_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if BufSize <= 0 {
		return 0
	}
	buff := make([]byte, BufSize)
	if IsNilProcApi(clipboard_GetTextBuf) {
		clipboard_GetTextBuf = liblcl.NewProc("Clipboard_GetTextBuf")
	}
	ret, _, _ := clipboard_GetTextBuf.Call(obj, uintptr(unsafe.Pointer(&buff[0])), uintptr(BufSize))
	if int(ret) < len(buff) {
		*Buffer = string(buff[:ret])
	}
	return int32(ret)
}

func Clipboard_GetAsText(obj uintptr) string {
	if IsNilProcApi(clipboard_GetAsText) {
		clipboard_GetAsText = liblcl.NewProc("Clipboard_GetAsText")
	}
	ret, _, _ := clipboard_GetAsText.Call(obj)
	return DStrToGoStr(ret)
}

func Clipboard_SetAsText(obj uintptr, value string) {
	if IsNilProcApi(clipboard_SetAsText) {
		clipboard_SetAsText = liblcl.NewProc("Clipboard_SetAsText")
	}
	clipboard_SetAsText.Call(obj, GoStrToDStr(value))
}

func Clipboard_GetAsHtml(obj uintptr, ExtractFragmentOnly bool) string {
	if IsNilProcApi(clipboard_GetAsHtml) {
		clipboard_GetAsHtml = liblcl.NewProc("Clipboard_GetAsHtml")
	}
	ret, _, _ := clipboard_GetAsHtml.Call(obj, GoBoolToDBool(ExtractFragmentOnly))
	return DStrToGoStr(ret)
}

func DSetClipboard(obj uintptr) uintptr {
	if IsNilProcApi(dSetClipboard) {
		dSetClipboard = liblcl.NewProc("DSetClipboard")
	}
	r, _, _ := dSetClipboard.Call(obj)
	return r
}

func DRegisterClipboardFormat(aFormat string) types.TClipboardFormat {
	if IsNilProcApi(dRegisterClipboardFormat) {
		dRegisterClipboardFormat = liblcl.NewProc("DRegisterClipboardFormat")
	}
	r, _, _ := dRegisterClipboardFormat.Call(GoStrToDStr(aFormat))
	return types.TClipboardFormat(r)
}

func DPredefinedClipboardFormat(aFormat types.TPredefinedClipboardFormat) types.TClipboardFormat {
	if IsNilProcApi(dPredefinedClipboardFormat) {
		dPredefinedClipboardFormat = liblcl.NewProc("DPredefinedClipboardFormat")
	}
	r, _, _ := dPredefinedClipboardFormat.Call(uintptr(aFormat))
	return types.TClipboardFormat(r)
}
