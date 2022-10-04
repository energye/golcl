package main

import "github.com/energye/golcl/lcl/win"

func GetCurrentThreadId() uintptr {
	return win.GetCurrentThreadId()
}
