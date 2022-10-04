package inits

import (
	"embed"
	"fmt"
	"github.com/energye/golcl/consts"
	"github.com/energye/golcl/emfs"
	"github.com/energye/golcl/lcl"
	api "github.com/energye/golcl/lcl/api"
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

func Init(libs *embed.FS, resources *embed.FS) bool {
	if libname.LibName == "" {
		libname.LibName = consts.HomeDir + consts.Separator + libname.GetDLLName()
	}
	//fmt.Println("emfs by libs:", libs, "\tresources:", resources)
	emfs.SetLibsFS(libs)
	emfs.SetResourcesFS(resources)

	tools.MkdirAll(consts.HomeDir)
	if !tools.IsExist(libname.LibName) {
		fmt.Println("liblcl", libname.LibName, " is not exist.")
		var path = libsPath + "/" + libname.GetDLLName()
		releaseLib(path, libname.LibName)
	}
	//if !tools.IsExist(resourcePath) {
	//	fmt.Println("resources is not exist.")
	//	releaseResource()
	//	fmt.Println("resources release success.")
	//}
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
	//i18n
	i18n.I18NInit()

}
