//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package main

import (
	"github.com/energye/golcl/lcl"
	_ "github.com/energye/golcl/lcl/locales/zh_CN"
	_ "github.com/energye/golcl/pkgs/winappres"
)

func main() {
	lcl.RunApp(&MainForm, &NewConnection)
}
