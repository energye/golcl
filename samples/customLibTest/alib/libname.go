//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package alib

import (
	"fmt"

	"github.com/energye/golcl/pkgs/libname"
)

func init() {
	fmt.Println("init")
	libname.LibName = "aaa"
}
