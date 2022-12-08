//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// energy extension
//
//govcl -> golcl/energy/inits/inits.go

package inits

import (
	"embed"
	"fmt"
	"github.com/energye/golcl/energy/consts"
	"github.com/energye/golcl/energy/emfs"
	"github.com/energye/golcl/energy/tools"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/i18n"
	"github.com/energye/golcl/lcl/locales/zh_CN"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/rtl/version"
	"github.com/energye/golcl/pkgs/libname"
	"github.com/energye/golcl/pkgs/macapp"
	"io"
	"io/fs"
	"os"
)

var (
	resourcePath = "resource"
	libsPath     = "libs"
)

func Init(libs *embed.FS, resources *embed.FS) {
	emfs.SetEMFS(libs, resources)
	libname.LibName = libname.LibPath()
	if libname.LibName == "" {
		libname.LibName = consts.HomeDir + consts.Separator + libname.GetDLLName()
		//lcllib都没有的情况
		//尝试在内置中获取-并释放到用户目录
		tools.MkdirAll(consts.HomeDir)
		releaseLib(fmt.Sprintf("%s/%s", libsPath, libname.GetDLLName()), libname.LibName)
		if tools.IsExist(libname.LibName) {
			fmt.Println(`Hint:
	Golcl dependency library liblcl was not found
	Please check whether liblcl exists locally
	If local liblcl exist, please put it in the specified location, if not please from the network to download a binary package (https://github.com/energye/energy/releases), or to compile.
	Configuration Location:
		1. Current program execution directory
		2. User root /golcl/
		3. Environment variables LCL_HOME or ENERGY_HOME
			environment variable LCL_HOME is configured preferentially in the non-energy framework
			environment variable ENERGY_HOME takes precedence in the Energy framework
			ENERGY_HOME environment variable is recommended
`)
		}
	}
	initAll()
}

func releaseLib(fsPath, out string) {
	if emfs.GetLibsFS() != nil {
		var err error
		var fsFile fs.File
		var file *os.File
		file, err = os.OpenFile(out, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		fsFile, err = emfs.GetLibsFS().Open(fsPath)
		if err == nil {
			defer fsFile.Close()
			var n int
			//读取数据
			buf := make([]byte, 4096)
			for {
				n, err = fsFile.Read(buf)
				if err == io.EOF {
					break
				}
				file.Write(buf[:n])
			}
			fmt.Println("release success.")
		}
	}
}

func initAll() {
	//macapp
	macapp.MacApp.Init()
	//api
	api.APIInit()
	//rtl
	rtl.RtlInit()
	//version
	version.VersionInit()
	//zh_cn
	zh_CN.ZH_CNInit()
	//win
	winInit()
	//lcl
	lcl.LCLInit()
	//i18n
	i18n.I18NInit()
	//
	//	lcl.ApplicationQueueAsyncCallInit()
}
