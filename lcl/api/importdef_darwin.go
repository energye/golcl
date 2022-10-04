//----------------------------------------
//
// Copyright © sxm. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build darwin
// +build darwin

package api

import "github.com/energye/dylib"

var (
	NSWindow_FromForm                      *dylib.LazyProc
	NSWindow_titleVisibility               *dylib.LazyProc
	NSWindow_setTitleVisibility            *dylib.LazyProc
	NSWindow_titlebarAppearsTransparent    *dylib.LazyProc
	NSWindow_setTitlebarAppearsTransparent *dylib.LazyProc
	NSWindow_styleMask                     *dylib.LazyProc
	NSWindow_setStyleMask                  *dylib.LazyProc
	NSWindow_setRepresentedURL             *dylib.LazyProc
	//NSWindow_release                       = liblcl.NewProc("NSWindow_release")
)
