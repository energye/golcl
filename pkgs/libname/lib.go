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
	"path"
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
	var dllName = GetDLLName()
	//当前目录
	var currentPathLibName = path.Join(consts.ExePath, dllName)
	if tools.IsExist(currentPathLibName) {
		return currentPathLibName
	}
	//用户目录
	var homePathLibName = path.Join(consts.HomeDir, dllName)
	if tools.IsExist(homePathLibName) {
		return homePathLibName
	}
	//环境变量 LCL_HOME
	var envPathLibName = path.Join(os.Getenv("LCL_HOME"), dllName)
	if tools.IsExist(envPathLibName) {
		return envPathLibName
	}
	//环境变量 ENERGY_HOME
	var energyPathLibName = path.Join(os.Getenv("ENERGY_HOME"), dllName)
	if tools.IsExist(energyPathLibName) {
		return energyPathLibName
	}
	return ""
}
