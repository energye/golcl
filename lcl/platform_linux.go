//----------------------------------------
//
// Copyright © sxm. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build linux
// +build linux

package lcl

import (
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

type GdkWindow uintptr

type GtkWidget uintptr

type GtkFixed uintptr

type XID uintptr

// PGtkWidget
func HandleToPlatformHandle(h types.HWND) GtkWidget {
	return GtkWidget(h)
}

func (f *TForm) PlatformWindow() GdkWindow {
	if api.GdkWindow_FromForm == nil {
		api.GdkWindow_FromForm = api.GetLibLCL().NewProc("GdkWindow_FromForm")
	}
	r, _, _ := api.GdkWindow_FromForm.Call(f.instance)
	return GdkWindow(r)
}

func (g GdkWindow) XID() (xid XID) {
	if api.GdkWindow_GetXId == nil {
		api.GdkWindow_GetXId = api.GetLibLCL().NewProc("GdkWindow_GetXId")
	}
	api.GdkWindow_GetXId.Call(uintptr(g), uintptr(unsafe.Pointer(&xid)))
	return
}

// lz中首先是一个widget，然后上面用了一个fixedWidget来处理的。
func (g GtkWidget) FixedWidget() GtkFixed {
	if api.GtkWidget_GetGtkFixed == nil {
		api.GtkWidget_GetGtkFixed = api.GetLibLCL().NewProc("GtkWidget_GetGtkFixed")
	}
	r, _, _ := api.GtkWidget_GetGtkFixed.Call(uintptr(g))
	return GtkFixed(r)
}

func (g GtkWidget) Window() GdkWindow {
	if api.GtkWidget_Window == nil {
		api.GtkWidget_Window = api.GetLibLCL().NewProc("GtkWidget_Window")
	}
	r, _, _ := api.GtkWidget_Window.Call(uintptr(g))
	return GdkWindow(r)
}
