package inits

import (
	"embed"
	"fmt"
	"github.com/energye/golcl/consts"
	"github.com/energye/golcl/emfs"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/i18n"
	"github.com/energye/golcl/lcl/locales/zh_CN"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/rtl/version"
	"github.com/energye/golcl/lcl/win"
	"github.com/energye/golcl/pkgs/libname"
	"github.com/energye/golcl/pkgs/macapp"
	"github.com/energye/golcl/tools"
	"io"
	"io/fs"
	"os"
)

var (
	resourcePath = "resource"
	libsPath     = "libs"
	isInit       = false
)

func libPath() string {
	//当前目录
	var currentPathLibName = consts.ExePath + consts.Separator + libname.GetDLLName()
	if tools.IsExist(currentPathLibName) {
		return currentPathLibName
	}
	//用户目录
	var homePathLibName = consts.HomeDir + consts.Separator + libname.GetDLLName()
	if tools.IsExist(currentPathLibName) {
		return homePathLibName
	}
	//环境变量
	var envPathLibName = os.Getenv("LCL_HOME") + consts.Separator + libname.GetDLLName()
	if tools.IsExist(envPathLibName) {
		return envPathLibName
	}
	var energyPathLibName = os.Getenv("ENERGY_HOME") + consts.Separator + libname.GetDLLName()
	if tools.IsExist(energyPathLibName) {
		return energyPathLibName
	}
	return ""
}

func Init(libs *embed.FS, resources *embed.FS) bool {
	emfs.SetLibsFS(libs)
	emfs.SetResourcesFS(resources)

	libname.LibName = libPath()
	if libname.LibName == "" {
		//lcllib都没有的情况
		//尝试在内置中获取-并释放到用户目录
		tools.MkdirAll(consts.HomeDir)
		releaseLib(fmt.Sprintf("%s/%s", libsPath, libname.GetDLLName()), libname.LibName)
	}
	initAll()
	return true
}

func releaseResource() {

}

func releaseLib(path, dllName string) {
	if emfs.GetLibsFS() != nil {
		var err error
		var fsFile fs.File
		var file *os.File
		file, err = os.OpenFile(dllName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			panic(err)
		}
		fsFile, err = emfs.GetLibsFS().Open(path)
		if err != nil {
			fmt.Println("emfs error:", err)
			file.Close()
			os.Remove(dllName)
		} else {
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
			fsFile.Close()
			file.Close()
			fmt.Println("liblcl", libname.LibName, "release success.")
		}
	}
}

func initAll() {
	if isInit {
		return
	}
	isInit = true
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
	win.Init()
	//lcl
	lcl.LCLInit()
	lcl.ApplicationQueueAsyncCallInit()
	//i18n
	i18n.I18NInit()

}
