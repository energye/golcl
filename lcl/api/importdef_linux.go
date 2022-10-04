//go:build linux
// +build linux

package api

import "github.com/energye/dylib"

var (
	GdkWindow_GetXId      *dylib.LazyProc
	GdkWindow_FromForm    *dylib.LazyProc
	GtkWidget_GetGtkFixed *dylib.LazyProc
	GtkWidget_Window      *dylib.LazyProc
)
