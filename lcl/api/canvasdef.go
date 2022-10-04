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

	. "github.com/energye/golcl/lcl/types"
)

func Canvas_BrushCopy(obj uintptr, dest TRect, bitmap uintptr, source TRect, color TColor) {
	if IsNilProcApi(canvas_BrushCopy) {
		canvas_BrushCopy = liblcl.NewProc("Canvas_BrushCopy")
	}
	canvas_BrushCopy.Call(obj, uintptr(unsafe.Pointer(&dest)), bitmap, uintptr(unsafe.Pointer(&source)), uintptr(color))
}

func Canvas_CopyRect(obj uintptr, dest TRect, canvas uintptr, source TRect) {
	if IsNilProcApi(canvas_CopyRect) {
		canvas_CopyRect = liblcl.NewProc("Canvas_CopyRect")
	}
	canvas_CopyRect.Call(obj, uintptr(unsafe.Pointer(&dest)), canvas, uintptr(unsafe.Pointer(&source)))
}

func Canvas_Draw1(obj uintptr, x, y int32, graphic uintptr) {
	if IsNilProcApi(canvas_Draw1) {
		canvas_Draw1 = liblcl.NewProc("Canvas_Draw1")
	}
	canvas_Draw1.Call(obj, uintptr(x), uintptr(y), graphic)
}

func Canvas_Draw2(obj uintptr, x, y int32, graphic uintptr, opacity uint8) {
	if IsNilProcApi(canvas_Draw2) {
		canvas_Draw2 = liblcl.NewProc("Canvas_Draw2")
	}
	canvas_Draw2.Call(obj, uintptr(x), uintptr(y), graphic, uintptr(opacity))
}

func Canvas_DrawFocusRect(obj uintptr, aRect TRect) {
	if IsNilProcApi(canvas_DrawFocusRect) {
		canvas_DrawFocusRect = liblcl.NewProc("Canvas_DrawFocusRect")
	}
	canvas_DrawFocusRect.Call(obj, uintptr(unsafe.Pointer(&aRect)))
}

func Canvas_FillRect(obj uintptr, aRect TRect) {
	if IsNilProcApi(canvas_FillRect) {
		canvas_FillRect = liblcl.NewProc("Canvas_FillRect")
	}
	canvas_FillRect.Call(obj, uintptr(unsafe.Pointer(&aRect)))
}

func Canvas_FrameRect(obj uintptr, aRect TRect) {
	if IsNilProcApi(canvas_FrameRect) {
		canvas_FrameRect = liblcl.NewProc("Canvas_FrameRect")
	}
	canvas_FrameRect.Call(obj, uintptr(unsafe.Pointer(&aRect)))
}

func Canvas_TextRect2(obj uintptr, aRect *TRect, text string, textFormat TTextFormat) {
	if IsNilProcApi(canvas_TextRect2) {
		canvas_TextRect2 = liblcl.NewProc("Canvas_TextRect2")
	}
	canvas_TextRect2.Call(obj, uintptr(unsafe.Pointer(aRect)), GoStrToDStr(text), uintptr(textFormat))
}

func Canvas_TextRect3(obj uintptr, aRect *TRect, text string, textFormat TTextFormat) {
	if IsNilProcApi(canvas_TextRect2) {
		canvas_TextRect2 = liblcl.NewProc("Canvas_TextRect3")
	}
	canvas_TextRect2.Call(obj, uintptr(unsafe.Pointer(aRect)), GoStrToDStr(text), 0, uintptr(textFormat))
}

func Canvas_TextRect1(obj uintptr, aRect TRect, x, y int32, text string) {
	if IsNilProcApi(canvas_TextRect1) {
		canvas_TextRect1 = liblcl.NewProc("Canvas_TextRect1")
	}
	canvas_TextRect1.Call(obj, uintptr(unsafe.Pointer(&aRect)), uintptr(x), uintptr(y), GoStrToDStr(text))
}

func Canvas_Polygon(obj uintptr, points []TPoint) {
	if len(points) == 0 {
		return
	}
	if IsNilProcApi(canvas_Polygon) {
		canvas_Polygon = liblcl.NewProc("Canvas_Polygon")
	}
	canvas_Polygon.Call(obj, uintptr(unsafe.Pointer(&points[0])), uintptr(len(points)))
}

func Canvas_Polyline(obj uintptr, points []TPoint) {
	if len(points) == 0 {
		return
	}
	if IsNilProcApi(canvas_Polyline) {
		canvas_Polyline = liblcl.NewProc("Canvas_Polyline")
	}
	canvas_Polyline.Call(obj, uintptr(unsafe.Pointer(&points[0])), uintptr(len(points)))
}

func Canvas_PolyBezier(obj uintptr, points []TPoint) {
	if len(points) == 0 {
		return
	}
	if IsNilProcApi(canvas_PolyBezier) {
		canvas_PolyBezier = liblcl.NewProc("Canvas_PolyBezier")
	}
	canvas_PolyBezier.Call(obj, uintptr(unsafe.Pointer(&points[0])), uintptr(len(points)))
}
