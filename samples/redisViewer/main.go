//----------------------------------------
//
// Copyright © sxm. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package main

import (
	"github.com/energye/golcl/inits"
	"github.com/energye/golcl/lcl"
	_ "github.com/energye/golcl/lcl/locales/zh_CN"
	_ "github.com/energye/golcl/pkgs/winappres"
)

func main() {
	inits.Init(nil, nil)
	lcl.RunApp(&MainForm, &NewConnection)
}
