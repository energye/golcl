//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// energy extension
//
//govcl/pkgs/libname/lib.go -> golcl/pkgs/libname/lib.go

package libname

import (
	"github.com/energye/golcl/energy/consts"
	"github.com/energye/golcl/energy/tools"
	"os"
	"runtime"
)

var (
	LibName          = ""
	platformExtNames = map[string]string{
		"windows": ".dll",
		"linux":   ".so",
		"darwin":  ".dylib",
	}
)

func GetDLLName() string {
	libName := "liblcl"
	if ext, ok := platformExtNames[runtime.GOOS]; ok {
		return libName + ext
	}
	return libName
}

func LibPath() string {
	//当前目录
	var currentPathLibName = consts.ExePath + consts.Separator + GetDLLName()
	if tools.IsExist(currentPathLibName) {
		return currentPathLibName
	}
	//用户目录
	var homePathLibName = consts.HomeDir + consts.Separator + GetDLLName()
	if tools.IsExist(homePathLibName) {
		return homePathLibName
	}
	//环境变量 LCL_HOME
	var envPathLibName = os.Getenv("LCL_HOME") + consts.Separator + GetDLLName()
	if tools.IsExist(envPathLibName) {
		return envPathLibName
	}
	//环境变量 ENERGY_HOME
	var energyPathLibName = os.Getenv("ENERGY_HOME") + consts.Separator + GetDLLName()
	if tools.IsExist(energyPathLibName) {
		return energyPathLibName
	}
	return ""
}
