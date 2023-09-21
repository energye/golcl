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
	CurrentExecuteDir string                       // 当前执行目录
	ExeDir            string                       // 执行文件所在目录
	ExePath           string                       // 执行文件完整目录
	ExeName           string                       // 执行文件名称
	HomeDir           string                       // 用户目录
	Separator         = string(filepath.Separator) // 目录结构分割符
	HomeGoLCLDir      string                       // 完整的用户目录
)

func init() {
	// 当前执行目录，在其它目录执行目标执行文件时，返回当前执行目录
	CurrentExecuteDir, _ = os.Getwd()
	// 执行文件所在完整目录
	ExePath = os.Args[0]
	// 拆分完整目录，返回目录名 + 文件名
	ExeDir, ExeName = filepath.Split(ExePath)
	// 用户目录
	HomeDir, _ = homedir.Dir()
	// golcl 目录
	HomeGoLCLDir = filepath.Join(HomeDir, RootPath)
}
