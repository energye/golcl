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
	tempDllDir       = ""
	platformExtNames = map[string]string{
		"windows": ".dll",
		"linux":   ".so",
		"darwin":  ".dylib",
	}
)

func SetTempDllDir(dir string) {
	tempDllDir = dir
}

func GetDLLName() string {
	libName := "liblcl"
	if ext, ok := platformExtNames[runtime.GOOS]; ok {
		return libName + ext
	}
	return libName
}

// LibPath
//	获取 liblcl 动态库目录
//  优先级
//	  1. 当前目录 > 用户目录 > LCL_HOME > ENERGY_HOME
//	推荐[当前目录]或指定[ENERGY_HOME]环境变量
func LibPath() string {
	var dllName = GetDLLName()

	//tempdll内置字节码设置目录
	var tempdllPathLibName = path.Join(tempDllDir, dllName)
	if tools.IsExist(tempdllPathLibName) {
		return tempdllPathLibName
	}
	//当前目录
	var currentPathLibName = path.Join(consts.ExeDir, dllName)
	if tools.IsExist(currentPathLibName) {
		return currentPathLibName
	}
	//用户目录
	var homePathLibName = path.Join(consts.HomeGoLCLDir, dllName)
	if tools.IsExist(homePathLibName) {
		return homePathLibName
	}
	//环境变量 LCL_HOME - 为区分单独使用GO LCL
	var envPathLibName = path.Join(os.Getenv("LCL_HOME"), dllName)
	if tools.IsExist(envPathLibName) {
		return envPathLibName
	}
	//环境变量 ENERGY_HOME -  在使用Energy框架
	var energyPathLibName = path.Join(os.Getenv("ENERGY_HOME"), dllName)
	if tools.IsExist(energyPathLibName) {
		return energyPathLibName
	}
	return ""
}
