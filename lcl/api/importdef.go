//----------------------------------------
//
// Copyright © sxm. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import "github.com/energye/golcl/dylib"

var (
	application_Instance   *dylib.LazyProc // = liblcl.NewProc("Application_Instance")
	application_CreateForm *dylib.LazyProc // = liblcl.NewProc("Application_CreateForm")
	application_Run        *dylib.LazyProc // = liblcl.NewProc("Application_Run")
	application_Initialize *dylib.LazyProc // = liblcl.NewProc("Application_Initialize")

	form_Create2      *dylib.LazyProc // = liblcl.NewProc("Form_Create2")
	form_SetOnWndProc *dylib.LazyProc // = liblcl.NewProc("Form_SetOnWndProc")
	form_SetGoPtr     *dylib.LazyProc // = liblcl.NewProc("Form_SetGoPtr")

	setEventCallback                   *dylib.LazyProc // = liblcl.NewProc("SetEventCallback")
	setMessageCallback                 *dylib.LazyProc // = liblcl.NewProc("SetMessageCallback")
	setThreadSyncCallback              *dylib.LazyProc // = liblcl.NewProc("SetThreadSyncCallback")
	setRequestCallCreateParamsCallback *dylib.LazyProc // = liblcl.NewProc("SetRequestCallCreateParamsCallback")

	dGetStringArrOf *dylib.LazyProc // = liblcl.NewProc("DGetStringArrOf")
	dStrLen         *dylib.LazyProc // = liblcl.NewProc("DStrLen")
	dMove           *dylib.LazyProc // = liblcl.NewProc("DMove")

	dShowMessage *dylib.LazyProc // = liblcl.NewProc("DShowMessage")
	dMessageDlg  *dylib.LazyProc // = liblcl.NewProc("DMessageDlg")

	mouse_Instance  *dylib.LazyProc // = liblcl.NewProc("Mouse_Instance")
	screen_Instance *dylib.LazyProc // = liblcl.NewProc("Screen_Instance")

	dSynchronize *dylib.LazyProc // = liblcl.NewProc("DSynchronize")

	// TMenuItem
	dTextToShortCut *dylib.LazyProc // = liblcl.NewProc("DTextToShortCut")
	dShortCutToText *dylib.LazyProc // = liblcl.NewProc("DShortCutToText")

	// TClipboard
	clipboard_Instance         *dylib.LazyProc // = liblcl.NewProc("Clipboard_Instance")
	clipboard_HasFormat        *dylib.LazyProc // = liblcl.NewProc("Clipboard_HasFormat")
	dSetClipboard              *dylib.LazyProc // = liblcl.NewProc("DSetClipboard")
	dRegisterClipboardFormat   *dylib.LazyProc // = liblcl.NewProc("DRegisterClipboardFormat")
	clipboard_GetAsHtml        *dylib.LazyProc // = liblcl.NewProc("Clipboard_GetAsHtml")
	clipboard_GetTextBuf       *dylib.LazyProc // = liblcl.NewProc("Clipboard_GetTextBuf")
	clipboard_GetAsText        *dylib.LazyProc // = liblcl.NewProc("Clipboard_GetAsText")
	clipboard_SetAsText        *dylib.LazyProc // = liblcl.NewProc("Clipboard_SetAsText")
	dPredefinedClipboardFormat *dylib.LazyProc // = liblcl.NewProc("DPredefinedClipboardFormat")

	// DSysOpen
	dSysOpen *dylib.LazyProc // = liblcl.NewProc("DSysOpen")

	// TMemoryStream
	memoryStream_Read  *dylib.LazyProc // = liblcl.NewProc("MemoryStream_Read")
	memoryStream_Write *dylib.LazyProc // = liblcl.NewProc("MemoryStream_Write")

	// TCanvas
	canvas_BrushCopy     *dylib.LazyProc // = liblcl.NewProc("Canvas_BrushCopy")
	canvas_CopyRect      *dylib.LazyProc // = liblcl.NewProc("Canvas_CopyRect")
	canvas_Draw1         *dylib.LazyProc // = liblcl.NewProc("Canvas_Draw1")
	canvas_Draw2         *dylib.LazyProc // = liblcl.NewProc("Canvas_Draw2")
	canvas_DrawFocusRect *dylib.LazyProc // = liblcl.NewProc("Canvas_DrawFocusRect")
	canvas_FillRect      *dylib.LazyProc // = liblcl.NewProc("Canvas_FillRect")
	canvas_FrameRect     *dylib.LazyProc // = liblcl.NewProc("Canvas_FrameRect")
	canvas_TextRect1     *dylib.LazyProc // = liblcl.NewProc("Canvas_TextRect1")
	canvas_TextRect2     *dylib.LazyProc // = liblcl.NewProc("Canvas_TextRect2")
	canvas_Polygon       *dylib.LazyProc // = liblcl.NewProc("Canvas_Polygon")
	canvas_Polyline      *dylib.LazyProc // = liblcl.NewProc("Canvas_Polyline")
	canvas_PolyBezier    *dylib.LazyProc // = liblcl.NewProc("Canvas_PolyBezier")

	// TImageList
	imageList_Draw1        *dylib.LazyProc // = liblcl.NewProc("ImageList_Draw1")
	imageList_Draw2        *dylib.LazyProc // = liblcl.NewProc("ImageList_Draw2")
	imageList_DrawOverlay1 *dylib.LazyProc // = liblcl.NewProc("ImageList_DrawOverlay1")
	imageList_DrawOverlay2 *dylib.LazyProc // = liblcl.NewProc("ImageList_DrawOverlay2")
	imageList_GetIcon1     *dylib.LazyProc // = liblcl.NewProc("ImageList_GetIcon1")
	imageList_GetIcon2     *dylib.LazyProc // = liblcl.NewProc("ImageList_GetIcon2")

	dExtractFilePath *dylib.LazyProc // = liblcl.NewProc("DExtractFilePath")
	dFileExists      *dylib.LazyProc // = liblcl.NewProc("DFileExists")

	dSelectDirectory1 *dylib.LazyProc // = liblcl.NewProc("DSelectDirectory1")
	dSelectDirectory2 *dylib.LazyProc // = liblcl.NewProc("DSelectDirectory2")
	dInputBox         *dylib.LazyProc // = liblcl.NewProc("DInputBox")
	dInputQuery       *dylib.LazyProc // = liblcl.NewProc("DInputQuery")
	dPasswordBox      *dylib.LazyProc // = liblcl.NewProc("DPasswordBox")
	dInputCombo       *dylib.LazyProc // = liblcl.NewProc("DInputCombo")
	dInputComboEx     *dylib.LazyProc // = liblcl.NewProc("DInputComboEx")

	// TSysLocaled
	dSysLocale *dylib.LazyProc // = liblcl.NewProc("DSysLocale")

	// Shortcut
	dCreateURLShortCut *dylib.LazyProc // = liblcl.NewProc("DCreateURLShortCut")
	dCreateShortCut    *dylib.LazyProc // = liblcl.NewProc("DCreateShortCut")

	// SetProperty
	dSetPropertyValue    *dylib.LazyProc // = liblcl.NewProc("DSetPropertyValue")
	dSetPropertySecValue *dylib.LazyProc // = liblcl.NewProc("DSetPropertySecValue")

	// Printer
	printer_Instance   *dylib.LazyProc // = liblcl.NewProc("Printer_Instance")
	printer_SetPrinter *dylib.LazyProc // = liblcl.NewProc("Printer_SetPrinter")

	// guid
	dGUIDToString *dylib.LazyProc // = liblcl.NewProc("DGUIDToString")
	dStringToGUID *dylib.LazyProc // = liblcl.NewProc("DStringToGUID")
	dCreateGUID   *dylib.LazyProc // = liblcl.NewProc("DCreateGUID")

	// libResources
	dGetLibResourceCount *dylib.LazyProc // = liblcl.NewProc("DGetLibResourceCount")
	dGetLibResourceItem  *dylib.LazyProc // = liblcl.NewProc("DGetLibResourceItem")
	dModifyLibResource   *dylib.LazyProc // = liblcl.NewProc("DModifyLibResource")
	dLibAbout            *dylib.LazyProc // = liblcl.NewProc("DLibAbout")

	// 库的信息
	dLibStringEncoding *dylib.LazyProc // = liblcl.NewProc("DLibStringEncoding")
	dLibVersion        *dylib.LazyProc // = liblcl.NewProc("DLibVersion")

	dMainThreadId    *dylib.LazyProc // = liblcl.NewProc("DMainThreadId")
	dCurrentThreadId *dylib.LazyProc // = liblcl.NewProc("DCurrentThreadId")

	dInitGoDll *dylib.LazyProc // = liblcl.NewProc("DInitGoDll")

	dFindControl           *dylib.LazyProc // = liblcl.NewProc("DFindControl")
	dFindLCLControl        *dylib.LazyProc // = liblcl.NewProc("DFindLCLControl")
	dFindOwnerControl      *dylib.LazyProc // = liblcl.NewProc("DFindOwnerControl")
	dFindControlAtPosition *dylib.LazyProc // = liblcl.NewProc("DFindControlAtPosition")
	dFindLCLWindow         *dylib.LazyProc // = liblcl.NewProc("DFindLCLWindow")
	dFindDragTarget        *dylib.LazyProc // = liblcl.NewProc("DFindDragTarget")
)
