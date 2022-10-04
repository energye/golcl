//----------------------------------------
//
// Copyright © sxm. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package libname

import (
	"runtime"
)

var (
	LibName          = ""
	platformExtNames = map[string]string{
		"windows": ".dll",
		"linux":   ".so",
		"darwin":  ".dylib",
	}
	platformExtArch = map[string]string{
		"amd64": "",
		"386":   "x86",
	}
)

func GetDLLName() string {
	libName := "liblcl"
	if ext, ok := platformExtNames[runtime.GOOS]; ok {
		if arch, ok := platformExtArch[runtime.GOARCH]; ok {
			libName += arch
		}
		return libName + ext
	}
	return libName
}
