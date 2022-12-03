//go:build windows
// +build windows

package main

import (
	"os"

	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/win"
	_ "github.com/energye/golcl/pkgs/winappres"
)

func main() {
	rtl.CreateURLShortCut(win.GetDesktopPath(), "govcl", "https://github.com/energye/golcl")

	rtl.CreateShortCut(win.GetDesktopPath(), "shortcut", os.Args[0], "", "描述", "-b -c")

	lcl.ShowMessage("Hello!")

}
