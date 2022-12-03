package consts

import (
	"fmt"
	"github.com/energye/golcl/energy/homedir"
	"os"
	"path/filepath"
)

const (
	RootPath = "golcl"
)

var (
	ExePath string //执行文件目录
	//用户目录
	DIR, err = homedir.Dir()
	//目录结构分割符
	Separator = string(filepath.Separator)
	//完整的用户目录
	HomeDir = fmt.Sprintf("%v%v%v", DIR, Separator, RootPath)
)

func init() {
	ExePath, err = os.Getwd()
	if err != nil {
		panic(err)
	}
}
