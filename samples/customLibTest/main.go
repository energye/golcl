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
	_ "github.com/energye/golcl/samples/customLibTest/alib"
)
import "github.com/energye/golcl/lcl"

type TMainForm struct {
	*lcl.TForm
}

var (
	mainForm *TMainForm
)

func main() {
	inits.Init(nil, nil)
	lcl.RunApp(&mainForm)
}
