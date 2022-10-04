//go:build !windows
// +build !windows

package api

import (
	"github.com/energye/dylib"
	"unsafe"

	. "github.com/energye/golcl/lcl/types"
)

var (
	dSendMessage         *dylib.LazyProc
	dPostMessage         *dylib.LazyProc
	dIsIconic            *dylib.LazyProc
	dIsWindow            *dylib.LazyProc
	dIsZoomed            *dylib.LazyProc
	dIsWindowVisible     *dylib.LazyProc
	dGetDC               *dylib.LazyProc
	dReleaseDC           *dylib.LazyProc
	dSetForegroundWindow *dylib.LazyProc
	dWindowFromPoint     *dylib.LazyProc
)

func DSendMessage(hWd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	if dSendMessage == nil {
		dSendMessage = liblcl.NewProc("DSendMessage")
	}
	r, _, _ := dSendMessage.Call(hWd, uintptr(msg), wParam, lParam)
	return r
}

func DPostMessage(hWd HWND, msg uint32, wParam, lParam uintptr) bool {
	if dPostMessage == nil {
		dPostMessage = liblcl.NewProc("DPostMessage")
	}
	r, _, _ := dPostMessage.Call(hWd, uintptr(msg), wParam, lParam)
	return r != 0
}

func DIsIconic(hWnd HWND) bool {
	if dIsIconic == nil {
		dIsIconic = liblcl.NewProc("DIsIconic")
	}
	r, _, _ := dIsIconic.Call(uintptr(hWnd))
	return r != 0
}

func DIsWindow(hWnd HWND) bool {
	if dIsWindow == nil {
		dIsWindow = liblcl.NewProc("DIsWindow")
	}
	r, _, _ := dIsWindow.Call(uintptr(hWnd))
	return r != 0
}

func DIsZoomed(hWnd HWND) bool {
	if dIsZoomed == nil {
		dIsZoomed = liblcl.NewProc("DIsZoomed")
	}
	r, _, _ := dIsZoomed.Call(uintptr(hWnd))
	return r != 0
}

func DIsWindowVisible(hWnd HWND) bool {
	if dIsWindowVisible == nil {
		dIsWindowVisible = liblcl.NewProc("DIsWindowVisible")
	}
	r, _, _ := dIsWindowVisible.Call(uintptr(hWnd))
	return r != 0
}

func DGetDC(hWnd HWND) HDC {
	if dGetDC == nil {
		dGetDC = liblcl.NewProc("DGetDC")
	}
	r, _, _ := dGetDC.Call(uintptr(hWnd))
	return HDC(r)
}

func DReleaseDC(hWnd HWND, dc HDC) int {
	if dReleaseDC == nil {
		dReleaseDC = liblcl.NewProc("DReleaseDC")
	}
	r, _, _ := dReleaseDC.Call(uintptr(hWnd), uintptr(dc))
	return int(r)
}

func DSetForegroundWindow(hWnd HWND) bool {
	if dSetForegroundWindow == nil {
		dSetForegroundWindow = liblcl.NewProc("DSetForegroundWindow")
	}
	r, _, _ := dSetForegroundWindow.Call(uintptr(hWnd))
	return r != 0
}

func DWindowFromPoint(point TPoint) HWND {
	if dWindowFromPoint == nil {
		dWindowFromPoint = liblcl.NewProc("DWindowFromPoint")
	}
	r, _, _ := dWindowFromPoint.Call(uintptr(unsafe.Pointer(&point)))
	return HWND(r)
}
