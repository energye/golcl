//----------------------------------------
//
// Copyright © sxm. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------
//go:build darwin
// +build darwin

package lcl

/*

   #cgo CFLAGS: -x objective-c
   #cgo LDFLAGS: -framework Cocoa

   #include <Cocoa/Cocoa.h>

  static NSWindow* toNSWindow(void* ptr) {
      return ((__bridge NSWindow*)ptr); //
  }

  // test
  static void NSWindow_setTitleVisibility(void *ptr) {
     NSWindow *win = toNSWindow(ptr);
     win.TitleVisibility = NSWindowTitleHidden;
  }

*/
//import "C"

import (
	. "github.com/energye/golcl/lcl/api"
	. "github.com/energye/golcl/lcl/types"
)

type (
	NSObject uintptr

	NSWindow uintptr

	NSURL uintptr
)

// NSObject
func HandleToPlatformHandle(h HWND) NSObject {
	return NSObject(h)
}

func (f *TForm) PlatformWindow() NSWindow {
	if NSWindow_FromForm == nil {
		NSWindow_FromForm = GetLibLCL().NewProc("NSWindow_FromForm")
	}
	r, _, _ := NSWindow_FromForm.Call(f.instance)
	return NSWindow(r)
}

func (n NSWindow) TitleVisibility() NSWindowTitleVisibility {
	if NSWindow_titleVisibility == nil {
		NSWindow_titleVisibility = GetLibLCL().NewProc("NSWindow_titleVisibility")
	}
	r, _, _ := NSWindow_titleVisibility.Call(uintptr(n))
	return NSWindowTitleVisibility(r)
}

func (n NSWindow) SetTitleVisibility(flag NSWindowTitleVisibility) {
	//C.NSWindow_setTitleVisibility(unsafe.Pointer(n))
	if NSWindow_setTitleVisibility == nil {
		NSWindow_setTitleVisibility = GetLibLCL().NewProc("NSWindow_setTitleVisibility")
	}
	NSWindow_setTitleVisibility.Call(uintptr(n), uintptr(flag))
}

func (n NSWindow) TitleBarAppearsTransparent() bool {
	if NSWindow_titlebarAppearsTransparent == nil {
		NSWindow_titlebarAppearsTransparent = GetLibLCL().NewProc("NSWindow_titlebarAppearsTransparent")
	}
	r, _, _ := NSWindow_titlebarAppearsTransparent.Call(uintptr(n))
	return DBoolToGoBool(r)
}

func (n NSWindow) SetTitleBarAppearsTransparent(flag bool) {
	if NSWindow_setTitlebarAppearsTransparent == nil {
		NSWindow_setTitlebarAppearsTransparent = GetLibLCL().NewProc("NSWindow_setTitlebarAppearsTransparent")
	}
	NSWindow_setTitlebarAppearsTransparent.Call(uintptr(n), GoBoolToDBool(flag))
}

func (n NSWindow) SetRepresentedURL(url NSURL) {
	if NSWindow_setRepresentedURL == nil {
		NSWindow_setRepresentedURL = GetLibLCL().NewProc("NSWindow_setRepresentedURL")
	}
	NSWindow_setRepresentedURL.Call(uintptr(n), uintptr(url))
}

func (n NSWindow) StyleMask() uint {
	if NSWindow_styleMask == nil {
		NSWindow_styleMask = GetLibLCL().NewProc("NSWindow_styleMask")
	}
	r, _, _ := NSWindow_styleMask.Call(uintptr(n))
	return uint(r)
}

func (n NSWindow) SetStyleMask(mask uint) {
	if NSWindow_setStyleMask == nil {
		NSWindow_setStyleMask = GetLibLCL().NewProc("NSWindow_setStyleMask")
	}
	NSWindow_setStyleMask.Call(uintptr(n), uintptr(mask))
}
