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
	"github.com/energye/golcl/samples/fileshelltree/src"
)
import _ "github.com/energye/golcl/pkgs/winappres"

func main() {
	inits.Init(nil, nil)
	lcl.RunApp(&src.Form1)
}
