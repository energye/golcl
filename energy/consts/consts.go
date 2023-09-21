//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package consts

import (
	"github.com/energye/golcl/energy/homedir"
	"os"
	"path/filepath"
)

const (
	RootPath = "golcl"
)

var (
	ExeDir       string                       // 执行文件所在目录
	ExePath      string                       // 执行文件完整目录
	ExeName      string                       // 执行文件名称
	HomeDir      string                       // 用户目录
	Separator    = string(filepath.Separator) // 目录结构分割符
	HomeGoLCLDir string                       // 完整的用户目录
)

func init() {
	ExePath = os.Args[0]
	ExeDir, ExeName = filepath.Split(ExePath)
	//用户目录
	HomeDir, _ = homedir.Dir()
	HomeGoLCLDir = filepath.Join(HomeDir, RootPath)
}
