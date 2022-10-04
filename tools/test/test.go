package main

import (
	"fmt"
	"github.com/energye/golcl/tools"
)

func main() {
	var conf conf
	tags := tools.NewFindStructTags(&conf, "json", "label", "default")
	tag := tags.GetFieldTag("AppName")
	fmt.Println(tag.Get("json"))
	fmt.Println(tag.Get("label"))
	fmt.Println(tag.Get("default"))
}

type conf struct {
	AppName     string `json:"app_name" label:"应用名称" default:"app"`
	Icon        string `json:"icon" label:"Icon图标" default:"app.ico"`
	Version     string `json:"version" label:"版本号" default:"1.0.0.0"`
	DesktopName string `json:"desktop_name" label:"桌面快捷方式名称" default:"app"`
	AppExeName  string `json:"app_exe_name" label:"应用执行文件名称" default:"app.exe"`
}
