//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package main

import _ "github.com/energye/golcl/samples/customLibTest/alib"
import "github.com/energye/golcl/lcl"

type TMainForm struct {
	*lcl.TForm
}

var (
	mainForm *TMainForm
)

func main() {

	lcl.RunApp(&mainForm)
}
